package definition

import (
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

// AppDefinition Application Definition
type AppDefinition struct {
	Type                      string
	ConfigVersion             float32 `yaml:"config_version"`
	Version                   float32
	Vendor                    string
	AppID                     string `yaml:"app_id"`
	GlobalAppID               string
	Group                     string
	Category                  string
	Priority                  int32
	AppType                   AppType `yaml:"app_type"`
	Name                      string
	Description               string
	icon                      string
	AdvancedNotificationsPath string `yaml:"advanced_notifications_path"`
	AdvancedConfigPath        string `yaml:"advanced_config_path"`
	Navigation                []AppNavigation
	Entities                  AppEntities
	Listener                  AppListener
	QuickActions              []AppQuickAction  `yaml:"quick_actions"`
	SearchActions             []AppSearchAction `yaml:"search_actions"`
	Queues                    []AppQueue
	Notifications             []AppNotification
	Roles                     []AppRole
	Config                    []AppConfig
}

// AppType Application Type
type AppType string

//App Types
const (
	// AppTypeEmployee Employee
	AppTypeEmployee AppType = "employee"
	// AppTypePublisher Publisher
	AppTypePublisher AppType = "publisher"
	// AppTypeCustomer Customer
	AppTypeCustomer AppType = "customer"
	// AppTypeDomainFeature Domain Feature
	AppTypeDomainFeature AppType = "domain.feature"
)

// AppNavigation Application Navigation ITem
type AppNavigation struct {
	ID          string
	Name        string
	Description string
	Icon        string
	Path        string
	Roles       []string
}

// AppEntities Entity values for building FIDs
type AppEntities struct {
	AppKey    string            `yaml:"app_key"`
	VendorKey string            `yaml:"vendor_key"`
	Entities  map[string]Entity `yaml:",inline"`
}

// Entity Definition of a single FDL data type
type Entity struct {
	Name        string
	Plural      string
	Description string
	Path        string
	Hovercard   string
}

// ListenerRepositoryType Service to listen to events on
type ListenerRepositoryType string

// Types of listeners
const (
	// Basic HTTP requests
	ListenerRepositoryTypeHTTP ListenerRepositoryType = "http"
	// Google Pub Sub
	ListenerRepositoryTypePubSub ListenerRepositoryType = "pubsub"
	// Amazon SQS
	ListenerRepositoryTypeSQS ListenerRepositoryType = "sqs"
)

// AppListener Definition of your app listener
type AppListener struct {
	Enabled    bool
	Repository ListenerRepositoryType
	Config     []AppListenerConfig
}

// AppListenerConfig Config items for listener
type AppListenerConfig struct {
	Name  string
	Value string
}

// QuickActionMode Launch mode for a quick action
type QuickActionMode string

const (
	// QuickActionModePage Redirect to a new page
	QuickActionModePage QuickActionMode = "page"
	// QuickActionModeDialog Open a dialog window
	QuickActionModeDialog QuickActionMode = "dialog"
	// QuickActionModeWindow Open in a new window
	QuickActionModeWindow QuickActionMode = "window"
)

// AppQuickAction Quick Action provided by your app
type AppQuickAction struct {
	ID    string
	Name  string
	Icon  string
	mode  QuickActionMode
	Path  string
	Roles []string
}

// AppSearchAction Search Action provided by your app
type AppSearchAction struct {
	ID          string
	Name        string
	Description string
	Icon        string
	Path        string
	Roles       []string
	Tokens      []string
}

// AppQueue Queue provided by your app
type AppQueue struct {
	ID    string
	Name  string
	Icon  string
	Path  string
	Roles []string
}

// AppNotification notification provided by your app
type AppNotification struct {
	ID          string
	Name        string
	Description string
	Icon        string
	Roles       []string
	Attributes  []AppNotificationAttribute
}

// AppNotificationAttributeType Type of notification attribute
type AppNotificationAttributeType string

const (
	// AppNotificationAttributeTypeString String Type
	AppNotificationAttributeTypeString AppNotificationAttributeType = "string"
	// AppNotificationAttributeTypeInteger Integer Type
	AppNotificationAttributeTypeInteger AppNotificationAttributeType = "integer"
	// AppNotificationAttributeTypeFloat Float Type
	AppNotificationAttributeTypeFloat AppNotificationAttributeType = "float"
	// AppNotificationAttributeTypeBoolean Boolean Type
	AppNotificationAttributeTypeBoolean AppNotificationAttributeType = "boolean"
)

// AppNotificationAttribute Attribute on your notification
type AppNotificationAttribute struct {
	Name string
	Type AppNotificationAttributeType
}

// AppRole Roles provided by your application
type AppRole struct {
	ID          string
	Name        string
	Description string
}

// AppConfigType - Type of config value
type AppConfigType string

const (
	// AppConfigTypeString String
	AppConfigTypeString AppConfigType = "string"
	// AppConfigTypeInteger Integer
	AppConfigTypeInteger AppConfigType = "integer"
	// AppConfigTypeFloat Float
	AppConfigTypeFloat AppConfigType = "float"
	// AppConfigTypeBoolean Boolean
	AppConfigTypeBoolean AppConfigType = "boolean"
	// AppConfigTypeJSON Json
	AppConfigTypeJSON AppConfigType = "json"
	// AppConfigTypeURI Uri
	AppConfigTypeURI AppConfigType = "uri"
	// AppConfigTypeOptions Options
	AppConfigTypeOptions AppConfigType = "options"
	// AppConfigTypeArrayComma ArrayComma
	AppConfigTypeArrayComma AppConfigType = "array:comma"
	// AppConfigTypeArrayLine ArrayLine
	AppConfigTypeArrayLine AppConfigType = "array:line"
)

// AppConfig Configurable item for your app per organisation
type AppConfig struct {
	ID          string
	Name        string
	Description string
	Help        string
	Type        AppConfigType
	Values      []map[string]string
}

// FromConfig Populates your definition based on your app-definition.yaml
func (d *AppDefinition) FromConfig(yamlFile string) error {
	yamlContent, err := ioutil.ReadFile(yamlFile)
	if err != nil {
		return err
	}

	err = yaml.Unmarshal(yamlContent, d)
	if err == nil {
		d.GlobalAppID = d.Vendor + "/" + d.AppID
	}
	return err
}
