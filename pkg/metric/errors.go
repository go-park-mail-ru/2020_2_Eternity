package metric

var ErrorMap = map[string]int{
	"user search error":   101,
	"pins search error":   102,
	"boards search error": 103,

	"CreateChat":         201,
	"GetChatById":        202,
	"GetUserChats":       203,
	"MarkMessagesRead":   204,
	"CreateMessage":      205,
	"DeleteMessage":      206,
	"GetLastNMessages":   207,
	"GetNMessagesBefore": 208,

	"invalid username":              301,
	"bad password":                  302,
	"can't create token":            303,
	"can't generate value":          304,
	"can't create csrf token value": 305,
	"parse token":                   306,
}

func fromErrorToCode(error string) int {
	code, ok := ErrorMap[error]
	if !ok {
		return 500
	}
	return code
}
