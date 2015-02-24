package tour

// Generated automatically by microgen: do not edit manually

import (
	"github.com/MarcGrol/microgen/myerrors"
	"github.com/MarcGrol/microgen/tourApp/events"
	"time"
)

// commands

type CreateTourCommand struct {
	Year int `json:"year" binding:"required"`
}

func (command CreateTourCommand) BasicValidate() error {

	// command.Year int

	return nil
}

type CreateCyclistCommand struct {
	Year int    `json:"year" binding:"required"`
	Id   int    `json:"id" binding:"required"`
	Name string `json:"name" binding:"required"`
	Team string `json:"team" binding:"required"`
}

func (command CreateCyclistCommand) BasicValidate() error {

	// command.Year int

	// command.Id int

	// command.Name string

	// command.Team string

	return nil
}

type CreateEtappeCommand struct {
	Year           int       `json:"year" binding:"required"`
	Id             int       `json:"id" binding:"required"`
	Date           time.Time `json:"date" binding:"required"`
	StartLocation  string    `json:"startLocation" binding:"required"`
	FinishLocation string    `json:"finishLocation" binding:"required"`
	Length         int       `json:"length" binding:"required"`
	Kind           int       `json:"kind" binding:"required"`
}

func (command CreateEtappeCommand) BasicValidate() error {

	// command.Year int

	// command.Id int

	// command.Date time.Time

	// command.StartLocation string

	// command.FinishLocation string

	// command.Length int

	// command.Kind int

	return nil
}

type CommandHandler interface {
	HandleCreateTourCommand(command CreateTourCommand) *myerrors.Error

	HandleCreateCyclistCommand(command CreateCyclistCommand) *myerrors.Error

	HandleCreateEtappeCommand(command CreateEtappeCommand) *myerrors.Error

	HandleGetTourQuery(year int) (*Tour, *myerrors.Error)
}

// events

type EventHandler interface {
}

type EventApplier interface {
	ApplyCyclistCreated(event events.CyclistCreated) *myerrors.Error
	ApplyEtappeCreated(event events.EtappeCreated) *myerrors.Error
	ApplyTourCreated(event events.TourCreated) *myerrors.Error
}
