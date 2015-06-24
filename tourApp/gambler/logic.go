package gambler

//go:generate gen

import (
	"errors"
	"fmt"
	"log"
	"strconv"

	"github.com/MarcGrol/microgen/infra"
	"github.com/MarcGrol/microgen/lib/envelope"
	"github.com/MarcGrol/microgen/lib/myerrors"
	"github.com/MarcGrol/microgen/lib/validation"

	"github.com/MarcGrol/microgen/tourApp/events"
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
		err := eventHandler.bus.Subscribe(eventType.String(), func(envelope *envelope.Envelope) error {
			return eventHandler.OnEvent(envelope)
		})
		if err != nil {
			return err
		}
	}

	return nil
}

func (eh *GamblerEventHandler) OnEvent(envelop *envelope.Envelope) error {
	return doStore(eh.store, []*envelope.Envelope{envelop})
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
	v := validation.Validator{}
	v.NotEmpty("GamblerUid", command.GamblerUid)
	v.NotEmpty("Name", command.Name)
	v.NotEmpty("Email", command.Email)

	return v.Err
}

func (ch *GamblerCommandHandler) HandleCreateGamblerCommand(command *CreateGamblerCommand) error {
	err := ch.validateCreateGamblerCommand(command)
	if err != nil {
		return myerrors.NewInvalidInputError(err)
	}
	gamblingContext, err := getGamblingContext(ch.store, command.GamblerUid, -1)
	if err != nil {
		return myerrors.NewInternalError(err)
	}

	_, found := gamblingContext.gamblersForTour[command.GamblerUid]
	if found == true {
		return myerrors.NewNotFoundErrorf("Gambler %s already exists",
			command.GamblerUid)
	}

	gamblingContext.gamblersForTour[command.GamblerUid] =
		&Gambler{
			Uid:      command.GamblerUid,
			Name:     command.Name,
			Email:    command.Email,
			Cyclists: make([]*Cyclist, 0, 10),
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
	v := validation.Validator{}
	v.GreaterThan("Year", 2014, command.Year)
	v.NotEmpty("GamblerUid", command.GamblerUid)
	v.MinSliceLength("CyclistIds", 10, command.CyclistIds)
	v.NoDuplicates("CyclistIds", command.CyclistIds)
	return v.Err
}

func (ch *GamblerCommandHandler) HandleCreateGamblerTeamCommand(command *CreateGamblerTeamCommand) error {
	err := ch.validateCreateGamblerTeamCommand(command)
	if err != nil {
		return myerrors.NewInvalidInputError(err)
	}
	gamblingContext, err := getGamblingContext(ch.store, command.GamblerUid, command.Year)
	if err != nil {
		return myerrors.NewInternalError(err)
	}
	if gamblingContext.Year == nil || *gamblingContext.Year != command.Year {
		return myerrors.NewNotFoundErrorf("Tour %d not found", command.Year)
	}
	_, found := gamblingContext.gamblersForTour[command.GamblerUid]
	if found == false {
		return myerrors.NewNotFoundErrorf("Gambler %s not found", command.GamblerUid)
	}

	err = cyclistsExist(gamblingContext.cyclistsForTour, command.CyclistIds)
	if err != nil {
		return myerrors.NewInvalidInputError(err)
	}

	// apply business logic
	gamblerTeamCreatedEvent := events.GamblerTeamCreated{
		GamblerUid:      command.GamblerUid,
		Year:            command.Year,
		GamblerCyclists: command.CyclistIds}

	return doStoreAndPublish(ch.store, ch.bus, []*envelope.Envelope{gamblerTeamCreatedEvent.Wrap()})
}

func cyclistsExist(allCyclists map[int]*Cyclist, cyclistIds []int) error {
	for _, id := range cyclistIds {
		_, exists := allCyclists[id]
		if exists == false {
			return fmt.Errorf("Cyclist %d does not exist", id)
		}
	}
	return nil
}

func doStore(store infra.Store, envelopes []*envelope.Envelope) error {
	for _, env := range envelopes {
		err := store.Store(env)
		if err != nil {
			log.Printf("Error storing event: %+v", err)
			return err
		}
		//log.Printf("Successfully stored event: %+v", env)
	}
	return nil
}

func doStoreAndPublish(store infra.Store, bus infra.PublishSubscriber, envelopes []*envelope.Envelope) error {
	err := doStore(store, envelopes)
	if err != nil {
		return myerrors.NewInternalError(err)
	}
	for _, envelop := range envelopes {
		err = bus.Publish(envelop)
		if err != nil {
			return myerrors.NewInternalError(err)
		}
	}
	return nil
}

func (ch *GamblerCommandHandler) HandleGetGamblerQuery(gamblerUid string, year int) (*Gambler, error) {
	// TODO validate input
	gamblingContext, err := getGamblingContext(ch.store, gamblerUid, year)
	if err != nil {
		return nil, myerrors.NewInternalError(err)
	}
	gambler, found := gamblingContext.gamblersForTour[gamblerUid]
	if found == false {
		return nil, myerrors.NewNotFoundErrorf("Gambler with uid %s not found", gamblerUid)
	}

	//log.Printf("HandleGetGamblerQuery.Gambler:%+v", gamblingContext.Gambler)

	return gambler, nil
}

func (ch *GamblerCommandHandler) HandleGetResultsQuery(year int) (*Results, error) {
	return nil, errors.New("HandleGetResultsQuery not implemented")
}

func getGamblingContext(store infra.Store, gamblerUid string, year int) (*GamblingContext, error) {
	context := NewGamblingContext()

	tourRelatedEvents, err := store.Get("tour", strconv.Itoa(year))
	if err != nil {
		return context, err
	}
	applyEvents(tourRelatedEvents, context)

	gamblerRelatedEvents, err := store.Get("gambler", gamblerUid)
	if err != nil {
		return context, err
	}
	applyEvents(gamblerRelatedEvents, context)

	return context, nil
}

type GamblingContext struct {
	Year            *int
	cyclistsForTour map[int]*Cyclist
	etappes         map[int]*Etappe
	gamblersForTour map[string]*Gambler
}

func NewGamblingContext() *GamblingContext {
	context := new(GamblingContext)
	context.Year = nil
	context.cyclistsForTour = make(map[int]*Cyclist)
	context.etappes = make(map[int]*Etappe)
	context.gamblersForTour = make(map[string]*Gambler)

	return context
}

// +gen slice:"SortBy,Where,Select[string],GroupBy[string]"
type Gambler struct {
	Uid      string
	Name     string
	Email    string
	Cyclists []*Cyclist
	Points   int
}

func NewGambler(uid string, name string, email string) *Gambler {
	gambler := new(Gambler)
	gambler.Uid = uid
	gambler.Name = name
	gambler.Email = email
	gambler.Cyclists = make([]*Cyclist, 0, 10)
	return gambler
}

// +gen slice:"SortBy,Where,Select[string],GroupBy[string]"
type Cyclist struct {
	Id     int
	Name   string
	Team   string
	Points int
}

type Etappe struct {
	id   int
	kind int
}

func (context *GamblingContext) ApplyTourCreated(event *events.TourCreated) {
	//log.Printf("ApplyTourCreated: context before: %+v, event: %+v", context, event)

	context.Year = new(int)
	*context.Year = event.Year

	log.Printf("ApplyTourCreated: context after: %+v", context)

	return
}

func (context *GamblingContext) ApplyCyclistCreated(event *events.CyclistCreated) {
	//log.Printf("ApplyCyclistCreated: context before: %+v, event: %+v", context, event)

	context.cyclistsForTour[event.CyclistId] =
		&Cyclist{
			Id:   event.CyclistId,
			Name: event.CyclistName,
			Team: event.CyclistTeam}

	log.Printf("ApplyCyclistCreated: context after: %+v", context)

	return
}

func (context *GamblingContext) ApplyGamblerCreated(event *events.GamblerCreated) {
	//log.Printf("ApplyGamblerCreated: context before: %+v, event: %+v", context, event)

	context.gamblersForTour[event.GamblerUid] =
		NewGambler(event.GamblerUid, event.GamblerName, event.GamblerEmail)

	log.Printf("ApplyGamblerCreated: context after: %+v", context)

	return
}

func (context *GamblingContext) ApplyGamblerTeamCreated(event *events.GamblerTeamCreated) {
	//log.Printf("ApplyGamblerTeamCreated: context: %+v, event: %+v", context, event)

	gambler, found := context.gamblersForTour[event.GamblerUid]
	if found {
		for _, cyclistId := range event.GamblerCyclists {
			cyclist, found := context.cyclistsForTour[cyclistId]
			if found {
				gambler.Cyclists = append(gambler.Cyclists, cyclist)
			}
		}
	}

	log.Printf("ApplyGamblerTeamCreated: context after: %+v", context)
	return
}

func (context *GamblingContext) ApplyEtappeCreated(event *events.EtappeCreated) {
	context.etappes[event.EtappeId] = &Etappe{
		id:   event.EtappeId,
		kind: event.EtappeKind}

	log.Printf("ApplyEtappeCreated: context after: %+v", context)

	return
}

type Classement int

const (
	ClassementUnknown Classement = iota
	ClassementDay
	ClassementAllround
	ClassementSprint
	ClassementClimb
)

func (context *GamblingContext) ApplyEtappeResultsCreated(event *events.EtappeResultsCreated) {
	etappe, found := context.etappes[event.LastEtappeId]
	if found {
		context.calculateCyclistPointsForEtappe(etappe, event.BestDayCyclistIds, ClassementDay)
		context.calculateCyclistPointsForEtappe(etappe, event.BestAllrounderCyclistIds, ClassementAllround)
		context.calculateCyclistPointsForEtappe(etappe, event.BestSprinterCyclistIds, ClassementSprint)
		context.calculateCyclistPointsForEtappe(etappe, event.BestClimberCyclistIds, ClassementClimb)

		/*
		 * calculate points for gamblers
		 */
		for _, gambler := range context.gamblersForTour {
			for _, cyclist := range gambler.Cyclists {
				gambler.Points += cyclist.Points
			}
		}

	}
	//log.Printf("ApplyEtappeResultsCreated: context after: %+v", context)
}

func (context *GamblingContext) calculateCyclistPointsForEtappe(etappe *Etappe, cyclistIds []int, classementType Classement) {
	for rank, cyclistId := range cyclistIds {
		cyclist, found := context.cyclistsForTour[cyclistId]
		if found {
			cyclist.Points += getPointsFor(etappe, rank, classementType, cyclist)
		}
	}
}

func getPointsFor(etappe *Etappe, rank int, classsementType Classement, cyclist *Cyclist) int {
	return 42
}

type Results struct {
	BestGamblers GamblerSlice
	BestCyclists CyclistSlice
}
