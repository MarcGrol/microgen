# Microgen

## Goal
Experiment with microservices using go

## Functionality
A tour-de-france gambling application:
- As an administrator: Create a tour for a particular year (1)
- As an administrator: Add cyclists to a tour (2)
- As an administrator: Add etappes to a tour (3)
- As a gambler: Create a profile (4)
- As a gambler: Compose your own team of cyclists for a particular year (5)
- As an administrator: Publish dayly result of etappes and calculate scores for cyclists and gamblers (6)
- As anybody: View tours with their cyclists, etappes and results (7)
- As anybody: View best cyclists of a tour (8)
- As anybody: View best gamblers of a tour (9)

## Devision of functions in services
### Tour-service
Responsible for managing tours with their etappes and cyclists (1,2,3,7,8,9)

### Gambler-service
Responsible for gamblers and their teams of cyclists (4,5)

### Results-service
Responsible for handling results calculating scores (6)

## Concept
An "application" consists of the following concepts:
 - "service": one or more loosely couples "services"
 - "commands" and "queries"(with "attributes"): Each service supports zero or more commands and quries. Browsers interact with the system via these commands and queries on a services.
 - "event" (with "attributes"): Each command emits zero or more events. An event is used to exchange of information between services in an async away.

##Technical solution
- Use a "dsl" to describe your application in terns of "services", "commands", "queries" and "events". Example: [application.go](./application.go)
- Generate events, interfaces based and a system-overview based on the "description" of application. Example: [events.go](./tourApp/events/events.go) and [interface.go](./tourApp/tour/interface.go). This to achieve consistent approach and ease error phrone tasks.
- Provide implementation for "bus" (=to exchange of events between services)
- Provide implementation of append-ony "store" for persistence
- Provide implementation of http handler to process commands.
- Provide a mechanism for starting and configuring services.
- Declarative way of testing services. Based on this test spec, documentation and relationships between services can be derived. Example: [logic_test.go](./tourApp/tour/logic_test.go)
- Provide clear and exact documentation that explains how services are related. Example: [logic_test.go](./tourApp/doc/graphviz.pdf)
- For each test scenario the "given", "when" and "expect" are recorded and written to file (./tourapp/doc/). This describes the scenario exactly in json.

## Obtaining, building, running and testing

    go get github.com/MarcGrol/microgen
    cd ${GOPATH}/src/github.com/MarcGrol/microgen
    go fmt ./...            # to format code
    go test ./...           # to run all unit tests
    go install              # to create executable
    ${GOPATH}/bin/microgen -tool=gen # to generate interfaces based in ./application.go
    
    # Start the bus
    ./bus/start_nsq.sh
    
    # Start the application
    ${GOPATH}/bin/microgen -service=tour    -port=8081 -base-dir=.
    ${GOPATH}/bin/microgen -service=gambler -port=8082 -base-dir=.
    ${GOPATH}/bin/microgen -service=score   -port=8083 -base-dir=.
    ${GOPATH}/bin/microgen -service=proxy   -port=8080 -base-dir=.
    
    # Fire comannds into the application
    curl -X POST --header "Content-type: application/json"  --header "Accept: application/json" --data '{"year":2015}' "http://localhost:8081/api/tour"
    curl -X GET  --header "Accept: application/json"  "http://localhost:8081/api/tour/2015"
    

