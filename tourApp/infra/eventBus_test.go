package infra

import (
	"github.com/MarcGrol/microgen/infra/bus"
	"github.com/MarcGrol/microgen/lib/envelope"
	"github.com/MarcGrol/microgen/tourApp/events"
	"github.com/stretchr/testify/assert"
	"sync"
	"testing"
)

func TestPublishSubscribe(t *testing.T) {
	bus := bus.NewEventBus("tourdefrance", "unittest", "127.0.0.1")

	wg := &sync.WaitGroup{}

	// subscribe to event
	received := make([]*envelope.Envelope, 0, 10)
	cb := func(envelop *envelope.Envelope) error {
		received = append(received, envelop)
		wg.Done()
		return nil
	}
	var tourCreatedType events.Type = events.TypeTourCreated
	bus.Subscribe(tourCreatedType.String(), cb)
	var cyclistCreatedType events.Type = events.TypeCyclistCreated
	bus.Subscribe(cyclistCreatedType.String(), cb)

	wg.Add(2)

	// publish 3 events
	{
		event := &events.TourCreated{Year: 2015}
		bus.Publish(event.Wrap())
	}
	{
		event := &events.CyclistCreated{
			Year:        2015,
			CyclistId:   42,
			CyclistName: "Lance",
			CyclistTeam: "Rabo"}
		bus.Publish(event.Wrap())
	}
	{
		event := &events.GamblerCreated{
			GamblerUid:   "myuid",
			GamblerName:  "myname",
			GamblerEmail: "myname@domain.com"}
		bus.Publish(event.Wrap())
	}

	// Block untill 2 events have been received
	wg.Wait()

	// verify received

	{
		assert.True(t, events.IsTourCreated(received[0]))
		actual := events.UnWrapTourCreated(received[0])
		assert.Equal(t, 2015, actual.Year)
	}

	{
		assert.True(t, events.IsCyclistCreated(received[1]))
		actual := events.UnWrapCyclistCreated(received[1])
		assert.Equal(t, 42, actual.CyclistId)
		assert.Equal(t, "Lance", actual.CyclistName)
		assert.Equal(t, "Rabo", actual.CyclistTeam)
	}
}
