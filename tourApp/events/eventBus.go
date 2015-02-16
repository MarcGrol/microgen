package events

import (
	"encoding/json"
	"github.com/xebia/microgen/bus"
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

func (bus *EventBus) Subscribe(eventType Type, userCallback EventHandlerFunc) error {
	callback := func(blob []byte) error {
		var envelope Envelope
		err := json.Unmarshal(blob, &envelope)
		if err != nil {
			log.Printf("Error unmarshalling json blob (%v)", err)
			return err
		} else {
			return userCallback(&envelope)
		}
	}
	return bus.nsqBus.Subscribe(bus.getTopicName(eventType), callback)
}

func (bus *EventBus) Publish(envelope *Envelope) error {
	jsonBlob, err := json.Marshal(envelope)
	if err != nil {
		log.Printf("Error marshalling event-envelope (%v)", err)
		return err
	}
	//log.Printf("Marshalled event of type %d (%s)", envelope.Type, jsonBlob)
	return bus.nsqBus.Publish(bus.getTopicName(envelope.Type), jsonBlob)
}

func (bus *EventBus) getTopicName(eventType Type) string {
	return bus.applicationName + "_" + eventType.String()
}
