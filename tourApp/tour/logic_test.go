package tour

import (
	"github.com/MarcGrol/microgen/tourApp/events"
	"github.com/MarcGrol/microgen/tourApp/test"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestCreateTourCommand(t *testing.T) {
	var service CommandHandler
	scenario := test.Scenario{
		Title: "Create new tour on clean system",
		Given: []*events.Envelope{},
		When: func(scenario *test.Scenario) error {
			service = NewTourCommandHandler(scenario.Bus, scenario.Store)
			return service.HandleCreateTourCommand(CreateTourCommand{Year: 2015})
		},
		Expect: []*events.Envelope{
			(&events.TourCreated{Year: 2015}).Wrap(),
		},
	}

	scenario.RunAndVerify(t)

	expected := scenario.Expect[0].TourCreated
	actual := scenario.Actual[0].TourCreated
	assert.Equal(t, expected.Year, actual.Year)

	// Test query
	tourOpaque, err := service.HandleGetTourQuery(expected.Year)
	assert.Nil(t, err)
	tour, ok := tourOpaque.(*Tour)
	assert.True(t, ok)
	assert.Equal(t, expected.Year, tour.Year)
	assert.Equal(t, 0, len(tour.Etappes))
	assert.Equal(t, 0, len(tour.Cyclists))
}

func TestCreateCyclistCommand(t *testing.T) {
	var service CommandHandler
	scenario := test.Scenario{
		Title: "Create new cyclist with existing tour",
		Given: []*events.Envelope{
			(&events.TourCreated{Year: 2015}).Wrap(),
		},
		When: func(scenario *test.Scenario) error {
			service = NewTourCommandHandler(scenario.Bus, scenario.Store)
			return service.HandleCreateCyclistCommand(
				CreateCyclistCommand{
					Year: 2015,
					Id:   42,
					Name: "My name",
					Team: "My team"})
		},
		Expect: []*events.Envelope{
			(&events.CyclistCreated{
				Year:        2015,
				CyclistId:   42,
				CyclistName: "My name",
				CyclistTeam: "My team"}).Wrap(),
		},
	}

	scenario.RunAndVerify(t)

	expected := scenario.Expect[0].CyclistCreated
	actual := scenario.Actual[0].CyclistCreated
	assert.Equal(t, expected.Year, actual.Year)
	assert.Equal(t, expected.CyclistId, actual.CyclistId)
	assert.Equal(t, expected.CyclistName, actual.CyclistName)
	assert.Equal(t, expected.CyclistTeam, actual.CyclistTeam)

	// Test query
	tourOpaque, err := service.HandleGetTourQuery(expected.Year)
	assert.Nil(t, err)
	tour, ok := tourOpaque.(*Tour)
	assert.True(t, ok)
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
	scenario := test.Scenario{
		Title: "Create new etappe with existing tour",
		Given: []*events.Envelope{
			(&events.TourCreated{Year: 2015}).Wrap(),
		},
		When: func(scenario *test.Scenario) error {
			service = NewTourCommandHandler(scenario.Bus, scenario.Store)
			return service.HandleCreateEtappeCommand(
				CreateEtappeCommand{
					Year:           2015,
					Id:             2,
					Date:           time.Date(2015, time.July, 14, 9, 0, 0, 0, time.Local),
					StartLocation:  "Parijs",
					FinishLocation: "Roubaix",
					Length:         255,
					Kind:           3})
		},
		Expect: []*events.Envelope{
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

	expected := scenario.Expect[0].EtappeCreated
	actual := scenario.Actual[0].EtappeCreated
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
	tourOpaque, err := service.HandleGetTourQuery(expected.Year)
	assert.Nil(t, err)
	tour, ok := tourOpaque.(*Tour)
	assert.True(t, ok)
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
