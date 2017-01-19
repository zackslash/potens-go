package domfeat

import (
	"net/http"

	"net/url"

	"github.com/cubex/proto-go/metaroute"
	"golang.org/x/net/context"
)

var UserContextValue = "cubex-user"

//GetUserInfo retrieves customer information from the inbound request
func GetUserInfo(r *http.Request) *metaroute.HTTPRequest_UserInformation {
	if r.Context().Value(UserContextValue) != nil {
		return r.Context().Value(UserContextValue).(*metaroute.HTTPRequest_UserInformation)
	}
	return nil
}

//HasUserInfo returns if the request is for a known customer
func HasUserInfo(r *http.Request) bool {
	return r.Context().Value(UserContextValue) != nil
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

	ct := context.WithValue(ctx, UserContextValue, in.User)
	return writer, request.WithContext(ct)
}
