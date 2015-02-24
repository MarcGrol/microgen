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

func (command CreateTourCommand) BasicValidate() error {

	// command.Year int

	return nil
}

type CreateCyclistCommand struct {
	Year int    `json:"year"`
	Id   int    `json:"id"`
	Name string `json:"name"`
	Team string `json:"team"`
}

func (command CreateCyclistCommand) BasicValidate() error {

	// command.Year int

	// command.Id int

	// command.Name string

	// command.Team string

	return nil
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
	ApplyTourCreated(event events.TourCreated) *myerrors.Error
	ApplyCyclistCreated(event events.CyclistCreated) *myerrors.Error
	ApplyEtappeCreated(event events.EtappeCreated) *myerrors.Error
}
