package telzarsms

import (
	"net/http"
	"os"
	"testing"
	"time"
)

var (
	userName, password   string
	fromNumber, toNumber string
	client               http.Client
	handler              HTTPHandler
)

func init() {
	userName = os.Getenv("TELZAR_USERNAME")
	password = os.Getenv("TELZAR_PASSWORD")
	fromNumber = os.Getenv("TELZAR_FROM_NUMBER")
	toNumber = os.Getenv("TELZAR_TO_NUMBER")
	client = http.Client{
		Timeout: time.Second * 15,
	}
	handler = HTTPHandler{
		Client: &client,
	}
}
func TestSendSMSHTTP(t *testing.T) {
	if testing.Short() {
		return
	}
	toPhones := []Phone{
		Phone{
			Phone: toNumber,
		},
	}
	xmlToSend := TelzarXML{
		User: UserAuth{
			Username: userName,
			Password: password,
		},
		Source: fromNumber,
		Destinations: DestinationXML{
			Phone: toPhones,
		},
		Message: "Test",
	}
	_, err := xmlToSend.SendSMS(handler)
	if err != nil {
		t.Errorf("Error sending SMS: %s", err)
	}
}
