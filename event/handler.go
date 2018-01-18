package event

const channelBufferLength = 100

// Handler defines an IPC event handler
type Handler struct {
	comm chan Event
}

// NewHandler instantiates a new Handler object
func NewHandler() *Handler {
	return &Handler{
		comm: make(chan Event, channelBufferLength),
	}
}

// Handle fetches events from the IPC queue and returns them as slices
func (h *Handler) Handle() []Event {
	var events []Event

	for {
		select {
		case event := <-h.comm:
			events = append(events, event)
		default:
			return events
		}
	}
}
