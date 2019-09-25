package telzarsms

import "encoding/xml"

// Phone represent the phone number as expected
type Phone struct {
	Phone string `xml:",chardata"`
	ID    string `xml:"id,attr"`
}

// DestinationXML represent the XML structure of destination
type DestinationXML struct {
	XMLName xml.Name `xml:"destination"`
	ClID    []string `xml:"cl_id,omitempty"`
	Phone   []Phone  `xml:"phone"`
}
