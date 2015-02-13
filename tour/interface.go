package tour

// Generated automatically by microgen: do not edit manually

import (
	"github.com/xebia/microgen/events"
	"time"
)

func StartApplication(bus events.PublishSubscriber, store events.Store, commandHandler CommandHandler, eventHandler EventHandler, model interface{}) error {
	return nil
}

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
	HandleCreateTourCommand(command CreateTourCommand) ([]*events.Envelope, error)

	HandleCreateCyclistCommand(command CreateCyclistCommand) ([]*events.Envelope, error)

	HandleCreateEtappeCommand(command CreateEtappeCommand) ([]*events.Envelope, error)
}

// events

type EventHandler interface {
}

type EventApplier interface {
	ApplyTourCreated(event events.TourCreated) error
	ApplyCyclistCreated(event events.CyclistCreated) error
	ApplyEtappeCreated(event events.EtappeCreated) error
}