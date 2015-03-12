package tour

import (
	"errors"
	"fmt"
	"github.com/MarcGrol/microgen/infra"
	"github.com/MarcGrol/microgen/lib/envelope"
	"github.com/MarcGrol/microgen/lib/myerrors"
	"github.com/MarcGrol/microgen/tourApp/events"
	"log"
	"strconv"
	"time"
)

type EtappeKind int

const (
	Flat      = 1
	Hilly     = 2
	Mountains = 3
	TimeTrial = 4
)

type TourEventHandler struct {
	bus   infra.PublishSubscriber
	store infra.Store
}

func NewTourEventHandler(bus infra.PublishSubscriber, store infra.Store) *TourEventHandler {
	handler := new(TourEventHandler)
	handler.bus = bus
	handler.store = store
	return handler
}

func (eventHandler *TourEventHandler) Start() error {
	// subscribe to events?
	return nil
}

type TourCommandHandler struct {
	bus   infra.PublishSubscriber
	store infra.Store
}

func NewTourCommandHandler(bus infra.PublishSubscriber, store infra.Store) CommandHandler {
	handler := new(TourCommandHandler)
	handler.bus = bus
	handler.store = store
	return handler
}

func (ch *TourCommandHandler) validateCreateTourCommand(command *CreateTourCommand) error {
	// TODO
	return nil
}

func (ch *TourCommandHandler) HandleCreateTourCommand(command *CreateTourCommand) error {
	err := ch.validateCreateTourCommand(command)
	if err != nil {
		return myerrors.NewInvalidInputError(err)
	}

	// get tour based on year
	_, found := getTourOnYear(ch.store, command.Year)
	if found == true {
		return myerrors.NewInvalidInputError(errors.New(fmt.Sprintf("Tour %d already exists", command.Year)))
	}

	tourCreatedEvent := events.TourCreated{command.Year}

	log.Printf("HandleCreateTourCommand completed:%+v -> %+v", command, tourCreatedEvent)

	// store and emit resulting event
	return ch.storeAndPublish([]*envelope.Envelope{tourCreatedEvent.Wrap()})
}

func (ch *TourCommandHandler) validateCreateCyclistCommand(command *CreateCyclistCommand) error {
	// TODO
	return nil
}

func (ch *TourCommandHandler) HandleCreateCyclistCommand(command *CreateCyclistCommand) error {
	err := ch.validateCreateCyclistCommand(command)
	if err != nil {
		return myerrors.NewInvalidInputError(err)
	}

	// get tour based on year
	_, found := getTourOnYear(ch.store, command.Year)
	if found == false {
		return myerrors.NewNotFoundError(errors.New(fmt.Sprintf("Tour %d does not exists", command.Year)))
	}

	// apply business logic
	cyclistCreatedEvent := events.CyclistCreated{Year: command.Year,
		CyclistId:   command.Id,
		CyclistName: command.Name,
		CyclistTeam: command.Team}

	//log.Printf("HandleCreateCyclistCommand completed:%+v -> %+v", command, cyclistCreatedEvent)

	// store and emit resulting event
	return ch.storeAndPublish([]*envelope.Envelope{cyclistCreatedEvent.Wrap()})
}

func (ch *TourCommandHandler) validateCreateEtappeCommand(command *CreateEtappeCommand) error {
	// TODO
	return nil
}

func (ch *TourCommandHandler) HandleCreateEtappeCommand(command *CreateEtappeCommand) error {
	log.Printf("create etappe command: %+v", command)
	err := ch.validateCreateEtappeCommand(command)
	if err != nil {
		return myerrors.NewInvalidInputError(err)
	}

	// get tour based on year
	_, found := getTourOnYear(ch.store, command.Year)
	if found == false {
		return myerrors.NewNotFoundError(errors.New(fmt.Sprintf("Tour %d does not exists", command.Year)))
	}

	// apply business logic
	etappeCreatedEvent := events.EtappeCreated{Year: command.Year,
		EtappeId:             command.Id,
		EtappeDate:           command.Date,
		EtappeStartLocation:  command.StartLocation,
		EtappeFinishLocation: command.FinishLocation,
		EtappeLength:         command.Length,
		EtappeKind:           command.Kind}

	//log.Printf("HandleCreateEtappeCommand completed:%+v -> %+v", command, etappeCreatedEvent)

	// store and emit resulting event
	return ch.storeAndPublish([]*envelope.Envelope{etappeCreatedEvent.Wrap()})
}

