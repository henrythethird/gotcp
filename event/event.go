package event

// Event describes a generic IPC event
type Event struct {
	Kind string
	Body interface{}
}
