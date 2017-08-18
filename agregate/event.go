package agregate

type Event interface {
	GetAgregateId() AgregateId
	Apply(Agregate)
}

type BasicEvent struct {
	AgregateId AgregateId
}

func (e BasicEvent) GetAgregateId() AgregateId {
	return e.AgregateId
}
