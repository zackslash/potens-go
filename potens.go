package potens

import (
	"crypto/tls"
	"errors"
	"flag"
	"log"
	"math/rand"
	"net"
	"path"
	"runtime"
	"strconv"
	"time"

	"github.com/fortifi/potens-go/definition"
	"github.com/fortifi/potens-go/identity"
	"github.com/fortifi/proto-go/discovery"
	"github.com/fortifi/proto-go/imperium"
	"github.com/satori/go.uuid"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

var (
	hostname            string
	port                = flag.Int("service-port", 0, "grpc service port")
	discoveryService    = flag.String("discovery-service", "discovery.fortifi.me:50056", "Fortifi App Discovery Service")
	discoveryConn       *grpc.ClientConn
	discoClient         discovery.DiscoveryClient
	imperiumService     = flag.String("imperium-service", "imperium.fortifi.me:50055", "Fortifi Imperium Service")
	imperiumCertificate []byte
	imperiumKey         []byte
	appDefinition       *definition.AppDefinition
	appIdentity         *identity.AppIdentity
	instanceID          = uuid.NewV4().String()
	currentStatus       = discovery.ServiceStatus_OFFLINE
)

func relPath(file string) string {
	_, filename, _, _ := runtime.Caller(2)
	return path.Join(path.Dir(filename), file)
}

// Start your service, retrieves tls Certificate to server, and registers with discovery service
func Start(appDef *definition.AppDefinition, appIdent *identity.AppIdentity) error {

	if appIdent == nil {
		appIdent = &identity.AppIdentity{}
		err := appIdent.FromJSONFile(relPath("app-identity.json"))
		if err != nil {
			return err
		}
	}

	if appDef == nil {
		appDef = &definition.AppDefinition{}
		err := appDef.FromConfig(relPath("app-definition.yaml"))
		if err != nil {
			return err
		}
	}

	if len(appDef.Vendor) < 2 {
		log.Fatal("The Vendor ID specified in your definition file is invalid")
	}

	if len(appDef.AppID) < 2 {
		log.Fatal("The App ID specified in your definition file is invalid")
	}

	if appIdent.AppID != appDef.GlobalAppID {
		log.Fatal("The App ID in your definition file does not match your identity file")
	}

	appDefinition = appDef
	appIdentity = appIdent

	log.Print("Starting App: " + appDef.GlobalAppID + " - " + appDef.Name)
	log.Print("Authing with: " + appIdent.IdentityID + " - " + appIdent.IdentityType)

	flag.Parse()
	log.SetFlags(0)

	usePort := *port
	if usePort == 0 {
		minPort := 50060
		maxPort := 55555
		rand.Seed(time.Now().UTC().UnixNano())
		usePort = rand.Intn(maxPort-minPort) + minPort
		port = &usePort
	}

	err := getCerts()
	if err != nil {
		return err
	}

	discoveryConn, err = grpc.Dial(*discoveryService, grpc.WithTransportCredentials(credentials.NewClientTLSFromCert(nil, "")))

	if err != nil {
		log.Fatal(err)
	}

	discoClient = discovery.NewDiscoveryClient(discoveryConn)
	regResult, err := discoClient.Register(context.Background(), &discovery.RegisterRequest{
		AppId:        appDef.GlobalAppID,
		InstanceUuid: instanceID,
		ServiceHost:  hostname,
		ServicePort:  int32(*port),
	})

	if err != nil {
		log.Fatal(err)
	}

	if !regResult.Recorded {
		log.Fatal("Failed to register with the discovery service")
	} else {
		log.Print("Registered with Discovery Service")
	}
	return nil
}

func heartBeat() {
	if currentStatus == discovery.ServiceStatus_ONLINE {
		for {
			discoClient.HeartBeat(context.Background(), &discovery.HeartBeatRequest{
				AppId:        appDefinition.GlobalAppID,
				InstanceUuid: instanceID,
			})
			time.Sleep(10 * time.Second)
		}
	}
}

// Online take your service online
func Online() error {
	statusResult, err := discoClient.Status(context.Background(), &discovery.StatusRequest{
		AppId:        appDefinition.GlobalAppID,
		InstanceUuid: instanceID,
		Status:       discovery.ServiceStatus_ONLINE,
		Target:       discovery.StatusTarget_BOTH,
	})

	if err != nil {
		return err
	}

	if !statusResult.Recorded {
		return errors.New("Failed to go online")
	}

	currentStatus = discovery.ServiceStatus_ONLINE

	go heartBeat()
	return nil
}

// Offline take your service offline
func Offline() error {
	statusResult, err := discoClient.Status(context.Background(), &discovery.StatusRequest{
		AppId:        appDefinition.GlobalAppID,
		InstanceUuid: instanceID,
		Status:       discovery.ServiceStatus_OFFLINE,
		Target:       discovery.StatusTarget_INSTANCE,
	})

	if err != nil {
		return err
	}

	if !statusResult.Recorded {
		return errors.New("Failed to go offline")
	}

	currentStatus = discovery.ServiceStatus_OFFLINE
	return nil
}

func getCerts() error {

	imperiumConnection, err := grpc.Dial(*imperiumService, grpc.WithTransportCredentials(credentials.NewClientTLSFromCert(nil, "")))
	if err != nil {
		return err
	}
	c := imperium.NewImperiumClient(imperiumConnection)
	response, err := c.Request(context.Background(), &imperium.CertificateRequest{
		AppId: appDefinition.GlobalAppID,
	})
	if err != nil {
		return err
	}

	hostname = response.Hostname
	imperiumCertificate = []byte(response.Certificate)
	imperiumKey = []byte(response.PrivateKey)

	log.Print("Received TLS Certificates for " + hostname)

	return nil
}

// CreateServer creates a gRPC server with your tls certificates
func CreateServer() (net.Listener, *grpc.Server, error) {

	lis, err := net.Listen("tcp", hostname+":"+strconv.FormatInt(int64(*port), 10))
	if err != nil {
		return nil, nil, err
	}

	cert, err := tls.X509KeyPair(imperiumCertificate, imperiumKey)
	if err != nil {
		return nil, nil, err
	}

	s := grpc.NewServer(grpc.Creds(credentials.NewServerTLSFromCert(&cert)))

	return lis, s, nil
}

// Identity retrieves your identity
func Identity() *identity.AppIdentity {
	return appIdentity
}

// Definition retrieves your definition
func Definition() *definition.AppDefinition {
	return appDefinition
}

// GetAppConnection grpc.dial a service based on the discovery service
func GetAppConnection(globalAppID string, opts ...grpc.DialOption) (*grpc.ClientConn, error) {

	locationResult, err := discoClient.GetLocation(context.Background(), &discovery.LocationRequest{AppId: globalAppID})

	if err != nil {
		return nil, err
	}

	opts = append(opts, grpc.WithTransportCredentials(credentials.NewClientTLSFromCert(nil, "")))

	return grpc.Dial(locationResult.ServiceHost+":"+strconv.FormatInt(int64(locationResult.ServicePort), 10), opts...)
}
