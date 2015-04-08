package envelope

import (
	"bytes"
	"encoding/json"
	"log"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestString(t *testing.T) {
	expected := Envelope{
		Uuid:           "123",
		SequenceNumber: 456,
		AggregateName:  "tour",
		AggregateUid:   "2015",
		Timestamp:      time.Now(),
		EventTypeName:  "dummy",
		EventData:      "1234567890abcdefghijklmnopqrstuvwxyz",
	}

	{
		actual := toBytesFromBytesJson(t, &expected)
		assert.NotNil(t, actual)

		compare(t, &expected, actual)
		assert.Equal(t, string(expected.EventData), string(actual.EventData))
	}
	{
		actual := toBytesFromBytesGob(t, &expected)
		assert.NotNil(t, actual)

		compare(t, &expected, actual)
		assert.Equal(t, string(expected.EventData), string(actual.EventData))
	}
}

type MyStruct struct {
	Id   uint64 `json:"id"`
	Name string `json:"name"`
}

func (envelope *MyStruct) ToJson() ([]byte, error) {
	blob, err := json.Marshal(envelope)
	if err != nil {
		return nil, err
	}
	return blob, nil
}

func (envelope *MyStruct) FromJson(blob []byte) error {
	err := json.Unmarshal(blob, envelope)
	if err != nil {
		return err
	}
	return nil
}

func TestStruct(t *testing.T) {
	expectedPayload := MyStruct{Id: 321, Name: "My name"}
	expectedBlob, err := expectedPayload.ToJson()
	assert.Nil(t, err)

	expected := Envelope{
		Uuid:           "123",
		SequenceNumber: 456,
		AggregateName:  "tour",
		AggregateUid:   "2015",
		Timestamp:      time.Now(),
		EventTypeName:  "dummy",
		EventData:      string(expectedBlob),
	}
	{
		actual := toBytesFromBytesJson(t, &expected)
		assert.NotNil(t, actual)

		compare(t, &expected, actual)

		actualPayload := new(MyStruct)
		err = actualPayload.FromJson([]byte(expected.EventData))
		assert.Nil(t, err)
		assert.Equal(t, expectedPayload.Id, actualPayload.Id)
		assert.Equal(t, expectedPayload.Name, actualPayload.Name)
	}

	{
		actual := toBytesFromBytesGob(t, &expected)
		assert.NotNil(t, actual)

		compare(t, &expected, actual)

		actualPayload := new(MyStruct)
		err = actualPayload.FromJson([]byte(expected.EventData))
		assert.Nil(t, err)
		assert.Equal(t, expectedPayload.Id, actualPayload.Id)
		assert.Equal(t, expectedPayload.Name, actualPayload.Name)

	}
}

func toBytesFromBytesJson(t *testing.T, expected *Envelope) *Envelope {
	blob, err := expected.ToJson()
	assert.Nil(t, err)
	assert.NotNil(t, blob)

	//log.Printf(string(blob))

	actual := new(Envelope)
	err = actual.FromJson(blob)
	assert.Nil(t, err)

	return actual
}

func toBytesFromBytesGob(t *testing.T, expected *Envelope) *Envelope {
	var buff bytes.Buffer
	err := expected.ToGob(&buff)
	assert.Nil(t, err)

	//log.Printf(string(blob))

	actual := new(Envelope)
	err = actual.FromGob(&buff)
	assert.Nil(t, err)

	return actual
}

func compare(t *testing.T, expected *Envelope, actual *Envelope) {
	assert.Equal(t, expected.Uuid, actual.Uuid)
	assert.Equal(t, expected.SequenceNumber, actual.SequenceNumber)
	assert.Equal(t, expected.AggregateName, actual.AggregateName)
	assert.Equal(t, expected.AggregateUid, actual.AggregateUid)
	assert.Equal(t, expected.Timestamp, actual.Timestamp)
	assert.Equal(t, expected.EventTypeName, actual.EventTypeName)
	assert.Equal(t, expected.EventData, actual.EventData)
}

func makeExpected() *Envelope {
	expected := Envelope{
		Uuid:           "123",
		SequenceNumber: 456,
		AggregateName:  "tour",
		AggregateUid:   "2015",
		Timestamp:      time.Now(),
		EventTypeName:  "dummy",
		EventData:      "1234567890abcdefghijklmnopqrstuvwxyz1234567890abcdefghijklmnopqrstuvwxyz1234567890abcdefghijklmnopqrstuvwxyz1234567890abcdefghijklmnopqrstuvwxyz",
	}
	return &expected
}

func BenchmarkJson(b *testing.B) {
	expected := makeExpected()

	for n := 0; n < b.N; n++ {
		blob, err := expected.ToJson()
		if err != nil {
			log.Fatalf("to-json error: %+v", err)
		} else {
			actual := new(Envelope)
			err = actual.FromJson(blob)
			if err != nil {
				log.Fatalf("to-json error: %+v", err)
			}
		}
	}
}

func BenchmarkGob(b *testing.B) {
	expected := makeExpected()

	var buff bytes.Buffer
	for n := 0; n < b.N; n++ {
		buff.Reset()
		err := expected.ToGob(&buff)
		if err != nil {
			log.Fatalf("to-gob error: %+v", err)
		} else {
			actual := new(Envelope)
			err = actual.FromGob(&buff)
			if err != nil {
				log.Fatalf("to-gob error: %+v", err)
			}
		}
	}
}
