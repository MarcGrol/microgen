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
	Start(listenPort int) error

	HandleCreateNewsItemCommand(command *CreateNewsItemCommand) error

	HandleGetNewsQuery(year int) (*News, error)
}

// events

type EventHandler interface {
	Start() error
	OnEnvelope(envelop *envelope.Envelope) error

	OnCyclistCreated(event *events.CyclistCreated) error
	OnEtappeResultsCreated(event *events.EtappeResultsCreated) error
	OnTourCreated(event *events.TourCreated) error
	OnEtappeCreated(event *events.EtappeCreated) error
}

type EventApplier interface {
	ApplyEtappeResultsCreated(event *events.EtappeResultsCreated)
	ApplyNewsItemCreated(event *events.NewsItemCreated)
	ApplyTourCreated(event *events.TourCreated)
	ApplyEtappeCreated(event *events.EtappeCreated)
	ApplyCyclistCreated(event *events.CyclistCreated)
}
