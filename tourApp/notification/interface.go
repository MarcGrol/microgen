package notification

// Generated automatically by microgen: do not edit manually

import (
	"github.com/MarcGrol/microgen/lib/envelope"
	"github.com/MarcGrol/microgen/tourApp/events"
)

// commands

type SubscribeToNotificationsCommand struct {
	Email string `json:"email" binding:"required"`
}

type CommandHandler interface {
	Start(listenPort int)

	HandleSubscribeToNotificationsCommand(command *SubscribeToNotificationsCommand) error
}

// events

type EventHandler interface {
	Start()
	OnEnvelope(envelop *envelope.Envelope) error

	OnNewsItemCreated(event *events.NewsItemCreated) error
	OnEtappeResultsCreated(event *events.EtappeResultsCreated) error
}

type EventApplier interface {
	ApplyEtappeResultsCreated(event *events.EtappeResultsCreated)
	ApplyNewsItemCreated(event *events.NewsItemCreated)
}
