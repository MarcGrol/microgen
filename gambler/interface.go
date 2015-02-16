package gambler

// Generated automatically by microgen: do not edit manually

import (
	"github.com/xebia/microgen/events"
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
}

type EventApplier interface {
	ApplyTourCreated(event events.TourCreated) error
	ApplyGamblerTeamCreated(event events.GamblerTeamCreated) error
	ApplyGamblerCreated(event events.GamblerCreated) error
}
