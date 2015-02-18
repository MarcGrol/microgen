package tour

import (
	"github.com/MarcGrol/microgen/tourApp/events"
	"github.com/MarcGrol/microgen/tourApp/test"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestCreateTourCommand(t *testing.T) {
	scenario := test.Scenario{
		Name:        "NewTourSuccess",
		Description: "Create new tour on clean system",
		Given:       []events.Envelope{},
		When: func(scenario *test.Scenario) error {
			service := NewTourCommandHandler(scenario.Bus, scenario.Store)
			return service.HandleCreateTourCommand(CreateTourCommand{Year: 2015})
		},
		Expect: []events.Envelope{
			{
				SequenceNumber: 1, AggregateName: "tour", AggregateUid: "2015",
				Type:        events.TypeTourCreated,
				TourCreated: &events.TourCreated{Year: 2015},
			},
		},
	}

	scenario.RunAndVerify(t)

	assert.Equal(t, scenario.Expect[0].TourCreated.Year, scenario.Actual[0].TourCreated.Year)
}

func TestCreateCyclistCommand(t *testing.T) {
	scenario := test.Scenario{
		Name:        "NewCyclistSuccess",
		Description: "Create new cyclist with existing tour",
		Given: []events.Envelope{
			{
				Type:        events.TypeTourCreated,
				TourCreated: &events.TourCreated{Year: 2015},
			},
		},
		When: func(scenario *test.Scenario) error {
			service := NewTourCommandHandler(scenario.Bus, scenario.Store)
			return service.HandleCreateCyclistCommand(
				CreateCyclistCommand{
					Year: 2015,
					Id:   42,
					Name: "My name",
					Team: "My team"})
		},
		Expect: []events.Envelope{
			{
				SequenceNumber: 2, AggregateName: "tour", AggregateUid: "2015",
				Type: events.TypeCyclistCreated,
				CyclistCreated: &events.CyclistCreated{
					Year:        2015,
					CyclistId:   42,
					CyclistName: "My name",
					CyclistTeam: "My team"},
			},
		},
	}
	
	scenario.RunAndVerify(t)

	assert.Equal(t, scenario.Expect[0].CyclistCreated.Year, scenario.Actual[0].CyclistCreated.Year)
	assert.Equal(t, scenario.Expect[0].CyclistCreated.CyclistId, scenario.Actual[0].CyclistCreated.CyclistId)
	assert.Equal(t, scenario.Expect[0].CyclistCreated.CyclistName, scenario.Actual[0].CyclistCreated.CyclistName)
	assert.Equal(t, scenario.Expect[0].CyclistCreated.CyclistTeam, scenario.Actual[0].CyclistCreated.CyclistTeam)
}

func TestCreateEtappeCommand(t *testing.T) {
	scenario := test.Scenario{
		Name:        "NewEtappeSuccess",
		Description: "Create new etappe with existing tour",
		Given: []events.Envelope{
			{
				Type:        events.TypeTourCreated,
				TourCreated: &events.TourCreated{Year: 2015},
			},
		},
		When: func(scenario *test.Scenario) error {
			service := NewTourCommandHandler(scenario.Bus, scenario.Store)
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
		Expect: []events.Envelope{
			{
				SequenceNumber: 2, AggregateName: "tour", AggregateUid: "2015",
				Type: events.TypeEtappeCreated,
				EtappeCreated: &events.EtappeCreated{
					Year:                 2015,
					EtappeId:             2,
					EtappeDate:           time.Date(2015, time.July, 14, 9, 0, 0, 0, time.Local),
					EtappeStartLocation:  "Parijs",
					EtappeLength:         255,
					EtappeFinishLocation: "Roubaix",
					EtappeKind:           3},
			},
		},
	}
	
	scenario.RunAndVerify(t)

	assert.Equal(t, scenario.Expect[0].EtappeCreated.Year, scenario.Actual[0].EtappeCreated.Year)
	assert.Equal(t, scenario.Expect[0].EtappeCreated.EtappeId, scenario.Actual[0].EtappeCreated.EtappeId)
	assert.Equal(t, scenario.Expect[0].EtappeCreated.EtappeDate.Year(), scenario.Actual[0].EtappeCreated.EtappeDate.Year())
	assert.Equal(t, scenario.Expect[0].EtappeCreated.EtappeDate.Month(), scenario.Actual[0].EtappeCreated.EtappeDate.Month())
	assert.Equal(t, scenario.Expect[0].EtappeCreated.EtappeDate.Day(), scenario.Actual[0].EtappeCreated.EtappeDate.Day())
	assert.Equal(t, scenario.Expect[0].EtappeCreated.EtappeStartLocation, scenario.Actual[0].EtappeCreated.EtappeStartLocation)
	assert.Equal(t, scenario.Expect[0].EtappeCreated.EtappeFinishLocation, scenario.Actual[0].EtappeCreated.EtappeFinishLocation)
	assert.Equal(t, scenario.Expect[0].EtappeCreated.EtappeLength, scenario.Actual[0].EtappeCreated.EtappeLength)
	assert.Equal(t, scenario.Expect[0].EtappeCreated.EtappeKind, scenario.Actual[0].EtappeCreated.EtappeKind)
}
