package score

// Generated automatically by microgen: do not edit manually

import (
	"github.com/MarcGrol/microgen/tourApp/events"
)

// commands

type CommandHandler interface {
	Start(listenPort int)

	HandleGetResultsQuery(gamblerUid string) (*Results, error)
}

// events

type EventHandler interface {
	Start()
	OnCyclistCreated(event *events.CyclistCreated) error
	OnEtappeCreated(event *events.EtappeCreated) error
	OnGamblerCreated(event *events.GamblerCreated) error
	OnGamblerTeamCreated(event *events.GamblerTeamCreated) error
	OnTourCreated(event *events.TourCreated) error
}

type EventApplier interface {
	ApplyTourCreated(event *events.TourCreated)
	ApplyCyclistCreated(event *events.CyclistCreated)
	ApplyEtappeCreated(event *events.EtappeCreated)
	ApplyGamblerCreated(event *events.GamblerCreated)
	ApplyGamblerTeamCreated(event *events.GamblerTeamCreated)
}
