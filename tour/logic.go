package tour

import (
	"errors"
	"fmt"
	"github.com/xebia/microgen/events"
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
	// get tour based on year
	_, found := getTourOnYear(tch.store, command.Year)
	if found == true {
		return errors.New(fmt.Sprintf("Tour %d already exists", command.Year))
	}

	// apply business logic
	tour := NewTour()
	tourCreatedEvent := events.TourCreated{command.Year}
	tour.ApplyTourCreated(tourCreatedEvent)

	// store and emit resulting event
	return tch.publishAndStore([]*events.Envelope{tourCreatedEvent.Wrap()})
}

func (tch *TourCommandHandler) HandleCreateCyclistCommand(command CreateCyclistCommand) error {
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

	// store and emit resulting event
	return tch.publishAndStore([]*events.Envelope{cyclistCreatedEvent.Wrap()})
}

func (tch *TourCommandHandler) HandleCreateEtappeCommand(command CreateEtappeCommand) error {
	// get tour based on year
	tour, found := getTourOnYear(tch.store, command.Year)
	if found == false {
		return errors.New(fmt.Sprintf("Tour %d does not exists", command.Year))
	}

	// apply business logic
	etappeCreatedEvent := events.EtappeCreated{Year: command.Year,
		EtaopeId:              command.Id,
		EtappeDate:            command.Date,
		EtappeStartLocation:   command.StartLocation,
		EtappeFinishtLocation: command.FinishLocation,
		EtappeLength:          command.Length,
		EtappeKind:            command.Kind}
	tour.ApplyEtappeCreated(etappeCreatedEvent)

	// store and emit resulting event
	return tch.publishAndStore([]*events.Envelope{etappeCreatedEvent.Wrap()})
}

func (tch *TourCommandHandler) publishAndStore([]*events.Envelope) error {
	return errors.New("publishAndStore not implemented")
}

func getTourOnYear(store events.Store, year int) (*Tour, bool) {
	var tourCreatedEvent *events.TourCreated = nil

	callback := func(envelope *events.Envelope) bool {
		if envelope.Type == events.TypeTourCreated && envelope.TourCreated != nil && envelope.TourCreated.Year == year {
			tourCreatedEvent = envelope.TourCreated
			return true
		}
		return false
	}
	store.Iterate(callback)

	if tourCreatedEvent == nil {
		return nil, false
	}

	tour := NewTour()
	tour.ApplyTourCreated(*tourCreatedEvent)
	return tour, true
}

type Tour struct {
	Year     *int
	Etappes  []*Etappe
	Cyclists []*Cyclist
}

type Cyclist struct {
	Number int
	Name   string
	Team   string
}

type Etappe struct {
	Id             int
	Date           time.Time
	StartLocation  string
	FinishLocation string
	Length         int
	Kind           int
}

func NewTour() *Tour {
	tour := new(Tour)
	tour.Etappes = make([]*Etappe, 0, 30)
	tour.Cyclists = make([]*Cyclist, 0, 250)
	return tour
}

func (t *Tour) ApplyTourCreated(event events.TourCreated) error {
	t.Year = new(int)
	*t.Year = event.Year
	return nil
}

func (t *Tour) ApplyCyclistCreated(event events.CyclistCreated) error {
	cyclist := new(Cyclist)
	cyclist.Number = event.CyclistId
	cyclist.Name = event.CyclistName
	cyclist.Team = event.CyclistTeam
	t.Cyclists = append(t.Cyclists, cyclist)
	return nil
}

func (t *Tour) ApplyEtappeCreated(event events.EtappeCreated) error {
	etappe := new(Etappe)
	etappe.Id = event.EtaopeId
	etappe.Date = event.EtappeDate
	etappe.StartLocation = event.EtappeStartLocation
	etappe.FinishLocation = event.EtappeFinishtLocation
	etappe.Length = event.EtappeLength
	etappe.Kind = event.EtappeKind
	t.Etappes = append(t.Etappes, etappe)
	return nil
}
