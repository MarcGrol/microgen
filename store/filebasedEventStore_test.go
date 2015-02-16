package store

import (
	"github.com/stretchr/testify/assert"
	"github.com/xebia/microgen/events"
	"os"
	"testing"
)

const (
	FILENAME = "test.db"
)

func TestStore(t *testing.T) {

	os.Remove(FILENAME)

	store := NewSimpleEventStore()
	err := store.Open(FILENAME)
	assert.Nil(t, err)

	{
		tourCreatedEvent := events.TourCreated{Year: 2015}
		store.Store(tourCreatedEvent.Wrap())

		cyclistCreatedEvent := events.CyclistCreated{
			Year:        2015,
			CyclistId:   42,
			CyclistName: "Lance",
			CyclistTeam: "Rabo"}
		store.Store(cyclistCreatedEvent.Wrap())

		envelopes := make([]*events.Envelope, 0, 2)
		cb := func(envelope *events.Envelope) bool {
			envelopes = append(envelopes, envelope)
			return false
		}
		store.Iterate(cb)
		assert.Equal(t, 2, len(envelopes))

		assert.Equal(t, uint64(1), envelopes[0].SequenceNumber)
		assert.Equal(t, "tour", envelopes[0].AggregateName)
		assert.Equal(t, "2015", envelopes[0].AggregateUid)
		assert.Equal(t, events.TypeTourCreated, envelopes[0].Type)
		assert.Equal(t, 2015, envelopes[0].TourCreated.Year)

		assert.Equal(t, uint64(2), envelopes[1].SequenceNumber)
		assert.Equal(t, "tour", envelopes[0].AggregateName)
		assert.Equal(t, "2015", envelopes[0].AggregateUid)
		assert.Equal(t, events.TypeCyclistCreated, envelopes[1].Type)
		assert.Equal(t, 42, envelopes[1].CyclistCreated.CyclistId)
		assert.Equal(t, "Lance", envelopes[1].CyclistCreated.CyclistName)
		assert.Equal(t, "Rabo", envelopes[1].CyclistCreated.CyclistTeam)
	}

	{
		envelopes := make([]*events.Envelope, 0, 1)
		cb := func(envelope *events.Envelope) bool {
			envelopes = append(envelopes, envelope)
			return true
		}
		store.Iterate(cb)
		assert.Equal(t, 1, len(envelopes))
	}

	store.Close()
}

func BenchmarkWrite(b *testing.B) {
	os.Remove(FILENAME)

	store := NewSimpleEventStore()
	store.Open(FILENAME)

	envelope := (&events.CyclistCreated{
		Year:        2015,
		CyclistId:   42,
		CyclistName: "Lance",
		CyclistTeam: "Rabo"}).Wrap()
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		store.Store(envelope)
		reader := func(envelope *events.Envelope) bool {
			return false
		}
		store.Iterate(reader)
	}

	store.Close()

}
