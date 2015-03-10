package gambler

import (
	"errors"
	"fmt"
	"github.com/MarcGrol/microgen/infra"
	"github.com/MarcGrol/microgen/lib/envelope"
	"github.com/MarcGrol/microgen/lib/myerrors"
	"github.com/MarcGrol/microgen/tourApp/events"
	"log"
	"strconv"
)

type GamblerEventHandler struct {
	bus   infra.PublishSubscriber
	store infra.Store
}

func NewGamblerEventHandler(bus infra.PublishSubscriber, store infra.Store) *GamblerEventHandler {
	handler := new(GamblerEventHandler)
	handler.bus = bus
	handler.store = store
	return handler
}

func (eventHandler *GamblerEventHandler) Start() error {
	for _, eventType := range events.GetTourEventTypes() {
		eventHandler.bus.Subscribe(eventType.String(), func(envelope *envelope.Envelope) error {
			return eventHandler.OnEnvelope(envelope)
		})
	}
	/*

		{
			var topic events.Type = events.TypeTourCreated
			eventHandler.bus.Subscribe(topic.String(), func(envelop *envelope.Envelope) error {
				event := events.UnWrapTourCreated(envelop)
				return eventHandler.OnTourCreated(event)
			})
		}
		{
			var topic events.Type = events.TypeCyclistCreated
			eventHandler.bus.Subscribe(topic.String(), func(envelop *envelope.Envelope) error {
				event := events.UnWrapCyclistCreated(envelop)
				return eventHandler.OnCyclistCreated(event)
			})
		}
		{
			var topic events.Type = events.TypeEtappeCreated
			eventHandler.bus.Subscribe(topic.String(), func(envelop *envelope.Envelope) error {
				event := events.UnWrapEtappeCreated(envelop)
				return eventHandler.OnEtappeCreated(event)
			})
		}
		{
			var topic events.Type = events.TypeEtappeResultsCreated
			eventHandler.bus.Subscribe(topic.String(), func(envelop *envelope.Envelope) error {
				event := events.UnWrapEtappeResultsCreated(envelop)
				return eventHandler.OnEtappeResultsCreated(event)
			})
		}
	*/
	return nil
}

func (eh *GamblerEventHandler) OnEnvelope(envelop *envelope.Envelope) error {
	return doStore(eh.store, []*envelope.Envelope{envelop})
}

func (eh *GamblerEventHandler) OnTourCreated(event *events.TourCreated) error {
	log.Printf("OnTourCreated: event: %+v", event)
	return nil
}

func (eh *GamblerEventHandler) OnCyclistCreated(event *events.CyclistCreated) error {

	log.Printf("OnCyclistCreated: event: %+v", event)
	return nil
}

func (eh *GamblerEventHandler) OnEtappeCreated(event *events.EtappeCreated) error {
	log.Printf("OnEtappeCreated: event: %+v", event)
	return nil
}

func (eh *GamblerEventHandler) OnEtappeResultsCreated(event *events.EtappeResultsCreated) error {
	log.Printf("OnEtappeResultsCreated: event: %+v", event)
	return nil
}

type GamblerCommandHandler struct {
	bus   infra.PublishSubscriber
	store infra.Store
}

func NewGamblerCommandHandler(bus infra.PublishSubscriber, store infra.Store) *GamblerCommandHandler {
	handler := new(GamblerCommandHandler)
	handler.bus = bus
	handler.store = store
	return handler
}

func (ch *GamblerCommandHandler) validateCreateGamblerCommand(command *CreateGamblerCommand) error {
	// TODO
	return nil
}

func (ch *GamblerCommandHandler) HandleCreateGamblerCommand(command *CreateGamblerCommand) error {
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
	return doStoreAndPublish(ch.store, ch.bus, []*envelope.Envelope{gamblerCreatedEvent.Wrap()})
}

func (ch *GamblerCommandHandler) validateCreateGamblerTeamCommand(command *CreateGamblerTeamCommand) error {
	// TODO
	return nil
}

