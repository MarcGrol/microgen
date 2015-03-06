package collector

// Generated automatically by microgen: do not edit manually

import (
	"github.com/MarcGrol/microgen/lib/envelope"
)

type CommandHandler interface {
	Start(listenPort int) error
	HandleSearchQuery(eventType string, aggregateType string, aggregateUid string) (*SearchResults, error)
}

type EventHandler interface {
	Start() error
	OnAnyEvent(event *envelope.Envelope) error
}
