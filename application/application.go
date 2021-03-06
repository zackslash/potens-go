package application

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
	"os"
	"path"
	"path/filepath"
	"strconv"
	"time"

	"github.com/cubex/portcullis-go/keys"
	"github.com/cubex/potens-go/adl"
	"github.com/cubex/potens-go/definition"
	"github.com/cubex/potens-go/i18n"
	"github.com/cubex/potens-go/identity"
	"github.com/cubex/potens-go/websocks"
	ad "github.com/cubex/proto-go/adl"
	"github.com/cubex/proto-go/discovery"
	"github.com/cubex/proto-go/imperium"
	"github.com/cubex/proto-go/undercroft"
	"github.com/dgrijalva/jwt-go"
	"github.com/fident/proto-go/fident"
	"github.com/opentracing/opentracing-go"
	zipkin "github.com/openzipkin/zipkin-go-opentracing"
	"github.com/satori/go.uuid"
	"github.com/uber-go/zap"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/metadata"
	yaml "gopkg.in/yaml.v2"
)

var (
	parseEnv = flag.Bool("parse-env", true, "Set to false to use production defaults")
)

//CubexApplication an instance of an application service
type CubexApplication struct {
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
	appVersion          discovery.AppVersion
	currentStatus       discovery.ServiceStatus
	CubexDomain         string
	adlClient           ad.AdlClient
	socketHandler       *websocks.Handler

	//Logger used for standard logging
	Logger zap.Logger

	//Tracer open tracer
	Tracer opentracing.Tracer

	parsedEnv         bool
	imperiumService   string
	discoveryService  string
	undercroftService string
	fidentService     string
	authToken         string
	pk                *rsa.PrivateKey
	kh                string
}

// DefaultCubexDomain default domain to connect to cubex services
const DefaultCubexDomain = "services.cubex.io"

func (s *CubexApplication) parseEnv() {
	defaultPort := "50051"
	cubexDomain := os.Getenv("CUBEX_DOMAIN")
	if cubexDomain != "" {
		s.CubexDomain = cubexDomain
	}

	s.discoveryService = os.Getenv("CUBEX_DISCOVERY_LOCATION")
	if s.discoveryService == "" {
		s.discoveryService = "discovery.cubex." + s.CubexDomain
	}

	discoPort := os.Getenv("CUBEX_DISCOVERY_PORT")
	if discoPort == "" {
		s.discoveryService += ":" + defaultPort
	} else {
		s.discoveryService += ":" + discoPort
	}

	s.imperiumService = os.Getenv("CUBEX_IMPERIUM_LOCATION")
	if s.imperiumService == "" {
		s.imperiumService = "imperium." + s.CubexDomain
	}

	imperiumPort := os.Getenv("CUBEX_IMPERIUM_PORT")
	if imperiumPort == "" {
		s.imperiumService += ":" + defaultPort
	} else {
		s.imperiumService += ":" + imperiumPort
	}

	s.undercroftService = os.Getenv("CUBEX_UNDERCROFT_LOCATION")
	if s.undercroftService == "" {
		s.undercroftService = "undercroft.cubex." + s.CubexDomain
	}

	undercroftPort := os.Getenv("CUBEX_UNDERCROFT_PORT")
	if undercroftPort == "" {
		s.undercroftService += ":" + defaultPort
	} else {
		s.undercroftService += ":" + undercroftPort
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

	version := os.Getenv("SERVICE_VERSION")
	s.appVersion = discovery.AppVersion_STABLE
	if version != "" {
		v, ok := discovery.AppVersion_value[version]
		if ok {
			s.appVersion = discovery.AppVersion(v)
		}
	}

	s.parsedEnv = true
}

//SetPort sets the gRPC port
func (s *CubexApplication) SetPort(port int32) {
	s.port = port
}

//GetPort gets the gRPC port
func (s *CubexApplication) GetPort() int32 {
	return s.port
}

//SetDiscoveryClient set a shared discovery client
func (s *CubexApplication) SetDiscoveryClient(discoClient discovery.DiscoveryClient) {
	s.discoClient = discoClient
}

func (s *CubexApplication) relPath(file string) string {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		s.Logger.Fatal(err.Error())
	}
	return path.Join(dir, file)
}

