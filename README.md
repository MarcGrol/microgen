# Microgen

## Goal
Experiment with microservices using go

## Functionality
A tour-de-france gambling application:
1. As an administrator: Create a tour for a particular year
2. As an administrator: Add cyclists to this tour
3. As an administrator: Add etappes to this tour
4. As a gambler: Create a profile 
5. As a gambler: Compose your own team of cyclists for a particular year
6. As an administrator: Publish dayly result of etappes and calculate scores for cyclists and gamblers
7. As anybody: View tours with their cyclists, etappes and results
8. As anybody: View best cyclists of a tour
9. As anybody: View best gamblers of a tour

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

