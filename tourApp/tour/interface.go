package tour

// Generated automatically by microgen: do not edit manually

import (
	"fmt"
	"github.com/MarcGrol/microgen/lib/envelope"
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
	Year           int       `json:"year" binding:"required"`
	Id             int       `json:"id" binding:"required"`
	Date           time.Time `json:"date" binding:"required"`
	StartLocation  string    `json:"startLocation" binding:"required"`
	FinishLocation string    `json:"finishLocation" binding:"required"`
	Length         int       `json:"length" binding:"required"`
	Kind           int       `json:"kind" binding:"required"`
}

type CreateEtappeResultsCommand struct {
	Year                   int   `json:"year" binding:"required"`
	EtappeId               int   `json:"etappeId" binding:"required"`
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

type EventApplier interface {
	ApplyTourCreated(event *events.TourCreated)
	ApplyCyclistCreated(event *events.CyclistCreated)
	ApplyEtappeCreated(event *events.EtappeCreated)
	ApplyEtappeResultsCreated(event *events.EtappeResultsCreated)
}

func applyEvents(envelopes []envelope.Envelope, aggregate EventApplier) error {
	for _, envelop := range envelopes {
		switch envelop.EventTypeName {
		case "EtappeResultsCreated":
			aggregate.ApplyEtappeResultsCreated(events.UnWrapEtappeResultsCreated(&envelop))
			break
		case "TourCreated":
			aggregate.ApplyTourCreated(events.UnWrapTourCreated(&envelop))
			break
		case "CyclistCreated":
			aggregate.ApplyCyclistCreated(events.UnWrapCyclistCreated(&envelop))
			break
		case "EtappeCreated":
			aggregate.ApplyEtappeCreated(events.UnWrapEtappeCreated(&envelop))
			break

		default:
			return fmt.Errorf("applyEvents: Unexpected event %s", envelop.EventTypeName)
		}
	}
	return nil
}
