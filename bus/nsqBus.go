package bus

import (
	"encoding/json"
	"github.com/bitly/go-nsq"
	"github.com/xebia/microgen/events"
	"log"
)

type NsqBus struct {
	applicationName string
	consumerName    string
	address         string
	config          *nsq.Config
	producer        *nsq.Producer
	consumers       []*nsq.Consumer
}

func NewNsqBus(applicationName string, consumerName string, address string) *NsqBus {
	bus := new(NsqBus)
	bus.applicationName = applicationName
	bus.consumerName = consumerName
	bus.address = address
	bus.config = nsq.NewConfig()
	bus.consumers = make([]*nsq.Consumer, 0, 10)
	return bus
}

func (bus *NsqBus) Subscribe(eventType events.Type, callback events.EventHandlerFunc) error {
	return bus.startConsumer(bus.getTopicName(eventType), callback)
}

func (bus *NsqBus) startConsumer(topic string, userCallback events.EventHandlerFunc) error {
	consumer, err := nsq.NewConsumer(topic, bus.consumerName, bus.config)
	if err != nil {
		log.Printf("Error creating nsq consumer %s/%s (%v)", topic, bus.consumerName, err)
		return err
	}

	callback := func(message *nsq.Message) error {
		envelope := events.Envelope{Type: events.TypeUnknown}
		err := json.Unmarshal(message.Body, &envelope)
		if err != nil {
			log.Printf("Error unmarshalling event-envelope (%v)", err)
			return nil
		}
		err = userCallback(&envelope)
		if err != nil {
			log.Printf("Error handling event-envelope (%v)", err)
		}
		return err
	}
	consumer.AddHandler(nsq.HandlerFunc(callback))

	err = consumer.ConnectToNSQLookupd(bus.address + ":4161")
	if err != nil {
		log.Printf("Error connecting to lookupd (%v)", err)
		return err
	}

	bus.consumers = append(bus.consumers, consumer)

	//log.Printf("Started consumer %s/%s", topic, bus.consumerName)

	return nil
}

func (bus *NsqBus) Publish(envelope *events.Envelope) error {
	if bus.producer == nil {
		err := bus.startProducer()
		if err != nil {
			return err
		}
	}
	jsonBlob, err := json.Marshal(envelope)
	if err != nil {
		log.Printf("Error marshalling event-envelope (%v)", err)
		return err
	}
	//log.Printf("Marshalled event of type %d (%s)", envelope.Type, jsonBlob)

	err = bus.producer.Publish(bus.getTopicName(envelope.Type), jsonBlob)
	if err != nil {
		log.Printf("Error publishing event-envelope (%v)", err)
		return err
	}

	//log.Printf("Published event of type %d (%s)", envelope.Type, jsonBlob)
	return nil
}

func (bus *NsqBus) startProducer() error {
	var err error
	bus.producer, err = nsq.NewProducer(bus.address+":4150", bus.config)
	if err != nil {
		log.Printf("Error creating nsq producer (%v)", err)
		return err
	}

	//log.Printf("Started producer")

	return nil
}

func (bus *NsqBus) getTopicName(eventType events.Type) string {
	return bus.applicationName + "_" + eventType.String()
}
