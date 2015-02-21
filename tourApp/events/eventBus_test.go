package events

import (
	"github.com/stretchr/testify/assert"
	"sync"
	"testing"
)

func TestPublishSubscribe(t *testing.T) {
	bus := NewEventBus("tourdefrance", "unittest", "127.0.0.1")

	// publish 3 events
	bus.Publish((&TourCreated{Year: 2015}).Wrap())
	bus.Publish((&CyclistCreated{
		Year:        2015,
		CyclistId:   42,
		CyclistName: "Lance",
		CyclistTeam: "Rabo"}).Wrap())
	bus.Publish((&GamblerCreated{
		GamblerUid:   "myuid",
		GamblerName:  "myname",
		GamblerEmail: "myname@domain.com"}).Wrap())

	wg := &sync.WaitGroup{}
	wg.Add(2)

	// subscribe to event
	received := make([]*Envelope, 0, 10)
	cb := func(envelope *Envelope) error {
		received = append(received, envelope)
		wg.Done()
		return nil
	}
	bus.Subscribe(TypeTourCreated, cb)
	bus.Subscribe(TypeCyclistCreated, cb)

	// Block untill 2 events have been received
	wg.Wait()

	// verify received
	assert.Equal(t, TypeTourCreated, received[0].Type)
	assert.Equal(t, 2015, received[0].TourCreated.Year)
	assert.Equal(t, TypeCyclistCreated, received[1].Type)
	assert.Equal(t, 42, received[1].CyclistCreated.CyclistId)
	assert.Equal(t, "Lance", received[1].CyclistCreated.CyclistName)
	assert.Equal(t, "Rabo", received[1].CyclistCreated.CyclistTeam)
}
