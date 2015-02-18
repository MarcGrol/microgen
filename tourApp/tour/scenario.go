package tour

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"github.com/MarcGrol/microgen/tourApp/events"
	"os"
	"log"
)

type Scenarios struct {
	Scenarios []Scenario
}

type ScenarioExecutorFunc func( scenario *Scenario) error


type Scenario struct {
	bus events.PublishSubscriber
	store events.Store

	Name string
	Description string
	Given []events.Envelope
	CallSubject  ScenarioExecutorFunc
	Expect []events.Envelope
	Actual []*events.Envelope
}

func (s *Scenario) RunAndVerify(t *testing.T) {

	s.store = NewFakeStore()
	s.bus = NewFakeBus()
	
	// store preconditions
	for _,given := range s.Given {
		s.store.Store(&given)
	}

	// catch published evemts
	s.Actual = make([]*events.Envelope,0,10)
	callback := func(envelope *events.Envelope) error {
		s.Actual = append(s.Actual, envelope)
		return nil
	}

	for _,expected := range s.Expect {
		s.bus.Subscribe(expected.Type, callback)
	}
	
	// execute operation on subject
	err := s.CallSubject(s)
	assert.Nil(t, err)

	// compare expected with actual
	assert.Equal(t,len(s.Expect),len(s.Actual))
	for idx,actual := range s.Actual {
		assert.Equal(t, s.Expect[idx].SequenceNumber, actual.SequenceNumber)
		assert.Equal(t, s.Expect[idx].AggregateName, actual.AggregateName)
		assert.Equal(t, s.Expect[idx].AggregateUid, actual.AggregateUid)
		assert.Equal(t, s.Expect[idx].Type, actual.Type)		
	}
}

type FakeBus struct {
	callbacks map[events.Type]events.EventHandlerFunc
	published map[events.Type][]events.Envelope 
	undeliverable map[events.Type][]events.Envelope 
}

func NewFakeBus() *FakeBus {
	bus := new(FakeBus)
	bus.callbacks = make(map[events.Type]events.EventHandlerFunc)
	bus.undeliverable = make(map[events.Type][]events.Envelope)
	return bus
}

func (bus *FakeBus) Subscribe(eventType events.Type, callback events.EventHandlerFunc) error {
	log.Printf( "FakeBus: subscribing to: %s", eventType.String())
	bus.callbacks[eventType] = callback
	return nil
}

func (bus *FakeBus) Publish(envelope *events.Envelope) error {
	callback, ok := bus.callbacks[envelope.Type]
	if ok == false {
		log.Printf( "FakeBus: undeliverable: %v", envelope)
		bus.undeliverable[envelope.Type] = append(bus.undeliverable[envelope.Type], *envelope)	
	} else {
		callback(envelope)
	}
	return nil
}

type FakeStore struct {
	stored []events.Envelope
}

func NewFakeStore() *FakeStore {
	store := new(FakeStore)
	store.stored = make([]events.Envelope,0,10)
	return store
}

func (store *FakeStore)  Store(envelope *events.Envelope) error {
	envelope.SequenceNumber = uint64(len(store.stored)+1)
	store.stored = append( store.stored, *envelope)
	log.Printf( "FakeStore: stored: %v", envelope)
	return  nil
}

func (store *FakeStore)  Iterate(callback events.StoredItemHandlerFunc) error {
	for _,envelope := range store.stored {
		callback(&envelope)
	}
	return nil
}

func createRealStore() (*events.EventStore, error) {
	os.Remove(FILENAME)
	store := events.NewEventStore()
	err := store.Open(FILENAME)
	if err != nil {
		return nil, err
	}
	return store, nil
}


func createRealBus(scenarioName string) *events.EventBus {
	return events.NewEventBus("scenarioTest", scenarioName+"Test" , "127.0.0.1")
}
