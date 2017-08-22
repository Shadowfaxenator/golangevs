package agregate

import (
	"strings"
)

var _EventRegistry EventRegistry = make(EventRegistry)

func RegisterEventType(es EventType, t Event) {
	_EventRegistry[es] = t
}
func ConstructEventFromReg(et EventType) Event {
	e := _EventRegistry[et]

	return e
}

type EventRegistry map[EventType]Event

type Event interface {
	Apply(Agregate)
}

type BasicEvent struct {
	Type EventType
}

func NewBasicEvent(t EventType) (e BasicEvent) {
	ss := strings.Split(string(t), ".")

	e = BasicEvent{Type: EventType(ss[len(ss)-1])}
	return
}
