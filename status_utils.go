package telzarsms

import "strconv"

// StrToStatus converts status that arrives as string to Status type.
// On error return -1
func StrToStatus(str string) Status {
	i, err := strconv.ParseInt(str, 10, 16)
	if err != nil {
		i = -1
	}
	return Status(i)
}
