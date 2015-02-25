package test

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/MarcGrol/microgen/myerrors"
	"github.com/MarcGrol/microgen/tourApp/events"
	"github.com/stretchr/testify/assert"
	"log"
	"os"
	"strings"
	"testing"
)

type Scenarios struct {
	Scenarios []Scenario
}

type ScenarioExecutorFunc func(scenario *Scenario) *myerrors.Error

type Scenario struct {
	Bus   events.PublishSubscriber `json:"-"`
	Store events.Store             `json:"-"`

	Title             string               `json:"title"`
	Given             []*events.Envelope   `json:"given"`
	When              ScenarioExecutorFunc `json:"-"`
	Command           interface{}          `json:"command"`
	Expect            []*events.Envelope   `json:"expect"`
	Actual            []*events.Envelope   `json:"actual"`
	Err               *myerrors.Error      `json:"err"`
	InvalidInputError bool                 `json:"invalidInputError"`
	NotFoundError     bool                 `json:"notFoundError"`
}

func (s *Scenario) RunAndVerify(t *testing.T) {

	s.Store = NewFakeStore()
	s.Bus = NewFakeBus()

	// store preconditions
	for _, given := range s.Given {
		s.Store.Store(given)
	}

	// subscribe to all expected topics to catch published evemts
	s.Actual = make([]*events.Envelope, 0, 10)
	callback := func(envelope *events.Envelope) error {
		s.Actual = append(s.Actual, envelope)
		return nil
	}
	for _, expected := range s.Expect {
		s.Bus.Subscribe(expected.Type, callback)
	}

	// execute operation on subject
	s.Err = s.When(s)

	if s.Err == nil {
		// basic ocmpare expected with actual
		assert.Equal(t, len(s.Expect), len(s.Actual))
		for idx, actual := range s.Actual {
			assert.Equal(t, s.Expect[idx].AggregateName, actual.AggregateName)
			assert.Equal(t, s.Expect[idx].AggregateUid, actual.AggregateUid)
			assert.Equal(t, s.Expect[idx].Type, actual.Type)
		}
	} else {
		s.InvalidInputError = myerrors.IsInvalidInputError(s.Err)
		s.NotFoundError = myerrors.IsNotFoundError(s.Err)
	}

	s.Dump(title2filename(s.Title))
}

func title2filename(title string) string {
	return "../doc/" + strings.Replace(title, " ", "_", -1) + ".json"
}

func (s *Scenario) Dump(filename string) error {
	w, err := os.Create(filename)
	if err != nil {
		log.Printf("Error opening json-file %s (%s)", filename, err.Error())
		return err
	}
	defer w.Close()

	jsondata, err := json.MarshalIndent(s, "", "  ")
	if err != nil {
		log.Printf("Error marshalling json %s", err.Error())
		return err
	}

	_, err = w.Write(jsondata)
	if err != nil {
		log.Printf("Error writing json to %s (%s)", filename, err.Error())
		return err
	}
	return nil
}

type FakeBus struct {
	callbacks     map[events.Type]events.EventHandlerFunc
	published     map[events.Type][]events.Envelope
	undeliverable map[events.Type][]events.Envelope
}

func NewFakeBus() *FakeBus {
	bus := new(FakeBus)
	bus.callbacks = make(map[events.Type]events.EventHandlerFunc)
	bus.undeliverable = make(map[events.Type][]events.Envelope)
	return bus
}

func (bus *FakeBus) Subscribe(eventType events.Type, callback events.EventHandlerFunc) error {
	bus.callbacks[eventType] = callback
	//log.Printf("FakeBus: subscribed to: %s", eventType.String())
	return nil
}

func (bus *FakeBus) Publish(envelope *events.Envelope) error {
	callback, ok := bus.callbacks[envelope.Type]
	if ok == false {
		bus.undeliverable[envelope.Type] = append(bus.undeliverable[envelope.Type], *envelope)
		//log.Printf("FakeBus: undeliverable: %v", envelope)
		return errors.New(fmt.Sprintf("Received event on non-subscribed channel %s", envelope.Type.String()))
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
	store.stored = make([]events.Envelope, 0, 10)
	return store
}

func (store *FakeStore) Store(envelope *events.Envelope) error {
	envelope.SequenceNumber = uint64(len(store.stored) + 1)
	store.stored = append(store.stored, *envelope)
	//log.Printf("FakeStore: stored: %v", envelope)
	return nil
}

func (store *FakeStore) Iterate(callback events.StoredItemHandlerFunc) error {
	for _, envelope := range store.stored {
		callback(&envelope)
	}
	return nil
}

func (store *FakeStore) Get(aggregateName string, aggregateUid string) ([]events.Envelope, error) {
	envelopes := make([]events.Envelope, 0, 10)

	callback := func(envelope *events.Envelope) {
		if envelope.AggregateName == aggregateName && envelope.AggregateUid == aggregateUid {
			envelopes = append(envelopes, *envelope)
		}
	}
	err := store.Iterate(callback)
	if err != nil {
		return nil, err
	}

	return envelopes, nil
}

const (
	DIRNAME  = "."
	FILENAME = "test.db"
)

func createRealStore() (*events.EventStore, error) {
	os.Remove(DIRNAME + "/" + FILENAME)
	store := events.NewEventStore(DIRNAME, FILENAME)
	err := store.Open()
	if err != nil {
		return nil, err
	}
	return store, nil
}

func createRealBus(scenarioName string) *events.EventBus {
	return events.NewEventBus("scenarioTest", scenarioName+"Test", "127.0.0.1")
}
