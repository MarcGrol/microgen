package gambler

// Generated automatically by microgen: do not edit manually

import (
	"github.com/MarcGrol/microgen/tourApp/events"
)

// commands

type CreateGamblerCommand struct {
	GamblerUid string `json:"gamblerUid" binding:"required"`
	Name       string `json:"name" binding:"required"`
	Email      string `json:"email" binding:"required"`
}

type CreateGamblerTeamCommand struct {
	GamblerUid string `json:"gamblerUid" binding:"required"`
	Year       int    `json:"year" binding:"required"`
	CyclistIds []int  `json:"cyclistIds" `
}

type CommandHandler interface {
	Start(listenPort int)

	HandleCreateGamblerCommand(command *CreateGamblerCommand) error

	HandleCreateGamblerTeamCommand(command *CreateGamblerTeamCommand) error

	HandleGetGamblerQuery(gamblerUid string, year int) (*Gambler, error)
}

// events

type EventHandler interface {
	Start()
	OnTourCreated(event *events.TourCreated) error
	OnCyclistCreated(event *events.CyclistCreated) error
}

type EventApplier interface {
	ApplyCyclistCreated(event *events.CyclistCreated)
	ApplyGamblerTeamCreated(event *events.GamblerTeamCreated)
	ApplyTourCreated(event *events.TourCreated)
	ApplyGamblerCreated(event *events.GamblerCreated)
}
