package gambler

import (
	"errors"
	"fmt"
	"github.com/MarcGrol/microgen/myerrors"
	"github.com/MarcGrol/microgen/tourApp/events"
	"log"
	"strconv"
)

type GamblerEventHandler struct {
	store events.Store
}

func NewGamblerEventHandler(store events.Store) EventHandler {
	handler := new(GamblerEventHandler)
	handler.store = store
	return handler
}

func (eh *GamblerEventHandler) OnTourCreated(event events.TourCreated) error {
	log.Printf("OnTourCreated: event: %+v", event)
	return doStore(eh.store, []*events.Envelope{event.Wrap()})
}

func (eh *GamblerEventHandler) OnCyclistCreated(event events.CyclistCreated) error {

	log.Printf("OnCyclistCreated: event: %+v", event)
	return doStore(eh.store, []*events.Envelope{event.Wrap()})
}

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

func (ch *GamblerCommandHandler) validateCreateGamblerCommand(command CreateGamblerCommand) error {
	// TODO
	return nil
}

func (ch *GamblerCommandHandler) HandleCreateGamblerCommand(command CreateGamblerCommand) error {
	err := ch.validateCreateGamblerCommand(command)
	if err != nil {
		return myerrors.NewInvalidInputError(err)
	}
	gamblerContext, err := getGamblerContext(ch.store, command.GamblerUid, -1)
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

	// store and emit resulting event
	return doStoreAndPublish(ch.store, ch.bus, []*events.Envelope{gamblerCreatedEvent.Wrap()})
}

func (ch *GamblerCommandHandler) validateCreateGamblerTeamCommand(command CreateGamblerTeamCommand) error {
	// TODO
	return nil
}

func (ch *GamblerCommandHandler) HandleCreateGamblerTeamCommand(command CreateGamblerTeamCommand) error {
	err := ch.validateCreateGamblerTeamCommand(command)
	if err != nil {
		return myerrors.NewInvalidInputError(err)
	}
	gamblerContext, err := getGamblerContext(ch.store, command.GamblerUid, command.Year)
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

	return doStoreAndPublish(ch.store, ch.bus, []*events.Envelope{gamblerTeamCreatedEvent.Wrap()})
}

func doStore(store events.Store, envelopes []*events.Envelope) error {
	for _, env := range envelopes {
		err := store.Store(env)
		if err != nil {
			log.Printf("Error storing event: %+v", err)
			return err
		}
		log.Printf("Successfully stored event: %+v", env)
	}
	return nil
}

func doStoreAndPublish(store events.Store, bus events.PublishSubscriber, envelopes []*events.Envelope) *myerrors.Error {
	err := doStore(store, envelopes)
	if err != nil {
		return myerrors.NewInternalError(err)
	}
	for _, env := range envelopes {
		err = bus.Publish(env)
		if err != nil {
			return myerrors.NewInternalError(err)
		}
	}
	return nil
}

func (ch *GamblerCommandHandler) HandleGetGamblerQuery(gamblerUid string, year int) (*Gambler, error) {
	// TODO validate input
	gamblerContext, err := getGamblerContext(ch.store, gamblerUid, year)
	if err != nil {
		return nil, myerrors.NewInternalError(err)
	}
	if gamblerContext.Gambler == nil {
		return nil, myerrors.NewNotFoundError(errors.New(fmt.Sprintf("Gambler with uid %s not found", gamblerUid)))
	}

	//log.Printf("HandleGetGamblerQuery.Gambler:%+v", gamblerContext.Gambler)

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

func (context *GamblerContext) ApplyTourCreated(event events.TourCreated) {
	//log.Printf("ApplyTourCreated: context before: %+v, event: %+v", context, event)

	context.Year = new(int)
	*context.Year = event.Year

	//log.Printf("ApplyTourCreated: context after: %+v", context)

	return
}

func (context *GamblerContext) ApplyCyclistCreated(event events.CyclistCreated) {
	//log.Printf("ApplyCyclistCreated: context before: %+v, event: %+v", context, event)

	context.cyclistsForTour[event.CyclistId] =
		Cyclist{Id: event.CyclistId, Name: event.CyclistName, Team: event.CyclistTeam}

	//log.Printf("ApplyCyclistCreated: context after: %+v", context)

	return
}

func (context *GamblerContext) ApplyGamblerCreated(event events.GamblerCreated) {
	//log.Printf("ApplyGamblerCreated: context before: %+v, event: %+v", context, event)

	context.Gambler = NewGambler(event.GamblerUid, event.GamblerName, event.GamblerEmail)

	//log.Printf("ApplyGamblerCreated: context after: %+v", context)

	return
}

func (context *GamblerContext) ApplyGamblerTeamCreated(event events.GamblerTeamCreated) {
	//log.Printf("ApplyGamblerTeamCreated: context: %+v, event: %+v", context, event)

	for _, cyclistId := range event.GamblerCyclists {
		cyclist, found := context.cyclistsForTour[cyclistId]
		if found {
			context.Gambler.Cyclists = append(context.Gambler.Cyclists, cyclist)
		}
	}

	//log.Printf("ApplyGamblerTeamCreated: context after: %+v", context)
	return
}
