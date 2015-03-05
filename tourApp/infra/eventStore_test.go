package infra

import (
	"github.com/MarcGrol/microgen/envelope"
	"github.com/MarcGrol/microgen/tourApp/events"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

const (
	DIRNAME  = "."
	FILENAME = "test.db"
)

func TestStore(t *testing.T) {

	os.Remove(DIRNAME + "/" + FILENAME)

	store := NewEventStore(DIRNAME, FILENAME)

	{
		// write and close
		err := store.Open()
		assert.Nil(t, err)
		tourCreatedEvent := &events.TourCreated{Year: 2015}
		store.Store(tourCreatedEvent.Wrap())
		store.Close()
	}

	{
		// write and close
		err := store.Open()
		assert.Nil(t, err)
		{
			cyclistCreatedEvent := &events.CyclistCreated{
				Year:        2015,
				CyclistId:   42,
				CyclistName: "Lance",
				CyclistTeam: "Rabo"}
			err = store.Store(cyclistCreatedEvent.Wrap())
			assert.Nil(t, err)
		}
		{
			cyclistCreatedEvent2 := &events.CyclistCreated{
				Year:        2016,
				CyclistId:   43,
				CyclistName: "Michael Boogerd",
				CyclistTeam: "Rabo"}
			err = store.Store(cyclistCreatedEvent2.Wrap())
			assert.Nil(t, err)
			store.Close()
		}
	}

	{
		// read all and close
		err := store.Open()
		assert.Nil(t, err)
		envelopes := make([]*envelope.Envelope, 0, 2)
		cb := func(envelope *envelope.Envelope) {
			envelopes = append(envelopes, envelope)
		}
		store.Iterate(cb)
		assert.Equal(t, 3, len(envelopes))

		assert.Equal(t, uint64(1), envelopes[0].SequenceNumber)
		assert.Equal(t, "tour", envelopes[0].AggregateName)
		assert.Equal(t, "2015", envelopes[0].AggregateUid)
		tourCreated, ok := events.GetIfIsTourCreated(envelopes[0])
		assert.True(t, ok)
		assert.NotNil(t, tourCreated)
		assert.Equal(t, 2015, tourCreated.Year)

		assert.Equal(t, uint64(2), envelopes[1].SequenceNumber)
		assert.Equal(t, "tour", envelopes[1].AggregateName)
		assert.Equal(t, "2015", envelopes[1].AggregateUid)
		cyclistCreated, ok := events.GetIfIsCyclistCreated(envelopes[1])
		assert.True(t, ok)
		assert.NotNil(t, cyclistCreated)
		assert.Equal(t, 42, cyclistCreated.CyclistId)
		assert.Equal(t, "Lance", cyclistCreated.CyclistName)
		assert.Equal(t, "Rabo", cyclistCreated.CyclistTeam)

		assert.Equal(t, uint64(3), envelopes[2].SequenceNumber)
		assert.Equal(t, "tour", envelopes[2].AggregateName)
		assert.Equal(t, "2016", envelopes[2].AggregateUid)
		assert.True(t, ok)
		cyclistCreated2, ok := events.GetIfIsCyclistCreated(envelopes[2])
		assert.NotNil(t, cyclistCreated2)
		assert.Equal(t, 43, cyclistCreated2.CyclistId)
		assert.Equal(t, "Michael Boogerd", cyclistCreated2.CyclistName)
		assert.Equal(t, "Rabo", cyclistCreated2.CyclistTeam)

		store.Close()
	}

	store.Close()
}

func BenchmarkWrite(b *testing.B) {
	os.Remove(DIRNAME + FILENAME)

	store := NewEventStore(DIRNAME, FILENAME)
	store.Open()

	event := &events.CyclistCreated{
		Year:        2015,
		CyclistId:   42,
		CyclistName: "Lance",
		CyclistTeam: "Rabo"}
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		store.Store(event.Wrap())
		reader := func(envelop *envelope.Envelope) {
		}
		store.Iterate(reader)
	}

	store.Close()

}
