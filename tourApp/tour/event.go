package tour

import "github.com/MarcGrol/microgen/infra"

type TourEventHandler struct {
	bus     infra.PublishSubscriber
	store   infra.Store
	context *tourContext
}

func NewTourEventHandler(bus infra.PublishSubscriber, store infra.Store, context *tourContext) *TourEventHandler {
	handler := new(TourEventHandler)
	handler.bus = bus
	handler.store = store
	handler.context = context
	return handler
}

func (eventHandler *TourEventHandler) Start() error {
	// do not subscribe to any events
	return nil
}
