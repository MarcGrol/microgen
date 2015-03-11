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

	OnEtappeResultsCreated(event *events.EtappeResultsCreated) error
	OnNewsItemCreated(event *events.NewsItemCreated) error
	OnTourCreated(event *events.TourCreated) error
	OnCyclistCreated(event *events.CyclistCreated) error
	OnEtappeCreated(event *events.EtappeCreated) error
}

type EventApplier interface {
	ApplyTourCreated(event *events.TourCreated)
	ApplyCyclistCreated(event *events.CyclistCreated)
	ApplyEtappeCreated(event *events.EtappeCreated)
	ApplyEtappeResultsCreated(event *events.EtappeResultsCreated)
	ApplyNewsItemCreated(event *events.NewsItemCreated)
}