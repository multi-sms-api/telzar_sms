package telzarsms

import (
	"database/sql/driver"
	"encoding/xml"
	"errors"
	"fmt"
	"reflect"
	"strconv"
)

// Bool represents the boolean value for XML
type Bool bool

// Unsubscribe holds the type of unsubscribing methods to use
type Unsubscribe uint8

// Status is the response status arrive back
type Status int

// DLRStatus holds the Delivery report (DLR) information
type DLRStatus int

func (b Bool) String() string {
	if b {
		return "true"
	}
	return "false"
}

func (u Unsubscribe) String() string {
	cause, ok := unsubscribeType[u]
	if !ok {
		return unsubscribeType[UnsubscribeOff]
	}
	return cause
}

func (s Status) String() string {
	status, ok := errorString[s]
	if !ok {
		status = "unknown"
	}
	return fmt.Sprintf("%d - %s", s, status)
}

func (dlr DLRStatus) String() string {
	status, ok := dlrStatusString[dlr]
	if !ok {
		status = "unknown"
	}
	return fmt.Sprintf("%d - %s", dlr, status)
}

// MarshalXML implement the interface for XML Marshaling
func (b Bool) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if b {
		return e.EncodeElement(1, start)
	}
	return e.EncodeElement(0, start)
}

// MarshalXML implement the interface for XML Marshaling
func (u Unsubscribe) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return e.EncodeElement(int(u), start)
}

// MarshalXML implement the interface for XML Marshaling
func (s Status) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return e.EncodeElement(int(s), start)
}

// MarshalXML implement the interface for XML Marshaling
func (dlr DLRStatus) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return e.EncodeElement(int(dlr), start)
}

// UnmarshalXML implements the interface for unmarshaling xml
func (b *Bool) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var i int
	err := d.DecodeElement(&i, &start)
	if err != nil {
		return err
	}

	*b = Bool(i > 0)
	return nil
}

// UnmarshalXML implements the interface for unmarshaling xml
func (u *Unsubscribe) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var i int
	err := d.DecodeElement(&i, &start)
	if err != nil {
		return err
	}

	*u = Unsubscribe(i)
	return nil
}

// UnmarshalXML implements the interface for unmarshaling xml
func (s *Status) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var i int
	err := d.DecodeElement(&i, &start)
	if err != nil {
		return err
	}

	*s = Status(i)
	return nil
}

// UnmarshalXML implements the interface for unmarshaling xml
func (dlr *DLRStatus) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var i int
	err := d.DecodeElement(&i, &start)
	if err != nil {
		return err
	}

	*dlr = DLRStatus(i)
	return nil
}

// Value implements the database interface of Value
func (b Bool) Value() (driver.Value, error) {
	return bool(b), nil
}

// Value implements the database interface of Value
func (u Unsubscribe) Value() (driver.Value, error) {
	return int(u), nil
}

// Value implements the database interface of Value
func (s Status) Value() (driver.Value, error) {
	return int(s), nil
}

// Value implements the database interface of Value
func (dlr DLRStatus) Value() (driver.Value, error) {
	return int(dlr), nil
}

// Scan implements the database interface for Scan
func (b *Bool) Scan(src interface{}) error {
	if src == nil {
		return errors.New("src cannot be nil")
	}

	switch src.(type) {
	case string:
		str := reflect.ValueOf(src).String()
		boolean, err := strconv.ParseBool(str)
		if err != nil {
			return err
		}
		*b = Bool(boolean)
	case int, int8, int16, int32, int64:
		*b = Bool(reflect.ValueOf(src).Int() > 0)
	case float32, float64:
		*b = Bool(int(reflect.ValueOf(src).Float()) > 0)
	case bool:
		*b = Bool(reflect.ValueOf(src).Bool())
	default:
		return fmt.Errorf("Invalid type of src: %T", src)
	}

	return nil
}

// Scan implements the database interface for Scan
func (u *Unsubscribe) Scan(src interface{}) error {
	if src == nil {
		return errors.New("src cannot be nil")
	}

	switch src.(type) {
	case string:
		str := reflect.ValueOf(src).String()
		i, err := strconv.ParseInt(str, 10, 16)
		if err != nil {
			return err
		}
		*u = Unsubscribe(i)
	case int, int8, int16, int32, int64:
		*u = Unsubscribe(reflect.ValueOf(src).Int())
	case float32, float64:
		*u = Unsubscribe(int(reflect.ValueOf(src).Float()))
	default:
		return fmt.Errorf("Invalid type of src: %T", src)
	}

	return nil
}

// Scan implements the database interface for Scan
func (s *Status) Scan(src interface{}) error {
	if src == nil {
		return errors.New("src cannot be nil")
	}

	switch src.(type) {
	case string:
		str := reflect.ValueOf(src).String()
		i, err := strconv.ParseInt(str, 10, 16)
		if err != nil {
			return err
		}
		*s = Status(i)
	case int, int8, int16, int32, int64:
		*s = Status(reflect.ValueOf(src).Int())
	case float32, float64:
		*s = Status(int(reflect.ValueOf(src).Float()))
	default:
		return fmt.Errorf("Invalid type of src: %T", src)
	}

	return nil
}

// Scan implements the database interface for Scan
func (dlr *DLRStatus) Scan(src interface{}) error {
	if src == nil {
		return errors.New("src cannot be nil")
	}

	switch src.(type) {
	case string:
		str := reflect.ValueOf(src).String()
		i, err := strconv.ParseInt(str, 10, 16)
		if err != nil {
			return err
		}
		*dlr = DLRStatus(i)
	case int, int8, int16, int32, int64:
		*dlr = DLRStatus(reflect.ValueOf(src).Int())
	case float32, float64:
		*dlr = DLRStatus(int(reflect.ValueOf(src).Float()))
	default:
		return fmt.Errorf("Invalid type of src: %T", src)
	}

	return nil
}
