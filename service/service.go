package service

import (
	"crypto/rsa"
	"crypto/tls"
	"crypto/x509"
	"encoding/pem"
	"errors"
	"flag"
	"log"
	"math/rand"
	"net"
	"strconv"
	"time"

	"golang.org/x/net/context"

	"path"

	"os"

	"path/filepath"

	"github.com/dgrijalva/jwt-go"
	"github.com/fident/proto-go/fident"
	"github.com/fortifi/portcullis-go/keys"
	"github.com/fortifi/potens-go/definition"
	"github.com/fortifi/potens-go/i18n"
	"github.com/fortifi/potens-go/identity"
	"github.com/fortifi/proto-go/discovery"
	"github.com/fortifi/proto-go/imperium"
	"github.com/fortifi/proto-go/undercroft"
	"github.com/satori/go.uuid"
	"github.com/uber-go/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/metadata"
	yaml "gopkg.in/yaml.v2"
)

var (
	parseEnv = flag.Bool("parse-env", true, "Set to false to use production defaults")
)

type FortifiService struct {
	appDefinition       *definition.AppDefinition
	appIdentity         *identity.AppIdentity
	port                int32
	discoClient         discovery.DiscoveryClient
	undercroftClient    undercroft.UndercroftClient
	fidentClient        fident.AuthClient
	imperiumCertificate []byte
	imperiumKey         []byte
	hostname            string
	instanceID          string
	currentStatus       discovery.ServiceStatus

	//Logger used for standard logging
	Logger zap.Logger

	parsedEnv        bool
	imperiumService  string
	discoveryService string
	registryService  string
	fidentService    string
	authToken        string
	pk               *rsa.PrivateKey
	kh               string
}

func (s *FortifiService) parseEnv() {
	defaultPort := "50051"
	fortDomain := os.Getenv("FORT_DOMAIN")
	if fortDomain == "" {
		fortDomain = "fortifi.services"
	}

	s.discoveryService = os.Getenv("FORT_DISCOVERY_LOCATION")
	if s.discoveryService == "" {
		s.discoveryService = "discovery-fortifi." + fortDomain
	}

	discoPort := os.Getenv("FORT_DISCOVERY_PORT")
	if discoPort == "" {
		s.discoveryService += ":" + defaultPort
	} else {
		s.discoveryService += ":" + discoPort
	}

	s.imperiumService = os.Getenv("FORT_IMPERIUM_LOCATION")
	if s.imperiumService == "" {
		s.imperiumService = "imperium-fortifi." + fortDomain
	}

	imperiumPort := os.Getenv("FORT_IMPERIUM_PORT")
	if imperiumPort == "" {
		s.imperiumService += ":" + defaultPort
	} else {
		s.imperiumService += ":" + imperiumPort
	}

	s.registryService = os.Getenv("FORT_REGISTRY_LOCATION")
	if s.registryService == "" {
		s.registryService = "registry-fortifi." + fortDomain
	}

	registryPort := os.Getenv("FORT_REGISTRY_PORT")
	if registryPort == "" {
		s.registryService += ":" + defaultPort
	} else {
		s.registryService += ":" + registryPort
	}

	s.fidentService = os.Getenv("FIDENT_LOCATION")
	if s.fidentService == "" {
		s.fidentService = "api.fident.io"
	}

	fidentPort := os.Getenv("FIDENT_PORT")
	if fidentPort == "" {
		s.fidentService += ":" + defaultPort
	} else {
		s.fidentService += ":" + fidentPort
	}

	s.parsedEnv = true
}

func (s *FortifiService) SetPort(port int32) {
	s.port = port
}

func (s *FortifiService) GetPort() int32 {
	return s.port
}

func (s *FortifiService) SetDiscoveryClient(discoClient discovery.DiscoveryClient) {
	s.discoClient = discoClient
}

func (s *FortifiService) relPath(file string) string {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		s.Logger.Fatal(err.Error())
	}
	return path.Join(dir, file)
}

