package main

import (
	"github.com/MarcGrol/microgen/spec"
	"log"
)

func main() {
	err := spec.GenerateApplication(application, ".")
	if err != nil {
		log.Fatalf("Error generating application %s", err)
	}
}