//New get a new instance of a service
func New(appDef *definition.AppDefinition, appIdent *identity.AppIdentity) (CubexApplication, error) {

	s := CubexApplication{}
	s.Logger = zap.New(zap.NewJSONEncoder())
	s.CubexDomain = DefaultCubexDomain

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
			return s, err
		}
	}

	if appDef == nil {
		appDef = &definition.AppDefinition{}
		err := appDef.FromConfig(s.relPath("app-definition.yaml"))
		if err != nil {
			return s, err
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
func (s *CubexApplication) Start(collector zipkin.Collector) error {

	var err error
	var span opentracing.Span

	s.Logger.Info("Starting App", zap.String("gaid", s.appDefinition.GlobalAppID), zap.String("name", i18n.NewTranslatable(s.appDefinition.Name).Get("en")))
	s.Logger.Debug("Authing App", zap.String("identity", s.appIdentity.IdentityID), zap.String("type", s.appIdentity.IdentityType))

	flag.Parse()
	log.SetFlags(0)

	if s.port == 0 {
		minPort := 50060
		maxPort := 55555
		rand.Seed(time.Now().UTC().UnixNano())
		s.port = int32(rand.Intn(maxPort-minPort) + minPort)
	}

	if collector != nil {
		recorder := zipkin.NewRecorder(collector, false, "0.0.0.0:"+strconv.FormatInt(int64(s.port), 10), s.appDefinition.GlobalAppID)
		s.Tracer, err = zipkin.NewTracer(recorder)
		if err != nil {
			return err
		}
		span = s.Tracer.StartSpan("ServiceStart")
		defer span.Finish()
	}

	if span != nil {
		span.LogEvent("Starting")
	}

	err = s.getCerts()
	if err != nil {
		return err
	}

	if span != nil {
		span.LogEvent("GotCerts")
	}

	if s.fidentClient == nil {
		authconn, err := grpc.Dial(s.fidentService, grpc.WithTransportCredentials(credentials.NewClientTLSFromCert(nil, "")))
		if err != nil {
			s.Logger.Fatal(err.Error())
		}
		s.fidentClient = fident.NewAuthClient(authconn)
	}

	// perform auth
	ac, err := s.fidentClient.GetAuthenticationChallenge(s.GetGrpcContext(), &fident.AuthChallengePayload{Username: s.appDefinition.GlobalAppID})
	if err != nil {
		s.Logger.Fatal(err.Error())
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

	if span != nil {
		span.LogEvent("Authed")
	}

	s.authToken = authres.Token

	if s.discoClient == nil {
		discoveryConn, err := grpc.Dial(s.discoveryService, grpc.WithTransportCredentials(credentials.NewClientTLSFromCert(nil, "")))

		if err != nil {
			s.Logger.Fatal(err.Error())
		}

		s.discoClient = discovery.NewDiscoveryClient(discoveryConn)
		if span != nil {
			span.LogEvent("ConnectDiscovery")
		}
	}

	if s.undercroftClient == nil {
		regConn, err := grpc.Dial(s.undercroftService, grpc.WithTransportCredentials(credentials.NewClientTLSFromCert(nil, "")))

		if err != nil {
			s.Logger.Fatal(err.Error())
		}

		s.undercroftClient = undercroft.NewUndercroftClient(regConn)
		if span != nil {
			span.LogEvent("ConnectUndercroft")
		}
	}

	//TODO: Remove this once CLI tools are available
	appDefYaml, err := yaml.Marshal(s.appDefinition)
	if err != nil {
		s.Logger.Fatal(err.Error())
	}

	s.undercroftClient.RegisterApp(s.GetGrpcContext(), &undercroft.AppRegisterRequest{
		VendorId:       s.appDefinition.Vendor,
		Id:             s.appDefinition.AppID,
		VendorSecret:   "",
		DefinitionYaml: string(appDefYaml),
	})

	if span != nil {
		span.LogEvent("undercroft.RegisterApp")
	}

	regResult, err := s.discoClient.Register(s.GetGrpcContext(), &discovery.RegisterRequest{
		AppId:        s.appDefinition.GlobalAppID,
		InstanceUuid: s.instanceID,
		ServiceHost:  s.hostname,
		Version:      s.appVersion,
		ServicePort:  s.port,
	})

	if err != nil {
		s.Logger.Fatal(err.Error())
	}

	if span != nil {
		span.LogEvent("discovery.Register")
	}

	if !regResult.Recorded {
		s.Logger.Fatal("Failed to register with the discovery service")
	} else {
		s.Logger.Debug("Registered with Discovery Service")
	}
	return nil
}

func (s *CubexApplication) heartBeat() {
	if s.currentStatus == discovery.ServiceStatus_ONLINE {
		for {
			s.discoClient.HeartBeat(s.GetGrpcContext(), &discovery.HeartBeatRequest{
				AppId:        s.appDefinition.GlobalAppID,
				InstanceUuid: s.instanceID,
				Version:      s.appVersion,
			})
			time.Sleep(10 * time.Second)
		}
	}
}

// Online take your service online
func (s *CubexApplication) Online() error {
	statusResult, err := s.discoClient.Status(s.GetGrpcContext(), &discovery.StatusRequest{
		AppId:        s.appDefinition.GlobalAppID,
		InstanceUuid: s.instanceID,
		Version:      s.appVersion,
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

// Close take your service offline, and if running locally, also shutdown
func (s *CubexApplication) Close() error {
	err := s.Offline()

	if s.CubexDomain != DefaultCubexDomain {
		err = s.Shutdown()

		s.undercroftClient.DeRegisterApp(s.GetGrpcContext(), &undercroft.AppRequest{
			VendorId: s.appDefinition.Vendor,
			Id:       s.appDefinition.AppID,
		})
	}

	return err
}

// Offline take your service offline
func (s *CubexApplication) Offline() error {
	statusResult, err := s.discoClient.Status(s.GetGrpcContext(), &discovery.StatusRequest{
		AppId:        s.appDefinition.GlobalAppID,
		InstanceUuid: s.instanceID,
		Version:      s.appVersion,
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

// Shutdown unregisters your service from discovery
func (s *CubexApplication) Shutdown() error {
	s.Logger.Info("Shutting Down App", zap.String("gaid", s.appDefinition.GlobalAppID), zap.String("name", i18n.NewTranslatable(s.appDefinition.Name).Get("en")))

	deregResult, err := s.discoClient.DeRegister(s.GetGrpcContext(), &discovery.DeRegisterRequest{
		AppId:        s.appDefinition.GlobalAppID,
		InstanceUuid: s.instanceID,
		Version:      s.appVersion,
	})

	if err != nil {
		s.Logger.Fatal(err.Error())
	}

	if !deregResult.Recorded {
		s.Logger.Fatal("Unable to deregister service")
	}

	return nil
}

func (s *CubexApplication) getCerts() error {

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

	s.Logger.Debug("Received TLS Certificates", zap.String("hostname", s.hostname))

	return nil
}

// CreateServer creates a gRPC server with your tls certificates
func (s *CubexApplication) CreateServer() (net.Listener, *grpc.Server, error) {

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
func (s *CubexApplication) Identity() *identity.AppIdentity {
	return s.appIdentity
}

// Definition retrieves your definition
func (s *CubexApplication) Definition() *definition.AppDefinition {
	return s.appDefinition
}

// ADL retrives ADL instance
func (s *CubexApplication) ADL(fid string) *adl.Entity {
	if s.adlClient == nil {
		con, err := s.GetAppConnection(adl.ADLGAID)
		if err != nil {
			s.Logger.Fatal("Unable to connect to ADL", zap.String("error", err.Error()))
		}
		s.adlClient = ad.NewAdlClient(con)
		ctx := s.GetGrpcContext()
		appID := s.Identity().AppID
		adl.SetContextAppID(ctx, appID)
	}
	return adl.Mutate(fid, &s.adlClient)
}

// Socket helper for sockets
func (s *CubexApplication) Socket() *websocks.Handler {
	if s.socketHandler == nil {
		con, err := s.GetAppConnection(websocks.SOCKETSGAID)
		if err != nil {
			s.Logger.Fatal("Unable to connect to Sockets Server", zap.String("error", err.Error()))
		}
		s.socketHandler = websocks.NewHandler(con, s.GetGrpcContext(), s.appDefinition.Vendor, s.appDefinition.AppID)
	}
	return s.socketHandler
}

// GetAppConnection grpc.dial a service based on the discovery service
func (s *CubexApplication) GetAppConnection(globalAppID string, opts ...grpc.DialOption) (*grpc.ClientConn, error) {

	locationResult, err := s.discoClient.GetLocation(s.GetGrpcContext(), &discovery.LocationRequest{AppId: globalAppID})

	if err != nil {
		return nil, err
	}

	opts = append(opts, grpc.WithTransportCredentials(credentials.NewClientTLSFromCert(nil, "")))

	return grpc.Dial(locationResult.ServiceHost+":"+strconv.FormatInt(int64(locationResult.ServicePort), 10), opts...)
}

// GetGrpcContext context to use when communicating with other services
func (s *CubexApplication) GetGrpcContext() context.Context {
	md := metadata.Pairs(
		keys.GetAppIDKey(), s.appDefinition.AppID,
		keys.GetAppVendorKey(), s.appDefinition.Vendor,
	)
	return metadata.NewContext(context.Background(), md)
}
