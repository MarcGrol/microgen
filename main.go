package main

import (
	"flag"
	"fmt"
	"github.com/MarcGrol/microgen/spec"
	"github.com/MarcGrol/microgen/tourApp/tour"
	"log"
	"os"
)

const (
	VERSION = "0.1"
)

var modus *string
var service *string
var httpPort *int
var busAddress *string

func printUsage() {
	fmt.Fprintf(os.Stderr, "\nUsage:\n")
	fmt.Fprintf(os.Stderr, " %s [flags]\n", os.Args[0])
	flag.PrintDefaults()
	fmt.Fprintf(os.Stderr, "\n")
	os.Exit(1)
}

func printVersion() {
	fmt.Fprintf(os.Stderr, "\nVersion: %s\n", VERSION)
	os.Exit(1)
}

func processArgs() {
	modus = flag.String("mode", "service", "Mode in which tools runs: tool or service")
	service = flag.String("service", "", "service to run: tour, gambler or result")
	httpPort = flag.Int("port", 8081, "listen port of http-server")
	busAddress = flag.String("bus-address", "localhost", "Hostname where nsq-bus is running")

	help := flag.Bool("help", false, "Usage information")
	version := flag.Bool("version", false, "Version information")

	flag.Parse()

	if help != nil && *help == true {
		printUsage()
	}
	if version != nil && *version == true {
		printVersion()
	}
}

func main() {
	processArgs()

	if *modus == "service" {
		if service == nil {
			fmt.Fprintf(os.Stderr, "Missing service name")
			printUsage()
		} else {
			if *service == "tour" {
				log.Printf("Starting tour")
				tour.Start(*httpPort, *busAddress)
			} else if *service == "gambler" {
				log.Printf("Starting gambler")
			} else if *service == "results" {
				log.Printf("Starting results")
			} else {
				fmt.Fprintf(os.Stderr, "Unrecognized service name %s", *service)
				printUsage()
			}
		}
	} else if *modus == "tool" {
		err := spec.GenerateApplication(application, ".")
		if err != nil {
			log.Fatalf("Error generating application %s", err)
		}
	}
}
