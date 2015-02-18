package tour

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"github.com/MarcGrol/microgen/tourApp/events"
)

const (
	FILENAME = "tour_test.db"
)


func TestCreateTourCommand(t *testing.T) {
	operationOnSubject := func(scenario *Scenario) error {
		service := NewTourCommandHandler(scenario.bus, scenario.store)	
		return service.HandleCreateTourCommand(CreateTourCommand{Year:2015})
	}
	scenario := Scenario {
		Name: "NewTourSuccess",
		Description: "Create new tour on clean system",
		Given:[]events.Envelope{},
		CallSubject: operationOnSubject,
		Expect:[]events.Envelope{
			{
				SequenceNumber:1, AggregateName:"tour", AggregateUid:"2015",Type:events.TypeTourCreated, 
				TourCreated:&events.TourCreated{Year:2015},
			},
		},
	}

	scenario.RunAndVerify(t)

	assert.Equal(t,2015,scenario.Actual[0].TourCreated.Year)
}

func TestCreateCyclistCommand(t *testing.T) {
	operationOnSubject := func(scenario *Scenario) error {
		service := NewTourCommandHandler(scenario.bus,scenario.store)
		return service.HandleCreateCyclistCommand(CreateCyclistCommand{Year:2015,Id:42,Name:"My name",Team:"My team"})
	}
	scenario := Scenario {
		Name: "NewCyclistSuccess",
		Description: "Create new cyclist with existing tour",
		Given:[]events.Envelope{ 
			{
				Type:events.TypeTourCreated, 
				TourCreated:&events.TourCreated{Year:2015},
			},
		},
		CallSubject: operationOnSubject,
		Expect:[]events.Envelope{
			{
				SequenceNumber:2, AggregateName:"tour", AggregateUid:"2015",Type:events.TypeCyclistCreated, 
				CyclistCreated:&events.CyclistCreated{Year:2015},
			},
		},
	}
	scenario.RunAndVerify(t)

	assert.Equal(t,2015,scenario.Actual[0].CyclistCreated.Year)
}


func TestCreateEtappeCommand(t *testing.T) {
	operationOnSubject := func(scenario *Scenario) error {
		service := NewTourCommandHandler(scenario.bus,scenario.store)
		return service.HandleCreateEtappeCommand(CreateEtappeCommand{Year:2015})
	}
	scenario := Scenario {
		Name: "NewEtappeSuccess",
		Description: "Create new etappe with existing tour",
		Given:[]events.Envelope{ 
			{
				Type:events.TypeTourCreated, 
				TourCreated:&events.TourCreated{Year:2015},
			},
		},
		CallSubject: operationOnSubject,
		Expect:[]events.Envelope{
			{
				SequenceNumber:2, AggregateName:"tour", AggregateUid:"2015",Type:events.TypeEtappeCreated, 
				EtappeCreated:&events.EtappeCreated{Year:2015},
			},
		},
	}
	scenario.RunAndVerify(t)

	assert.Equal(t,2015,scenario.Actual[0].EtappeCreated.Year)
}

