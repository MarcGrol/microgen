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

	HandleGetGamblerQuery(gamblerUid string) (*Gambler, *myerrors.Error)
}

// events

type EventHandler interface {
	OnTourCreated(event events.TourCreated) *myerrors.Error
	OnCyclistCreated(event events.CyclistCreated) *myerrors.Error
}

type EventApplier interface {
	ApplyCyclistCreated(event events.CyclistCreated) *myerrors.Error
	ApplyGamblerTeamCreated(event events.GamblerTeamCreated) *myerrors.Error
	ApplyTourCreated(event events.TourCreated) *myerrors.Error
	ApplyGamblerCreated(event events.GamblerCreated) *myerrors.Error
}
