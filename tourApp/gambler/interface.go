package gambler

// Generated automatically by microgen: do not edit manually

import (
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
	HandleCreateGamblerCommand(command CreateGamblerCommand) error

	HandleCreateGamblerTeamCommand(command CreateGamblerTeamCommand) error
}

// events

type EventHandler interface {
	OnTourCreated(event events.TourCreated) error
	OnCyclistCreated(event events.CyclistCreated) error
}

type EventApplier interface {
	ApplyGamblerTeamCreated(event events.GamblerTeamCreated) error
	ApplyTourCreated(event events.TourCreated) error
	ApplyGamblerCreated(event events.GamblerCreated) error
	ApplyCyclistCreated(event events.CyclistCreated) error
}
