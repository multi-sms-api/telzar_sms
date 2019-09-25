package telzarsms

import "encoding/xml"

// SMSBulkXML holds SMS message for bulk messages
type SMSBulkXML struct {
	XMLName     xml.Name       `xml:"sms"`
	Source      string         `xml:"source"`
	Destination DestinationXML `xml:"destination"`
	Message     string         `xml:"message"`
}

// MessagesBulkXML holds a list of SMS messages for bulk sending
type MessagesBulkXML struct {
	XMLName  xml.Name     `xml:"messages"`
	SMS      []SMSBulkXML `xml:"sms"`
	Timing   string       `xml:"timing,omitempty"`
	Response Bool         `xml:"response"`
}

// TelzarBulkXML holds full bulk XML
type TelzarBulkXML struct {
	XMLName  xml.Name `xml:"bulk"`
	User     UserAuth
	Messages MessagesBulkXML
}
