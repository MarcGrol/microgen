package results

// Generated automatically by microgen: do not edit manually

import (
	"github.com/xebia/microgen/events"
)

func StartApplication(bus events.PublishSubscriber, store events.Store, commandHandler CommandHandler, eventHandler EventHandler, model interface{}) error {
	return nil
}

// commands

type CreateDayResultsCommand struct {
	Year                   int   `json:"year"`
	Id                     int   `json:"id"`
	BestDayCyclistIds      []int `json:"bestDayCyclistIds"`
	BestAllroundCyclistIds []int `json:"bestAllroundCyclistIds"`
	BestClimbCyclistIds    []int `json:"bestClimbCyclistIds"`
	BestSprintCyclistIds   []int `json:"bestSprintCyclistIds"`
}

type CommandHandler interface {
	HandleCreateDayResultsCommand(command CreateDayResultsCommand) ([]*events.Envelope, error)
}

// events

type EventHandler interface {
	OnTourCreated(event events.TourCreated) ([]*events.Envelope, error)
	OnEtappeCreated(event events.EtappeCreated) ([]*events.Envelope, error)
	OnCyclistCreated(event events.CyclistCreated) ([]*events.Envelope, error)
	OnGamblerCreated(event events.GamblerCreated) ([]*events.Envelope, error)
	OnGamblerTeamCreated(event events.GamblerTeamCreated) ([]*events.Envelope, error)
}

type EventApplier interface {
	ApplyGamblerTeamCreated(event events.GamblerTeamCreated) error
	ApplyEtappeResultsAvailable(event events.EtappeResultsAvailable) error
	ApplyCyclistScoreCalculated(event events.CyclistScoreCalculated) error
	ApplyGamblerScoreCalculated(event events.GamblerScoreCalculated) error
	ApplyTourCreated(event events.TourCreated) error
	ApplyEtappeCreated(event events.EtappeCreated) error
	ApplyCyclistCreated(event events.CyclistCreated) error
	ApplyGamblerCreated(event events.GamblerCreated) error
}
