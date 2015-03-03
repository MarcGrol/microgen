# Microgen

## Goal
Experiment with microservices using go. 

## Why?
The core of "micro" is to keep services small so you do not have to understand the full big monolyth before you can be effective. Another advantage of this approach that natural boundaries between services make it easy to keep the system modular. Finally, the application is easy to scale over multiple machines.

However, in the end you still have to understand the "whole". Clear overview of dependencies between services could help here. Also eventual consistancy introduces new complexity.

## Approach
Specification effort:
* Define your screens 
* Based on screens, determine your commands and queries
* Group related commands and queries in a service if this promotes cohesion
* Define your events, that exchange information between services.
* Specify which events a command emits
* Specify on which events a service depends

Develop non-functional components:
* append-only event-store
* publish-subscribe bus
* Scenario based testing: "given-when-expect"

Develop functional services:
 * create unit tests
 * implement business-logic: handle commands and events
 * attach service to http and event-bus

Build screens from the provided services:
 * create end-to-end tests
 * develop screens using services

## Functionality
A tour-de-france gambling application:
- As an administrator: 
    - Create a tour for a particular year (1)
    - Add cyclists to a tour (2)
    - Add etappes to a tour (3)
    - Publish dayly result of etappes and calculate scores for cyclists and gamblers (4)
    - Keep track of everything that has happened within the system (5)
- As a gambler: 
    - Create a profile (21)
    - Compose your own team of cyclists for a particular year (22)
    - View scores of my cylists (23)
- As anybody: 
    - View tours with their cyclists (31)
    - View etappes and results (32)
    - View best cyclists of a tour (33)
    - View best gamblers of a tour (34)
- As a system: 
    - Notify gamblers of new day results (41)

## Devision of functions in services
### Tour-service
Responsible for managing tours with their etappes and cyclists (1,2,3,4)

### Gambler-service
Responsible for gamblers and their teams of cyclists (21,22,23)

### Score-service
Responsible for calculating and presenting results and scores (31,32, 33,34,41)

### Collector-service
Non functional service that keeps track of everything that has happened. (5)

### Proxy-service
Responsible for serving the web-UI and hiding all services behind a single http-endpoint. Could be the place to apply your non-functionals (ssl-offloading, security, logging, statistics, etc).

## Concept
An "application" consists of the following concepts:
 - "service": An application cosnsts of one or more loosely couples "services". each service is responsible for keeping is own data.
 - "commands" and "queries"(with "attributes"): Each service supports zero or more commands and queries. Browsers interact with the system via these commands and queries on a services.
 - "event" (with "attributes"): Each command emits zero or more events. An event is used to exchange of information between services in an async away.

##Technical solution
- Use a "dsl" to describe your application in terns of "services", "commands", "queries" and "events". Example: [application.go](./application.go). Based and this system-description some parts of the system are generated. Example: [events.go](./tourApp/events/events.go) and [interface.go](./tourApp/gambler/interface.go). This approach is chosen to achieve consistent approach and ease error phrone tasks.
- Use event sourcing. Store events and rebuild current state by replaying events ralated to an aggregate in the order of arrival. Currently the following aggregate are recognized: Tour and Gambler.
- Could easily use command-Query seperation if scale requires it.
- Provided implementation for "bus" (=to exchange of events between services). Current implementatio is based on NSQ.
- Provided implementation of append-ony "store" for persistence. Current implementation uses the filesystem.
- Provided implementation of http handler to process commands. Current solution is based on gin-gonic.
- Provide a mechanism for starting and configuring services. Current solution compiles into a single executable. This executable can be configured (using command-line flags) to acts a service a, b or c. This to ease the distribution and deployment.
- Declarative way of testing services. Example: [logic_test.go](./tourApp/tour/logic_test.go). Based on this test spec, documentation and relationships between services can be derived. In addition to this, each test scenario (= "given", "when" and "expect") is recorded and written to file. This file describes the scenario exactly in json. Example: [scenario_example.json](./tourApp/doc/example_Create_new_gambler_success.txt)
- Provide clear, exact and up to date documentation that explains how services are related. Example:  [graphviz.dot](./tourApp/doc/graphviz.dot) and [graphviz.pdf](./tourApp/doc/graphviz.pdf)

## Obtaining, building, running and testing

    # prepare (has following dependencies (use go list -f {{.Deps}}) )
    go get github.com/bitly/go-nsq
    # go get github.com/bitly/go-simplejson
    go get github.com/gin-gonic/gin
    # go get github.com/gin-gonic/gin/binding
    # go get github.com/gin-gonic/gin/render
    # go get github.com/julienschmidt/httprouter
    # go get github.com/mreiferson/go-snappystream
    go get code.google.com/p/go-uuid/uuid
    # go get code.google.com/p/snappy-go/snappy

    go get github.com/MarcGrol/microgen
    cd ${GOPATH}/src/github.com/MarcGrol/microgen
    go test ./...           # to run all unit tests
    go install              # to create executable
    
    # Sync source-code with application-dsl
    ${GOPATH}/bin/microgen -tool=gen -base-dir=. # to generate interfaces based on ./application.go
    go fmt ./...            # to re-format generated code
    
    # Start the bus
    ./bus/start_nsq.sh
    
    # Start the application
    ${GOPATH}/bin/microgen -service=tour      -port=8081 -base-dir=.
    ${GOPATH}/bin/microgen -service=gambler   -port=8082 -base-dir=.
    ${GOPATH}/bin/microgen -service=score     -port=8083 -base-dir=.
    ${GOPATH}/bin/microgen -service=collector -port=8084 -base-dir=.
    ${GOPATH}/bin/microgen -service=proxy     -port=8080 -base-dir=.
    
    # Fire commands into the application
    curl -X POST --header "Content-type: application/json"  --header "Accept: application/json" --data '{"year":2015}' "http://localhost:8080/api/tour"
    curl -X GET  --header "Accept: application/json"  "http://localhost:8080/api/tour/2015"
    

