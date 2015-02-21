package events

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/MarcGrol/microgen/store"
	"log"
	"sync"
)

type EventStore struct {
	dirname            string
	filename		   string
	store              store.SimpleEventStore
	mutex              sync.RWMutex
	lastSequenceNumber uint64
}

func NewEventStore( dirname string, filename string ) *EventStore {
	store := new(EventStore)
	store.dirname =dirname
	store.filename = filename
	return store
}

func (store *EventStore) Open() error {
	store.mutex.Lock()
	defer store.mutex.Unlock()

	err := store.store.Open(store.dirname, store.filename)
	if err != nil {
		return err
	}
	store.lastSequenceNumber = store.getLastSequenceNumber()
	return nil
}

func (store *EventStore) Store(envelope *Envelope) error {
	store.mutex.Lock()
	defer store.mutex.Unlock()

	store.assignSequenceNumber(envelope)

	return store.writeEvent(envelope)
}

func (store *EventStore) writeEvent(envelope *Envelope) error {
	//log.Printf("write event: %v\n", envelope)

	// serialize event to json
	jsonBlob, err := json.Marshal(envelope)
	if err != nil {
		return errors.New(fmt.Sprintf("Error marshalling event (%v)", err))
	}
	//log.Printf("Marshalled envelope of type %d into %d bytes", envelope.Type, len(jsonBlob))

	return store.store.Append(jsonBlob)
}

func (store *EventStore) Iterate(handlerFunc StoredItemHandlerFunc) error {
	store.mutex.RLock()
	defer store.mutex.RUnlock()

	return store.iterate(handlerFunc)
}

func (store *EventStore) iterate(handlerFunc StoredItemHandlerFunc) error {
	callback := func(blob []byte) {
		var envelope Envelope
		err := json.Unmarshal(blob, &envelope)
		if err != nil {
			log.Printf("Error unmarshalling json blob (%v)", err)
		} else {
			//log.Printf("read event: %v\n", envelope)
			handlerFunc(&envelope)
		}
	}
	return store.store.Iterate(callback)
}

func (store *EventStore) assignSequenceNumber(envelope *Envelope) {
	store.lastSequenceNumber = store.lastSequenceNumber + 1
	envelope.SequenceNumber = store.lastSequenceNumber
}

func (store *EventStore) getLastSequenceNumber() uint64 {
	var lastIndex uint64 = 0

	callback := func(envelope *Envelope) {
		lastIndex++
	}
	store.iterate(callback)

	return lastIndex
}

func (store *EventStore) Close() {
	store.mutex.Lock()
	defer store.mutex.Unlock()
	store.store.Close()
}

func (store *EventStore) Get(aggregateName string, aggregateUid string) ([]Envelope, error) {
	envelopes := make([]Envelope, 0, 10)

	callback := func(envelope *Envelope) {
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
