package gambler

import (
	"errors"
	"fmt"
	"github.com/MarcGrol/microgen/myerrors"
	"github.com/MarcGrol/microgen/tourApp/events"
	"log"
	"strconv"
)

type GamblerCommandHandler struct {
	bus   events.PublishSubscriber
	store events.Store
}

func NewGamblerCommandHandler(bus events.PublishSubscriber, store events.Store) CommandHandler {
	handler := new(GamblerCommandHandler)
	handler.bus = bus
	handler.store = store
	return handler
}

func (gch *GamblerCommandHandler) HandleCreateGamblerCommand(command CreateGamblerCommand) *myerrors.Error {
	gamblerContext, err := getGamblerContext(gch.store, command.GamblerUid, -1)
	if err != nil {
		return myerrors.NewInternalError(err)
	}

	if gamblerContext.Gambler != nil {
		return myerrors.NewInvalidInputError(errors.New(fmt.Sprintf("gambler %s already exists", command.GamblerUid)))
	}

	// apply business logic
	gamblerCreatedEvent := events.GamblerCreated{
		GamblerUid:   command.GamblerUid,
		GamblerName:  command.Name,
		GamblerEmail: command.Email}
	gamblerContext.ApplyGamblerCreated(gamblerCreatedEvent)

	// store and emit resulting event
	return gch.storeAndPublish([]*events.Envelope{gamblerCreatedEvent.Wrap()})
}

func (gch *GamblerCommandHandler) HandleCreateGamblerTeamCommand(command CreateGamblerTeamCommand) *myerrors.Error {
	gamblerContext, err := getGamblerContext(gch.store, command.GamblerUid, command.Year)
	if err != nil {
		return myerrors.NewInternalError(err)
	}
	if gamblerContext.Year == nil || *gamblerContext.Year != command.Year {
		return myerrors.NewNotFoundError(errors.New(fmt.Sprintf("Tour %d not found", command.Year)))
	}
	if gamblerContext.Gambler == nil {
		return myerrors.NewNotFoundError(errors.New(fmt.Sprintf("Gambler %s not found", command.GamblerUid)))
	}

	// apply business logic
	gamblerTeamCreatedEvent := events.GamblerTeamCreated{
		GamblerUid:      command.GamblerUid,
		Year:            command.Year,
		GamblerCyclists: command.CyclistIds}
	gamblerContext.ApplyGamblerTeamCreated(gamblerTeamCreatedEvent)

	return gch.storeAndPublish([]*events.Envelope{gamblerTeamCreatedEvent.Wrap()})
}

func (gch *GamblerCommandHandler) storeAndPublish(envelopes []*events.Envelope) *myerrors.Error {
	for _, env := range envelopes {
		err := gch.store.Store(env)
		if err != nil {
			return myerrors.NewInternalError(err)
		}
		err = gch.bus.Publish(env)
		if err != nil {
			return myerrors.NewInternalError(err)
		}
	}
	return nil
}

func (gch *GamblerCommandHandler) HandleGetGamblerQuery(gamblerUid string, year int) (*Gambler, *myerrors.Error) {
	// TODO validate input
	gamblerContext, err := getGamblerContext(gch.store, gamblerUid, year)
	if err != nil {
		return nil, myerrors.NewInternalError(err)
	}
	if gamblerContext.Gambler == nil {
		return nil, myerrors.NewNotFoundError(errors.New(fmt.Sprintf("Gambler with uid %s not found", gamblerUid)))
	}

	log.Printf("HandleGetGamblerQuery.Gambler:%+v", gamblerContext.Gambler)

	return gamblerContext.Gambler, nil
}

func getGamblerContext(store events.Store, gamblerUid string, year int) (*GamblerContext, error) {
	context := NewGamblerContext()

	tourRelatedEvents, err := store.Get("tour", strconv.Itoa(year))
	if err != nil {
		return context, err
	}

	gamblerRelatedEvents, err := store.Get("gambler", gamblerUid)
	if err != nil {
		return context, err
	}

	for _, envelope := range tourRelatedEvents {
		if envelope.Type == events.TypeTourCreated {
			context.ApplyTourCreated(*envelope.TourCreated)
		} else if envelope.Type == events.TypeCyclistCreated {
			context.ApplyCyclistCreated(*envelope.CyclistCreated)
		} else {
			log.Panicf("getGamblerOnUid(tour): Unexpected event %s", envelope.Type.String())
		}
	}

	for _, envelope := range gamblerRelatedEvents {
		if envelope.Type == events.TypeGamblerCreated {
			context.ApplyGamblerCreated(*envelope.GamblerCreated)
		} else if envelope.Type == events.TypeGamblerTeamCreated {
			context.ApplyGamblerTeamCreated(*envelope.GamblerTeamCreated)
		} else {
			log.Panicf("getGamblerOnUid(gambler): Unexpected event %s", envelope.Type.String())
		}
	}

	return context, nil
}

type GamblerContext struct {
	Year            *int
	cyclistsForTour map[int]Cyclist
	Gambler         *Gambler
}

func NewGamblerContext() *GamblerContext {
	context := new(GamblerContext)
	context.Year = nil
	context.cyclistsForTour = make(map[int]Cyclist)
	context.Gambler = nil
	return context
}

type Gambler struct {
	Uid      string
	Name     string
	Email    string
	Cyclists []Cyclist
}

func NewGambler(uid string, name string, email string) *Gambler {
	gambler := new(Gambler)
	gambler.Uid = uid
	gambler.Name = name
	gambler.Email = email
	gambler.Cyclists = make([]Cyclist, 0, 10)
	return gambler
}

type Cyclist struct {
	Id   int
	Name string
	Team string
}

func (context *GamblerContext) ApplyTourCreated(event events.TourCreated) *myerrors.Error {
	log.Printf("ApplyTourCreated: context before: %+v, event: %+v", context, event)

	context.Year = new(int)
	*context.Year = event.Year

	log.Printf("ApplyTourCreated: context after: %+v", context)

	return nil
}

func (context *GamblerContext) ApplyCyclistCreated(event events.CyclistCreated) *myerrors.Error {
	log.Printf("ApplyCyclistCreated: context before: %+v, event: %+v", context, event)

	if context.Year == nil || *context.Year != event.Year {
		return myerrors.NewNotFoundError(errors.New(fmt.Sprintf("Tour %d not found", event.Year)))
	}

	context.cyclistsForTour[event.CyclistId] =
		Cyclist{Id: event.CyclistId, Name: event.CyclistName, Team: event.CyclistTeam}

	log.Printf("ApplyCyclistCreated: context after: %+v", context)

	return nil
}

func (context *GamblerContext) ApplyGamblerCreated(event events.GamblerCreated) *myerrors.Error {
	log.Printf("ApplyGamblerCreated: context before: %+v, event: %+v", context, event)

	context.Gambler = NewGambler(event.GamblerUid, event.GamblerName, event.GamblerEmail)

	log.Printf("ApplyGamblerCreated: context after: %+v", context)

	return nil
}

func (context *GamblerContext) ApplyGamblerTeamCreated(event events.GamblerTeamCreated) *myerrors.Error {
	log.Printf("ApplyGamblerTeamCreated: context: %+v, event: %+v", context, event)

	for _, cyclistId := range event.GamblerCyclists {
		cyclist, found := context.cyclistsForTour[cyclistId]
		if found {
			context.Gambler.Cyclists = append(context.Gambler.Cyclists, cyclist)
		}
	}

	log.Printf("ApplyGamblerTeamCreated: context after: %+v", context)
	return nil
}
