package gambler

import (
	"log"
	"strconv"

	"github.com/MarcGrol/microgen/infra"
	"github.com/MarcGrol/microgen/lib/envelope"
	"github.com/MarcGrol/microgen/tourApp/events"
)

//go:generate gen

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

func (context *GamblingContext) ApplyAll(envelopes []envelope.Envelope) {
	applyEvents(envelopes, context)
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
