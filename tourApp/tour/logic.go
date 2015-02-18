package tour

import (
	"errors"
	"fmt"
	"github.com/MarcGrol/microgen/tourApp/events"
	//	"log"
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

func (tch *TourCommandHandler) HandleCreateTourCommand(command CreateTourCommand) error {
	// TODO validate input

	// get tour based on year
	_, found := getTourOnYear(tch.store, command.Year)
	if found == true {
		return errors.New(fmt.Sprintf("Tour %d already exists", command.Year))
	}

	// apply business logic
	tour := NewTour()
	tourCreatedEvent := events.TourCreated{command.Year}
	tour.ApplyTourCreated(tourCreatedEvent)

	//log.Printf("HandleCreateTourCommand completed:%v -> %v", command, tourCreatedEvent)

	// store and emit resulting event
	return tch.publishAndStore([]*events.Envelope{tourCreatedEvent.Wrap()})
}

func (tch *TourCommandHandler) HandleCreateCyclistCommand(command CreateCyclistCommand) error {
	// TODO validate input

	// get tour based on year
	tour, found := getTourOnYear(tch.store, command.Year)
	if found == false {
		return errors.New(fmt.Sprintf("Tour %d does not exists", command.Year))
	}

	// apply business logic
	cyclistCreatedEvent := events.CyclistCreated{Year: command.Year,
		CyclistId:   command.Id,
		CyclistName: command.Name,
		CyclistTeam: command.Team}
	tour.ApplyCyclistCreated(cyclistCreatedEvent)

	//log.Printf("HandleCreateCyclistCommand completed:%v -> %v", command, cyclistCreatedEvent)

	// store and emit resulting event
	return tch.publishAndStore([]*events.Envelope{cyclistCreatedEvent.Wrap()})
}

func (tch *TourCommandHandler) HandleCreateEtappeCommand(command CreateEtappeCommand) error {
	// TODO validate input

	// get tour based on year
	tour, found := getTourOnYear(tch.store, command.Year)
	if found == false {
		return errors.New(fmt.Sprintf("Tour %d does not exists", command.Year))
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
	return tch.publishAndStore([]*events.Envelope{etappeCreatedEvent.Wrap()})
}

func (tch *TourCommandHandler) publishAndStore(envelopes []*events.Envelope) error {
	for _, env := range envelopes {
		//log.Printf("publishAndStore:%v", env)
		err := tch.store.Store(env)
		if err != nil {
			return err
		}
		err = tch.bus.Publish(env)
		if err != nil {
			return err
		}
	}
	return nil
}

func getTourOnYear(store events.Store, year int) (*Tour, bool) {
	tourRelatedEvents := make([]*events.Envelope, 0, 10)

	callback := func(envelope *events.Envelope) {
		if envelope.AggregateName == "tour" && envelope.AggregateUid == strconv.Itoa(year) {
			tourRelatedEvents = append(tourRelatedEvents, envelope)
		}
	}
	store.Iterate(callback)

	if len(tourRelatedEvents) == 0 {
		return nil, false
	}

	tour := NewTour()
	for _, envelope := range tourRelatedEvents {
		var err error
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
	year     *int
	etappes  []*Etappe
	cyclists []*Cyclist
}

type Cyclist struct {
	number int
	name   string
	team   string
}

type Etappe struct {
	id             int
	date           time.Time
	startLocation  string
	finishLocation string
	length         int
	kind           int
}

func NewTour() *Tour {
	tour := new(Tour)
	tour.etappes = make([]*Etappe, 0, 30)
	tour.cyclists = make([]*Cyclist, 0, 250)
	return tour
}

func (t *Tour) ApplyTourCreated(event events.TourCreated) error {
	//log.Printf("ApplyTourCreated:%v", event)

	t.year = new(int)
	*t.year = event.Year
	return nil
}

func (t *Tour) ApplyCyclistCreated(event events.CyclistCreated) error {
	//log.Printf("ApplyCyclistCreated:%v", event)

	cyclist := new(Cyclist)
	cyclist.number = event.CyclistId
	cyclist.name = event.CyclistName
	cyclist.team = event.CyclistTeam
	t.cyclists = append(t.cyclists, cyclist)
	return nil
}

func (t *Tour) ApplyEtappeCreated(event events.EtappeCreated) error {
	//log.Printf("ApplyEtappeCreated:%v", event)

	etappe := new(Etappe)
	etappe.id = event.EtappeId
	etappe.date = event.EtappeDate
	etappe.startLocation = event.EtappeStartLocation
	etappe.finishLocation = event.EtappeFinishLocation
	etappe.length = event.EtappeLength
	etappe.kind = event.EtappeKind
	t.etappes = append(t.etappes, etappe)
	return nil
}

type TourQueryHandler struct {
	bus   events.PublishSubscriber
	store events.Store
}

func NewTourQueryHandler(bus events.PublishSubscriber, store events.Store) *TourQueryHandler {
	handler := new(TourQueryHandler)
	handler.bus = bus
	handler.store = store
	return handler
}

func (tqh *TourQueryHandler) GetTour(year int) (*Tour, error) {
	// TODO validate input

	// get tour based on year
	tour, found := getTourOnYear(tqh.store, year)
	if found == false {
		return nil, errors.New(fmt.Sprintf("Tour %d not found", year))
	}
	return tour, nil
}
