package tour

//go:generate gen

import (
	"errors"
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/MarcGrol/microgen/infra"
	"github.com/MarcGrol/microgen/lib/envelope"
	"github.com/MarcGrol/microgen/lib/myerrors"
	"github.com/MarcGrol/microgen/lib/validation"
	"github.com/MarcGrol/microgen/tourApp/events"
)

type EtappeKind int

const (
	Flat = 1 + iota
	Hilly
	Mountains
	TimeTrial
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
	v := validation.Validator{}
	v.GreaterThan("Year", 2010, command.Year)
	return v.Err
}

func (ch *TourCommandHandler) HandleCreateTourCommand(command *CreateTourCommand) error {
	err := ch.validateCreateTourCommand(command)
	if err != nil {
		return myerrors.NewInvalidInputError(err)
	}

	// get tour based on year
	_, found := getTourOnYear(ch.store, command.Year)
	if found == true {
		return myerrors.NewInvalidInputErrorf(fmt.Sprintf("Tour %d already exists", command.Year))
	}

	// create event
	tourCreatedEvent := events.TourCreated{Year: command.Year}

	log.Printf("HandleCreateTourCommand completed:%+v -> %+v", command, tourCreatedEvent)

	// store and emit resulting event
	return ch.storeAndPublish([]*envelope.Envelope{tourCreatedEvent.Wrap()})
}

func (ch *TourCommandHandler) validateCreateCyclistCommand(command *CreateCyclistCommand) error {
	v := validation.Validator{}
	v.GreaterThan("Year", 2010, command.Year)
	v.GreaterThan("Id", 0, command.Id)
	v.NotEmpty("Name", command.Name)
	v.NotEmpty("Team", command.Team)
	return v.Err
}

func (ch *TourCommandHandler) HandleCreateCyclistCommand(command *CreateCyclistCommand) error {
	err := ch.validateCreateCyclistCommand(command)
	if err != nil {
		return myerrors.NewInvalidInputError(err)
	}

	// get tour based on year
	tour, found := getTourOnYear(ch.store, command.Year)
	if found == false {
		return myerrors.NewNotFoundErrorf(fmt.Sprintf("Tour %d does not exist", command.Year))
	}

	// verify if cyclist already exists
	if tour.hasCyclist(command.Id) {
		return myerrors.NewInvalidInputErrorf(fmt.Sprintf("Cyclist with %d already exists", command.Id))
	}

	// create event
	cyclistCreatedEvent := events.CyclistCreated{Year: command.Year,
		CyclistId:   command.Id,
		CyclistName: command.Name,
		CyclistTeam: command.Team}

	log.Printf("HandleCreateCyclistCommand completed:%+v -> %+v", command, cyclistCreatedEvent)

	// store and emit resulting event
	return ch.storeAndPublish([]*envelope.Envelope{cyclistCreatedEvent.Wrap()})
}

func (ch *TourCommandHandler) validateCreateEtappeCommand(command *CreateEtappeCommand) error {
	v := validation.Validator{}
	v.GreaterThan("Year", 2010, command.Year)
	v.GreaterThan("Id", 0, command.Id)
	v.NotEmpty("StartLocation", command.StartLocation)
	v.NotEmpty("FinishLocation", command.FinishLocation)
	v.GreaterThan("Length", 0, command.Length)
	v.GreaterThan("Kind", -1, command.Kind)
	v.After("Date", "2015-07-01T00:00:00Z", command.Date)

	return v.Err
}

