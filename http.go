package telzarsms

import (
	"bytes"
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
	"strings"
)

// HTTPHandler perform HTTP actions, and implement
type HTTPHandler struct {
	Client *http.Client
}

// DoHTTP sends an HTTP Request for sending an SMS
func (h HTTPHandler) DoHTTP(
	method, contentType, address string, fields url.Values, body []byte) (resp *http.Response, err error) {

	var request *http.Request
	var bodyReader *bytes.Reader

	fullAddress := fmt.Sprintf("%s", address)

	if body != nil {
		bodyReader = bytes.NewReader(body)
	}

	switch method {
	case http.MethodGet:
		fullAddress = fmt.Sprintf("%s?%s", fullAddress, fields.Encode())
		request, err = http.NewRequest(http.MethodGet, fullAddress, bodyReader)
	case http.MethodPost:
		request, err = http.NewRequest(http.MethodPost, fullAddress, bodyReader)
	}

	request.Header.Set("Content-Type", XMLMimeType)

	if err != nil {
		return nil, err
	}

	if contentType != "" {
		request.Header.Set("Content-Type", contentType)
	}
	request.Close = true

	ctx, cancel := context.WithTimeout(request.Context(), h.Client.Timeout)
	defer cancel()
	defer h.Client.CloseIdleConnections()

	resp, err = h.Client.Do(request.WithContext(ctx))

	if err != nil {
		if strings.Contains(os.Getenv("SMSHTTPDEBUG"), "dump=true") {
			fmt.Printf("Error was given: %s", err)
		}
		return
	}

	if resp == nil {
		err = fmt.Errorf("resp is nil")
		if strings.Contains(os.Getenv("SMSHTTPDEBUG"), "dump=true") {
			fmt.Println("resp is nil")
		}
		return
	}

	if strings.Contains(os.Getenv("SMSHTTPDEBUG"), "dump=true") {
		dump, err := httputil.DumpRequestOut(request, true)
		fmt.Printf(">>>> dump request: %s \nerr: %s\n", dump, err)

		dump, err = httputil.DumpResponse(resp, true)
		fmt.Printf(">>>> dump response: %s \nerr: %s\n", dump, err)
	}

	if resp.Body == nil {
		err = fmt.Errorf("resp.body is nil")
		return
	}

	var respBody []byte
	respBody, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}

	if resp.StatusCode == http.StatusOK {
		var status XMLResponse
		status, err = FromXMLResponse(respBody)
		if err != nil {
			return
		}
		if status.Status != ErrorOK {
			err = ToError(status)
		}
	}

	return
}