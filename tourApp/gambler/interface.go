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
	OnEnvelope(envelop *envelope.Envelope) error
}

type EventApplier interface {
	ApplyEtappeResultsCreated(event *events.EtappeResultsCreated)
	ApplyTourCreated(event *events.TourCreated)
	ApplyGamblerCreated(event *events.GamblerCreated)
	ApplyCyclistCreated(event *events.CyclistCreated)
	ApplyGamblerTeamCreated(event *events.GamblerTeamCreated)
	ApplyEtappeCreated(event *events.EtappeCreated)
}

func applyEvents(envelopes []envelope.Envelope, aggregate EventApplier) error {
	for _, envelop := range envelopes {
		switch envelop.EventTypeName {
		case "CyclistCreated":
			aggregate.ApplyCyclistCreated(events.UnWrapCyclistCreated(&envelop))
			break
		case "GamblerTeamCreated":
			aggregate.ApplyGamblerTeamCreated(events.UnWrapGamblerTeamCreated(&envelop))
			break
		case "EtappeCreated":
			aggregate.ApplyEtappeCreated(events.UnWrapEtappeCreated(&envelop))
			break
		case "EtappeResultsCreated":
			aggregate.ApplyEtappeResultsCreated(events.UnWrapEtappeResultsCreated(&envelop))
			break
		case "TourCreated":
			aggregate.ApplyTourCreated(events.UnWrapTourCreated(&envelop))
			break
		case "GamblerCreated":
			aggregate.ApplyGamblerCreated(events.UnWrapGamblerCreated(&envelop))
			break

		default:
			return fmt.Errorf("applyEvents: Unexpected event %s", envelop.EventTypeName)
		}
	}
	return nil
}
