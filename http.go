package telzarsms

import (
	"net/http"
	"net/url"
)

// HTTPHandler perform HTTP actions, and implement
type HTTPHandler struct {
	Client *http.Client
}

// DoHTTP sends an HTTP Request for sending an SMS
func (h HTTPHandler) DoHTTP(
	method, contentType, address string, fields url.Values, body []byte) (resp *http.Response, err error) {
	return smshandler.DoHTTP(h.Client, method, contentType, address, fields, body)
}

// OnGettingSMS is an HTTP server handler when incoming SMS arrives.
// If mux exists, it will use it for a server, otherwise it will
// use http.HandleFunc.
func (h HTTPHandler) OnGettingSMS(path string, mux *http.ServeMux, httpHandler http.HandlerFunc) {
	if mux != nil {
		mux.HandleFunc(path, httpHandler)
		return
	}

	http.HandleFunc(path, httpHandler)
}