func (ch *TourCommandHandler) HandleCreateEtappeResultsCommand(command *CreateEtappeResultsCommand) error {
	return errors.New("HandleCreateEtappeResultsCommand not implemented")
}

func (ch *TourCommandHandler) HandleGetTourQuery(year int) (*Tour, error) {
	// TODO validate input
	tour, found := getTourOnYear(ch.store, year)
	if found == false {
		return nil, myerrors.NewNotFoundError(errors.New(fmt.Sprintf("Tour %d not found", year)))
	}
	log.Printf("GetTour:%+v", tour)

	return tour, nil
}

func (ch *TourCommandHandler) storeAndPublish(envelopes []*envelope.Envelope) error {
	for _, env := range envelopes {
		err := ch.store.Store(env)
		if err != nil {
			return myerrors.NewInternalError(err)
		}
		err = ch.bus.Publish(env)
		if err != nil {
			return myerrors.NewInternalError(err)
		}
	}
	return nil
}

func getTourOnYear(store infra.Store, year int) (*Tour, bool) {
	tourRelatedEvents, err := store.Get("tour", strconv.Itoa(year))
	if err != nil || len(tourRelatedEvents) == 0 {
		return nil, false
	}

	tour := NewTour()
	applyEvents(tourRelatedEvents, tour)
	return tour, true
}

type Tour struct {
	Year     int       `json:"year"`
	Etappes  []Etappe  `json:"etappes"`
	Cyclists []Cyclist `json:"cyclists"`
}

type Cyclist struct {
	Number int    `json:"number"`
	Name   string `json:"name"`
	Team   string `json:"team"`
}

type Etappe struct {
	Id             int       `json:"id"`
	Date           time.Time `json:"date"`
	StartLocation  string    `json:"startLocation"`
	FinishLocation string    `json:"finishLocation"`
	Length         int       `json:"length"`
	Kind           int       `json:"kind"`
}

func NewTour() *Tour {
	tour := new(Tour)
	tour.Etappes = make([]Etappe, 0, 30)
	tour.Cyclists = make([]Cyclist, 0, 250)
	return tour
}

func (t *Tour) ApplyTourCreated(event *events.TourCreated) {

	t.Year = event.Year

	//log.Printf("ApplyTourCreated after:%+v -> %+v", event, t)

	return
}

func (t *Tour) ApplyCyclistCreated(event *events.CyclistCreated) {

	cyclist := new(Cyclist)
	cyclist.Number = event.CyclistId
	cyclist.Name = event.CyclistName
	cyclist.Team = event.CyclistTeam
	t.Cyclists = append(t.Cyclists, *cyclist)

	//log.Printf("ApplyCyclistCreated after:%+v -> %+v", event, t)

	return
}

func (t *Tour) ApplyEtappeCreated(event *events.EtappeCreated) {

	etappe := new(Etappe)

	etappe.Id = event.EtappeId
	etappe.Date = event.EtappeDate
	etappe.StartLocation = event.EtappeStartLocation
	etappe.FinishLocation = event.EtappeFinishLocation
	etappe.Length = event.EtappeLength
	etappe.Kind = event.EtappeKind
	t.Etappes = append(t.Etappes, *etappe)

	//log.Printf("ApplyEtappeCreated after:%+v -> %+v", event, t)

	return
}

func (t *Tour) ApplyEtappeResultsCreated(event *events.EtappeResultsCreated) {
	log.Fatal("ApplyEtappeResultsCreated not implemented")

}
