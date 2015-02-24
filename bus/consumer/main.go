package main

import (
	"log"
	"sync"
	"github.com/bitly/go-nsq"
)

func main() {
	wg := &sync.WaitGroup{}
	wg.Add(1)
	config := nsq.NewConfig()
	q, _ := nsq.NewConsumer("tourApp_CyclistCreated", "gambler", config)
	q.AddHandler(nsq.HandlerFunc(func(message *nsq.Message) error {
		log.Printf("Got a message: %v", message)
	      wg.Done()
		return nil
	}))

	err := q.ConnectToNSQLookupd("127.0.0.1:4161")
	if err != nil {
		log.Panic("Could not connect")
	}

	wg.Wait()
}
