package news

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
			service = NewNewsEventHandler(scenario.Bus, scenario.Store)
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
			service = NewNewsEventHandler(scenario.Bus, scenario.Store)
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
			service = NewNewsEventHandler(scenario.Bus, scenario.Store)
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
			service = NewNewsEventHandler(scenario.Bus, scenario.Store)
			return service.OnEvent(scenario.Envelop)
		},
		Expect: []*envelope.Envelope{
			givenTour, givenEtappe,
			givenCyclist1, givenCyclist2, givenCyclist3,
			input},
	}

	scenario.RunAndVerify(t)

	assert.Nil(t, scenario.ErrMsg)
}

func TestCreateNewsItemCommand(t *testing.T) {
	var service CommandHandler
	scenario := test.CommandScenario{
		Title: "Create new gambler success",
		Given: []*envelope.Envelope{
			(&events.TourCreated{Year: 2015}).Wrap(),
			(&events.EtappeCreated{
				Year:                 2015,
				EtappeId:             2,
				EtappeDate:           time.Date(2015, time.July, 13, 9, 0, 0, 0, time.Local),
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
			(&events.EtappeResultsCreated{
				Year:                     2015,
				LastEtappeId:             2,
				BestDayCyclistIds:        []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
				BestAllrounderCyclistIds: []int{1, 2, 3, 4, 5},
				BestClimberCyclistIds:    []int{1, 2, 3, 4, 5},
				BestSprinterCyclistIds:   []int{1, 2, 3, 4, 5}}).Wrap(),
		},
		Command: &CreateNewsItemCommand{
			Year:      2015,
			Timestamp: time.Date(2015, time.July, 14, 9, 0, 0, 0, time.Local),
			Message:   "Hi there",
			Sender:    "Marc"},
		When: func(scenario *test.CommandScenario) error {
			service = NewNewsCommandHandler(scenario.Bus, scenario.Store)
			return service.HandleCreateNewsItemCommand(scenario.Command.(*CreateNewsItemCommand))
		},
		Expect: []*envelope.Envelope{
			(&events.NewsItemCreated{
				Year:      2015,
				Timestamp: time.Date(2015, time.July, 14, 9, 0, 0, 0, time.Local),
				Message:   "Hi there",
				Sender:    "Marc"}).Wrap(),
		},
	}

	scenario.RunAndVerify(t)

	assert.Nil(t, scenario.ErrMsg)

	expected, ok := events.GetIfIsNewsItemCreated(scenario.Expect[0])
	assert.True(t, ok)
	assert.NotNil(t, expected)
	actual, ok := events.GetIfIsNewsItemCreated(&scenario.Actual[0])
	assert.True(t, ok)
	assert.NotNil(t, actual)
	assert.Equal(t, expected.Year, actual.Year)
	assert.Equal(t, expected.Timestamp, actual.Timestamp)
	assert.Equal(t, expected.Message, actual.Message)
	assert.Equal(t, expected.Sender, actual.Sender)

	// Test query to verify
	news, err := service.HandleGetNewsQuery(2015)
	assert.Nil(t, err)
	assert.NotNil(t, news)
	assert.Equal(t, 3, len(news.newsItems))

	assert.Equal(t, "admin", news.newsItems[1].Sender)
	assert.Equal(t, "The tour of 2015 is about to start", news.newsItems[0].Message)

	assert.Equal(t, "admin", news.newsItems[1].Sender)
	assert.Equal(t, "Etappe 2 from Parijs to Roubaix has finished.\nEtappe result:\n-1- 1\n-2- 2\n-3- 3\n",
		news.newsItems[1].Message)

	assert.Equal(t, "Marc", news.newsItems[2].Sender)
	assert.Equal(t, "Hi there", news.newsItems[2].Message)

}

func TestCreateNewsItemCommandInvalidInput(t *testing.T) {
	var service CommandHandler
	scenario := test.CommandScenario{
		Title: "Create new news-item invalid input",
		Given: []*envelope.Envelope{},
		Command: &CreateNewsItemCommand{
			Year:      2015,
			Timestamp: time.Date(2015, time.July, 14, 9, 0, 0, 0, time.Local),
			// no message content
			Sender: "Marc"},
		When: func(scenario *test.CommandScenario) error {
			service = NewNewsCommandHandler(scenario.Bus, scenario.Store)
			return service.HandleCreateNewsItemCommand(scenario.Command.(*CreateNewsItemCommand))
		},
		Expect: []*envelope.Envelope{},
	}

	scenario.RunAndVerify(t)

	assert.Equal(t, "Missing parameter Message", *scenario.ErrMsg)
	assert.True(t, scenario.InvalidInputError)
}
