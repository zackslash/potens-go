package definition

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

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

/**
 * Application Type
 */

type AppType string

const (
	AppTypeEmployee      AppType = "employee"
	AppTypePublisher     AppType = "publisher"
	AppTypeCustomer      AppType = "customer"
	AppTypeDomainFeature AppType = "domain.feature"
)

type AppNavigation struct {
	ID          string
	Name        string
	Description string
	Icon        string
	Path        string
	Roles       []string
}

type AppEntities struct {
	AppKey    string            `yaml:"app_key"`
	VendorKey string            `yaml:"vendor_key"`
	Entities  map[string]Entity `yaml:",inline"`
}

type Entity struct {
	Name        string
	Plural      string
	Description string
	Path        string
	Hovercard   string
}

/**
 * Listener Repository Type
 */

type ListenerRepositoryType string

const (
	ListenerRepositoryTypeHTTP   ListenerRepositoryType = "http"
	ListenerRepositoryTypePubSub ListenerRepositoryType = "pubsub"
	ListenerRepositoryTypeSQS    ListenerRepositoryType = "sqs"
)

type AppListener struct {
	Enabled    bool
	Repository ListenerRepositoryType
	Config     []AppListenerConfig
}

type AppListenerConfig struct {
	Name  string
	Value string
}

/**
 * Quick Action Mode
 */

type QuickActionMode string

const (
	QuickActionModePage   QuickActionMode = "page"
	QuickActionModeDialog QuickActionMode = "dialog"
	QuickActionModeWindow QuickActionMode = "window"
)

type AppQuickAction struct {
	ID    string
	Name  string
	Icon  string
	mode  QuickActionMode
	Path  string
	Roles []string
}

type AppSearchAction struct {
	ID          string
	Name        string
	Description string
	Icon        string
	Path        string
	Roles       []string
	Tokens      []string
}

type AppQueue struct {
	ID    string
	Name  string
	Icon  string
	Path  string
	Roles []string
}

type AppNotification struct {
	ID          string
	Name        string
	Description string
	Icon        string
	Roles       []string
	Attributes  []AppNotificationAttribute
}

/**
 * App Notification Attribute Type
 */

type AppNotificationAttributeType string

const (
	AppNotificationAttributeTypeString  AppNotificationAttributeType = "string"
	AppNotificationAttributeTypeInteger AppNotificationAttributeType = "integer"
	AppNotificationAttributeTypeFloat   AppNotificationAttributeType = "float"
	AppNotificationAttributeTypeBoolean AppNotificationAttributeType = "boolean"
)

type AppNotificationAttribute struct {
	Name string
	Type AppNotificationAttributeType
}

type AppRole struct {
	ID          string
	Name        string
	Description string
}

/**
 * App Config Type
 */

type AppConfigType string

const (
	AppConfigTypeString     AppConfigType = "string"
	AppConfigTypeInteger    AppConfigType = "integer"
	AppConfigTypeFloat      AppConfigType = "float"
	AppConfigTypeBoolean    AppConfigType = "boolean"
	AppConfigTypeJson       AppConfigType = "json"
	AppConfigTypeUri        AppConfigType = "uri"
	AppConfigTypeOptions    AppConfigType = "options"
	AppConfigTypeArrayComma AppConfigType = "array:comma"
	AppConfigTypeArrayLine  AppConfigType = "array:line"
)

type AppConfig struct {
	ID          string
	Name        string
	Description string
	Help        string
	Type        AppConfigType
	Values      []map[string]string
}

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
