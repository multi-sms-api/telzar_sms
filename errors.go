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

// ToError converts XMLResponse to SMSError. If everything is ok, it will return
// nil
func ToError(returnStatus XMLResponse) *SMSError {
	if returnStatus.Status == ErrorOK {
		return nil
	}
	result := SMSError{
		Status:     returnStatus.Status,
		Message:    returnStatus.Message,
		ShipmentID: returnStatus.ShipmentID,
	}

	return &result
}
