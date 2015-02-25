package infra

import (
	"github.com/MarcGrol/microgen/tourApp/events"
	"github.com/stretchr/testify/assert"
	"sync"
	"testing"
)

func TestPublishSubscribe(t *testing.T) {
	bus := NewEventBus("tourdefrance", "unittest", "127.0.0.1")

	wg := &sync.WaitGroup{}

	// subscribe to event
	received := make([]*events.Envelope, 0, 10)
	cb := func(envelope *events.Envelope) error {
		received = append(received, envelope)
		wg.Done()
		return nil
	}
	bus.Subscribe(events.TypeTourCreated, cb)
	bus.Subscribe(events.TypeCyclistCreated, cb)

	wg.Add(2)

	// publish 3 events
	bus.Publish((&events.TourCreated{Year: 2015}).Wrap())
	bus.Publish((&events.CyclistCreated{
		Year:        2015,
		CyclistId:   42,
		CyclistName: "Lance",
		CyclistTeam: "Rabo"}).Wrap())
	bus.Publish((&events.GamblerCreated{
		GamblerUid:   "myuid",
		GamblerName:  "myname",
		GamblerEmail: "myname@domain.com"}).Wrap())

	// Block untill 2 events have been received
	wg.Wait()

	// verify received
	assert.Equal(t, events.TypeTourCreated, received[0].Type)
	assert.Equal(t, 2015, received[0].TourCreated.Year)
	assert.Equal(t, events.TypeCyclistCreated, received[1].Type)
	assert.Equal(t, 42, received[1].CyclistCreated.CyclistId)
	assert.Equal(t, "Lance", received[1].CyclistCreated.CyclistName)
	assert.Equal(t, "Rabo", received[1].CyclistCreated.CyclistTeam)
}
