package event

// The Dispatcher dispatches events through the communication channel
type Dispatcher struct {
	comm chan<- Event
}

// NewDispatcher instantiates a new dispatcher
func NewDispatcher(handler *Handler) *Dispatcher {
	return &Dispatcher{handler.comm}
}

// Emit sends an event to the receiver
func (d *Dispatcher) Emit(event Event) {
	d.comm <- event
}
