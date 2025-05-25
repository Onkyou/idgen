package idprefix

const (
	Achievement     = "ach"
	Event           = "evt"
	Message         = "msg"
	PointOfInterest = "poi"
	Transaction     = "txn"
	User            = "usr"
)

var PrefixToEntity = map[string]string{
	Achievement:     "Achievement",
	Event:           "Event",
	Message:         "Message",
	PointOfInterest: "PointOfInterest",
	Transaction:     "Transaction",
	User:            "User",
}
