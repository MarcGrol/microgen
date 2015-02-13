package main

import (
	"github.com/xebia/microgen/gen"
	"log"
)

func main() {
	err := gen.GenerateApplication(application, ".")
	if err != nil {
		log.Fatalf("Error generating application %s", err)
	}
}
