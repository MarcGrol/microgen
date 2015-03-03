package main

import (
	"flag"
	"fmt"
	"github.com/MarcGrol/microgen/dsl"
	"github.com/MarcGrol/microgen/tourApp/collector"
	"github.com/MarcGrol/microgen/tourApp/gambler"
	"github.com/MarcGrol/microgen/tourApp/proxy"
	"github.com/MarcGrol/microgen/tourApp/score"
	"github.com/MarcGrol/microgen/tourApp/tour"
	"log"
	"os"
)

const (
	VERSION = "0.1"
)

var tool *string
var service *string
var httpPort *int
var busAddress *string
var baseDir *string
var templateDir *string

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
	tool = flag.String("tool", "", "Run in 'tool-mode: 'gen'")
	templateDir = flag.String("template-dir", ".", "For 'tool'-mode: Directory where templates are located")
	baseDir = flag.String("base-dir", ".", "For modus 'tool': Base directory used in both 'tool' and 'service'-modus")
	service = flag.String("service", "", "For modus 'service': service to run: 'tour', 'gambler','score', 'proxy' or 'collector'")
	httpPort = flag.Int("port", 8081, "For modus 'service': listen port of http-server")
	busAddress = flag.String("bus-address", "localhost", "For modus 'service': Hostname where nsq-bus is running")

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
	// We use a single executable that can, based on cli-args, do everything
	// from running as service, proxy or act as code-generation tool
	// Advantage is that application wiuth all its services ships as a single executable
	if len(*service) > 0 {
		if *service == "tour" {
			err := tour.Start(*httpPort, *busAddress, *baseDir)
			if err != nil {
				log.Fatalf("Error starting 'tour'-service on port %d, bus-address:%s and base-dir: %s (%+v)",
					*httpPort, *busAddress, *baseDir, err)
			}
		} else if *service == "gambler" {
			err := gambler.Start(*httpPort, *busAddress, *baseDir)
			if err != nil {
				log.Fatalf("Error starting 'gambler'-service on port %d, bus-address:%s and base-dir: %s (%+v)",
					*httpPort, *busAddress, *baseDir, err)
			}
		} else if *service == "score" {
			err := score.Start(*httpPort, *busAddress, *baseDir)
			if err != nil {
				log.Fatalf("Error starting 'score'-service on port %d, bus-address:%s and base-dir: %s (%+v)",
					*httpPort, *busAddress, *baseDir, err)
			}
		} else if *service == "collector" {
			err := collector.Start(*httpPort, *busAddress, *baseDir)
			if err != nil {
				log.Fatalf("Error starting 'collector'-service on port %d, bus-address:%s and base-dir: %s (%+v)",
					*httpPort, *busAddress, *baseDir, err)
			}
		} else if *service == "proxy" {
			err := proxy.Start(*baseDir, *httpPort, 8081, 8082, 8083, 8084)
			if err != nil {
				log.Fatalf("Error starting 'proxy'-service on port %d", *httpPort, err)
			}
		} else {
			fmt.Fprintf(os.Stderr, "Unrecognized service name %s", *service)
			printUsage()
		}
	} else if len(*tool) > 0 {
		if *tool == "gen" {
			err := spec.GenerateApplication(application, *baseDir)
			if err != nil {
				log.Fatalf("Error generating application: %s", err)
			}
		} else {
			fmt.Fprintf(os.Stderr, "Unrecognized tool name: %s", *tool)
			printUsage()
		}
	} else {
		fmt.Fprintf(os.Stderr, "Unrecognized command")
		printUsage()
	}
}