//New get a new instance of a service
func New(appDef *definition.AppDefinition, appIdent *identity.AppIdentity) (FortifiService, error) {

	s := FortifiService{}
	s.Logger = zap.New(zap.NewJSONEncoder())

	if !s.parsedEnv && *parseEnv {
		s.parseEnv()
		s.Logger.Debug("Parsed Environment")
	}

	s.instanceID = uuid.NewV4().String()
	s.currentStatus = discovery.ServiceStatus_OFFLINE

	if appIdent == nil {
		appIdent = &identity.AppIdentity{}
		err := appIdent.FromJSONFile(s.relPath("app-identity.json"))
		if err != nil {
			return nil, err
		}
	}

	if appDef == nil {
		appDef = &definition.AppDefinition{}
		err := appDef.FromConfig(s.relPath("app-definition.yaml"))
		if err != nil {
			return nil, err
		}
	}

	if len(appDef.Vendor) < 2 {
		s.Logger.Fatal("The Vendor ID specified in your definition file is invalid")
	}

	if len(appDef.AppID) < 2 {
		s.Logger.Fatal("The App ID specified in your definition file is invalid")
	}

	if appIdent.AppID != appDef.GlobalAppID {
		s.Logger.Fatal("The App ID in your definition file does not match your identity file")
	}

	s.appDefinition = appDef
	s.appIdentity = appIdent

	block, _ := pem.Decode([]byte(appIdent.PrivateKey))
	if block == nil {
		s.Logger.Fatal("No RSA private key found")
	}

	var key *rsa.PrivateKey
	if block.Type == "RSA PRIVATE KEY" {
		rsapk, err := x509.ParsePKCS1PrivateKey(block.Bytes)
		if err != nil {
			s.Logger.Fatal("Unable to read RSA private key")
		}
		key = rsapk
	}
	s.pk = key
	s.kh = s.appIdentity.KeyHandle

	return s, nil
}

// Start your service, retrieves tls Certificate to server, and registers with discovery service
func (s *FortifiService) Start() error {

	s.Logger.Info("Starting App: " + s.appDefinition.GlobalAppID + " - " + i18n.NewTranslatable(s.appDefinition.Name).Get("en"))
	s.Logger.Info("Authing with: " + s.appIdentity.IdentityID + " - " + s.appIdentity.IdentityType)

	flag.Parse()
	log.SetFlags(0)

	if s.port == 0 {
		minPort := 50060
		maxPort := 55555
		rand.Seed(time.Now().UTC().UnixNano())
		s.port = int32(rand.Intn(maxPort-minPort) + minPort)
	}

	err := s.getCerts()
	if err != nil {
		return err
	}

	if s.fidentClient == nil {
		authconn, err := grpc.Dial(s.fidentService, grpc.WithTransportCredentials(credentials.NewClientTLSFromCert(nil, "")))
		if err != nil {
			s.Logger.Fatal(err)
		}
		s.fidentClient = fident.NewAuthClient(authconn)
	}

	// perform auth
	ac, err := s.fidentClient.GetAuthenticationChallenge(s.GetGrpcContext(), &fident.AuthChallengePayload{Username: s.appDefinition.GlobalAppID})
	if err != nil {
		s.Logger.Fatal(err)
	}

	// TODO: Verify challenge is from Fident using fident public key - !TO BE DISTRIBUTED! (?)
	/*token, err := jwt.Parse(ac.Challenge, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		pubKey, err := ioutil.ReadFile(rsaPubKeyLocation)
		if err != nil {
			return nil, fmt.Errorf("Unable to read public key")
		}

		key, err := jwt.ParseRSAPublicKeyFromPEM(pubKey)
		if err != nil {
			return nil, fmt.Errorf("Unable to parse public key")
		}
		return key, nil
	})*/

	// Sign challenge
	challengeResponseToken := jwt.NewWithClaims(jwt.SigningMethodRS256, jwt.MapClaims{
		"challenge_token": ac.Challenge,
	})

	response, err := challengeResponseToken.SignedString(s.pk)
	if err != nil {
		s.Logger.Fatal("Unable to generate challenge response")
	}

	authres, err := s.fidentClient.PerformAuthentication(s.GetGrpcContext(), &fident.PerformAuthPayload{Username: s.appDefinition.GlobalAppID, ChallengeResponse: response})
	if err != nil {
		return err
	}

	s.authToken = authres.Token

	if s.discoClient == nil {
		discoveryConn, err := grpc.Dial(s.discoveryService, grpc.WithTransportCredentials(credentials.NewClientTLSFromCert(nil, "")))

		if err != nil {
			s.Logger.Fatal(err)
		}

		s.discoClient = discovery.NewDiscoveryClient(discoveryConn)
	}

	if s.undercroftClient == nil {
		regConn, err := grpc.Dial(s.registryService, grpc.WithTransportCredentials(credentials.NewClientTLSFromCert(nil, "")))

		if err != nil {
			s.Logger.Fatal(err)
		}

		s.undercroftClient = undercroft.NewUndercroftClient(regConn)
	}

	//TODO: Remove this once CLI tools are available
	appDefYaml, err := yaml.Marshal(s.appDefinition)
	if err != nil {
		s.Logger.Fatal(err)
	}

	s.undercroftClient.RegisterApp(s.GetGrpcContext(), &undercroft.AppRegisterRequest{
		VendorId:       s.appDefinition.Vendor,
		Id:             s.appDefinition.AppID,
		VendorSecret:   "",
		DefinitionYaml: string(appDefYaml),
	})

	regResult, err := s.discoClient.Register(s.GetGrpcContext(), &discovery.RegisterRequest{
		AppId:        s.appDefinition.GlobalAppID,
		InstanceUuid: s.instanceID,
		ServiceHost:  s.hostname,
		ServicePort:  s.port,
	})

	if err != nil {
		s.Logger.Fatal(err)
	}

	if !regResult.Recorded {
		s.Logger.Fatal("Failed to register with the discovery service")
	} else {
		s.Logger.Info("Registered with Discovery Service")
	}
	return nil
}

