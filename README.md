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
 - "command" (with "attributes"): Each service supports zero or more commands. Actors interact with the system via these commands (and queries) on services.
 - "event" (with "attributes"): Each commands emits zero or more events. An event is used to exchange of information between services in an async away

##Technical solution
- Describe application in terns of "services", "commands" and "events". Example: [application.go](./application.go)
- Generate events and interfaces based on "description" of application. Example: [events.go](./tourApp/events/events.go) and [interface.go](./tourApp/tour/interface.go). This to achieve consistent approach and ease error phrone tasks.
- Provide implementation for "bus" (=to exchange of events between services)
- Provide implementation of append-ony "store" for persistence
- Provide implementation of http handler to process commands.
- Provide a mechanism for starting and configuring services.
- Declarative way of testing services. Based on this test spec, documentation and relationships between services can be derived. Example: [logic_test.go](./tourApp/tour/logic_test.go)
- Provide clear and exact documentation that explains how services are related.

## Running and testing

    go fmt ./...
    go test ./...
    go install
    ${GOPATH}/bin/microgen -service=tour
    
    curl -X POST --header "Content-type: application/json" --data '{"year":2015}' "http://localhost:8081/api/tour"
    

