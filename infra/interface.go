package infra

import (
	"github.com/MarcGrol/microgen/lib/envelope"
)

type EventHandlerFunc func(Envelope *envelope.Envelope) error

type PublishSubscriber interface {
	Subscribe(eventTypeName string, callback EventHandlerFunc) error
	Publish(envelope *envelope.Envelope) error
}

type StoredItemHandlerFunc func(envelope *envelope.Envelope)

type Store interface {
	Store(envelope *envelope.Envelope) error
	Iterate(callback StoredItemHandlerFunc) error
	Get(aggregateName string, aggregateUid string) ([]envelope.Envelope, error)
}
