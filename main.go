package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/MarcGrol/microgen/tool/dsl"
	"github.com/MarcGrol/microgen/tourApp/collector"
	"github.com/MarcGrol/microgen/tourApp/gambler"
	"github.com/MarcGrol/microgen/tourApp/news"
	"github.com/MarcGrol/microgen/tourApp/prov"
	"github.com/MarcGrol/microgen/tourApp/proxy"
	"github.com/MarcGrol/microgen/tourApp/tour"
)

const (
	VERSION = "0.1"
)

var tool *string
var service *string
var httpPort *int
var address *string
var baseDir *string
var targetHostPort *string

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
	tool = flag.String("tool", "", "Run in 'tool-mode: 'gen' pr 'prov'")
	baseDir = flag.String("base-dir", ".", "For modus 'tool': Base directory used in both 'tool' and 'service'-modus")
	service = flag.String("service", "", "For modus 'service': service to run: 'tour', 'gambler','news', 'proxy' or 'collector'")
	httpPort = flag.Int("port", 8081, "For modus 'service': listen port of http-server")
	address = flag.String("address", "localhost", "For modus 'service': Hostname where application running")
	targetHostPort = flag.String("target-host", "localhost:8080", "For tool 'prov': Hostname where the application is running")

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

		switch *service {

		case "tour":
			err := tour.Start(*httpPort, *address, *baseDir)
			if err != nil {
				log.Fatalf("Error starting 'tour'-service on port %d, address:%s and base-dir: %s (%+v)",
					*httpPort, *address, *baseDir, err)
			}

		case "gambler":
			err := gambler.Start(*httpPort, *address, *baseDir)
			if err != nil {
				log.Fatalf("Error starting 'gambler'-service on port %d, address:%s and base-dir: %s (%+v)",
					*httpPort, *address, *baseDir, err)
			}

		case "news":
			err := news.Start(*httpPort, *address, *baseDir)
			if err != nil {
				log.Fatalf("Error starting 'news'-service on port %d, address:%s and base-dir: %s (%+v)",
					*httpPort, *address, *baseDir, err)
			}

		case "collector":
			err := collector.Start(*httpPort, *address, *baseDir)
			if err != nil {
				log.Fatalf("Error starting 'collector'-service on port %d, bus-address:%s and base-dir: %s (%+v)",
					*httpPort, *address, *baseDir, err)
			}

		case "proxy":
			err := proxy.Start(*baseDir, *httpPort, *address, 8081, 8082, 8083, 8084)
			if err != nil {
				log.Fatalf("Error starting 'proxy'-service on port %d (%+v)", *httpPort, err)
			}

		default:
			fmt.Fprintf(os.Stderr, "Unrecognized service name %s", *service)
			printUsage()
		}

	} else if len(*tool) > 0 {

		switch *tool {

		case "gen":
			err := dsl.GenerateApplication(application, *baseDir)
			if err != nil {
				log.Fatalf("Error generating application: %s", err)
			}

		case "prov":
			err := prov.Start(*targetHostPort)
			if err != nil {
				log.Fatalf("Error provisioning application: %s", err)
			}

		default:
			fmt.Fprintf(os.Stderr, "Unrecognized tool name: %s", *tool)
			printUsage()

		}

	} else {
		fmt.Fprintf(os.Stderr, "Unrecognized command")
		printUsage()
	}
}
