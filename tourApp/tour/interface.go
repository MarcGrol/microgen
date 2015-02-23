package tour

// Generated automatically by microgen: do not edit manually

import (
	"github.com/MarcGrol/microgen/myerrors"
	"github.com/MarcGrol/microgen/tourApp/events"
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
	HandleCreateTourCommand(command CreateTourCommand) *myerrors.Error

	HandleCreateCyclistCommand(command CreateCyclistCommand) *myerrors.Error

	HandleCreateEtappeCommand(command CreateEtappeCommand) *myerrors.Error

	HandleGetTourQuery(year int) (interface{}, *myerrors.Error)
}

// events

type EventHandler interface {
}

type EventApplier interface {
	ApplyTourCreated(event events.TourCreated) *myerrors.Error
	ApplyCyclistCreated(event events.CyclistCreated) *myerrors.Error
	ApplyEtappeCreated(event events.EtappeCreated) *myerrors.Error
}
