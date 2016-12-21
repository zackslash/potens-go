package definition

import (
	"io/ioutil"
	"strings"

	"github.com/fortifi/potens-go/i18n"

	yaml "gopkg.in/yaml.v2"
)

// AppDefinition Application Definition
type AppDefinition struct {
	Type                      string
	ConfigVersion             float32 `yaml:"config_version"`
	Version                   float32
	Vendor                    string
	TrustedVendor             bool
	AppID                     string `yaml:"app_id"`
	GlobalAppID               string
	Group                     string
	Category                  string
	Priority                  int32
	AppType                   AppType `yaml:"app_type"`
	Name                      i18n.Translations
	Description               i18n.Translations
	Icon                      string
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
	Permissions               []AppPermission `yaml:"permissions"`
	Config                    []AppConfig
	Integrations              AppIntegrations
}

// AppType Application Type
type AppType string

//App Types
const (
	// AppTypeEmployee Employee
	AppTypeEmployee AppType = "employee"
	// AppTypeUser User
	AppTypeUser AppType = "user"
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
	Name        i18n.Translations
	Description i18n.Translations
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
	Name        i18n.Translations
	Plural      i18n.Translations
	Description i18n.Translations
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
	Name  i18n.Translations
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
	Name  i18n.Translations
	Icon  string
	mode  QuickActionMode
	Path  string
	Roles []string
}

// AppSearchAction Search Action provided by your app
type AppSearchAction struct {
	ID          string
	Name        i18n.Translations
	Description i18n.Translations
	Icon        string
	Path        string
	Roles       []string
	Tokens      []string
}

// AppQueue Queue provided by your app
type AppQueue struct {
	ID    string
	Name  i18n.Translations
	Icon  string
	Path  string
	Roles []string
}

// AppNotification notification provided by your app
type AppNotification struct {
	ID          string
	Name        i18n.Translations
	Description i18n.Translations
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
	Name        i18n.Translations
	Description i18n.Translations
}

// VendorID Retrieves the vendor ID for this role, empty for a global role
func (role *AppRole) VendorID(appDef *AppDefinition) string {
	if role.IsBuiltIn() {
		return ""
	}
	roleSplit := strings.SplitN(role.ID, "/", 3)
	if len(roleSplit) == 3 && len(roleSplit[0]) > 0 {
		return roleSplit[0]
	}

	return appDef.Vendor
}

// AppID Retrieves the application ID for this role, empty for a global role
func (role *AppRole) AppID(appDef *AppDefinition) string {
	if role.IsBuiltIn() {
		return ""
	}
	roleSplit := strings.SplitN(role.ID, "/", 3)
	if len(roleSplit) == 3 && len(roleSplit[1]) > 0 {
		return roleSplit[1]
	}

	return appDef.AppID
}

// IsBuiltIn returns true for global roles, e.g. owner
func (role *AppRole) IsBuiltIn() bool {
	return !strings.Contains(role.ID, "/")
}

// IsSameVendor returns true if the vendor for the role matches the vendor in the provided definition
func (role *AppRole) IsSameVendor(appDef *AppDefinition) bool {
	return role.VendorID(appDef) == appDef.Vendor
}

// AppPermissionMode Permission Request
type AppPermissionMode string

const (
	// AppConfigModeString Require Permission
	AppPermissionModeRequired AppPermissionMode = "require"
	// AppPermissionModeOptional Optionally Require
	AppPermissionModeOptional AppPermissionMode = "optional"
)

type AppPermission struct {
	GlobalAppID string `yaml:"gaid"`
	RPC         string
	Mode        AppPermissionMode
	Reason      i18n.Translations
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
	Name        i18n.Translations
	Description i18n.Translations
	Note        i18n.Translations
	Help        i18n.Translations
	Type        AppConfigType
	Values      []map[string]i18n.Translations
}

type AppIntegrationPanel struct {
	ID   string
	Hook string
	Path string
}

// AppIntegrationMenuItemMode Launch mode for a integration menu item
type AppIntegrationMenuItemMode string

const (
	// AppIntegrationMenuItemModeFull Redirect to a new page
	AppIntegrationMenuItemModeFull AppIntegrationMenuItemMode = "full"
	// AppIntegrationMenuItemModeIntegrated Open within the content area of the entity page
	AppIntegrationMenuItemModeIntegrated AppIntegrationMenuItemMode = "integrated"
)

type AppIntegrationMenuItem struct {
	ID          string
	Hook        string
	Path        string
	Mode        AppIntegrationMenuItemMode
	Title       i18n.Translations
	Description i18n.Translations
}

// AppIntegrationActionMode Launch mode for a integration action
type AppIntegrationActionMode string

const (
	// AppIntegrationActionModePage Redirect to a new page
	AppIntegrationActionModePage AppIntegrationActionMode = "page"
	// AppIntegrationActionModeDialog Open in a dialog
	AppIntegrationActionModeDialog AppIntegrationActionMode = "dialog"
	// AppIntegrationActionModeWindow Open in a new window
	AppIntegrationActionModeWindow AppIntegrationActionMode = "window"
	// AppIntegrationActionModeIntegrated Open within the content area of the entity page
	AppIntegrationActionModeIntegrated AppIntegrationActionMode = "integrated"
)

type AppIntegrationAction struct {
	ID          string
	Hook        string
	Path        string
	Icon        string
	Mode        AppIntegrationActionMode
	Title       i18n.Translations
	Description i18n.Translations
}

type AppIntegrations struct {
	Panels          []AppIntegrationPanel
	HeaderMenuItems []AppIntegrationMenuItem `yaml:"header_menu_items"`
	PageMenuItems   []AppIntegrationMenuItem `yaml:"page_menu_items"`
	HeaderActions   []AppIntegrationAction   `yaml:"header_actions"`
	PageActions     []AppIntegrationAction   `yaml:"page_actions"`
}

// FromConfig Populates your definition based on your app-definition.yaml
func (d *AppDefinition) FromConfig(yamlFile string) error {
	yamlContent, err := ioutil.ReadFile(yamlFile)
	if err != nil {
		return err
	}

	return d.FromYamlString(string(yamlContent))
}

func (d *AppDefinition) FromYamlString(yamlContent string) error {
	err := yaml.Unmarshal([]byte(yamlContent), d)
	if err == nil {
		d.GlobalAppID = d.Vendor + "/" + d.AppID
	}
	return err
}
