package news

// Generated automatically by microgen: do not edit manually

import (
	"fmt"
	"time"

	"github.com/MarcGrol/microgen/lib/envelope"
	"github.com/MarcGrol/microgen/tourApp/events"
)

// commands

type CreateNewsItemCommand struct {
	Year      int       `json:"year" binding:"required"`
	Timestamp time.Time `json:"timestamp" binding:"required"`
	Message   string    `json:"message" binding:"required"`
	Sender    string    `json:"sender" binding:"required"`
}

type CommandHandler interface {
	Start(listenPort int) error

	HandleCreateNewsItemCommand(command *CreateNewsItemCommand) error

	HandleGetNewsQuery(year int) (*News, error)
}

// events

type EventHandler interface {
	Start() error
	OnEvent(envelop *envelope.Envelope) error
}

type AggregateRoot interface {
	ApplyAll(envelopes []envelope.Envelope)
	ApplyTourCreated(event *events.TourCreated)
	ApplyEtappeCreated(event *events.EtappeCreated)
	ApplyCyclistCreated(event *events.CyclistCreated)
	ApplyEtappeResultsCreated(event *events.EtappeResultsCreated)
	ApplyNewsItemCreated(event *events.NewsItemCreated)
}

func applyEvent(envelop envelope.Envelope, aggregateRoot AggregateRoot) error {
	switch envelop.EventTypeName {
	case "EtappeCreated":
		aggregateRoot.ApplyEtappeCreated(events.UnWrapEtappeCreated(&envelop))
		break
	case "CyclistCreated":
		aggregateRoot.ApplyCyclistCreated(events.UnWrapCyclistCreated(&envelop))
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

	default:
		return fmt.Errorf("applyEvents: Unexpected event %s", envelop.EventTypeName)
	}
	return nil
}

func applyEvents(envelopes []envelope.Envelope, aggregateRoot AggregateRoot) error {
	var err error
	for _, envelop := range envelopes {
		err = applyEvent(envelop, aggregateRoot)
		if err != nil {
			break
		}
	}
	return err
}
