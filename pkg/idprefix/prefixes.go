package idprefix

const (
	User        = "usr"
	Transaction = "txn"
	Event       = "evt"
	Achievement = "ach"
)

var PrefixToEntity = map[string]string{
	User:        "User",
	Transaction: "Transaction",
	Event:       "Event",
	Achievement: "Achievement",
}
