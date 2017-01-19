package domfeat

import (
	"net/http"

	"github.com/cubex/proto-go/metaroute"
)

type MetaResponseWriter struct {
	Response metaroute.HTTPResponse
	Headers  http.Header
}

func (w *MetaResponseWriter) Header() http.Header {
	return w.Headers
}

func (w *MetaResponseWriter) WriteHeader(code int) {
	w.Response.StatusCode = int32(code)
}

func (w *MetaResponseWriter) Write(data []byte) (int, error) {
	w.Response.Body += string(data)
	return len(data), nil
}
