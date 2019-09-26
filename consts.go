package telzarsms

// Basic Addresses for contacting Telzar
const (
	ProdAPIAddress  = `https://www.019sms.co.il/api`
	DevAPIAddress   = `https://www.019sms.co.il:8090/api/test`
	HTTPContentType = `text/xml; charset=utf-8`
	HTTPMethod      = "POST"
)

// Supported types of Unsubscribe
const (
	UnsubscribeOff Unsubscribe = iota + 1
	UnsubscribeByReturnSMS
	UnsubscribeByLink
)

var unsubscribeType = map[Unsubscribe]string{
	UnsubscribeOff:         "off",
	UnsubscribeByReturnSMS: "remove by return sms",
	UnsubscribeByLink:      "remove by link",
}

// Error Codes
const (
	ErrorOK                              Status = 0
	ErrorXMLParsing                      Status = 1
	ErrorMissingRequiredXMLField         Status = 2
	ErrorIncorrectUserNameOrPassword     Status = 3
	ErrorNotEnoughCredit                 Status = 4
	ErrorNoPermissionForSMS              Status = 5
	ErrorInvalidPhone                    Status = 933
	ErrorSomeNumbersNotBlacklisted       Status = 944
	ErrorCamppaignAlreadyCanceled        Status = 955
	ErrorCampaignAlreadySent             Status = 966
	ErrorInvalidGivenCampaign            Status = 977
	ErrorCantUseMultipleContactLists     Status = 987
	ErrorContactListNotExist             Status = 988
	ErrorInvalidMessageLength            Status = 989
	ErrorAmountIsTooHighForCurrentCredit Status = 990
	ErrorNonIntegerValuesForAmount       Status = 991
	ErrorInvalidSourceLength             Status = 992
	ErrorInvalidPasswordLength           Status = 993
	ErrorUsernameAlreadyExists           Status = 994
	ErrorInvalidUsernameLength           Status = 995
	ErrorInvalidNameLength               Status = 996
	ErrorInvalidCommand                  Status = 997
	ErrorUnknownErrorInRequest           Status = 998
	ErrorContactSupport                  Status = 999
)

var errorString = map[Status]string{
	ErrorOK:                              "OK",
	ErrorXMLParsing:                      "There was problem parsing your XML",
	ErrorMissingRequiredXMLField:         "Missing required XML field",
	ErrorIncorrectUserNameOrPassword:     "Username or password is incorrect",
	ErrorNotEnoughCredit:                 "Not enough credit",
	ErrorNoPermissionForSMS:              "No permission to send SMS at this time",
	ErrorInvalidPhone:                    "Invalid phone number",
	ErrorSomeNumbersNotBlacklisted:       "Some numbers are not in blacklist",
	ErrorCamppaignAlreadyCanceled:        "Campaign already canceled",
	ErrorCampaignAlreadySent:             "Campaign already sent",
	ErrorInvalidGivenCampaign:            "Campaign does not belong to customers or does not exists",
	ErrorCantUseMultipleContactLists:     "Can't use more then one contact list of singles phone numbers when using dynamic fields",
	ErrorContactListNotExist:             "Contact list does not exists",
	ErrorInvalidMessageLength:            "The message length is too short or too long",
	ErrorAmountIsTooHighForCurrentCredit: "Amount must be smaller then credit",
	ErrorNonIntegerValuesForAmount:       "Amount must contain only digits",
	ErrorInvalidSourceLength:             "source is too long or too short",
	ErrorInvalidPasswordLength:           "password is too long or too short",
	ErrorUsernameAlreadyExists:           "Username already exists",
	ErrorInvalidUsernameLength:           "Username is too long or too short",
	ErrorInvalidNameLength:               "Name is too long or too short",
	ErrorInvalidCommand:                  "Invalid command was sent",
	ErrorUnknownErrorInRequest:           "Unknown error on request",
	ErrorContactSupport:                  "Contact support",
}

// DLR status codes
const (
	DLRStatusArrived                DLRStatus = 0
	DLRStatusFailed1                DLRStatus = 1
	DLRStatusTimeout                DLRStatus = 2
	DLRStatusFailed2                DLRStatus = 3
	DLRStatusCellularFailed1        DLRStatus = 4
	DLRStatusFailed3                DLRStatus = 5
	DLRStatusFailed4                DLRStatus = 6
	DLRStatusCellularFailed2        DLRStatus = 14
	DLRStatusKosherNumber           DLRStatus = 15
	DLRStatusDeliveryTimePermission DLRStatus = 16
	DLRStatusNotArrived1            DLRStatus = 101
	DLRStatusArrivedDestination     DLRStatus = 102
	DLRStatusExpired                DLRStatus = 103
	DLRStatusDeleted                DLRStatus = 104
	DLRStatusNotArrived2            DLRStatus = 105
	DLRStatusNotArrived3            DLRStatus = 106
	DLRStatusNotArrived4            DLRStatus = 107
	DLRStatusRejected               DLRStatus = 108
	DLRStatusNotArrived5            DLRStatus = 109
	DLRStatusBlockedOnRequest       DLRStatus = 201
	DLRStatusPermissionDenied       DLRStatus = 998
	DLRStatusUnknownCause           DLRStatus = 999
)

var dlrStatusString = map[DLRStatus]string{
	DLRStatusArrived:                "Arrived to destination",
	DLRStatusFailed1:                "Failed 1",
	DLRStatusTimeout:                "Timeout",
	DLRStatusFailed2:                "Failed 2",
	DLRStatusCellularFailed1:        "Cellular failed",
	DLRStatusFailed3:                "Failed 3",
	DLRStatusFailed4:                "Failed 4",
	DLRStatusCellularFailed2:        "Cellular failed",
	DLRStatusKosherNumber:           "Kosher Number",
	DLRStatusDeliveryTimePermission: "No permission to send scheduled message",
	DLRStatusNotArrived1:            "Not arrived 1",
	DLRStatusArrivedDestination:     "Arrived destination",
	DLRStatusExpired:                "Expired",
	DLRStatusDeleted:                "Deleted",
	DLRStatusNotArrived2:            "Not arrived 2",
	DLRStatusNotArrived3:            "Not arrived 3",
	DLRStatusNotArrived4:            "Not arrived 4",
	DLRStatusRejected:               "Rejected",
	DLRStatusNotArrived5:            "Not arrived 5",
	DLRStatusBlockedOnRequest:       "Blocked on request",
	DLRStatusPermissionDenied:       "Permission denied",
	DLRStatusUnknownCause:           "Unknown error",
}
