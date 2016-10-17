package webui

import (
	"github.com/fortifi/potens-go/webui/breadcrumb"
	"github.com/fortifi/proto-go/platform"
)

//CreateResponse creates a new initialised response
func CreateResponse(PageTitle string) *platform.HTTPResponse {
	response := &platform.HTTPResponse{}
	response.Headers = make(map[string]*platform.HTTPResponse_HTTPHeaderParameter)
	response.Headers["x-fort-title"] = &platform.HTTPResponse_HTTPHeaderParameter{Values: []string{PageTitle}}
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
