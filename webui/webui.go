package webui

import (
	"net/url"

	"github.com/cubex/potens-go/webui/breadcrumb"
	"github.com/cubex/proto-go/applications"
)

//CreateResponse creates a new initialised response
func CreateResponse() *applications.HTTPResponse {
	response := &applications.HTTPResponse{}
	response.Headers = make(map[string]*applications.HTTPResponse_HTTPHeaderParameter)
	return response
}

//SetBreadcrumb set the breadcrumb on the response
func SetBreadcrumb(response *applications.HTTPResponse, breadcrumb breadcrumb.Breadcrumb) {
	response.Headers["x-cube-breadcrumb"] = &applications.HTTPResponse_HTTPHeaderParameter{Values: []string{breadcrumb.Json()}}
}

//SetPageTitle set the page title on the response
func SetPageTitle(response *applications.HTTPResponse, PageTitle string) {
	response.Headers["x-cube-title"] = &applications.HTTPResponse_HTTPHeaderParameter{Values: []string{PageTitle}}
}

//SetBackPath set the back path on the response, relative to app route
func SetBackPath(response *applications.HTTPResponse, BackPath string) {
	response.Headers["x-cube-back-path"] = &applications.HTTPResponse_HTTPHeaderParameter{Values: []string{BackPath}}
}

//SetPageIcon set the icon url/code on the response
func SetPageIcon(response *applications.HTTPResponse, Icon string) {
	response.Headers["x-cube-icon"] = &applications.HTTPResponse_HTTPHeaderParameter{Values: []string{Icon}}
}

//SetPageFID set the FID for the entity being shown on the page
func SetPageFID(response *applications.HTTPResponse, FID string) {
	response.Headers["x-cube-page-fid"] = &applications.HTTPResponse_HTTPHeaderParameter{Values: []string{FID}}
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
func SetPageIntegrations(response *applications.HTTPResponse, IntegrationType PageIntergrationType) {
	response.Headers["x-cube-integrations"] = &applications.HTTPResponse_HTTPHeaderParameter{Values: []string{string(IntegrationType)}}
}

func GetUrl(request *applications.HTTPRequest) *url.URL {
	return &url.URL{
		Scheme:     "https",
		Host:       "apps.cubex.io",
		Path:       request.Path,
		RawPath:    request.Path,
		ForceQuery: false,
		RawQuery:   request.QueryString,
	}
}
