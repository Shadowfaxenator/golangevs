package agregate

import (
	"sync"
)

type AgregateId int

type Agregate interface {
	getId() AgregateId
	incrementVersion()
	appendPendingEvent(e Event)
}

type BasicAgregate struct {
	Id              AgregateId
	ExpectedVersion int
	Changes         []Event
	Mu              *sync.RWMutex
}

func (a *BasicAgregate) incrementVersion() {
	a.Mu.Lock()
	defer a.Mu.Unlock()
	a.ExpectedVersion++
}
func (a *BasicAgregate) getId() AgregateId {
	a.Mu.RLock()
	defer a.Mu.RUnlock()
	return a.Id
}

func (a *BasicAgregate) appendPendingEvent(e Event) {
	a.Mu.Lock()
	defer a.Mu.Unlock()
	a.Changes = append(a.Changes, e)
}

// func (state *BasicAgregate) trackChange(e Event) {
// 	state.changes = append(state.changes, e)
// 	state.transition(e)
// }

func Construct(ag Agregate, events []Event) {

	for _, e := range events {

		if e.GetAgregateId() == ag.getId() {
			e.Apply(ag)
			ag.incrementVersion()

		}

	}
}

func AddChange(ag Agregate, e Event) {
	if e.GetAgregateId() == ag.getId() {
		e.Apply(ag)
		ag.appendPendingEvent(e)
	}
}

// func NewBasicAgregate() *BasicAgregate {
// 	b := &BasicAgregate{Mu: new(sync.RWMutex)}
// 	return
// }
