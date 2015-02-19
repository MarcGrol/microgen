package gambler

// Generated automatically by microgen: do not edit manually

import (
	"github.com/MarcGrol/microgen/myerrors"
	"github.com/MarcGrol/microgen/tourApp/events"
)

// commands

type CreateGamblerCommand struct {
	GamblerUid string `json:"gamblerUid"`
	Name       string `json:"name"`
	Email      string `json:"email"`
}

type CreateGamblerTeamCommand struct {
	GamblerUid string `json:"gamblerUid"`
	Year       int    `json:"year"`
	CyclistIds []int  `json:"cyclistIds"`
}

type CommandHandler interface {
	HandleCreateGamblerCommand(command CreateGamblerCommand) *myerrors.Error

	HandleCreateGamblerTeamCommand(command CreateGamblerTeamCommand) *myerrors.Error
}

// events

type EventHandler interface {
	OnCyclistCreated(event events.CyclistCreated) *myerrors.Error
	OnTourCreated(event events.TourCreated) *myerrors.Error
}

type EventApplier interface {
	ApplyGamblerTeamCreated(event events.GamblerTeamCreated) *myerrors.Error
	ApplyTourCreated(event events.TourCreated) *myerrors.Error
	ApplyGamblerCreated(event events.GamblerCreated) *myerrors.Error
	ApplyCyclistCreated(event events.CyclistCreated) *myerrors.Error
}
