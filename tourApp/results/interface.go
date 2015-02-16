package results

// Generated automatically by microgen: do not edit manually

import (
	"github.com/xebia/microgen/tourApp/events"
)

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
	HandleCreateDayResultsCommand(command CreateDayResultsCommand) error
}

// events

type EventHandler interface {
	OnCyclistCreated(event events.CyclistCreated) error
	OnGamblerCreated(event events.GamblerCreated) error
	OnGamblerTeamCreated(event events.GamblerTeamCreated) error
	OnTourCreated(event events.TourCreated) error
	OnEtappeCreated(event events.EtappeCreated) error
}

type EventApplier interface {
	ApplyGamblerScoreCalculated(event events.GamblerScoreCalculated) error
	ApplyTourCreated(event events.TourCreated) error
	ApplyEtappeCreated(event events.EtappeCreated) error
	ApplyCyclistCreated(event events.CyclistCreated) error
	ApplyGamblerCreated(event events.GamblerCreated) error
	ApplyGamblerTeamCreated(event events.GamblerTeamCreated) error
	ApplyEtappeResultsAvailable(event events.EtappeResultsAvailable) error
	ApplyCyclistScoreCalculated(event events.CyclistScoreCalculated) error
}
