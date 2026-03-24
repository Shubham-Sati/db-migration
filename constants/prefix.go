package constants

// Prefix contains all table-specific PID prefixes for the chat-analytics system
// Each table has a unique prefix to identify records across services
var Prefix = struct {
	USER string

	ROOM               string
	ROOM_MEMBER        string
	MESSAGE            string
	MESSAGE_REACTION   string
	MESSAGE_ATTACHMENT string

	EVENT        string
	USER_SESSION string
	DAILY_METRIC string
	ROOM_METRIC  string
	USER_METRIC  string

	TRACE_ID string
}{

	USER: "usr",

	ROOM:               "room",
	ROOM_MEMBER:        "rmbr",
	MESSAGE:            "msg",
	MESSAGE_REACTION:   "msgreac",
	MESSAGE_ATTACHMENT: "msgatt",

	EVENT:        "evt",
	USER_SESSION: "usrsess",
	DAILY_METRIC: "dmet",
	ROOM_METRIC:  "rmet",
	USER_METRIC:  "umet",

	TRACE_ID: "trace",
}
