package telzarsms

import (
	"net/url"
	"time"
)

// DLRPush holds information arrived from DLR sent as Push notification
type DLRPush struct {
	ExternalID string
	Status     Status
	HeMessage  string
	EnMessage  string
	Date       time.Time
	Phone      string
	Operator   string
	ShipmentID string
}

// URLToDLRPush translate url.Values to DLRPush struct
func URLToDLRPush(values url.Values) *DLRPush {
	date, err := time.Parse(DateFormat, values.Get("date"))
	if err != nil {
		date = time.Now()
	}
	result := DLRPush{
		ExternalID: values.Get("external_id"),
		Status:     StrToStatus(values.Get("status")),
		HeMessage:  values.Get("he_message"),
		EnMessage:  values.Get("en_message"),
		Date:       date,
		Phone:      values.Get("phone"),
		Operator:   values.Get("operator"),
		ShipmentID: values.Get("shipment_id"),
	}

	return &result
}
