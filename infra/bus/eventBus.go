package bus

import (
	"encoding/json"
	"github.com/MarcGrol/microgen/infra"
	"github.com/MarcGrol/microgen/lib/envelope"
	"log"
)

type EventBus struct {
	nsqBus          *NsqBus
	applicationName string
}

func NewEventBus(applicationName string, consumerName string, address string) *EventBus {
	mybus := new(EventBus)
	mybus.nsqBus = NewNsqBus(consumerName, address)
	mybus.applicationName = applicationName
	return mybus
}

func (bus *EventBus) Subscribe(eventTypeName string, userCallback infra.EventHandlerFunc) error {
	envelop := new(envelope.Envelope)
	callback := func(blob []byte) error {
		err := json.Unmarshal(blob, envelop)
		if err != nil {
			log.Printf("Error unmarshalling json blob (%+v)", err)
			return err
		}

		log.Printf("**** Received event (%+v)", envelop)
		return userCallback(envelop)
	}
	return bus.nsqBus.Subscribe(bus.getTopicName(eventTypeName), callback)
}

func (bus *EventBus) Publish(envelope *envelope.Envelope) error {
	jsonBlob, err := json.Marshal(envelope)
	if err != nil {
		log.Printf("Error marshalling event-envelope (%+v)", err)
		return err
	}
	//log.Printf("Marshalled event of type %d (%s)", envelope.Type, jsonBlob)
	err = bus.nsqBus.Publish(bus.getTopicName(envelope.EventTypeName), jsonBlob)
	if err != nil {
		log.Printf("Error publishing event-envelope (%+v)", err)
		return err
	}

	log.Printf("**** Published event (%+v)", envelope)
	return nil
}

func (bus *EventBus) getTopicName(eventName string) string {
	return bus.applicationName + "_" + eventName
}
