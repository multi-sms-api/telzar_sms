package telzarsms

import "encoding/xml"

// XMLResponse holds information about the given response back from the server
type XMLResponse struct {
	XMLName        xml.Name    `xml:"sms"`
	Status         Status      `xml:"status"`
	Message        string      `xml:"message"`
	ShipmentID     string      `xml:"shipment_id"`
	AddUnsubscribe Unsubscribe `xml:"add_unsubscribe,omitempty"`
}

// FromXMLResponse turns the body to XMLResponse, or error if something bad
// happened
func FromXMLResponse(buf []byte) (XMLResponse, error) {
	var response XMLResponse

	err := xml.Unmarshal(buf, &response)
	return response, err
}
