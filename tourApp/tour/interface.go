package tour

// Generated automatically by microgen: do not edit manually

import (
	"github.com/MarcGrol/microgen/tourApp/events"
	"time"
)

// commands

type CreateTourCommand struct {
	Year int `json:"year" binding:"required"`
}

type CreateCyclistCommand struct {
	Year int    `json:"year" binding:"required"`
	Id   int    `json:"id" binding:"required"`
	Name string `json:"name" binding:"required"`
	Team string `json:"team" binding:"required"`
}

type CreateEtappeCommand struct {
	Year           int       `json:"year"`
	Id             int       `json:"id"`
	Date           time.Time `json:"thedate"`
	StartLocation  string    `json:"startLocation"`
	FinishLocation string    `json:"finishLocation"`
	Length         int       `json:"length"`
	Kind           int       `json:"kind"`
}

type CommandHandler interface {
	Start(listenPort int)

	HandleCreateTourCommand(command *CreateTourCommand) error

	HandleCreateCyclistCommand(command *CreateCyclistCommand) error

	HandleCreateEtappeCommand(command *CreateEtappeCommand) error

	HandleGetTourQuery(year int) (*Tour, error)
}

// events

type EventHandler interface {
	Start()
}

type EventApplier interface {
	ApplyTourCreated(event *events.TourCreated)
	ApplyCyclistCreated(event *events.CyclistCreated)
	ApplyEtappeCreated(event *events.EtappeCreated)
}
