package test

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/MarcGrol/microgen/infra"
	"github.com/MarcGrol/microgen/lib/envelope"
	"github.com/MarcGrol/microgen/lib/myerrors"
	"github.com/stretchr/testify/assert"
	"log"
	"os"
	"strings"
	"testing"
)

type CommandScenarioExecutorFunc func(scenario *CommandScenario) error

type CommandScenario struct {
	Bus   infra.PublishSubscriber `json:"-"`
	Store infra.Store             `json:"-"`

	Title             string                      `json:"title"`
	Given             []*envelope.Envelope        `json:"given"`
	When              CommandScenarioExecutorFunc `json:"-"`
	Command           interface{}                 `json:"command"`
	Expect            []*envelope.Envelope        `json:"expect"`
	Actual            []envelope.Envelope         `json:"actual"`
	ErrMsg            *string                     `json:"errMsg"`
	InvalidInputError bool                        `json:"invalidInputError"`
	NotFoundError     bool                        `json:"notFoundError"`
}

func (s *CommandScenario) RunAndVerify(t *testing.T) {

	s.Store = NewFakeStore()
	s.Bus = NewFakeBus()

	// store preconditions
	for _, given := range s.Given {
		s.Store.Store(given)
	}

	// subscribe to all expected topics to catch published evemts
	s.Actual = make([]envelope.Envelope, 0, 10)
	callback := func(envelope *envelope.Envelope) error {
		s.Actual = append(s.Actual, *envelope)
		return nil
	}
	for _, expected := range s.Expect {
		s.Bus.Subscribe(expected.EventTypeName, callback)
	}

	// execute operation on subject
	err := s.When(s)

	if err == nil {
		// basic ocmpare expected with actual
		assert.Equal(t, len(s.Expect), len(s.Actual))
		for idx, actual := range s.Actual {
			assert.Equal(t, s.Expect[idx].AggregateName, actual.AggregateName)
			assert.Equal(t, s.Expect[idx].AggregateUid, actual.AggregateUid)
			assert.Equal(t, s.Expect[idx].EventTypeName, actual.EventTypeName)
		}
	} else {
		s.ErrMsg = new(string)
		*s.ErrMsg = err.Error()
		s.InvalidInputError = myerrors.IsInvalidInputError(err)
		s.NotFoundError = myerrors.IsNotFoundError(err)
	}

	Dump(s, title2filename(s.Title))
}

type EventScenarioExecutorFunc func(scenario *EventScenario) error

type EventScenario struct {
	Bus   infra.PublishSubscriber `json:"-"`
	Store infra.Store             `json:"-"`

	Title   string                    `json:"title"`
	Given   []*envelope.Envelope      `json:"given"`
	When    EventScenarioExecutorFunc `json:"-"`
	Envelop *envelope.Envelope        `json:"envelope"`
	Expect  []*envelope.Envelope      `json:"expect"`
	Actual  []envelope.Envelope       `json:"actual"`
	ErrMsg  *string                   `json:"errMsg"`
}

func (s *EventScenario) RunAndVerify(t *testing.T) {

	s.Store = NewFakeStore()
	s.Bus = NewFakeBus()

	// store preconditions
	for _, given := range s.Given {
		s.Store.Store(given)
	}

	// subscribe to all expected topics to catch published evemts
	s.Actual = make([]envelope.Envelope, 0, 10)
	callback := func(envelope *envelope.Envelope) error {
		s.Actual = append(s.Actual, *envelope)
		return nil
	}
	for _, expected := range s.Expect {
		s.Bus.Subscribe(expected.EventTypeName, callback)
	}

	// execute operation on subject
	err := s.When(s)

	// get messages from store
	idx := 0
	s.Store.Iterate(func(envelop *envelope.Envelope) {
		s.Actual = append(s.Actual, *envelop)
		idx++
	})

	if err == nil {
		// basic ocmpare expected with actual
		assert.Equal(t, len(s.Expect), len(s.Actual))
		for idx := 0; idx < len(s.Expect); idx++ {
			expected := s.Expect[idx]
			actual := s.Actual[idx]
			assert.Equal(t, expected.Uuid, actual.Uuid)
			assert.Equal(t, expected.Timestamp, actual.Timestamp)
			assert.Equal(t, expected.SequenceNumber, actual.SequenceNumber)
			assert.Equal(t, expected.AggregateName, actual.AggregateName)
			assert.Equal(t, expected.AggregateUid, actual.AggregateUid)
			assert.Equal(t, expected.EventTypeName, actual.EventTypeName)
		}
	} else {
		s.ErrMsg = new(string)
		*s.ErrMsg = err.Error()
	}

	Dump(s, title2filename(s.Title))
}

func title2filename(title string) string {
	return "../doc/" + strings.Replace(title, " ", "_", -1) + ".json"
}

func Dump(scenario interface{}, filename string) error {
	w, err := os.Create(filename)
	if err != nil {
		log.Printf("Error opening json-file %s (%s)", filename, err.Error())
		return err
	}
	defer w.Close()

	jsondata, err := json.MarshalIndent(scenario, "", "  ")
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
	callbacks     map[string]infra.EventHandlerFunc
	published     map[string][]envelope.Envelope
	undeliverable map[string][]envelope.Envelope
}

func NewFakeBus() *FakeBus {
	bus := new(FakeBus)
	bus.callbacks = make(map[string]infra.EventHandlerFunc)
	bus.undeliverable = make(map[string][]envelope.Envelope)
	return bus
}

func (bus *FakeBus) Subscribe(eventTypeName string, callback infra.EventHandlerFunc) error {
	bus.callbacks[eventTypeName] = callback
	//log.Printf("FakeBus: subscribed to: %s", eventType.String())
	return nil
}

func (bus *FakeBus) Publish(envelop *envelope.Envelope) error {
	callback, ok := bus.callbacks[envelop.EventTypeName]
	if ok == false {
		bus.undeliverable[envelop.EventTypeName] = append(bus.undeliverable[envelop.EventTypeName], *envelop)
		//log.Printf("FakeBus: undeliverable: %v", envelope)
		return errors.New(fmt.Sprintf("Received event on non-subscribed channel %s", envelop.EventTypeName))
	} else {
		callback(envelop)
	}
	return nil
}

type FakeStore struct {
	stored []envelope.Envelope
}

func NewFakeStore() *FakeStore {
	store := new(FakeStore)
	store.stored = make([]envelope.Envelope, 0, 10)
	return store
}

func (store *FakeStore) Store(envelope *envelope.Envelope) error {
	envelope.SequenceNumber = uint64(len(store.stored) + 1)
	store.stored = append(store.stored, *envelope)
	return nil
}

func (store *FakeStore) Iterate(callback infra.StoredItemHandlerFunc) error {
	for _, envelope := range store.stored {
		callback(&envelope)
	}
	return nil
}

func (store *FakeStore) Get(aggregateName string, aggregateUid string) ([]envelope.Envelope, error) {
	envelopes := make([]envelope.Envelope, 0, 10)

	callback := func(envelop *envelope.Envelope) {
		if envelop.AggregateName == aggregateName && envelop.AggregateUid == aggregateUid {
			envelopes = append(envelopes, *envelop)
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
