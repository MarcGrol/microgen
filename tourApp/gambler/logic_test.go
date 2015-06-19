package gambler

import (
	"testing"
	"time"

	"github.com/MarcGrol/microgen/lib/envelope"
	"github.com/MarcGrol/microgen/lib/test"
	"github.com/MarcGrol/microgen/tourApp/events"
	"github.com/stretchr/testify/assert"
)

func TestTourCreatedEvent(t *testing.T) {
	var service EventHandler
	input := (&events.TourCreated{Year: 2015}).Wrap()

	scenario := test.EventScenario{
		Title:   "Handle tour created event",
		Given:   []*envelope.Envelope{},
		Envelop: input,
		When: func(scenario *test.EventScenario) error {
			service = NewGamblerEventHandler(scenario.Bus, scenario.Store)
			return service.OnEvent(scenario.Envelop)
		},
		Expect: []*envelope.Envelope{input},
	}

	scenario.RunAndVerify(t)

	assert.Nil(t, scenario.ErrMsg)

}

func TestCyclistCreatedEvent(t *testing.T) {
	var service EventHandler
	given := (&events.TourCreated{Year: 2015}).Wrap()
	input := (&events.CyclistCreated{Year: 2015, CyclistId: 1, CyclistName: "Lance", CyclistTeam: "Shack"}).Wrap()

	scenario := test.EventScenario{
		Title:   "Handle cyclist created event",
		Given:   []*envelope.Envelope{given},
		Envelop: input,
		When: func(scenario *test.EventScenario) error {
			service = NewGamblerEventHandler(scenario.Bus, scenario.Store)
			return service.OnEvent(scenario.Envelop)
		},
		Expect: []*envelope.Envelope{given, input},
	}

	scenario.RunAndVerify(t)

	assert.Nil(t, scenario.ErrMsg)
}

func TestEtappeCreatedEvent(t *testing.T) {
	var service EventHandler
	given := (&events.TourCreated{Year: 2015}).Wrap()
	input := (&events.EtappeCreated{
		Year:                 2015,
		EtappeId:             1,
		EtappeDate:           time.Now(),
		EtappeStartLocation:  "Luik",
		EtappeFinishLocation: "Bastenaken",
		EtappeLength:         256,
		EtappeKind:           1}).Wrap()

	scenario := test.EventScenario{
		Title:   "Handle etappe created event",
		Given:   []*envelope.Envelope{given},
		Envelop: input,
		When: func(scenario *test.EventScenario) error {
			service = NewGamblerEventHandler(scenario.Bus, scenario.Store)
			return service.OnEvent(scenario.Envelop)
		},
		Expect: []*envelope.Envelope{given, input},
	}

	scenario.RunAndVerify(t)

	assert.Nil(t, scenario.ErrMsg)
}

func TestEtappeResultsEvent(t *testing.T) {
	var service EventHandler
	givenTour := (&events.TourCreated{Year: 2015}).Wrap()

	givenEtappe := (&events.EtappeCreated{
		Year:                 2015,
		EtappeId:             1,
		EtappeDate:           time.Now(),
		EtappeStartLocation:  "Luik",
		EtappeFinishLocation: "Bastenaken",
		EtappeLength:         256,
		EtappeKind:           1}).Wrap()

	givenCyclist1 := (&events.CyclistCreated{
		Year:        2015,
		CyclistId:   1,
		CyclistName: "Lance",
		CyclistTeam: "Shack"}).Wrap()

	givenCyclist2 := (&events.CyclistCreated{
		Year:        2015,
		CyclistId:   2,
		CyclistName: "Boogerd",
		CyclistTeam: "Rabo"}).Wrap()

	givenCyclist3 := (&events.CyclistCreated{
		Year:        2015,
		CyclistId:   3,
		CyclistName: "Pantani",
		CyclistTeam: "Lampre"}).Wrap()

	input := (&events.EtappeResultsCreated{
		Year:                     2015,
		LastEtappeId:             3,
		BestDayCyclistIds:        []int{1, 2},
		BestAllrounderCyclistIds: []int{1, 2, 3},
		BestSprinterCyclistIds:   []int{3, 2, 1},
		BestClimberCyclistIds:    []int{3, 2},
	}).Wrap()

	scenario := test.EventScenario{
		Title:   "Handle etappe results created event",
		Given:   []*envelope.Envelope{givenTour, givenEtappe, givenCyclist1, givenCyclist2, givenCyclist3},
		Envelop: input,
		When: func(scenario *test.EventScenario) error {
			service = NewGamblerEventHandler(scenario.Bus, scenario.Store)
			return service.OnEvent(scenario.Envelop)
		},
		Expect: []*envelope.Envelope{givenTour, givenEtappe, givenCyclist1, givenCyclist2, givenCyclist3, input},
	}

	scenario.RunAndVerify(t)

	assert.Nil(t, scenario.ErrMsg)
}

