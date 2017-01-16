package webui

import (
	"net/url"

	"github.com/fortifi/potens-go/webui/breadcrumb"
	"github.com/fortifi/proto-go/platform"
)

//CreateResponse creates a new initialised response
func CreateResponse() *platform.HTTPResponse {
	response := &platform.HTTPResponse{}
	response.Headers = make(map[string]*platform.HTTPResponse_HTTPHeaderParameter)
	return response
}

//SetBreadcrumb set the breadcrumb on the response
func SetBreadcrumb(response *platform.HTTPResponse, breadcrumb breadcrumb.Breadcrumb) {
	response.Headers["x-fort-breadcrumb"] = &platform.HTTPResponse_HTTPHeaderParameter{Values: []string{breadcrumb.Json()}}
}

//SetPageTitle set the page title on the response
func SetPageTitle(response *platform.HTTPResponse, PageTitle string) {
	response.Headers["x-fort-title"] = &platform.HTTPResponse_HTTPHeaderParameter{Values: []string{PageTitle}}
}

//SetBackPath set the back path on the response, relative to app route
func SetBackPath(response *platform.HTTPResponse, BackPath string) {
	response.Headers["x-fort-back-path"] = &platform.HTTPResponse_HTTPHeaderParameter{Values: []string{BackPath}}
}

//SetPageIcon set the icon url/code on the response
func SetPageIcon(response *platform.HTTPResponse, Icon string) {
	response.Headers["x-fort-icon"] = &platform.HTTPResponse_HTTPHeaderParameter{Values: []string{Icon}}
}

// PageIntergrationType
type PageIntergrationType string

//Page Intergration Types
const (
	// PageIntergrationTypeDefault Default
	PageIntergrationTypeDefault PageIntergrationType = "default"
	// PageIntergrationTypeNone None
	PageIntergrationTypeNone PageIntergrationType = "none"
)

//SetPageIcon set the icon url/code on the response
func SetPageIntegrations(response *platform.HTTPResponse, IntegrationType PageIntergrationType) {
	response.Headers["x-fort-integrations"] = &platform.HTTPResponse_HTTPHeaderParameter{Values: []string{string(IntegrationType)}}
}

func GetUrl(request *platform.HTTPRequest) *url.URL {
	return &url.URL{
		Scheme:     "https",
		Host:       "apps.fortifi.io",
		Path:       request.Path,
		RawPath:    request.Path,
		ForceQuery: false,
		RawQuery:   request.QueryString,
	}
}
