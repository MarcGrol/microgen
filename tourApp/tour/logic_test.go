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

func TestCreateTourCommandTourExists(t *testing.T) {
	var service CommandHandler
	scenario := test.CommandScenario{
		Title: "Create tour with existing tour",
		Given: []*envelope.Envelope{
			(&events.TourCreated{Year: 2015}).Wrap(),
		},
		Command: &CreateTourCommand{Year: 2015},
		When: func(scenario *test.CommandScenario) error {
			service = NewTourCommandHandler(scenario.Bus, scenario.Store)
			return service.HandleCreateTourCommand(scenario.Command.(*CreateTourCommand))
		},
		Expect: []*envelope.Envelope{},
	}

	scenario.RunAndVerify(t)

	assert.NotNil(t, scenario.ErrMsg)
	assert.Equal(t, "Tour 2015 already exists", *scenario.ErrMsg)
	assert.True(t, scenario.InvalidInputError)
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

func TestCreateCyclistCommandUnknownTour(t *testing.T) {
	var service CommandHandler
	scenario := test.CommandScenario{
		Title: "Create cyclist with unknown tour",
		Given: []*envelope.Envelope{},
		Command: &CreateCyclistCommand{
			Year: 2015,
			Id:   42,
			Name: "My name",
			Team: "My team"},
		When: func(scenario *test.CommandScenario) error {
			service = NewTourCommandHandler(scenario.Bus, scenario.Store)
			return service.HandleCreateCyclistCommand(scenario.Command.(*CreateCyclistCommand))
		},
		Expect: []*envelope.Envelope{},
	}

	scenario.RunAndVerify(t)

	assert.NotNil(t, scenario.ErrMsg)
	assert.Equal(t, "Tour 2015 does not exist", *scenario.ErrMsg)
	assert.True(t, scenario.NotFoundError)
}

func TestCreateCyclistCommandInvalidCyclist(t *testing.T) {
	var service CommandHandler
	scenario := test.CommandScenario{
		Title: "Create invalid new cyclist",
		Given: []*envelope.Envelope{
			(&events.TourCreated{Year: 2015}).Wrap(),
		},
		Command: &CreateCyclistCommand{
			Year: 2015,
			Name: "My name",
			Team: "My team"},
		When: func(scenario *test.CommandScenario) error {
			service = NewTourCommandHandler(scenario.Bus, scenario.Store)
			return service.HandleCreateCyclistCommand(scenario.Command.(*CreateCyclistCommand))
		},
		Expect: []*envelope.Envelope{},
	}

	scenario.RunAndVerify(t)

	assert.NotNil(t, scenario.ErrMsg)
	assert.Equal(t, "Invalid parameter Id", *scenario.ErrMsg)
	assert.True(t, scenario.InvalidInputError)
}

func TestCreateCyclistCommandDuplicateCyclist(t *testing.T) {
	var service CommandHandler
	scenario := test.CommandScenario{
		Title: "Create duplicate new cyclist",
		Given: []*envelope.Envelope{
			(&events.TourCreated{Year: 2015}).Wrap(),
			(&events.CyclistCreated{
				Year:        2015,
				CyclistId:   42,
				CyclistName: "My name",
				CyclistTeam: "My team"}).Wrap(),
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
		Expect: []*envelope.Envelope{},
	}

	scenario.RunAndVerify(t)

	assert.NotNil(t, scenario.ErrMsg)
	assert.Equal(t, "Cyclist with id 42 already exists", *scenario.ErrMsg)
	assert.True(t, scenario.InvalidInputError)
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

func TestCreateEtappeVommandUnknownTour(t *testing.T) {
	var service CommandHandler
	scenario := test.CommandScenario{
		Title: "Create etappe with unknown tour",
		Given: []*envelope.Envelope{
		//(&events.TourCreated{Year: 2013}).Wrap(),
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
		Expect: []*envelope.Envelope{},
	}

	scenario.RunAndVerify(t)

	assert.NotNil(t, scenario.ErrMsg)
	assert.Equal(t, "Tour 2015 does not exist", *scenario.ErrMsg)
	assert.True(t, scenario.NotFoundError)
}

func TestCreateEtappeCommandInvalidEtappe(t *testing.T) {
	var service CommandHandler
	scenario := test.CommandScenario{
		Title: "Create invalid etappe",
		Given: []*envelope.Envelope{
			(&events.TourCreated{Year: 2015}).Wrap(),
		},
		Command: &CreateEtappeCommand{
			Year:           2015,
			Id:             2,
			StartLocation:  "Parijs",
			FinishLocation: "Roubaix",
			Length:         255,
			Kind:           3},
		When: func(scenario *test.CommandScenario) error {
			service = NewTourCommandHandler(scenario.Bus, scenario.Store)
			return service.HandleCreateEtappeCommand(scenario.Command.(*CreateEtappeCommand))
		},
		Expect: []*envelope.Envelope{},
	}

	scenario.RunAndVerify(t)

	assert.NotNil(t, scenario.ErrMsg)
	assert.Equal(t, "Invalid parameter Date", *scenario.ErrMsg)
	assert.True(t, scenario.InvalidInputError)
}

func TestCreateEtappeCommandDuplicateEtappe(t *testing.T) {
	var service CommandHandler
	scenario := test.CommandScenario{
		Title: "Create duplicate etappe",
		Given: []*envelope.Envelope{
			(&events.TourCreated{Year: 2015}).Wrap(),
			(&events.EtappeCreated{
				Year:                 2015,
				EtappeId:             2,
				EtappeDate:           time.Date(2015, time.July, 14, 9, 0, 0, 0, time.Local),
				EtappeStartLocation:  "Parijs",
				EtappeFinishLocation: "Roubaix",
				EtappeLength:         255,
				EtappeKind:           3}).Wrap(),
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
		Expect: []*envelope.Envelope{},
	}

	scenario.RunAndVerify(t)

	assert.NotNil(t, scenario.ErrMsg)
	assert.Equal(t, "Etappe with id 2 already exists", *scenario.ErrMsg)
	assert.True(t, scenario.InvalidInputError)
}

func TestCreateEtappeResultsCommandInvalidRequest(t *testing.T) {
	var service CommandHandler
	scenario := test.CommandScenario{
		Title: "Create new etappe result invalid request",
		Given: []*envelope.Envelope{},
		Command: &CreateEtappeResultsCommand{
			Year:                   2015,
			EtappeId:               2,
			BestAllroundCyclistIds: []int{1, 2, 3, 4, 5},
			BestClimbCyclistIds:    []int{1, 2, 3, 4, 5},
			BestSprintCyclistIds:   []int{1, 2, 3, 4, 5},
		},
		When: func(scenario *test.CommandScenario) error {
			service = NewTourCommandHandler(scenario.Bus, scenario.Store)
			return service.HandleCreateEtappeResultsCommand(scenario.Command.(*CreateEtappeResultsCommand))
		},
		Expect: []*envelope.Envelope{},
	}

	scenario.RunAndVerify(t)

	assert.NotNil(t, scenario.ErrMsg)
	assert.Equal(t, "Invalid parameter BestDayCyclistIds", *scenario.ErrMsg)
	assert.True(t, scenario.InvalidInputError)
}

func TestCreateEtappeResultsUnknownTourCommand(t *testing.T) {
	var service CommandHandler
	scenario := test.CommandScenario{
		Title: "Create new etappe result unknown tour",
		Given: []*envelope.Envelope{},
		Command: &CreateEtappeResultsCommand{
			Year:                   2015,
			EtappeId:               2,
			BestDayCyclistIds:      []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			BestAllroundCyclistIds: []int{1, 2, 3, 4, 5},
			BestClimbCyclistIds:    []int{1, 2, 3, 4, 5},
			BestSprintCyclistIds:   []int{1, 2, 3, 4, 5},
		},
		When: func(scenario *test.CommandScenario) error {
			service = NewTourCommandHandler(scenario.Bus, scenario.Store)
			return service.HandleCreateEtappeResultsCommand(scenario.Command.(*CreateEtappeResultsCommand))
		},
		Expect: []*envelope.Envelope{},
	}

	scenario.RunAndVerify(t)

	assert.NotNil(t, scenario.ErrMsg)
	assert.Equal(t, "Tour 2015 does not exist", *scenario.ErrMsg)
	assert.True(t, scenario.NotFoundError)
}

func TestCreateEtappeResultsCommandUnknownEtappe(t *testing.T) {
	var service CommandHandler
	scenario := test.CommandScenario{
		Title: "Create new etappe result unknown etappe",
		Given: []*envelope.Envelope{
			(&events.TourCreated{Year: 2015}).Wrap(),
		},
		Command: &CreateEtappeResultsCommand{
			Year:                   2015,
			EtappeId:               2,
			BestDayCyclistIds:      []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			BestAllroundCyclistIds: []int{1, 2, 3, 4, 5},
			BestClimbCyclistIds:    []int{1, 2, 3, 4, 5},
			BestSprintCyclistIds:   []int{1, 2, 3, 4, 5},
		},
		When: func(scenario *test.CommandScenario) error {
			service = NewTourCommandHandler(scenario.Bus, scenario.Store)
			return service.HandleCreateEtappeResultsCommand(scenario.Command.(*CreateEtappeResultsCommand))
		},
		Expect: []*envelope.Envelope{},
	}

	scenario.RunAndVerify(t)

	assert.NotNil(t, scenario.ErrMsg)
	assert.Equal(t, "Etappe with id 2 does not exist", *scenario.ErrMsg)
	assert.True(t, scenario.NotFoundError)
}

func TestCreateEtappeResultsUnknownCyclistCommand(t *testing.T) {
	var service CommandHandler
	scenario := test.CommandScenario{
		Title: "Create new etappe result unknown cyclist",
		Given: []*envelope.Envelope{
			(&events.TourCreated{Year: 2015}).Wrap(),
			(&events.EtappeCreated{
				Year:                 2015,
				EtappeId:             2,
				EtappeDate:           time.Date(2015, time.July, 14, 9, 0, 0, 0, time.Local),
				EtappeStartLocation:  "Parijs",
				EtappeFinishLocation: "Roubaix",
				EtappeLength:         255,
				EtappeKind:           3}).Wrap(),
		},
		Command: &CreateEtappeResultsCommand{
			Year:                   2015,
			EtappeId:               2,
			BestDayCyclistIds:      []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			BestAllroundCyclistIds: []int{1, 2, 3, 4, 5},
			BestClimbCyclistIds:    []int{1, 2, 3, 4, 5},
			BestSprintCyclistIds:   []int{1, 2, 3, 4, 5},
		},
		When: func(scenario *test.CommandScenario) error {
			service = NewTourCommandHandler(scenario.Bus, scenario.Store)
			return service.HandleCreateEtappeResultsCommand(scenario.Command.(*CreateEtappeResultsCommand))
		},
		Expect: []*envelope.Envelope{},
	}

	scenario.RunAndVerify(t)

	assert.NotNil(t, scenario.ErrMsg)
	assert.Equal(t, "BestDayCyclistIds: Cyclist with id 1 does not exist", *scenario.ErrMsg)
	assert.True(t, scenario.NotFoundError)
}

func TestCreateEtappeResultsDuplicateCyclistCommand(t *testing.T) {
	var service CommandHandler
	scenario := test.CommandScenario{
		Title: "Create new etappe result duplicate cyclist",
		Given: []*envelope.Envelope{
			(&events.TourCreated{Year: 2015}).Wrap(),
			(&events.EtappeCreated{
				Year:                 2015,
				EtappeId:             2,
				EtappeDate:           time.Date(2015, time.July, 14, 9, 0, 0, 0, time.Local),
				EtappeStartLocation:  "Parijs",
				EtappeFinishLocation: "Roubaix",
				EtappeLength:         255,
				EtappeKind:           3}).Wrap(),
			(&events.CyclistCreated{
				Year:        2015,
				CyclistId:   1,
				CyclistName: "1",
				CyclistTeam: "My team"}).Wrap(),
			(&events.CyclistCreated{
				Year:        2015,
				CyclistId:   2,
				CyclistName: "2",
				CyclistTeam: "My team"}).Wrap(),
			(&events.CyclistCreated{
				Year:        2015,
				CyclistId:   3,
				CyclistName: "3",
				CyclistTeam: "My team"}).Wrap(),
			(&events.CyclistCreated{
				Year:        2015,
				CyclistId:   4,
				CyclistName: "4",
				CyclistTeam: "My team"}).Wrap(),
			(&events.CyclistCreated{
				Year:        2015,
				CyclistId:   5,
				CyclistName: "5",
				CyclistTeam: "My team"}).Wrap(),
			(&events.CyclistCreated{
				Year:        2015,
				CyclistId:   6,
				CyclistName: "6",
				CyclistTeam: "My team"}).Wrap(),
			(&events.CyclistCreated{
				Year:        2015,
				CyclistId:   7,
				CyclistName: "7",
				CyclistTeam: "My team"}).Wrap(),
			(&events.CyclistCreated{
				Year:        2015,
				CyclistId:   8,
				CyclistName: "8",
				CyclistTeam: "My team"}).Wrap(),
			(&events.CyclistCreated{
				Year:        2015,
				CyclistId:   9,
				CyclistName: "9",
				CyclistTeam: "My team"}).Wrap(),
			(&events.CyclistCreated{
				Year:        2015,
				CyclistId:   10,
				CyclistName: "10",
				CyclistTeam: "My team"}).Wrap(),
		},
		Command: &CreateEtappeResultsCommand{
			Year:                   2015,
			EtappeId:               2,
			BestDayCyclistIds:      []int{1, 1, 3, 4, 5, 6, 7, 8, 9, 10},
			BestAllroundCyclistIds: []int{1, 2, 3, 4, 5},
			BestClimbCyclistIds:    []int{1, 2, 3, 4, 5},
			BestSprintCyclistIds:   []int{1, 2, 3, 4, 5},
		},
		When: func(scenario *test.CommandScenario) error {
			service = NewTourCommandHandler(scenario.Bus, scenario.Store)
			return service.HandleCreateEtappeResultsCommand(scenario.Command.(*CreateEtappeResultsCommand))
		},
		Expect: []*envelope.Envelope{},
	}

	scenario.RunAndVerify(t)

	assert.NotNil(t, scenario.ErrMsg)
	assert.Equal(t, "BestDayCyclistIds contains duplicates", *scenario.ErrMsg)
}

func TestCreateEtappeResultsSuccessCommand(t *testing.T) {
	var service CommandHandler
	scenario := test.CommandScenario{
		Title: "Create new etappe result success",
		Given: []*envelope.Envelope{
			(&events.TourCreated{Year: 2015}).Wrap(),
			(&events.EtappeCreated{
				Year:                 2015,
				EtappeId:             2,
				EtappeDate:           time.Date(2015, time.July, 14, 9, 0, 0, 0, time.Local),
				EtappeStartLocation:  "Parijs",
				EtappeFinishLocation: "Roubaix",
				EtappeLength:         255,
				EtappeKind:           3}).Wrap(),
			(&events.CyclistCreated{
				Year:        2015,
				CyclistId:   1,
				CyclistName: "1",
				CyclistTeam: "My team"}).Wrap(),
			(&events.CyclistCreated{
				Year:        2015,
				CyclistId:   2,
				CyclistName: "2",
				CyclistTeam: "My team"}).Wrap(),
			(&events.CyclistCreated{
				Year:        2015,
				CyclistId:   3,
				CyclistName: "3",
				CyclistTeam: "My team"}).Wrap(),
			(&events.CyclistCreated{
				Year:        2015,
				CyclistId:   4,
				CyclistName: "4",
				CyclistTeam: "My team"}).Wrap(),
			(&events.CyclistCreated{
				Year:        2015,
				CyclistId:   5,
				CyclistName: "5",
				CyclistTeam: "My team"}).Wrap(),
			(&events.CyclistCreated{
				Year:        2015,
				CyclistId:   6,
				CyclistName: "6",
				CyclistTeam: "My team"}).Wrap(),
			(&events.CyclistCreated{
				Year:        2015,
				CyclistId:   7,
				CyclistName: "7",
				CyclistTeam: "My team"}).Wrap(),
			(&events.CyclistCreated{
				Year:        2015,
				CyclistId:   8,
				CyclistName: "8",
				CyclistTeam: "My team"}).Wrap(),
			(&events.CyclistCreated{
				Year:        2015,
				CyclistId:   9,
				CyclistName: "9",
				CyclistTeam: "My team"}).Wrap(),
			(&events.CyclistCreated{
				Year:        2015,
				CyclistId:   10,
				CyclistName: "10",
				CyclistTeam: "My team"}).Wrap(),
		},
		Command: &CreateEtappeResultsCommand{
			Year:                   2015,
			EtappeId:               2,
			BestDayCyclistIds:      []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			BestAllroundCyclistIds: []int{1, 2, 3, 4, 5},
			BestClimbCyclistIds:    []int{1, 2, 3, 4, 5},
			BestSprintCyclistIds:   []int{1, 2, 3, 4, 5},
		},
		When: func(scenario *test.CommandScenario) error {
			service = NewTourCommandHandler(scenario.Bus, scenario.Store)
			return service.HandleCreateEtappeResultsCommand(scenario.Command.(*CreateEtappeResultsCommand))
		},
		Expect: []*envelope.Envelope{
			(&events.EtappeResultsCreated{
				Year:                     2015,
				LastEtappeId:             2,
				BestDayCyclistIds:        []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
				BestAllrounderCyclistIds: []int{1, 2, 3, 4, 5},
				BestClimberCyclistIds:    []int{1, 2, 3, 4, 5},
				BestSprinterCyclistIds:   []int{1, 2, 3, 4, 5}}).Wrap(),
		},
	}

	scenario.RunAndVerify(t)

	assert.Nil(t, scenario.ErrMsg)

	expected := events.UnWrapEtappeResultsCreated(scenario.Expect[0])
	actual := events.UnWrapEtappeResultsCreated(&scenario.Actual[0])
	assert.Equal(t, expected.Year, actual.Year)
	assert.Equal(t, expected.LastEtappeId, actual.LastEtappeId)
	assert.Equal(t, expected.BestDayCyclistIds, actual.BestDayCyclistIds)
	assert.Equal(t, expected.BestAllrounderCyclistIds, actual.BestAllrounderCyclistIds)
	assert.Equal(t, expected.BestSprinterCyclistIds, actual.BestSprinterCyclistIds)
	assert.Equal(t, expected.BestClimberCyclistIds, actual.BestClimberCyclistIds)

	// Test query
	tour, err := service.HandleGetTourQuery(expected.Year)
	assert.Nil(t, err)
	assert.NotNil(t, tour)
	assert.Equal(t, 2015, tour.Year)
	assert.Equal(t, 1, len(tour.Etappes))
	assert.Equal(t, 10, len(tour.Cyclists))
	assert.Equal(t, expected.LastEtappeId, tour.Etappes[0].Id)

	assert.Equal(t, expected.BestDayCyclistIds[0], tour.Etappes[0].Results.BestDayCyclists[0].Number)
	assert.Equal(t, expected.BestDayCyclistIds[9], tour.Etappes[0].Results.BestDayCyclists[9].Number)

	assert.Equal(t, expected.BestAllrounderCyclistIds[0], tour.Etappes[0].Results.BestAllrounderCyclists[0].Number)
	assert.Equal(t, expected.BestAllrounderCyclistIds[4], tour.Etappes[0].Results.BestAllrounderCyclists[4].Number)

	assert.Equal(t, expected.BestSprinterCyclistIds[0], tour.Etappes[0].Results.BestSprinterCyclists[0].Number)
	assert.Equal(t, expected.BestSprinterCyclistIds[4], tour.Etappes[0].Results.BestSprinterCyclists[4].Number)

	assert.Equal(t, expected.BestClimberCyclistIds[0], tour.Etappes[0].Results.BestClimberCyclists[0].Number)
	assert.Equal(t, expected.BestClimberCyclistIds[4], tour.Etappes[0].Results.BestClimberCyclists[4].Number)

}
