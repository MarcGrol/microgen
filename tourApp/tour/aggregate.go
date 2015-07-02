package tour

//go:generate gen

import (
	"log"
	"time"

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

type tourContext struct {
	tours map[int]*Tour
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

func newTourContext() *tourContext {
	return &tourContext{
		tours: make(map[int]*Tour),
	}
}

func (tc *tourContext) ApplyAll(envelopes []envelope.Envelope) {
	applyEvents(envelopes, tc)
}

func (tc *tourContext) ApplyTourCreated(event *events.TourCreated) {
	_, exists := tc.tours[event.Year]
	if exists == false {
		tc.tours[event.Year] = NewTour(event.Year)

		log.Printf("ApplyTourCreated after:%+v -> %+v", event, tc)
	}

	return
}

func (tc *tourContext) ApplyCyclistCreated(event *events.CyclistCreated) {
	tour, exists := tc.tours[event.Year]
	if exists {
		tour.Cyclists = append(tour.Cyclists,
			Cyclist{
				Number: event.CyclistId,
				Name:   event.CyclistName,
				Team:   event.CyclistTeam})

		log.Printf("ApplyCyclistCreated after:%+v -> %+v", event, tour)
	}

	return
}

func (tc *tourContext) ApplyEtappeCreated(event *events.EtappeCreated) {
	tour, exists := tc.tours[event.Year]
	if exists {
		tour.Etappes = append(tour.Etappes,
			Etappe{
				Id:             event.EtappeId,
				Date:           event.EtappeDate,
				StartLocation:  event.EtappeStartLocation,
				FinishLocation: event.EtappeFinishLocation,
				Length:         event.EtappeLength,
				Kind:           event.EtappeKind})

		log.Printf("ApplyEtappeCreated after:%+v -> %+v", event, tour)
	}

	return
}

func (tc *tourContext) ApplyEtappeResultsCreated(event *events.EtappeResultsCreated) {
	tour, exists := tc.tours[event.Year]
	if exists {
		etappe, found := tour.findEtappe(event.LastEtappeId)
		if found {
			etappe.Results = &Result{
				BestDayCyclists:        tour.cyclistsForIds(event.BestDayCyclistIds),
				BestAllrounderCyclists: tour.cyclistsForIds(event.BestAllrounderCyclistIds),
				BestSprinterCyclists:   tour.cyclistsForIds(event.BestSprinterCyclistIds),
				BestClimberCyclists:    tour.cyclistsForIds(event.BestClimberCyclistIds)}

			log.Printf("ApplyEtappeResultsCreated after: %+v -> %+v", event, tour)
		}
	}
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

func NewTour(year int) *Tour {
	tour := new(Tour)
	tour.Year = year
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
