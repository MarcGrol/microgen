package news

import (
	"log"

	"github.com/MarcGrol/microgen/infra"
	"github.com/MarcGrol/microgen/lib/envelope"
	"github.com/MarcGrol/microgen/tourApp/events"
)

type NewsEventHandler struct {
	bus         infra.PublishSubscriber
	store       infra.Store
	newsContext *NewsContext
}

func NewNewsEventHandler(bus infra.PublishSubscriber, store infra.Store, newsContext *NewsContext) *NewsEventHandler {
	handler := new(NewsEventHandler)
	handler.bus = bus
	handler.store = store
	handler.newsContext = newsContext

	return handler
}

func (eventHandler *NewsEventHandler) Start() error {

	for _, eventType := range events.GetTourEventTypes() {
		err := eventHandler.bus.Subscribe(eventType.String(), func(envelope *envelope.Envelope) error {
			return eventHandler.OnEvent(envelope)
		})
		if err != nil {
			return err
		}
	}

	return nil
}

func (eh *NewsEventHandler) OnEvent(envelop *envelope.Envelope) error {
	// store event on disk to testore
	err := doStore(eh.store, []*envelope.Envelope{envelop})
	if err != nil {
		return err
	}
	// apply event in memory
	applyEvent(*envelop, eh.newsContext)

	return nil
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
