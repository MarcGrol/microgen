package gambler

// Generated automatically by microgen: do not edit manually

import (
	"fmt"

	"github.com/MarcGrol/microgen/lib/envelope"
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
	Start(listenPort int) error

	HandleCreateGamblerCommand(command *CreateGamblerCommand) error

	HandleCreateGamblerTeamCommand(command *CreateGamblerTeamCommand) error

	HandleGetGamblerQuery(gamblerUid string, year int) (*Gambler, error)

	HandleGetResultsQuery(year int) (*Results, error)
}

// events

type EventHandler interface {
	Start() error
	OnEvent(envelop *envelope.Envelope) error
}

type AggregateRoot interface {
	ApplyTourCreated(event *events.TourCreated)
	ApplyGamblerCreated(event *events.GamblerCreated)
	ApplyCyclistCreated(event *events.CyclistCreated)
	ApplyGamblerTeamCreated(event *events.GamblerTeamCreated)
	ApplyEtappeCreated(event *events.EtappeCreated)
	ApplyEtappeResultsCreated(event *events.EtappeResultsCreated)
}

func applyEvents(envelopes []envelope.Envelope, aggregateRoot AggregateRoot) error {
	for _, envelop := range envelopes {
		switch envelop.EventTypeName {
		case "CyclistCreated":
			aggregateRoot.ApplyCyclistCreated(events.UnWrapCyclistCreated(&envelop))
			break
		case "GamblerTeamCreated":
			aggregateRoot.ApplyGamblerTeamCreated(events.UnWrapGamblerTeamCreated(&envelop))
			break
		case "EtappeCreated":
			aggregateRoot.ApplyEtappeCreated(events.UnWrapEtappeCreated(&envelop))
			break
		case "EtappeResultsCreated":
			aggregateRoot.ApplyEtappeResultsCreated(events.UnWrapEtappeResultsCreated(&envelop))
			break
		case "TourCreated":
			aggregateRoot.ApplyTourCreated(events.UnWrapTourCreated(&envelop))
			break
		case "GamblerCreated":
			aggregateRoot.ApplyGamblerCreated(events.UnWrapGamblerCreated(&envelop))
			break

		default:
			return fmt.Errorf("applyEvents: Unexpected event %s", envelop.EventTypeName)
		}
	}
	return nil
}
