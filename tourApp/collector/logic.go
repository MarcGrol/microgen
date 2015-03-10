package collector

import (
	"errors"
	"github.com/MarcGrol/microgen/infra"
	"github.com/MarcGrol/microgen/lib/envelope"
	"github.com/MarcGrol/microgen/lib/myerrors"
	"github.com/MarcGrol/microgen/tourApp/events"
	"log"
)

type CollectorCommandHandler struct {
	bus   infra.PublishSubscriber
	store infra.Store
}

func NewCollectorCommandHandler(bus infra.PublishSubscriber, store infra.Store) CommandHandler {
	handler := new(CollectorCommandHandler)
	handler.bus = bus
	handler.store = store
	return handler
}

type CollectorEventHandler struct {
	bus   infra.PublishSubscriber
	store infra.Store
}

func NewCollectorEventHandler(bus infra.PublishSubscriber, store infra.Store) EventHandler {
	handler := new(CollectorEventHandler)
	handler.bus = bus
	handler.store = store
	return handler
}

func (eventHandler *CollectorEventHandler) Start() error {

	for _, eventType := range events.GetAllEventTypes() {
		eventHandler.bus.Subscribe(eventType.String(), func(envelope *envelope.Envelope) error {
			return eventHandler.OnAnyEvent(envelope)
		})
	}

	return nil
}

func (eh *CollectorEventHandler) OnAnyEvent(envelop *envelope.Envelope) error {
	log.Printf("OnEvent: envelope: %+v", envelop)
	return doStore(eh.store, []*envelope.Envelope{envelop})
}

func doStore(store infra.Store, envelopes []*envelope.Envelope) error {
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

func (ch *CollectorCommandHandler) HandleSearchQuery(eventType string, aggregateType string, aggregateUid string) (*SearchResults, error) {
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
	Events []envelope.Envelope `json:"events"`
}
