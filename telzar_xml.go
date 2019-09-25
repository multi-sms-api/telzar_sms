package telzarsms

import "encoding/xml"

// TelzarXML Holds a single SMS message to send
type TelzarXML struct {
	XMLName        xml.Name       `xml:"sms"`
	User           UserAuth       `xml:"user"`
	Source         string         `xml:"source"`
	Destination    DestinationXML `xml:"destination"`
	Tag            string         `xml:"tag,omitempty"`
	Message        string         `xml:"message"`
	AddDynamic     Bool           `xml:"add_dynamic,omitempty"`
	Timing         string         `xml:"timing,omitempty"`
	AddUnsubscribe Bool           `xml:"add_unsubscribe"`
	Response       Bool           `xml:"response,omitempty"`
}
