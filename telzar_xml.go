package telzarsms

import (
	"encoding/xml"
	"fmt"
	"net/http"
	"os"
	"strings"
)

// TelzarXML Holds a single SMS message to send
type TelzarXML struct {
	XMLName        xml.Name       `xml:"sms"`
	User           UserAuth       `xml:"user"`
	Source         string         `xml:"source"`
	Destinations   DestinationXML `xml:"destinations"`
	Tag            string         `xml:"tag,omitempty"`
	Message        string         `xml:"message"`
	AddDynamic     Bool           `xml:"add_dynamic,omitempty"`
	Timing         string         `xml:"timing,omitempty"`
	AddUnsubscribe Unsubscribe    `xml:"add_unsubscribe"`
	Response       Bool           `xml:"response,omitempty"`
}

// SendSMS sends the given SMS to InforU based on HTTP client
func (x TelzarXML) SendSMS(h HTTPHandler) (*http.Response, error) {
	buf, err := xml.Marshal(x)
	if err != nil {
		return nil, err
	}
	if strings.Contains(os.Getenv("SMSHTTPDEBUG"), "dump=true") {
		fmt.Printf(">>>> dump TelzarXML: %s\n", buf)
	}
	address := ProdAPIAddress
	if strings.Contains(os.Getenv("SMSHTTPDEBUG"), "test=on") {
		address = DevAPIAddress
	}
	resp, err := h.DoHTTP(HTTPMethod, HTTPContentType, address, nil, buf)
	return resp, err
}
