package tour

// Generated automatically by microgen: do not edit manually

import (
	"fmt"
	"time"

	"github.com/MarcGrol/microgen/lib/envelope"
	"github.com/MarcGrol/microgen/tourApp/events"
)

// commands

type CreateTourCommand struct {
	Year int `json:"year" binding:"required"`
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

type CreateEtappeResultsCommand struct {
	Year                   int   `json:"year"`
	EtappeId               int   `json:"etappeId"`
	BestDayCyclistIds      []int `json:"bestDayCyclistIds" `
	BestAllroundCyclistIds []int `json:"bestAllroundCyclistIds" `
	BestClimbCyclistIds    []int `json:"bestClimbCyclistIds" `
	BestSprintCyclistIds   []int `json:"bestSprintCyclistIds" `
}

type CommandHandler interface {
	Start(listenPort int) error

	HandleCreateTourCommand(command *CreateTourCommand) error

	HandleCreateCyclistCommand(command *CreateCyclistCommand) error

	HandleCreateEtappeCommand(command *CreateEtappeCommand) error

	HandleCreateEtappeResultsCommand(command *CreateEtappeResultsCommand) error

	HandleGetTourQuery(year int) (*Tour, error)
}

// events

type EventHandler interface {
	Start() error
}

type AggregateRoot interface {
	ApplyTourCreated(event *events.TourCreated)
	ApplyCyclistCreated(event *events.CyclistCreated)
	ApplyEtappeCreated(event *events.EtappeCreated)
	ApplyEtappeResultsCreated(event *events.EtappeResultsCreated)
}

func applyEvents(envelopes []envelope.Envelope, aggregateRoot AggregateRoot) error {
	for _, envelop := range envelopes {
		switch envelop.EventTypeName {
		case "TourCreated":
			aggregateRoot.ApplyTourCreated(events.UnWrapTourCreated(&envelop))
			break
		case "CyclistCreated":
			aggregateRoot.ApplyCyclistCreated(events.UnWrapCyclistCreated(&envelop))
			break
		case "EtappeCreated":
			aggregateRoot.ApplyEtappeCreated(events.UnWrapEtappeCreated(&envelop))
			break
		case "EtappeResultsCreated":
			aggregateRoot.ApplyEtappeResultsCreated(events.UnWrapEtappeResultsCreated(&envelop))
			break

		default:
			return fmt.Errorf("applyEvents: Unexpected event %s", envelop.EventTypeName)
		}
	}
	return nil
}
