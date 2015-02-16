package main

import (
	"github.com/xebia/microgen/spec"
	"log"
)

func main() {
	err := spec.GenerateApplication(application, ".")
	if err != nil {
		log.Fatalf("Error generating application %s", err)
	}
}
