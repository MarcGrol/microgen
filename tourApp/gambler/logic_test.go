package gambler

import (
	"github.com/MarcGrol/microgen/myerrors"
	"github.com/MarcGrol/microgen/tourApp/events"
	"github.com/MarcGrol/microgen/tourApp/test"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCreateGamblerCommand(t *testing.T) {
	var service CommandHandler
	scenario := test.Scenario{
		Title: "Create new gambler success",
		Given: []*events.Envelope{
			(&events.TourCreated{Year: 2015}).Wrap(),
		},
		Command: CreateGamblerCommand{GamblerUid: "my uid", Name: "My name", Email: "me@home.nl"},
		When: func(scenario *test.Scenario) *myerrors.Error {
			service = NewGamblerCommandHandler(scenario.Bus, scenario.Store)
			return service.HandleCreateGamblerCommand(scenario.Command.(CreateGamblerCommand))
		},
		Expect: []*events.Envelope{
			(&events.GamblerCreated{GamblerUid: "my uid", GamblerName: "My name", GamblerEmail: "me@home.nl"}).Wrap(),
		},
	}

	scenario.RunAndVerify(t)

	assert.Nil(t, scenario.Err)

	expected := scenario.Expect[0].GamblerCreated
	actual := scenario.Actual[0].GamblerCreated
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
	scenario := test.Scenario{
		Title: "Create new gambler team success",
		Given: []*events.Envelope{
			(&events.TourCreated{Year: 2015}).Wrap(),
			(&events.CyclistCreated{Year: 2015, CyclistId: 1, CyclistName: "cyclist 1", CyclistTeam: "team 1"}).Wrap(),
			(&events.CyclistCreated{Year: 2015, CyclistId: 2, CyclistName: "cyclist 2", CyclistTeam: "team 2"}).Wrap(),
			(&events.GamblerCreated{GamblerUid: "my uid", GamblerName: "My name", GamblerEmail: "me@home.nl"}).Wrap(),
		},
		Command: CreateGamblerTeamCommand{GamblerUid: "my uid", Year: 2015, CyclistIds: []int{1, 2}},
		When: func(scenario *test.Scenario) *myerrors.Error {
			service = NewGamblerCommandHandler(scenario.Bus, scenario.Store)
			return service.HandleCreateGamblerTeamCommand(scenario.Command.(CreateGamblerTeamCommand))
		},
		Expect: []*events.Envelope{
			(&events.GamblerTeamCreated{GamblerUid: "my uid", Year: 2015, GamblerCyclists: []int{1, 2}}).Wrap(),
		},
	}

	scenario.RunAndVerify(t)

	assert.Nil(t, scenario.Err)

	expected := scenario.Expect[0].GamblerTeamCreated
	actual := scenario.Actual[0].GamblerTeamCreated
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
