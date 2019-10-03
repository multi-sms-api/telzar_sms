package telzarsms

import (
	"fmt"
)

// SMSError holds information about sending information
type SMSError struct {
	Status     Status
	Message    string
	ShipmentID string
}

func (e SMSError) Error() string {
	return fmt.Sprintf("%d: %s", e.Status, e.Message)
}
