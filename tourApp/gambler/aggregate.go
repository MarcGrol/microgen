package gambler

import (
	"log"

	"github.com/MarcGrol/microgen/lib/envelope"
	"github.com/MarcGrol/microgen/tourApp/events"
)

//go:generate gen

type Classement int

const (
	ClassementUnknown Classement = iota
	ClassementDay
	ClassementAllround
	ClassementSprint
	ClassementClimb
)

type GamblingContext struct {
	years    map[int]*GamblingYear
	gamblers map[string]Gambler
}

func NewGamblingContext() *GamblingContext {
	return &GamblingContext{
		years:    make(map[int]*GamblingYear),
		gamblers: make(map[string]Gambler),
	}
}

type GamblingYear struct {
	Year            int
	cyclistsForTour map[int]*Cyclist
	etappes         map[int]*Etappe
	gamblersForTour map[string]*Gambler
}

func NewGamblingYear(year int) *GamblingYear {
	return &GamblingYear{
		Year:            year,
		cyclistsForTour: make(map[int]*Cyclist),
		etappes:         make(map[int]*Etappe),
		gamblersForTour: make(map[string]*Gambler),
	}
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
	return &Gambler{
		Uid:      uid,
		Name:     name,
		Email:    email,
		Cyclists: make([]*Cyclist, 0, 10),
	}
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

func (context *GamblingContext) ApplyAll(envelopes []envelope.Envelope) {
	applyEvents(envelopes, context)
}

func (context *GamblingContext) ApplyTourCreated(event *events.TourCreated) {
	log.Printf("ApplyTourCreated: context before: %+v, event: %+v", context, event)

	_, exists := context.years[event.Year]
	if exists == false {
		context.years[event.Year] = NewGamblingYear(event.Year)

		log.Printf("ApplyTourCreated: context after: %+v", context)
	}

	return
}

func (context *GamblingContext) ApplyCyclistCreated(event *events.CyclistCreated) {
	log.Printf("ApplyCyclistCreated: context before: %+v, event: %+v", context, event)

	gamblingYear, exists := context.years[event.Year]
	if exists {
		gamblingYear.cyclistsForTour[event.CyclistId] =
			&Cyclist{
				Id:   event.CyclistId,
				Name: event.CyclistName,
				Team: event.CyclistTeam}

		log.Printf("ApplyCyclistCreated: context after: %+v", context)
	}

	return
}

func (context *GamblingContext) ApplyGamblerCreated(event *events.GamblerCreated) {
	log.Printf("ApplyGamblerCreated: context before: %+v, event: %+v", context, event)

	_, exists := context.gamblers[event.GamblerUid]
	if exists == false {
		context.gamblers[event.GamblerUid] =
			*NewGambler(event.GamblerUid, event.GamblerName, event.GamblerEmail)

		log.Printf("ApplyGamblerCreated: context after: %+v", context)
	}

	return
}

func (context *GamblingContext) ApplyGamblerTeamCreated(event *events.GamblerTeamCreated) {
	log.Printf("ApplyGamblerTeamCreated: context: %+v, event: %+v", context, event)

	gamblingYear, exists := context.years[event.Year]
	if exists {
		basicGambler, exists := context.gamblers[event.GamblerUid]
		if exists {
			yearGambler := &Gambler{
				Uid:      basicGambler.Uid,
				Name:     basicGambler.Name,
				Email:    basicGambler.Email,
				Points:   0,
				Cyclists: make([]*Cyclist, 0, len(event.GamblerCyclists)),
			}
			gamblingYear.gamblersForTour[event.GamblerUid] = yearGambler
			for _, cyclistId := range event.GamblerCyclists {
				cyclist, found := gamblingYear.cyclistsForTour[cyclistId]
				log.Printf("cyclist: %+v", cyclist)
				if found {
					yearGambler.Cyclists = append(yearGambler.Cyclists, cyclist)
				}
			}
			log.Printf("ApplyGamblerTeamCreated: context after: %+v", context)
		}
	}
	return
}

func (context *GamblingContext) ApplyEtappeCreated(event *events.EtappeCreated) {
	gamblingYear, exists := context.years[event.Year]
	if exists {

		gamblingYear.etappes[event.EtappeId] = &Etappe{
			id:   event.EtappeId,
			kind: event.EtappeKind,
		}

		log.Printf("ApplyEtappeCreated: context after: %+v", context)
	}

	return
}

func (context *GamblingContext) ApplyEtappeResultsCreated(event *events.EtappeResultsCreated) {

	log.Printf("ApplyEtappeResultsCreated: event: %+v, context before: %+v", event, context)

	gamblingYear, exists := context.years[event.Year]
	if exists {
		etappe, found := gamblingYear.etappes[event.LastEtappeId]
		if found {
			gamblingYear.calculateCyclistPointsForEtappe(etappe, event.BestDayCyclistIds, ClassementDay)
			gamblingYear.calculateCyclistPointsForEtappe(etappe, event.BestAllrounderCyclistIds, ClassementAllround)
			gamblingYear.calculateCyclistPointsForEtappe(etappe, event.BestSprinterCyclistIds, ClassementSprint)
			gamblingYear.calculateCyclistPointsForEtappe(etappe, event.BestClimberCyclistIds, ClassementClimb)

			/*
			 * calculate points for gamblers
			 */
			for _, gambler := range gamblingYear.gamblersForTour {
				for _, cyclist := range gambler.Cyclists {
					gambler.Points += cyclist.Points
				}
			}

		}
		log.Printf("ApplyEtappeResultsCreated: context after: %+v", context)
	}
}

func (gamblingYear *GamblingYear) calculateCyclistPointsForEtappe(etappe *Etappe, cyclistIds []int, classementType Classement) {
	for rank, cyclistId := range cyclistIds {
		cyclist, found := gamblingYear.cyclistsForTour[cyclistId]
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
