package vex

type VEventHandler struct {
	events map[string][]func()
}

func (e *VEventHandler) RegisterEvent(event string, runnable func()) {
	if _, ok := e.events[event]; !ok {
		e.events[event] = make([]func(), 0)
	}

	e.events[event] = append(e.events[event], runnable)
}

func (e *VEventHandler) FireEvent(event string) {
	if _, ok := e.events[event]; ok {
		for _, eventRunnable := range e.events[event] {
			eventRunnable()
		}
	}
}

func NewEventHandler() *VEventHandler {
	return &VEventHandler{events: make(map[string][]func())}
}
