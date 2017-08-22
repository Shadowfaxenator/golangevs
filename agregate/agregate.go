package agregate

import (
	"bytes"
	"encoding/base64"
	"encoding/gob"
	"log"
	"sync"

	"gopkg.in/mgo.v2"

	"gopkg.in/mgo.v2/bson"
)

type AgregateId bson.ObjectId
type EventType string
type AgregateType string

type Agregate interface {
	incrementVersion()
	appendPendingEvent(e Event)
	Self() *BasicAgregate
}

type BasicAgregate struct {
	ID            bson.ObjectId
	Version       int
	PendingEvents []Event
	// Mutex to lock
	*sync.RWMutex `bson:"-"`
	Type          AgregateType
}

func (a BasicAgregate) GetBSON() (interface{}, error) {
	var buf bytes.Buffer

	enc := gob.NewEncoder(&buf)

	if err := enc.Encode(a.PendingEvents); err != nil {
		return nil, err
	}
	return struct {
		ID      bson.ObjectId `bson:"agregateID"`
		Version int
		Type    string
		Events  string
	}{a.ID, a.Version, string(a.Type), base64.StdEncoding.EncodeToString(buf.Bytes())}, nil
}

func (a *BasicAgregate) SetBSON(raw bson.Raw) error {
	temp := &struct {
		ID      bson.ObjectId `bson:"agregateID"`
		Version int
		Type    string
		Events  string
	}{}
	if err := raw.Unmarshal(temp); err != nil {
		return err
	}
	b, err := base64.StdEncoding.DecodeString(temp.Events)
	if err != nil {
		return err
	}
	dec := gob.NewDecoder(bytes.NewReader(b))
	var events []Event
	if err := dec.Decode(&events); err != nil {
		return err
	}

	a.ID = temp.ID
	a.PendingEvents = events
	a.Type = AgregateType(temp.Type)
	a.Version = temp.Version
	a.RWMutex = new(sync.RWMutex)
	return nil
}

func (a *BasicAgregate) Self() *BasicAgregate {
	return a
}
func (a *BasicAgregate) incrementVersion() {
	a.Lock()
	defer a.Unlock()
	a.Version++
}

func (a *BasicAgregate) appendPendingEvent(e Event) {
	a.Lock()
	defer a.Unlock()
	a.PendingEvents = append(a.PendingEvents, e)
}

// func (state *BasicAgregate) trackChange(e Event) {
// 	state.changes = append(state.changes, e)
// 	state.transition(e)
// }

func Construct(ag Agregate) error {
	ses, err := NewRepo()
	defer ses.Close()
	if err != nil {
		return err
	}
	c := ses.DB("").C(string(ag.Self().Type))
	iter := c.Find(bson.M{"agregateID": ag.Self().ID}).Iter()
	defer func() {
		if err := iter.Close(); err != nil {
			log.Fatal(err)
		}
	}()
	var item BasicAgregate

	for iter.Next(&item) {
		for _, e := range item.PendingEvents {

			e.Apply(ag)
			ag.incrementVersion()
		}
	}
	return nil

}

func AddChange(ag Agregate, e Event) {

	e.Apply(ag)
	ag.appendPendingEvent(e)

}

func NewBasicAgregate(id AgregateId, t AgregateType) *BasicAgregate {

	b := &BasicAgregate{}

	b.ID = bson.ObjectId(id)

	b.Type = t
	b.RWMutex = new(sync.RWMutex)
	return b
}

func Commit(ag Agregate) error {

	ses, err := NewRepo()
	defer ses.Close()
	if err != nil {
		return err
	}
	c := ses.DB("").C(string(ag.Self().Type))
	ind := mgo.Index{

		Key:    []string{"version", "agregateID"},
		Unique: true,
	}

	err = c.EnsureIndex(ind)
	if err != nil {
		return err
	}
	if err := c.Insert(ag.Self()); err != nil {

		return err
	}
	return nil

}

func NewAgregateID() AgregateId {
	return AgregateId(bson.NewObjectId())
}
