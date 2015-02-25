package collector

import (
	"errors"
	"github.com/MarcGrol/microgen/myerrors"
	"github.com/MarcGrol/microgen/tourApp/events"
	"log"
)

type CollectorEventHandler struct {
	store events.Store
}

func NewCollectorEventHandler(store events.Store) EventHandler {
	handler := new(CollectorEventHandler)
	handler.store = store
	return handler
}

func (eh *CollectorEventHandler) OnAnyEvent(envelope *events.Envelope) error {
	log.Printf("OnEvent: envelope: %+v", envelope)
	return doStore(eh.store, []*events.Envelope{envelope})
}

type CollectorCommandHandler struct {
	bus   events.PublishSubscriber
	store events.Store
}

func NewCollectorCommandHandler(bus events.PublishSubscriber, store events.Store) CommandHandler {
	handler := new(CollectorCommandHandler)
	handler.bus = bus
	handler.store = store
	return handler
}

func doStore(store events.Store, envelopes []*events.Envelope) error {
	for _, env := range envelopes {
		err := store.Store(env)
		if err != nil {
			log.Printf("Error storing event: %+v", err)
			return err
		}
		log.Printf("Successfully stored event: %+v", env)
	}
	return nil
}

func (ch *CollectorCommandHandler) HandleSearchQuery(eventType string, aggregateType string, aggregateUid string) (*SearchResults, *myerrors.Error) {
	var err *myerrors.Error = nil

	results := new(SearchResults)
	if len(aggregateType) > 0 && len(aggregateUid) > 0 {
		results.Events, _ = ch.store.Get(aggregateType, aggregateUid)
	} else {
		err = myerrors.NewInternalError(errors.New("Search on event-type not supported"))
	}

	return nil, err
}

type SearchResults struct {
	Events []events.Envelope `json:"events"`
}
