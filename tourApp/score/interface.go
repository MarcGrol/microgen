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

func (command CreateDayResultsCommand) BasicValidate() error {

	// command.Year int

	// command.EtappeId int

	// command.BestDayCyclistIds int

	// command.BestAllroundCyclistIds int

	// command.BestClimbCyclistIds int

	// command.BestSprintCyclistIds int

	return nil
}

type CommandHandler interface {
	HandleCreateDayResultsCommand(command CreateDayResultsCommand) *myerrors.Error

	HandleGetResultsQuery(gamblerUid string) (*Results, *myerrors.Error)
}

// events

type EventHandler interface {
	OnTourCreated(event events.TourCreated) *myerrors.Error
	OnEtappeCreated(event events.EtappeCreated) *myerrors.Error
	OnCyclistCreated(event events.CyclistCreated) *myerrors.Error
	OnGamblerCreated(event events.GamblerCreated) *myerrors.Error
	OnGamblerTeamCreated(event events.GamblerTeamCreated) *myerrors.Error
}

type EventApplier interface {
	ApplyCyclistCreated(event events.CyclistCreated) *myerrors.Error
	ApplyGamblerCreated(event events.GamblerCreated) *myerrors.Error
	ApplyGamblerTeamCreated(event events.GamblerTeamCreated) *myerrors.Error
	ApplyEtappeResultsAvailable(event events.EtappeResultsAvailable) *myerrors.Error
	ApplyCyclistScoreCalculated(event events.CyclistScoreCalculated) *myerrors.Error
	ApplyGamblerScoreCalculated(event events.GamblerScoreCalculated) *myerrors.Error
	ApplyTourCreated(event events.TourCreated) *myerrors.Error
	ApplyEtappeCreated(event events.EtappeCreated) *myerrors.Error
}
