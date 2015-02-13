package tour

// Generated automatically by microgen: do not edit manually

import (
	"github.com/xebia/microgen/events"
	"time"
)

// commands

type CreateTourCommand struct {
	Year int `json:"year"`
}

type CreateCyclistCommand struct {
	Year int    `json:"year"`
	Id   int    `json:"id"`
	Name string `json:"name"`
	Team string `json:"team"`
}

type CreateEtappeCommand struct {
	Year           int       `json:"year"`
	Id             int       `json:"id"`
	Date           time.Time `json:"date"`
	StartLocation  string    `json:"startLocation"`
	FinishLocation string    `json:"finishLocation"`
	Length         int       `json:"length"`
	Kind           int       `json:"kind"`
}

type CommandHandler interface {
	HandleCreateTourCommand(command CreateTourCommand) error

	HandleCreateCyclistCommand(command CreateCyclistCommand) error

	HandleCreateEtappeCommand(command CreateEtappeCommand) error
}

// events

type EventHandler interface {
}

type EventApplier interface {
	ApplyTourCreated(event events.TourCreated) error
	ApplyCyclistCreated(event events.CyclistCreated) error
	ApplyEtappeCreated(event events.EtappeCreated) error
}