func TestCreateGamblerCommand(t *testing.T) {
	var service CommandHandler
	scenario := test.CommandScenario{
		Title: "Create new gambler success",
		Given: []*envelope.Envelope{
			(&events.TourCreated{Year: 2015}).Wrap(),
		},
		Command: &CreateGamblerCommand{GamblerUid: "my uid", Name: "My name", Email: "me@home.nl"},
		When: func(scenario *test.CommandScenario) error {
			service = NewGamblerCommandHandler(scenario.Bus, scenario.Store)
			return service.HandleCreateGamblerCommand(scenario.Command.(*CreateGamblerCommand))
		},
		Expect: []*envelope.Envelope{
			(&events.GamblerCreated{GamblerUid: "my uid", GamblerName: "My name", GamblerEmail: "me@home.nl"}).Wrap(),
		},
	}

	scenario.RunAndVerify(t)

	assert.Nil(t, scenario.ErrMsg)

	expected, ok := events.GetIfIsGamblerCreated(scenario.Expect[0])
	assert.True(t, ok)
	assert.NotNil(t, expected)
	actual, ok := events.GetIfIsGamblerCreated(&scenario.Actual[0])
	assert.True(t, ok)
	assert.NotNil(t, actual)
	assert.Equal(t, expected.GamblerUid, actual.GamblerUid)
	assert.Equal(t, expected.GamblerName, actual.GamblerName)
	assert.Equal(t, expected.GamblerEmail, actual.GamblerEmail)

	// Test query
	gambler, err := service.HandleGetGamblerQuery(expected.GamblerUid, -1)
	assert.Nil(t, err)
	assert.NotNil(t, gambler)
	assert.Equal(t, expected.GamblerUid, gambler.Uid)
	assert.Equal(t, expected.GamblerName, gambler.Name)
	assert.Equal(t, expected.GamblerEmail, gambler.Email)
}

func TestCreateGamblerTeamCommand(t *testing.T) {
	var service CommandHandler
	scenario := test.CommandScenario{
		Title: "Create new gambler team success",
		Given: []*envelope.Envelope{
			(&events.TourCreated{Year: 2015}).Wrap(),
			(&events.CyclistCreated{Year: 2015, CyclistId: 1, CyclistName: "cyclist 1", CyclistTeam: "team 1"}).Wrap(),
			(&events.CyclistCreated{Year: 2015, CyclistId: 2, CyclistName: "cyclist 2", CyclistTeam: "team 2"}).Wrap(),
			(&events.GamblerCreated{GamblerUid: "my uid", GamblerName: "My name", GamblerEmail: "me@home.nl"}).Wrap(),
		},
		Command: &CreateGamblerTeamCommand{GamblerUid: "my uid", Year: 2015, CyclistIds: []int{1, 2}},
		When: func(scenario *test.CommandScenario) error {
			service = NewGamblerCommandHandler(scenario.Bus, scenario.Store)
			return service.HandleCreateGamblerTeamCommand(scenario.Command.(*CreateGamblerTeamCommand))
		},
		Expect: []*envelope.Envelope{
			(&events.GamblerTeamCreated{GamblerUid: "my uid", Year: 2015, GamblerCyclists: []int{1, 2}}).Wrap(),
		},
	}

	scenario.RunAndVerify(t)

	assert.Nil(t, scenario.ErrMsg)

	expected := events.UnWrapGamblerTeamCreated(scenario.Expect[0])
	actual := events.UnWrapGamblerTeamCreated(&scenario.Actual[0])
	assert.Equal(t, expected.Year, actual.Year)
	assert.Equal(t, expected.GamblerUid, actual.GamblerUid)
	assert.Equal(t, 2, len(actual.GamblerCyclists))

	// Test query
	gambler, err := service.HandleGetGamblerQuery(expected.GamblerUid, expected.Year)
	assert.Nil(t, err)
	assert.NotNil(t, gambler)
	assert.Equal(t, expected.GamblerUid, gambler.Uid)
	assert.Equal(t, "My name", gambler.Name)
	assert.Equal(t, "me@home.nl", gambler.Email)

	assert.Equal(t, 2, len(gambler.Cyclists))
	assert.Equal(t, 1, gambler.Cyclists[0].Id)
	assert.Equal(t, "cyclist 1", gambler.Cyclists[0].Name)
	assert.Equal(t, "team 1", gambler.Cyclists[0].Team)
	assert.Equal(t, 2, gambler.Cyclists[1].Id)
	assert.Equal(t, "cyclist 2", gambler.Cyclists[1].Name)
	assert.Equal(t, "team 2", gambler.Cyclists[1].Team)

}
