package tour

import (
	"errors"
	"fmt"
	"github.com/MarcGrol/microgen/myerrors"
	"github.com/MarcGrol/microgen/tourApp/events"
	//"log"
	"strconv"
	"time"
)

type TourCommandHandler struct {
	bus   events.PublishSubscriber
	store events.Store
}

func NewTourCommandHandler(bus events.PublishSubscriber, store events.Store) CommandHandler {
	handler := new(TourCommandHandler)
	handler.bus = bus
	handler.store = store
	return handler
}

func (tch *TourCommandHandler) HandleCreateTourCommand(command CreateTourCommand) *myerrors.Error {
	// TODO validate input

	// get tour based on year
	_, found := getTourOnYear(tch.store, command.Year)
	if found == true {
		return myerrors.NewInvalidInputError(errors.New(fmt.Sprintf("Tour %d already exists", command.Year)))
	}

	// apply business logic
	tour := NewTour()
	tourCreatedEvent := events.TourCreated{command.Year}
	tour.ApplyTourCreated(tourCreatedEvent)

	//log.Printf("HandleCreateTourCommand completed:%v -> %v", command, tourCreatedEvent)

	// store and emit resulting event
	return tch.storeAndPublish([]*events.Envelope{tourCreatedEvent.Wrap()})
}

func (tch *TourCommandHandler) HandleCreateCyclistCommand(command CreateCyclistCommand) *myerrors.Error {
	// TODO validate input

	// get tour based on year
	tour, found := getTourOnYear(tch.store, command.Year)
	if found == false {
		return myerrors.NewNotFoundError(errors.New(fmt.Sprintf("Tour %d does not exists", command.Year)))
	}

	// apply business logic
	cyclistCreatedEvent := events.CyclistCreated{Year: command.Year,
		CyclistId:   command.Id,
		CyclistName: command.Name,
		CyclistTeam: command.Team}
	tour.ApplyCyclistCreated(cyclistCreatedEvent)

	//log.Printf("HandleCreateCyclistCommand completed:%v -> %v", command, cyclistCreatedEvent)

	// store and emit resulting event
	return tch.storeAndPublish([]*events.Envelope{cyclistCreatedEvent.Wrap()})
}

func (tch *TourCommandHandler) HandleCreateEtappeCommand(command CreateEtappeCommand) *myerrors.Error {
	// TODO validate input

	// get tour based on year
	tour, found := getTourOnYear(tch.store, command.Year)
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
	tour.ApplyEtappeCreated(etappeCreatedEvent)

	//log.Printf("HandleCreateEtappeCommand completed:%v -> %v", command, etappeCreatedEvent)

	// store and emit resulting event
	return tch.storeAndPublish([]*events.Envelope{etappeCreatedEvent.Wrap()})
}

func (tch *TourCommandHandler) HandleGetTourQuery(year int) (interface{}, *myerrors.Error) {
	// TODO validate input
	tour, found := getTourOnYear(tch.store, year)
	if found == false {
		return nil, myerrors.NewNotFoundError(errors.New(fmt.Sprintf("Tour %d not found", year)))
	}
	//log.Printf("GetTour:%v", tour)

	return tour, nil
}

func (tch *TourCommandHandler) storeAndPublish(envelopes []*events.Envelope) *myerrors.Error {
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

func getTourOnYear(store events.Store, year int) (*Tour, bool) {
	tourRelatedEvents, err := store.Get("tour", strconv.Itoa(year))
	if err != nil || len(tourRelatedEvents) == 0 {
		return nil, false
	}

	tour := NewTour()
	for _, envelope := range tourRelatedEvents {
		var err *myerrors.Error
		if envelope.Type == events.TypeTourCreated {
			err = tour.ApplyTourCreated(*envelope.TourCreated)
		} else if envelope.Type == events.TypeEtappeCreated {
			err = tour.ApplyEtappeCreated(*envelope.EtappeCreated)
		} else if envelope.Type == events.TypeCyclistCreated {
			err = tour.ApplyCyclistCreated(*envelope.CyclistCreated)
		}
		if err != nil {
			break
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

func (t *Tour) ApplyTourCreated(event events.TourCreated) *myerrors.Error {

	t.Year = event.Year

	//log.Printf("ApplyTourCreated after:%v -> %v", event, t)

	return nil
}

func (t *Tour) ApplyCyclistCreated(event events.CyclistCreated) *myerrors.Error {

	cyclist := new(Cyclist)
	cyclist.Number = event.CyclistId
	cyclist.Name = event.CyclistName
	cyclist.Team = event.CyclistTeam
	t.Cyclists = append(t.Cyclists, *cyclist)

	//log.Printf("ApplyCyclistCreated after:%v -> %v", event, t)

	return nil
}

func (t *Tour) ApplyEtappeCreated(event events.EtappeCreated) *myerrors.Error {

	etappe := new(Etappe)

	etappe.Id = event.EtappeId
	etappe.Date = event.EtappeDate
	etappe.StartLocation = event.EtappeStartLocation
	etappe.FinishLocation = event.EtappeFinishLocation
	etappe.Length = event.EtappeLength
	etappe.Kind = event.EtappeKind
	t.Etappes = append(t.Etappes, *etappe)

	//log.Printf("ApplyEtappeCreated after:%v -> %v", event, t)

	return nil
}
