package events

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

const (
	DIRNAME = "."
	FILENAME = "test.db"
)

func TestStore(t *testing.T) {

	os.Remove(DIRNAME+"/"+FILENAME)

	store := NewEventStore(DIRNAME,FILENAME)

	{
		// write and close
		err := store.Open()
		assert.Nil(t, err)
		tourCreatedEvent := TourCreated{Year: 2015}
		store.Store(tourCreatedEvent.Wrap())
		store.Close()
	}

	{
		// write and close
		err := store.Open()
		assert.Nil(t, err)
		cyclistCreatedEvent := CyclistCreated{
			Year:        2015,
			CyclistId:   42,
			CyclistName: "Lance",
			CyclistTeam: "Rabo"}
		err = store.Store(cyclistCreatedEvent.Wrap())
		assert.Nil(t, err)
		cyclistCreatedEvent2 := CyclistCreated{
			Year:        2016,
			CyclistId:   43,
			CyclistName: "Michael Boogerd",
			CyclistTeam: "Rabo"}
		err = store.Store(cyclistCreatedEvent2.Wrap())
		assert.Nil(t, err)
		store.Close()
	}

	{
		// read all and close
		err := store.Open()
		assert.Nil(t, err)
		envelopes := make([]*Envelope, 0, 2)
		cb := func(envelope *Envelope) {
			envelopes = append(envelopes, envelope)
		}
		store.Iterate(cb)
		assert.Equal(t, 3, len(envelopes))

		assert.Equal(t, uint64(1), envelopes[0].SequenceNumber)
		assert.Equal(t, "tour", envelopes[0].AggregateName)
		assert.Equal(t, "2015", envelopes[0].AggregateUid)
		assert.Equal(t, TypeTourCreated, envelopes[0].Type)
		assert.Equal(t, 2015, envelopes[0].TourCreated.Year)

		assert.Equal(t, uint64(2), envelopes[1].SequenceNumber)
		assert.Equal(t, "tour", envelopes[1].AggregateName)
		assert.Equal(t, "2015", envelopes[1].AggregateUid)
		assert.Equal(t, TypeCyclistCreated, envelopes[1].Type)
		assert.Equal(t, 42, envelopes[1].CyclistCreated.CyclistId)
		assert.Equal(t, "Lance", envelopes[1].CyclistCreated.CyclistName)
		assert.Equal(t, "Rabo", envelopes[1].CyclistCreated.CyclistTeam)

		assert.Equal(t, uint64(3), envelopes[2].SequenceNumber)
		assert.Equal(t, "tour", envelopes[2].AggregateName)
		assert.Equal(t, "2016", envelopes[2].AggregateUid)
		assert.Equal(t, TypeCyclistCreated, envelopes[2].Type)
		assert.Equal(t, 43, envelopes[2].CyclistCreated.CyclistId)
		assert.Equal(t, "Michael Boogerd", envelopes[2].CyclistCreated.CyclistName)
		assert.Equal(t, "Rabo", envelopes[2].CyclistCreated.CyclistTeam)

		store.Close()
	}

	store.Close()
}

func BenchmarkWrite(b *testing.B) {
	os.Remove(DIRNAME+FILENAME)

	store := NewEventStore(DIRNAME,FILENAME)
	store.Open()

	envelope := (&CyclistCreated{
		Year:        2015,
		CyclistId:   42,
		CyclistName: "Lance",
		CyclistTeam: "Rabo"}).Wrap()
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		store.Store(envelope)
		reader := func(envelope *Envelope) {
		}
		store.Iterate(reader)
	}

	store.Close()

}
