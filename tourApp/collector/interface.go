package collector

// Generated automatically by microgen: do not edit manually

import (
	"github.com/MarcGrol/microgen/envelope"
	"github.com/MarcGrol/microgen/myerrors"
)

type CommandHandler interface {
	HandleSearchQuery(eventType string, aggregateType string, aggregateUid string) (*SearchResults, *myerrors.Error)
}

type EventHandler interface {
	OnAnyEvent(event *envelope.Envelope) error
}
