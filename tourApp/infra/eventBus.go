package infra

import (
	"encoding/json"
	"github.com/MarcGrol/microgen/bus"
	"github.com/MarcGrol/microgen/tourApp/events"
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

func (bus *EventBus) Subscribe(eventType events.Type, userCallback events.EventHandlerFunc) error {
	callback := func(blob []byte) error {
		var envelope events.Envelope
		err := json.Unmarshal(blob, &envelope)
		if err != nil {
			log.Printf("Error unmarshalling json blob (%+v)", err)
			return err
		} else {
			log.Printf("**** Received event (%+v)", envelope)
			return userCallback(&envelope)
		}
	}
	return bus.nsqBus.Subscribe(bus.getTopicName(eventType), callback)
}

func (bus *EventBus) Publish(envelope *events.Envelope) error {
	jsonBlob, err := json.Marshal(envelope)
	if err != nil {
		log.Printf("Error marshalling event-envelope (%+v)", err)
		return err
	}
	//log.Printf("Marshalled event of type %d (%s)", envelope.Type, jsonBlob)
	err = bus.nsqBus.Publish(bus.getTopicName(envelope.Type), jsonBlob)
	if err != nil {
		log.Printf("Error publishing event-envelope (%+v)", err)
	}
	log.Printf("**** Published event (%+v)", envelope)
	return err
}

func (bus *EventBus) getTopicName(eventType events.Type) string {
	return bus.applicationName + "_" + eventType.String()
}
