package score

// Generated automatically by microgen: do not edit manually

import (
	"github.com/MarcGrol/microgen/tourApp/events"
)

// commands

type CreateDayResultsCommand struct {
	Year                   int   `json:"year" binding:"required"`
	EtappeId               int   `json:"etappeId" binding:"required"`
	BestDayCyclistIds      []int `json:"bestDayCyclistIds" `
	BestAllroundCyclistIds []int `json:"bestAllroundCyclistIds" `
	BestClimbCyclistIds    []int `json:"bestClimbCyclistIds" `
	BestSprintCyclistIds   []int `json:"bestSprintCyclistIds" `
}

type CommandHandler interface {
	HandleCreateDayResultsCommand(command *CreateDayResultsCommand) error

	HandleGetResultsQuery(gamblerUid string) (*Results, error)
}

// events

type EventHandler interface {
	OnTourCreated(event *events.TourCreated) error
	OnEtappeCreated(event *events.EtappeCreated) error
	OnCyclistCreated(event *events.CyclistCreated) error
	OnGamblerCreated(event *events.GamblerCreated) error
	OnGamblerTeamCreated(event *events.GamblerTeamCreated) error
}

type EventApplier interface {
	ApplyEtappeCreated(event *events.EtappeCreated)
	ApplyCyclistCreated(event *events.CyclistCreated)
	ApplyGamblerCreated(event *events.GamblerCreated)
	ApplyGamblerTeamCreated(event *events.GamblerTeamCreated)
	ApplyEtappeResultsAvailable(event *events.EtappeResultsAvailable)
	ApplyCyclistScoreCalculated(event *events.CyclistScoreCalculated)
	ApplyGamblerScoreCalculated(event *events.GamblerScoreCalculated)
	ApplyTourCreated(event *events.TourCreated)
}
