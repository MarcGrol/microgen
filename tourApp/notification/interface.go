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
	OnEvent(envelop *envelope.Envelope) error
}

type AggregateRoot interface {
	ApplyTourCreated(event *events.TourCreated)
	ApplyCyclistCreated(event *events.CyclistCreated)
	ApplyEtappeCreated(event *events.EtappeCreated)
	ApplyEtappeResultsCreated(event *events.EtappeResultsCreated)
	ApplyNewsItemCreated(event *events.NewsItemCreated)
}

func applyEvents(envelopes []envelope.Envelope, aggregateRoot AggregateRoot) error {
	for _, envelop := range envelopes {
		switch envelop.EventTypeName {
		case "EtappeCreated":
			aggregateRoot.ApplyEtappeCreated(events.UnWrapEtappeCreated(&envelop))
			break
		case "EtappeResultsCreated":
			aggregateRoot.ApplyEtappeResultsCreated(events.UnWrapEtappeResultsCreated(&envelop))
			break
		case "NewsItemCreated":
			aggregateRoot.ApplyNewsItemCreated(events.UnWrapNewsItemCreated(&envelop))
			break
		case "TourCreated":
			aggregateRoot.ApplyTourCreated(events.UnWrapTourCreated(&envelop))
			break
		case "CyclistCreated":
			aggregateRoot.ApplyCyclistCreated(events.UnWrapCyclistCreated(&envelop))
			break

		default:
			return fmt.Errorf("applyEvents: Unexpected event %s", envelop.EventTypeName)
		}
	}
	return nil
}
