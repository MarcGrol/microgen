# Microgen

## Goal
Experiment with microservices using go

## Functionality
A tour-de-france gambling application:
- As an administrator: Create a tour for a particular year (1)
- As an administrator: Add cyclists to a tour (2)
- As an administrator: Add etappes to a tour (3)
- As an administrator: Publish dayly result of etappes and calculate scores for cyclists and gamblers (4)
- As an administrator: Keep track of everything that has happened within the system (5)
- As a gambler: Create a profile (21)
- As a gambler: Compose your own team of cyclists for a particular year (22)
- As a gambler: View scores of my cylists (23)
- As anybody: View tours with their cyclists (31)
- As anybody: View etappes and results (32)
- As anybody: View best cyclists of a tour (33)
- As anybody: View best gamblers of a tour (34)
- As a system: Notify gambler of new day results (41)

## Devision of functions in services
### Tour-service
Responsible for managing tours with their etappes and cyclists (1,2,3,31)

### Gambler-service
Responsible for gamblers and their teams of cyclists (21,22,23)

### Results-service
Responsible for calculating and presenting results and scores (4,32,33,34,41)

### Collector-service
Non functional service that keeps track of everything that has happened. (5)


## Concept
An "application" consists of the following concepts:
 - "service": one or more loosely couples "services"
 - "commands" and "queries"(with "attributes"): Each service supports zero or more commands and quries. Browsers interact with the system via these commands and queries on a services.
 - "event" (with "attributes"): Each command emits zero or more events. An event is used to exchange of information between services in an async away.

##Technical solution
- Use a "dsl" to describe your application in terns of "services", "commands", "queries" and "events". Example: [application.go](./application.go)
- Generate events, interfaces based and a system-overview based on the "description" of application. Example: [events.go](./tourApp/events/events.go) and [interface.go](./tourApp/gambler/interface.go). This to achieve consistent approach and ease error phrone tasks.
- Provide implementation for "bus" (=to exchange of events between services)
- Provide implementation of append-ony "store" for persistence
- Provide implementation of http handler to process commands.
- Provide a mechanism for starting and configuring services.
- Declarative way of testing services. Based on this test spec, documentation and relationships between services can be derived. Example: [logic_test.go](./tourApp/tour/logic_test.go)
- Provide clear and exact documentation that explains how services are related. Example: [graphviz.dot](./tourApp/doc/graphviz.pdf)
- For each test scenario the "given", "when" and "expect" are recorded and written file. This describes the scenario exactly in json. Example: [scenario_example.json](./tourApp/doc/example_Create_new_gambler_success.txt)

## Obtaining, building, running and testing

    # prepare
    go get github.com/MarcGrol/microgen
    cd ${GOPATH}/src/github.com/MarcGrol/microgen
    go fmt ./...            # to format code
    go test ./...           # to run all unit tests
    go install              # to create executable
    
    # Sync source-code with application-dsl
    ${GOPATH}/bin/microgen -tool=gen -base-dir=. # to generate interfaces based in ./application.go
    
    # Start the bus
    ./bus/start_nsq.sh
    
    # Start the application
    ${GOPATH}/bin/microgen -service=tour      -port=8081 -base-dir=.
    ${GOPATH}/bin/microgen -service=gambler   -port=8082 -base-dir=.
    ${GOPATH}/bin/microgen -service=score     -port=8083 -base-dir=.
    ${GOPATH}/bin/microgen -service=collector -port=8084 -base-dir=.
    ${GOPATH}/bin/microgen -service=proxy     -port=8080 -base-dir=.
    
    # Fire commands into the application
    curl -X POST --header "Content-type: application/json"  --header "Accept: application/json" --data '{"year":2015}' "http://localhost:8081/api/tour"
    curl -X GET  --header "Accept: application/json"  "http://localhost:8081/api/tour/2015"
    