func (ch *TourCommandHandler) HandleCreateEtappeCommand(command *CreateEtappeCommand) error {
	err := ch.validateCreateEtappeCommand(command)
	if err != nil {
		return myerrors.NewInvalidInputError(err)
	}

	// get tour based on year
	tour, found := getTourOnYear(ch.store, command.Year)
	if found == false {
		return myerrors.NewNotFoundError(errors.New(fmt.Sprintf("Tour %d does not exist", command.Year)))
	}

	// verify if etappe already exists
	if tour.hasEtappe(command.Id) {
		return myerrors.NewInvalidInputErrorf(fmt.Sprintf("Etappe with %d already exists", command.Id))
	}

	// create event
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

func (ch *TourCommandHandler) validateCreateEtappeResultsCommand(command *CreateEtappeResultsCommand) error {
	v := validation.Validator{}
	v.GreaterThan("Year", 2010, command.Year)
	v.GreaterThan("EtappeId", 0, command.EtappeId)
	v.SliceLength("BestDayCyclistIds", 10, command.BestDayCyclistIds)
	v.SliceLength("BestAllroundCyclistIds", 5, command.BestAllroundCyclistIds)
	v.SliceLength("BestClimbCyclistIds", 5, command.BestClimbCyclistIds)
	v.SliceLength("BestSprintCyclistIds", 5, command.BestSprintCyclistIds)

	return v.Err
}

func (ch *TourCommandHandler) HandleCreateEtappeResultsCommand(command *CreateEtappeResultsCommand) error {
	err := ch.validateCreateEtappeResultsCommand(command)
	if err != nil {
		return myerrors.NewInvalidInputError(err)
	}

	// get tour based on year
	tour, found := getTourOnYear(ch.store, command.Year)
	if found == false {
		return myerrors.NewNotFoundError(errors.New(fmt.Sprintf("Tour %d does not exist", command.Year)))
	}

	// verify that etappe already exists
	if tour.hasEtappe(command.EtappeId) == false {
		return myerrors.NewInvalidInputErrorf(fmt.Sprintf("Etappe %d does not exist", command.EtappeId))
	}

	// verify that referenced cyclists already exists
	for _, id := range command.BestDayCyclistIds {
		if tour.hasCyclist(id) == false {
			return myerrors.NewInvalidInputErrorf(fmt.Sprintf("BestDayCyclistIds: Cyclist %d does not exist", id))
		}
	}
	for _, id := range command.BestAllroundCyclistIds {
		if tour.hasCyclist(id) == false {
			return myerrors.NewInvalidInputErrorf(fmt.Sprintf("BestAllroundCyclistIds: Cyclist %d does not exist", id))
		}
	}
	for _, id := range command.BestSprintCyclistIds {
		if tour.hasCyclist(id) == false {
			return myerrors.NewInvalidInputErrorf(fmt.Sprintf("BestSprintCyclistIds: Cyclist %d does not exist", id))
		}
	}
	for _, id := range command.BestClimbCyclistIds {
		if tour.hasCyclist(id) == false {
			return myerrors.NewInvalidInputErrorf(fmt.Sprintf("BestClimbCyclistIds: Cyclist %d does not exist", id))
		}
	}

	//check for duplicate cyclists per category
	if len(uniq(command.BestDayCyclistIds)) < len(command.BestDayCyclistIds) {
		return myerrors.NewInvalidInputErrorf("BestDayCyclistIds contains duplicates")
	}
	if len(uniq(command.BestAllroundCyclistIds)) < len(command.BestAllroundCyclistIds) {
		return myerrors.NewInvalidInputErrorf("BestAllroundCyclistIds contains duplicates")
	}
	if len(uniq(command.BestSprintCyclistIds)) < len(command.BestSprintCyclistIds) {
		return myerrors.NewInvalidInputErrorf("BestSprintCyclistIds contains duplicates")
	}
	if len(uniq(command.BestClimbCyclistIds)) < len(command.BestClimbCyclistIds) {
		return myerrors.NewInvalidInputErrorf("BestClimbCyclistIds contains duplicates")
	}

	// compose event
	etappeResultCreatedEvent := events.EtappeResultsCreated{
		Year:                     command.Year,
		LastEtappeId:             command.EtappeId,
		BestDayCyclistIds:        command.BestDayCyclistIds,
		BestAllrounderCyclistIds: command.BestAllroundCyclistIds,
		BestSprinterCyclistIds:   command.BestSprintCyclistIds,
		BestClimberCyclistIds:    command.BestClimbCyclistIds}

	log.Printf("HandleCreateEtappeResultsCommand completed:%+v -> %+v", command, etappeResultCreatedEvent)

	// store and emit resulting event
	return ch.storeAndPublish([]*envelope.Envelope{etappeResultCreatedEvent.Wrap()})
}

func uniq(list []int) []int {
	unique_set := make(map[int]int, len(list))
	i := 0
	for _, x := range list {
		if _, there := unique_set[x]; !there {
			unique_set[x] = i
			i++
		}
	}
	result := make([]int, len(unique_set))
	for x, i := range unique_set {
		result[i] = x
	}
	return result
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
	Year     int          `json:"year"`
	Etappes  EtappeSlice  `json:"etappes"`
	Cyclists CyclistSlice `json:"cyclists"`
}

// +gen slice:"SortBy,Where,Select[string],GroupBy[string],Any,First"
type Cyclist struct {
	Number int    `json:"number"`
	Name   string `json:"name"`
	Team   string `json:"team"`
}

// +gen slice:"SortBy,Where,Select[string],Any,First"
type Etappe struct {
	Id             int       `json:"id"`
	Date           time.Time `json:"date"`
	StartLocation  string    `json:"startLocation"`
	FinishLocation string    `json:"finishLocation"`
	Length         int       `json:"length"`
	Kind           int       `json:"kind"`
	Results        *Result   `json:"results"`
}

type Result struct {
	BestDayCyclists        []*Cyclist
	BestAllrounderCyclists []*Cyclist
	BestSprinterCyclists   []*Cyclist
	BestClimberCyclists    []*Cyclist
}

func NewTour() *Tour {
	tour := new(Tour)
	tour.Etappes = make([]Etappe, 0, 30)
	tour.Cyclists = make([]Cyclist, 0, 250)
	return tour
}

func (t Tour) hasEtappe(id int) bool {
	return t.Etappes.Any(func(e Etappe) bool {
		return e.Id == id
	})
}

func (t Tour) findEtappe(id int) (*Etappe, bool) {
	found := false
	var etappe *Etappe = nil
	for _, e := range t.Etappes {
		if e.Id == id {
			etappe = &e
			found = true
			break
		}
	}
	return etappe, found
}

func (t Tour) hasCyclist(id int) bool {
	return t.Cyclists.Any(func(c Cyclist) bool {
		return c.Number == id
	})
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
	for idx, etappe := range t.Etappes {
		if etappe.Id == event.LastEtappeId {
			// access the slice directly otherwise settings pointer doesn't stick
			t.Etappes[idx].Results = &Result{
				BestDayCyclists:        t.CyclistsForIds(event.BestDayCyclistIds),
				BestAllrounderCyclists: t.CyclistsForIds(event.BestAllrounderCyclistIds),
				BestSprinterCyclists:   t.CyclistsForIds(event.BestSprinterCyclistIds),
				BestClimberCyclists:    t.CyclistsForIds(event.BestClimberCyclistIds)}
			break
		}
	}
	//log.Printf("ApplyEtappeResultsCreated after: %+v -> %+v", event, t)
}

func (t *Tour) CyclistsForIds(ids []int) []*Cyclist {
	cyclists := make([]*Cyclist, 0, len(ids))
	for _, id := range ids {
		c, err := t.Cyclists.First(func(c Cyclist) bool {
			return c.Number == id
		})
		if err == nil {
			cyclists = append(cyclists, &c)
		}
	}
	return cyclists
}
