package telzarsms

import "encoding/xml"

// UserAuth holds authentication info for a request
type UserAuth struct {
	XMLName  xml.Name `xml:"user"`
	Username string   `xml:"username"`
	Password string   `xml:"password"`
}
