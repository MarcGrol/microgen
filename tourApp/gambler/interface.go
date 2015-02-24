package gambler

// Generated automatically by microgen: do not edit manually

import (
	"github.com/MarcGrol/microgen/myerrors"
	"github.com/MarcGrol/microgen/tourApp/events"
)

// commands

type CreateGamblerCommand struct {
	GamblerUid string `json:"gamblerUid" binding:"required"`
	Name       string `json:"name" binding:"required"`
	Email      string `json:"email" binding:"required"`
}

func (command CreateGamblerCommand) BasicValidate() error {

	// command.GamblerUid string

	// command.Name string

	// command.Email string

	return nil
}

type CreateGamblerTeamCommand struct {
	GamblerUid string `json:"gamblerUid" binding:"required"`
	Year       int    `json:"year" binding:"required"`
	CyclistIds []int  `json:"cyclistIds" `
}

func (command CreateGamblerTeamCommand) BasicValidate() error {

	// command.GamblerUid string

	// command.Year int

	// command.CyclistIds int

	return nil
}

type CommandHandler interface {
	HandleCreateGamblerCommand(command CreateGamblerCommand) *myerrors.Error

	HandleCreateGamblerTeamCommand(command CreateGamblerTeamCommand) *myerrors.Error

	HandleGetGamblerQuery(gamblerUid string, year int) (*Gambler, *myerrors.Error)
}

// events

type EventHandler interface {
	OnCyclistCreated(event events.CyclistCreated) *myerrors.Error
	OnTourCreated(event events.TourCreated) *myerrors.Error
}

type EventApplier interface {
	ApplyGamblerCreated(event events.GamblerCreated) *myerrors.Error
	ApplyCyclistCreated(event events.CyclistCreated) *myerrors.Error
	ApplyGamblerTeamCreated(event events.GamblerTeamCreated) *myerrors.Error
	ApplyTourCreated(event events.TourCreated) *myerrors.Error
}
