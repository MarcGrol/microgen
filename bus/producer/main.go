package main


import (

  "log"

  "fmt"

  "github.com/bitly/go-nsq"

)


func main() {

  config := nsq.NewConfig()

  w, _ := nsq.NewProducer("127.0.0.1:4150", config)


for i:=0; i<10000; i++ {

  err := w.Publish("mytopic", []byte(fmt.Sprintf("this is test message %d", i)))

  if err != nil {

      log.Panic("Could not connect")

  }

}


  w.Stop()

}