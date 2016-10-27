package domfeat

import (
	"net/http"

	"net/url"

	"github.com/fortifi/proto-go/metaroute"
	"golang.org/x/net/context"
)

var CustomerContextValue = "fort-customer"

//GetCustomerInfo retrieves customer information from the inbound request
func GetCustomerInfo(r *http.Request) *metaroute.HTTPRequest_CustomerInformation {
	if r.Context().Value(CustomerContextValue) != nil {
		return r.Context().Value(CustomerContextValue).(*metaroute.HTTPRequest_CustomerInformation)
	}
	return nil
}

//HasCustomerInfo returns if the request is for a known customer
func HasCustomerInfo(r *http.Request) bool {
	return r.Context().Value(CustomerContextValue) != nil
}

//NewRequestResponse create a new HTTP Request / ResponseWriter for handling HTTP Requests
func NewRequestResponse(ctx context.Context, in *metaroute.HTTPRequest) (MetaResponseWriter, *http.Request) {
	requestUrl := &url.URL{
		Scheme:     "https",
		Host:       in.Host,
		Path:       in.Path,
		RawPath:    in.Path,
		ForceQuery: false,
		RawQuery:   in.QueryString,
	}

	request := &http.Request{
		Method: in.Method,
		URL:    requestUrl,
	}

	writer := MetaResponseWriter{
		Response: metaroute.HTTPResponse{},
		Headers:  make(http.Header),
	}

	ct := context.WithValue(ctx, CustomerContextValue, in.Customer)
	return writer, request.WithContext(ct)
}
