package idprefix

const (
	Achievement = "ach"
	Event       = "evt"
	Message     = "msg"
	Transaction = "txn"
	User        = "usr"
)

var PrefixToEntity = map[string]string{
	Achievement: "Achievement",
	Event:       "Event",
	Message:     "Message",
	Transaction: "Transaction",
	User:        "User",
}