func (ch *GamblerCommandHandler) HandleCreateGamblerTeamCommand(command *CreateGamblerTeamCommand) error {
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

	return doStoreAndPublish(ch.store, ch.bus, []*envelope.Envelope{gamblerTeamCreatedEvent.Wrap()})
}

func doStore(store infra.Store, envelopes []*envelope.Envelope) error {
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

func doStoreAndPublish(store infra.Store, bus infra.PublishSubscriber, envelopes []*envelope.Envelope) error {
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

func (ch *GamblerCommandHandler) HandleGetResultsQuery(year int) (*Results, error) {
	return nil, errors.New("HandleGetResultsQuery not implemented")
}

func getGamblerContext(store infra.Store, gamblerUid string, year int) (*GamblerContext, error) {
	context := NewGamblerContext()

	tourRelatedEvents, err := store.Get("tour", strconv.Itoa(year))
	if err != nil {
		return context, err
	}

	gamblerRelatedEvents, err := store.Get("gambler", gamblerUid)
	if err != nil {
		return context, err
	}

	for _, envelop := range tourRelatedEvents {
		if events.IsTourCreated(&envelop) {
			context.ApplyTourCreated(events.UnWrapTourCreated(&envelop))
		} else if events.IsCyclistCreated(&envelop) {
			context.ApplyCyclistCreated(events.UnWrapCyclistCreated(&envelop))
		} else if events.IsEtappeCreated(&envelop) {
			context.ApplyEtappeCreated(events.UnWrapEtappeCreated(&envelop))
		} else if events.IsEtappeResultsCreated(&envelop) {
			context.ApplyEtappeResultsCreated(events.UnWrapEtappeResultsCreated(&envelop))
		} else {
			log.Panicf("getGamblerOnUid(tour): Unexpected event %s", envelop.EventTypeName)
		}
	}

	for _, envelop := range gamblerRelatedEvents {
		if events.IsGamblerCreated(&envelop) {
			event := events.UnWrapGamblerCreated(&envelop)
			context.ApplyGamblerCreated(event)
		} else if events.IsGamblerTeamCreated(&envelop) {
			event := events.UnWrapGamblerTeamCreated(&envelop)
			context.ApplyGamblerTeamCreated(event)
		} else {
			log.Panicf("getGamblerOnUid(gambler): Unexpected event %s", envelop.EventTypeName)
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

func (context *GamblerContext) ApplyTourCreated(event *events.TourCreated) {
	//log.Printf("ApplyTourCreated: context before: %+v, event: %+v", context, event)

	context.Year = new(int)
	*context.Year = event.Year

	//log.Printf("ApplyTourCreated: context after: %+v", context)

	return
}

func (context *GamblerContext) ApplyCyclistCreated(event *events.CyclistCreated) {
	//log.Printf("ApplyCyclistCreated: context before: %+v, event: %+v", context, event)

	context.cyclistsForTour[event.CyclistId] =
		Cyclist{Id: event.CyclistId, Name: event.CyclistName, Team: event.CyclistTeam}

	//log.Printf("ApplyCyclistCreated: context after: %+v", context)

	return
}

func (context *GamblerContext) ApplyGamblerCreated(event *events.GamblerCreated) {
	//log.Printf("ApplyGamblerCreated: context before: %+v, event: %+v", context, event)

	context.Gambler = NewGambler(event.GamblerUid, event.GamblerName, event.GamblerEmail)

	//log.Printf("ApplyGamblerCreated: context after: %+v", context)

	return
}

func (context *GamblerContext) ApplyGamblerTeamCreated(event *events.GamblerTeamCreated) {
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

func (context *GamblerContext) ApplyEtappeCreated(event *events.EtappeCreated) {
	log.Fatal("gambler.ApplyEtappeCreated not implemented")
}

func (context *GamblerContext) ApplyEtappeResultsCreated(event *events.EtappeResultsCreated) {
	log.Fatal("gambler.ApplyEtappeResultsCreated not implemented")
}

type Results struct {
	BestGamblers []Gambler
	BestCyclists []Cyclist
}
