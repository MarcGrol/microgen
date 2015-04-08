package tour

import (
	"testing"
	"time"

	"github.com/MarcGrol/microgen/lib/envelope"
	"github.com/MarcGrol/microgen/lib/test"
	"github.com/MarcGrol/microgen/tourApp/events"
	"github.com/stretchr/testify/assert"
)

func TestCreateTourCommand(t *testing.T) {
	var service CommandHandler
	scenario := test.CommandScenario{
		Title:   "Create new tour success",
		Given:   []*envelope.Envelope{},
		Command: &CreateTourCommand{Year: 2015},
		When: func(scenario *test.CommandScenario) error {
			service = NewTourCommandHandler(scenario.Bus, scenario.Store)
			return service.HandleCreateTourCommand(scenario.Command.(*CreateTourCommand))
		},
		Expect: []*envelope.Envelope{
			(&events.TourCreated{Year: 2015}).Wrap(),
		},
	}

	scenario.RunAndVerify(t)

	assert.Nil(t, scenario.ErrMsg)

	expected := events.UnWrapTourCreated(scenario.Expect[0])
	actual := events.UnWrapTourCreated(&scenario.Actual[0])
	assert.Equal(t, expected.Year, actual.Year)

	// Test query
	tour, err := service.HandleGetTourQuery(expected.Year)
	assert.Nil(t, err)
	assert.NotNil(t, tour)
	assert.Equal(t, expected.Year, tour.Year)
	assert.Equal(t, 0, len(tour.Etappes))
	assert.Equal(t, 0, len(tour.Cyclists))
}

func TestCreateCyclistCommand(t *testing.T) {
	var service CommandHandler
	scenario := test.CommandScenario{
		Title: "Create new cyclist success",
		Given: []*envelope.Envelope{
			(&events.TourCreated{Year: 2015}).Wrap(),
		},
		Command: &CreateCyclistCommand{
			Year: 2015,
			Id:   42,
			Name: "My name",
			Team: "My team"},
		When: func(scenario *test.CommandScenario) error {
			service = NewTourCommandHandler(scenario.Bus, scenario.Store)
			return service.HandleCreateCyclistCommand(scenario.Command.(*CreateCyclistCommand))
		},
		Expect: []*envelope.Envelope{
			(&events.CyclistCreated{
				Year:        2015,
				CyclistId:   42,
				CyclistName: "My name",
				CyclistTeam: "My team"}).Wrap(),
		},
	}

	scenario.RunAndVerify(t)

	assert.Nil(t, scenario.ErrMsg)

	expected := events.UnWrapCyclistCreated(scenario.Expect[0])
	actual := events.UnWrapCyclistCreated(&scenario.Actual[0])
	assert.Equal(t, expected.Year, actual.Year)
	assert.Equal(t, expected.CyclistId, actual.CyclistId)
	assert.Equal(t, expected.CyclistName, actual.CyclistName)
	assert.Equal(t, expected.CyclistTeam, actual.CyclistTeam)

	// Test query
	tour, err := service.HandleGetTourQuery(expected.Year)
	assert.Nil(t, err)
	assert.NotNil(t, tour)
	assert.Equal(t, 2015, tour.Year)
	assert.Equal(t, 0, len(tour.Etappes))
	assert.Equal(t, 1, len(tour.Cyclists))
	assert.Equal(t, expected.Year, tour.Year)
	assert.Equal(t, expected.CyclistId, tour.Cyclists[0].Number)
	assert.Equal(t, expected.CyclistName, tour.Cyclists[0].Name)
	assert.Equal(t, expected.CyclistTeam, tour.Cyclists[0].Team)

}

func TestCreateEtappeCommand(t *testing.T) {
	var service CommandHandler
	scenario := test.CommandScenario{
		Title: "Create new etappe success",
		Given: []*envelope.Envelope{
			(&events.TourCreated{Year: 2015}).Wrap(),
		},
		Command: &CreateEtappeCommand{
			Year:           2015,
			Id:             2,
			Date:           time.Date(2015, time.July, 14, 9, 0, 0, 0, time.Local),
			StartLocation:  "Parijs",
			FinishLocation: "Roubaix",
			Length:         255,
			Kind:           3},
		When: func(scenario *test.CommandScenario) error {
			service = NewTourCommandHandler(scenario.Bus, scenario.Store)
			return service.HandleCreateEtappeCommand(scenario.Command.(*CreateEtappeCommand))
		},
		Expect: []*envelope.Envelope{
			(&events.EtappeCreated{
				Year:                 2015,
				EtappeId:             2,
				EtappeDate:           time.Date(2015, time.July, 14, 9, 0, 0, 0, time.Local),
				EtappeStartLocation:  "Parijs",
				EtappeFinishLocation: "Roubaix",
				EtappeLength:         255,
				EtappeKind:           3}).Wrap(),
		},
	}

	scenario.RunAndVerify(t)

	assert.Nil(t, scenario.ErrMsg)

	expected := events.UnWrapEtappeCreated(scenario.Expect[0])
	actual := events.UnWrapEtappeCreated(&scenario.Actual[0])
	assert.Equal(t, expected.Year, actual.Year)
	assert.Equal(t, expected.EtappeId, actual.EtappeId)
	assert.Equal(t, expected.EtappeDate.Year(), actual.EtappeDate.Year())
	assert.Equal(t, expected.EtappeDate.Month(), actual.EtappeDate.Month())
	assert.Equal(t, expected.EtappeDate.Day(), actual.EtappeDate.Day())
	assert.Equal(t, expected.EtappeStartLocation, actual.EtappeStartLocation)
	assert.Equal(t, expected.EtappeFinishLocation, actual.EtappeFinishLocation)
	assert.Equal(t, expected.EtappeLength, actual.EtappeLength)
	assert.Equal(t, expected.EtappeKind, actual.EtappeKind)

	// Test query
	tour, err := service.HandleGetTourQuery(expected.Year)
	assert.Nil(t, err)
	assert.NotNil(t, tour)
	assert.Equal(t, 2015, tour.Year)
	assert.Equal(t, 1, len(tour.Etappes))
	assert.Equal(t, 0, len(tour.Cyclists))
	assert.Equal(t, expected.EtappeId, tour.Etappes[0].Id)
	assert.Equal(t, expected.EtappeDate.Year(), tour.Etappes[0].Date.Year())
	assert.Equal(t, expected.EtappeDate.Month(), tour.Etappes[0].Date.Month())
	assert.Equal(t, expected.EtappeDate.Day(), tour.Etappes[0].Date.Day())
	assert.Equal(t, expected.EtappeStartLocation, tour.Etappes[0].StartLocation)
	assert.Equal(t, expected.EtappeFinishLocation, tour.Etappes[0].FinishLocation)
	assert.Equal(t, expected.EtappeLength, tour.Etappes[0].Length)
	assert.Equal(t, expected.EtappeKind, tour.Etappes[0].Kind)
}
