# Microgen

## Goal
Experiment with microservices using go

## Functionality
A tour-de-france gambling application:
- As an administrator: Create a tour for a particular year
- As an administrator: Add cyclists to this tour
- As an administrator: Add etappes to this tour
- As an administrator: Publish dayly result of etappe
- As a gambler: Create a profile 
- As a gambler: Compose your own team of cyclists for a particular year
- As a gambler: View etappes results
- As a gambler: View best cyclists
- As a gambler: View best gamblers

## Concept
An "application" consists of the following concepts:
 - "service": one or more loosely couples "services"
 - "command" (with "attributes"): Each service supports zero or more commands. Actors interact with the system via these commands (and queries) on services.
 - "event" (with "attributes"): Each commands emits zero or more events. An event is used to exchange of information between services in an async away

##Technical solution
- Describe application in terns of "services", "commands" and "events". Example: [https://github.com/MarcGrol/microgen/blob/master/application.go]
- generate events and interfaces based on "description" of application. Example: [https://github.com/MarcGrol/microgen/blob/master/tourApp/events/events.go] and [https://github.com/MarcGrol/microgen/blob/master/tourApp/tour/interface.go]
- Provide implementation for "bus" (=exchange of events between services)
- Provide implementation of append-ony "store" for persistence
- Declarative way of testing services. Based on this test spec, documentation and relationships between services can be derived. Example: [https://github.com/MarcGrol/microgen/blob/master/tourApp/tour/logic_test.go]
- Provide clear and exact documentation that explains how services are related.
