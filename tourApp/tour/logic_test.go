package tour

import (
	"github.com/MarcGrol/microgen/tourApp/events"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestCreateTourCommand(t *testing.T) {
	operationOnSubject := func(scenario *Scenario) error {
		service := NewTourCommandHandler(scenario.bus, scenario.store)
		return service.HandleCreateTourCommand(CreateTourCommand{Year: 2015})
	}
	scenario := Scenario{
		Name:        "NewTourSuccess",
		Description: "Create new tour on clean system",
		Given:       []events.Envelope{},
		CallSubject: operationOnSubject,
		Expect: []events.Envelope{
			{
				SequenceNumber: 1, AggregateName: "tour", AggregateUid: "2015", Type: events.TypeTourCreated,
				TourCreated: &events.TourCreated{Year: 2015},
			},
		},
	}

	scenario.RunAndVerify(t)

	assert.Equal(t, 2015, scenario.Actual[0].TourCreated.Year)
}

func TestCreateCyclistCommand(t *testing.T) {
	operationOnSubject := func(scenario *Scenario) error {
		service := NewTourCommandHandler(scenario.bus, scenario.store)
		return service.HandleCreateCyclistCommand(CreateCyclistCommand{Year: 2015, Id: 42, Name: "My name", Team: "My team"})
	}
	scenario := Scenario{
		Name:        "NewCyclistSuccess",
		Description: "Create new cyclist with existing tour",
		Given: []events.Envelope{
			{
				Type:        events.TypeTourCreated,
				TourCreated: &events.TourCreated{Year: 2015},
			},
		},
		CallSubject: operationOnSubject,
		Expect: []events.Envelope{
			{
				SequenceNumber: 2, AggregateName: "tour", AggregateUid: "2015", Type: events.TypeCyclistCreated,
				CyclistCreated: &events.CyclistCreated{Year: 2015},
			},
		},
	}
	scenario.RunAndVerify(t)

	assert.Equal(t, 2015, scenario.Actual[0].CyclistCreated.Year)
	assert.Equal(t, 42, scenario.Actual[0].CyclistCreated.CyclistId)
	assert.Equal(t, "My name", scenario.Actual[0].CyclistCreated.CyclistName)
	assert.Equal(t, "My team", scenario.Actual[0].CyclistCreated.CyclistTeam)
}

func TestCreateEtappeCommand(t *testing.T) {
	operationOnSubject := func(scenario *Scenario) error {
		service := NewTourCommandHandler(scenario.bus, scenario.store)
		return service.HandleCreateEtappeCommand(CreateEtappeCommand{Year: 2015, Id:2, Date:time.Date(2015, time.July, 14, 9, 0, 0, 0, time.Local),StartLocation:"Parijs",FinishLocation:"Roubaix",Length:255,Kind:3 })
	}
	scenario := Scenario{
		Name:        "NewEtappeSuccess",
		Description: "Create new etappe with existing tour",
		Given: []events.Envelope{
			{
				Type:        events.TypeTourCreated,
				TourCreated: &events.TourCreated{Year: 2015},
			},
		},
		CallSubject: operationOnSubject,
		Expect: []events.Envelope{
			{
				SequenceNumber: 2, AggregateName: "tour", AggregateUid: "2015", Type: events.TypeEtappeCreated,
				EtappeCreated: &events.EtappeCreated{Year: 2015},
			},
		},
	}
	scenario.RunAndVerify(t)

	assert.Equal(t, 2015, scenario.Actual[0].EtappeCreated.Year)
	assert.Equal(t, 2015, scenario.Actual[0].EtappeCreated.EtappeDate.Year())
	assert.Equal(t, time.July, scenario.Actual[0].EtappeCreated.EtappeDate.Month())
	assert.Equal(t, 14, scenario.Actual[0].EtappeCreated.EtappeDate.Day())
	assert.Equal(t, "Parijs", scenario.Actual[0].EtappeCreated.EtappeStartLocation)
	assert.Equal(t, "Roubaix", scenario.Actual[0].EtappeCreated.EtappeFinishtLocation)
	assert.Equal(t, 255, scenario.Actual[0].EtappeCreated.EtappeLength)
	assert.Equal(t, 3, scenario.Actual[0].EtappeCreated.EtappeKind)
	assert.Equal(t, 2015, scenario.Actual[0].EtappeCreated.Year)
}
