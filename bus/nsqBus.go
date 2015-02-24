package bus

import (
	"github.com/bitly/go-nsq"
	"log"
)

type NsqBus struct {
	consumerName string
	address      string
	config       *nsq.Config
	producer     *nsq.Producer
	consumers    []*nsq.Consumer
}

func NewNsqBus(consumerName string, address string) *NsqBus {
	bus := new(NsqBus)
	bus.consumerName = consumerName
	bus.address = address
	bus.config = nsq.NewConfig()
	bus.consumers = make([]*nsq.Consumer, 0, 10)
	return bus
}

type BlobHandlerFunc func(blob []byte) error

func (bus *NsqBus) Subscribe(topic string, callback BlobHandlerFunc) error {
	return bus.startConsumer(topic, callback)
}

func (bus *NsqBus) startConsumer(topic string, userCallback BlobHandlerFunc) error {
	log.Printf("Connecting using topic %s anc channel %s", topic, bus.consumerName)
	consumer, err := nsq.NewConsumer(topic, bus.consumerName, bus.config)
	if err != nil {
		log.Printf("Error creating nsq consumer %s/%s (%+v)", topic, bus.consumerName, err)
		return err
	}

	consumer.AddHandler(nsq.HandlerFunc(func(message *nsq.Message) error {
		err := userCallback(message.Body)
		if err != nil {
			log.Printf("Error handling event-envelope (%+v)", err)
		} else {
			log.Printf("Successfully handled event-envelope (%v+)", err)
		}
		return err
	}))

	err = consumer.ConnectToNSQLookupd(bus.address + ":4161")
	if err != nil {
		log.Printf("Error connecting to lookupd (%+v)", err)
		return err
	}

	bus.consumers = append(bus.consumers, consumer)

	//log.Printf("Started consumer %s/%s", topic, bus.consumerName)

	return nil
}

func (bus *NsqBus) Publish(topic string, blob []byte) error {
	if bus.producer == nil {
		err := bus.startProducer()
		if err != nil {
			return err
		}
	}

	err := bus.producer.Publish(topic, blob)
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
