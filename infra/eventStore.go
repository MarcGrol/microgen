package infra

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/MarcGrol/microgen/envelope"
	"github.com/MarcGrol/microgen/store"
	"log"
	"sync"
)

type EventStore struct {
	store              *store.FileBlobStore
	mutex              sync.RWMutex
	lastSequenceNumber uint64
}

func NewEventStore(dirname string, filename string) *EventStore {
	s := new(EventStore)
	s.store = store.NewFileBlobStore(dirname, filename)
	return s
}

func (s *EventStore) Open() error {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	err := s.store.Open()
	if err != nil {
		return err
	}
	s.lastSequenceNumber = s.getLastSequenceNumber()
	return nil
}

func (s *EventStore) Store(envelope *envelope.Envelope) error {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	s.assignSequenceNumber(envelope)

	return s.writeEvent(envelope)
}

func (s *EventStore) writeEvent(envelope *envelope.Envelope) error {
	log.Printf("write event: %v\n", envelope)

	// serialize event to json
	jsonBlob, err := json.Marshal(envelope)
	if err != nil {
		return errors.New(fmt.Sprintf("Error marshalling event (%+v)", err))
	}
	//log.Printf("Marshalled envelope of type %d into %d bytes", envelope.Type, len(jsonBlob))

	return s.store.Append(jsonBlob)
}

func (s *EventStore) Iterate(handlerFunc StoredItemHandlerFunc) error {
	s.mutex.RLock()
	defer s.mutex.RUnlock()

	return s.iterate(handlerFunc)
}

func (s *EventStore) iterate(handlerFunc StoredItemHandlerFunc) error {
	callback := func(blob []byte) {
		var envelope envelope.Envelope
		err := json.Unmarshal(blob, &envelope)
		if err != nil {
			log.Printf("Error unmarshalling json blob (%+v)", err)
			return
		}
		//log.Printf("read event: %v\n", envelope)
		handlerFunc(&envelope)
	}
	return s.store.Iterate(callback)
}

func (s *EventStore) assignSequenceNumber(envelope *envelope.Envelope) {
	s.lastSequenceNumber = s.lastSequenceNumber + 1
	envelope.SequenceNumber = s.lastSequenceNumber
}

func (s *EventStore) getLastSequenceNumber() uint64 {
	var lastIndex uint64 = 0

	callback := func(envelope *envelope.Envelope) {
		lastIndex++
	}
	s.iterate(callback)

	return lastIndex
}

func (s *EventStore) Close() {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	s.store.Close()
}

func (s *EventStore) Get(aggregateName string, aggregateUid string) ([]envelope.Envelope, error) {
	envelopes := make([]envelope.Envelope, 0, 10)

	callback := func(envelope *envelope.Envelope) {
		if envelope.AggregateName == aggregateName && envelope.AggregateUid == aggregateUid {
			envelopes = append(envelopes, *envelope)
		}
	}
	err := s.Iterate(callback)
	if err != nil {
		return nil, err
	}

	return envelopes, nil
}
