package notification

// Generated automatically by microgen: do not edit manually

import (
	"fmt"

	"github.com/MarcGrol/microgen/lib/envelope"
	"github.com/MarcGrol/microgen/tourApp/events"
)

// commands

type SubscribeToNotificationsCommand struct {
	Email string `json:"email" binding:"required"`
}

type CommandHandler interface {
	Start(listenPort int) error

	HandleSubscribeToNotificationsCommand(command *SubscribeToNotificationsCommand) error
}

// events

type EventHandler interface {
	Start() error
	OnEnvelope(envelop *envelope.Envelope) error
}

type EventApplier interface {
	ApplyNewsItemCreated(event *events.NewsItemCreated)
	ApplyTourCreated(event *events.TourCreated)
	ApplyCyclistCreated(event *events.CyclistCreated)
	ApplyEtappeCreated(event *events.EtappeCreated)
	ApplyEtappeResultsCreated(event *events.EtappeResultsCreated)
}

func applyEvents(envelopes []envelope.Envelope, aggregate EventApplier) error {
	for _, envelop := range envelopes {
		switch envelop.EventTypeName {
		case "TourCreated":
			aggregate.ApplyTourCreated(events.UnWrapTourCreated(&envelop))
			break
		case "CyclistCreated":
			aggregate.ApplyCyclistCreated(events.UnWrapCyclistCreated(&envelop))
			break
		case "EtappeCreated":
			aggregate.ApplyEtappeCreated(events.UnWrapEtappeCreated(&envelop))
			break
		case "EtappeResultsCreated":
			aggregate.ApplyEtappeResultsCreated(events.UnWrapEtappeResultsCreated(&envelop))
			break
		case "NewsItemCreated":
			aggregate.ApplyNewsItemCreated(events.UnWrapNewsItemCreated(&envelop))
			break

		default:
			return fmt.Errorf("applyEvents: Unexpected event %s", envelop.EventTypeName)
		}
	}
	return nil
}
