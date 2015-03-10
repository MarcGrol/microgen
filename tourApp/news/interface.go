package news

// Generated automatically by microgen: do not edit manually

import (
	"github.com/MarcGrol/microgen/lib/envelope"
	"github.com/MarcGrol/microgen/tourApp/events"
	"time"
)

// commands

type CreateNewsItemCommand struct {
	Year      int       `json:"year" binding:"required"`
	Timestamp time.Time `json:"timestamp" binding:"required"`
	Message   string    `json:"message" binding:"required"`
	Sender    string    `json:"sender" binding:"required"`
}

type CommandHandler interface {
	Start(listenPort int)

	HandleCreateNewsItemCommand(command *CreateNewsItemCommand) error

	HandleGetNewsQuery(year int) (*News, error)
}

// events

type EventHandler interface {
	Start()
	OnEnvelope(envelop *envelope.Envelope) error

	OnTourCreated(event *events.TourCreated) error
	OnEtappeCreated(event *events.EtappeCreated) error
	OnCyclistCreated(event *events.CyclistCreated) error
	OnEtappeResultsCreated(event *events.EtappeResultsCreated) error
}

type EventApplier interface {
	ApplyEtappeCreated(event *events.EtappeCreated)
	ApplyCyclistCreated(event *events.CyclistCreated)
	ApplyEtappeResultsCreated(event *events.EtappeResultsCreated)
	ApplyNewsItemCreated(event *events.NewsItemCreated)
	ApplyTourCreated(event *events.TourCreated)
}
