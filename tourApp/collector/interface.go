package collector

// Generated automatically by microgen: do not edit manually

import (
	"github.com/MarcGrol/microgen/lib/envelope"
)

type CommandHandler interface {
	HandleSearchQuery(eventType string, aggregateType string, aggregateUid string) (*SearchResults, error)
}

type EventHandler interface {
	OnAnyEvent(event *envelope.Envelope) error
}
