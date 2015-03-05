package infra

import (
	"encoding/json"
	"github.com/MarcGrol/microgen/bus"
	"github.com/MarcGrol/microgen/envelope"
	"log"
)

type EventBus struct {
	nsqBus          *bus.NsqBus
	applicationName string
}

func NewEventBus(applicationName string, consumerName string, address string) *EventBus {
	mybus := new(EventBus)
	mybus.nsqBus = bus.NewNsqBus(consumerName, address)
	mybus.applicationName = applicationName
	return mybus
}

func (bus *EventBus) Subscribe(eventTypeName string, userCallback EventHandlerFunc) error {
	var envelop envelope.Envelope
	callback := func(blob []byte) error {
		err := json.Unmarshal(blob, &envelop)
		if err != nil {
			log.Printf("Error unmarshalling json blob (%+v)", err)
			return err
		}

		log.Printf("**** Received event (%+v)", envelop)
		return userCallback(&envelop)
	}
	return bus.nsqBus.Subscribe(bus.getTopicName(&envelop), callback)
}

func (bus *EventBus) Publish(envelope *envelope.Envelope) error {
	jsonBlob, err := json.Marshal(envelope)
	if err != nil {
		log.Printf("Error marshalling event-envelope (%+v)", err)
		return err
	}
	//log.Printf("Marshalled event of type %d (%s)", envelope.Type, jsonBlob)
	err = bus.nsqBus.Publish(bus.getTopicName(envelope), jsonBlob)
	if err != nil {
		log.Printf("Error publishing event-envelope (%+v)", err)
		return err
	}

	log.Printf("**** Published event (%+v)", envelope)
	return nil
}

func (bus *EventBus) getTopicName(envelope *envelope.Envelope) string {
	return bus.applicationName + "_" + envelope.EventTypeName
}
