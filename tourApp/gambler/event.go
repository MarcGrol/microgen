package gambler

import (
	"github.com/MarcGrol/microgen/infra"
	"github.com/MarcGrol/microgen/lib/envelope"
	"github.com/MarcGrol/microgen/tourApp/events"
)

type GamblerEventHandler struct {
	bus     infra.PublishSubscriber
	store   infra.Store
	context *GamblingContext
}

func NewGamblerEventHandler(bus infra.PublishSubscriber, store infra.Store, context *GamblingContext) *GamblerEventHandler {
	handler := new(GamblerEventHandler)
	handler.bus = bus
	handler.store = store
	handler.context = context
	return handler
}

func (eventHandler *GamblerEventHandler) Start() error {
	envelopes, err := eventHandler.store.GetAll()
	if err != nil {
		return err
	}
	eventHandler.context.ApplyAll(envelopes)

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

func (eh *GamblerEventHandler) OnEvent(envelop *envelope.Envelope) error {
	err := doStore(eh.store, []*envelope.Envelope{envelop})
	if err != nil {
		return err
	}
	// apply event in memory
	applyEvent(*envelop, eh.context)

	return nil
}
