package score

// Generated automatically by microgen: do not edit manually

import (
	"github.com/MarcGrol/microgen/myerrors"
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
	HandleCreateDayResultsCommand(command CreateDayResultsCommand) *myerrors.Error

	HandleGetResultsQuery(gamblerUid string) (*Results, *myerrors.Error)
}

// events

type EventHandler interface {
	OnGamblerTeamCreated(event events.GamblerTeamCreated) *myerrors.Error
	OnTourCreated(event events.TourCreated) *myerrors.Error
	OnEtappeCreated(event events.EtappeCreated) *myerrors.Error
	OnCyclistCreated(event events.CyclistCreated) *myerrors.Error
	OnGamblerCreated(event events.GamblerCreated) *myerrors.Error
}

type EventApplier interface {
	ApplyEtappeResultsAvailable(event events.EtappeResultsAvailable) *myerrors.Error
	ApplyCyclistScoreCalculated(event events.CyclistScoreCalculated) *myerrors.Error
	ApplyGamblerScoreCalculated(event events.GamblerScoreCalculated) *myerrors.Error
	ApplyTourCreated(event events.TourCreated) *myerrors.Error
	ApplyEtappeCreated(event events.EtappeCreated) *myerrors.Error
	ApplyCyclistCreated(event events.CyclistCreated) *myerrors.Error
	ApplyGamblerCreated(event events.GamblerCreated) *myerrors.Error
	ApplyGamblerTeamCreated(event events.GamblerTeamCreated) *myerrors.Error
}
