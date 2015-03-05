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

func (tch *TourCommandHandler) validateCreateTourCommand(command *CreateTourCommand) error {
	// TODO
	return nil
}

func (tch *TourCommandHandler) HandleCreateTourCommand(command *CreateTourCommand) error {
	err := tch.validateCreateTourCommand(command)
	if err != nil {
		return myerrors.NewInvalidInputError(err)
	}

	// get tour based on year
	_, found := getTourOnYear(tch.store, command.Year)
	if found == true {
		return myerrors.NewInvalidInputError(errors.New(fmt.Sprintf("Tour %d already exists", command.Year)))
	}

	// apply business logic
	tourCreatedEvent := events.TourCreated{command.Year}

	log.Printf("HandleCreateTourCommand completed:%+v -> %+v", command, tourCreatedEvent)

	// store and emit resulting event
	return tch.storeAndPublish([]*envelope.Envelope{tourCreatedEvent.Wrap()})
}

func (tch *TourCommandHandler) validateCreateCyclistCommand(command *CreateCyclistCommand) error {
	// TODO
	return nil
}

func (tch *TourCommandHandler) HandleCreateCyclistCommand(command *CreateCyclistCommand) error {
	err := tch.validateCreateCyclistCommand(command)
	if err != nil {
		return myerrors.NewInvalidInputError(err)
	}

	// get tour based on year
	_, found := getTourOnYear(tch.store, command.Year)
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
	return tch.storeAndPublish([]*envelope.Envelope{cyclistCreatedEvent.Wrap()})
}

func (tch *TourCommandHandler) validateCreateEtappeCommand(command *CreateEtappeCommand) error {
	// TODO
	return nil
}

func (tch *TourCommandHandler) HandleCreateEtappeCommand(command *CreateEtappeCommand) error {
	err := tch.validateCreateEtappeCommand(command)
	if err != nil {
		return myerrors.NewInvalidInputError(err)
	}

	// get tour based on year
	_, found := getTourOnYear(tch.store, command.Year)
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
	return tch.storeAndPublish([]*envelope.Envelope{etappeCreatedEvent.Wrap()})
}

func (tch *TourCommandHandler) HandleGetTourQuery(year int) (*Tour, error) {
	// TODO validate input
	tour, found := getTourOnYear(tch.store, year)
	if found == false {
		return nil, myerrors.NewNotFoundError(errors.New(fmt.Sprintf("Tour %d not found", year)))
	}
	log.Printf("GetTour:%+v", tour)

	return tour, nil
}

func (tch *TourCommandHandler) storeAndPublish(envelopes []*envelope.Envelope) error {
	for _, env := range envelopes {
		err := tch.store.Store(env)
		if err != nil {
			return myerrors.NewInternalError(err)
		}
		err = tch.bus.Publish(env)
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
	for _, envelop := range tourRelatedEvents {
		if events.IsTourCreated(&envelop) {
			event := events.UnWrapTourCreated(&envelop)
			tour.ApplyTourCreated(event)
		} else if events.IsEtappeCreated(&envelop) {
			event := events.UnWrapEtappeCreated(&envelop)
			tour.ApplyEtappeCreated(event)
		} else if events.IsCyclistCreated(&envelop) {
			event := events.UnWrapCyclistCreated(&envelop)
			tour.ApplyCyclistCreated(event)
		} else {
			log.Panicf("getTourOnYear: Unexpected event %s", envelop.EventTypeName)
		}
	}

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