func (s *FortifiService) heartBeat() {
	if s.currentStatus == discovery.ServiceStatus_ONLINE {
		for {
			s.discoClient.HeartBeat(s.GetGrpcContext(), &discovery.HeartBeatRequest{
				AppId:        s.appDefinition.GlobalAppID,
				InstanceUuid: s.instanceID,
			})
			time.Sleep(10 * time.Second)
		}
	}
}

// Online take your service online
func (s *FortifiService) Online() error {
	statusResult, err := s.discoClient.Status(s.GetGrpcContext(), &discovery.StatusRequest{
		AppId:        s.appDefinition.GlobalAppID,
		InstanceUuid: s.instanceID,
		Status:       discovery.ServiceStatus_ONLINE,
		Target:       discovery.StatusTarget_BOTH,
	})

	if err != nil {
		return err
	}

	if !statusResult.Recorded {
		return errors.New("Failed to go online")
	}

	s.currentStatus = discovery.ServiceStatus_ONLINE

	go s.heartBeat()
	return nil
}

// Offline take your service offline
func (s *FortifiService) Offline() error {
	statusResult, err := s.discoClient.Status(s.GetGrpcContext(), &discovery.StatusRequest{
		AppId:        s.appDefinition.GlobalAppID,
		InstanceUuid: s.instanceID,
		Status:       discovery.ServiceStatus_OFFLINE,
		Target:       discovery.StatusTarget_INSTANCE,
	})

	if err != nil {
		return err
	}

	if !statusResult.Recorded {
		return errors.New("Failed to go offline")
	}

	s.currentStatus = discovery.ServiceStatus_OFFLINE
	return nil
}

func (s *FortifiService) getCerts() error {

	imperiumConnection, err := grpc.Dial(s.imperiumService, grpc.WithTransportCredentials(credentials.NewClientTLSFromCert(nil, "")))
	if err != nil {
		return err
	}
	c := imperium.NewImperiumClient(imperiumConnection)
	response, err := c.Request(s.GetGrpcContext(), &imperium.CertificateRequest{
		AppId: s.appDefinition.GlobalAppID,
	})
	if err != nil {
		return err
	}

	s.hostname = response.Hostname
	s.imperiumCertificate = []byte(response.Certificate)
	s.imperiumKey = []byte(response.PrivateKey)

	s.Logger.Info("Received TLS Certificates", zap.String("hostname", s.hostname))

	return nil
}

// CreateServer creates a gRPC server with your tls certificates
func (s *FortifiService) CreateServer() (net.Listener, *grpc.Server, error) {

	lis, err := net.Listen("tcp", s.hostname+":"+strconv.FormatInt(int64(s.port), 10))
	if err != nil {
		return nil, nil, err
	}

	cert, err := tls.X509KeyPair(s.imperiumCertificate, s.imperiumKey)
	if err != nil {
		return nil, nil, err
	}

	serv := grpc.NewServer(grpc.Creds(credentials.NewServerTLSFromCert(&cert)))

	return lis, serv, nil
}

// Identity retrieves your identity
func (s *FortifiService) Identity() *identity.AppIdentity {
	return s.appIdentity
}

// Definition retrieves your definition
func (s *FortifiService) Definition() *definition.AppDefinition {
	return s.appDefinition
}

// GetAppConnection grpc.dial a service based on the discovery service
func (s *FortifiService) GetAppConnection(globalAppID string, opts ...grpc.DialOption) (*grpc.ClientConn, error) {

	locationResult, err := s.discoClient.GetLocation(s.GetGrpcContext(), &discovery.LocationRequest{AppId: globalAppID})

	if err != nil {
		return nil, err
	}

	opts = append(opts, grpc.WithTransportCredentials(credentials.NewClientTLSFromCert(nil, "")))

	return grpc.Dial(locationResult.ServiceHost+":"+strconv.FormatInt(int64(locationResult.ServicePort), 10), opts...)
}

// GetGrpcContext context to use when communicating with other services
func (s *FortifiService) GetGrpcContext() context.Context {
	md := metadata.Pairs(
		keys.GetAppIDKey(), s.appDefinition.AppID,
		keys.GetAppVendorKey(), s.appDefinition.Vendor,
	)
	return metadata.NewContext(context.Background(), md)
}
