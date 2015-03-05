package envelope

import (
	"bytes"
	"encoding/gob"
	"encoding/json"
	"log"
	"time"
)

type Envelope struct {
	Uuid           string    `json:"uuid"`
	SequenceNumber uint64    `json:"sequenceNumber"`
	Timestamp      time.Time `json:"timestamp"`
	AggregateName  string    `json:"aggregateName"`
	AggregateUid   string    `json:"aggregateUid"`
	EventTypeName  string    `json:"eventTypeName"`
	EventData      string    `json:"eventData"`
}

func (envelope *Envelope) ToJson() (string, error) {
	blob, err := json.Marshal(envelope)
	if err != nil {
		log.Printf("Error marshalling json blob (%+v)", err)
		return "sss", err
	}
	return string(blob), nil
}

func (envelope *Envelope) FromJson(blob string) error {
	err := json.Unmarshal([]byte(blob), envelope)
	if err != nil {
		log.Printf("Error unmarshalling json blob (%+v)", err)
		return err
	}
	return nil
}

func (envelope *Envelope) ToGob(buf *bytes.Buffer) error {
	enc := gob.NewEncoder(buf)
	err := enc.Encode(envelope)
	if err != nil {
		return err
	}
	return nil
}

func (envelope *Envelope) FromGob(buf *bytes.Buffer) error {
	dec := gob.NewDecoder(buf)
	err := dec.Decode(envelope)
	if err != nil {
		log.Printf("Error unmarshalling gob blob (%+v)", err)
		return err
	}
	return nil
}
