package tour

//go:generate gen

import (
	"strconv"
	"time"

	"github.com/MarcGrol/microgen/infra"
	"github.com/MarcGrol/microgen/lib/envelope"
	"github.com/MarcGrol/microgen/tourApp/events"
)

type EtappeKind int

const (
	Flat = 1 + iota
	Hilly
	Mountains
	TimeTrial
)

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
	for idx, e := range t.Etappes {
		if e.Id == id {
			// access the slice directly otherwise settings pointer doesn't stick
			return &t.Etappes[idx], true
		}
	}
	return nil, false
}

func (t Tour) hasCyclist(id int) bool {
	return t.Cyclists.Any(func(c Cyclist) bool {
		return c.Number == id
	})
}

func (t *Tour) ApplyAll(envelopes []envelope.Envelope) {
	applyEvents(envelopes, t)
}

func (t *Tour) ApplyTourCreated(event *events.TourCreated) {

	t.Year = event.Year

	//log.Printf("ApplyTourCreated after:%+v -> %+v", event, t)

	return
}

func (t *Tour) ApplyCyclistCreated(event *events.CyclistCreated) {

	t.Cyclists = append(t.Cyclists,
		Cyclist{
			Number: event.CyclistId,
			Name:   event.CyclistName,
			Team:   event.CyclistTeam})

	//log.Printf("ApplyCyclistCreated after:%+v -> %+v", event, t)

	return
}

func (t *Tour) ApplyEtappeCreated(event *events.EtappeCreated) {
	t.Etappes = append(t.Etappes,
		Etappe{
			Id:             event.EtappeId,
			Date:           event.EtappeDate,
			StartLocation:  event.EtappeStartLocation,
			FinishLocation: event.EtappeFinishLocation,
			Length:         event.EtappeLength,
			Kind:           event.EtappeKind})

	//log.Printf("ApplyEtappeCreated after:%+v -> %+v", event, t)

	return
}

func (t *Tour) ApplyEtappeResultsCreated(event *events.EtappeResultsCreated) {
	etappe, found := t.findEtappe(event.LastEtappeId)
	if found {
		etappe.Results = &Result{
			BestDayCyclists:        t.cyclistsForIds(event.BestDayCyclistIds),
			BestAllrounderCyclists: t.cyclistsForIds(event.BestAllrounderCyclistIds),
			BestSprinterCyclists:   t.cyclistsForIds(event.BestSprinterCyclistIds),
			BestClimberCyclists:    t.cyclistsForIds(event.BestClimberCyclistIds)}
	}
	//log.Printf("ApplyEtappeResultsCreated after: %+v -> %+v", event, t)
}

func (t *Tour) cyclistsForIds(ids []int) []*Cyclist {
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
