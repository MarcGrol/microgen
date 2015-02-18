# Microgen

## Goal
Experiment with microservices using go

## Functionality
A tour-de-france gambling application:
- As an administrator: Create a tour for a particular year
- As an administrator: Add cyclists to this tour
- As an administrator: Add etappes to this tour
- As an administrator: Publish dayly result of etappe
- As a gambler: create a profile 
- As a gambler: compose your own team of cyclists for a particular year
- As a gambler: View etappes results
- As a gambler: View best cyclists
- As a gambler: View best gamblers

## Concept
An "application" consists of the following concepts:
 - "service": one or more loosely couples "services"
 - "command": each service supports zero or more commands. Actors interact with the system via these commands (and queries) on services.
 - "event" (with "attributes"): each commands emits zero or more events. An event is used to exchange of information between services in an async away

##Technical solution
- describe application in terns of "services", "commands" and "events". Example: [https://github.com/MarcGrol/microgen/blob/master/application.go]
- generate events and interfaces based on "description" of application. Example: tourApp/events/events.go and tourApp/tour/interface.go
- provide implementation for "bus" (=exchange of events between services)
- provide implementation of append-ony "store" for persistence
- declarative way of testing services. Based on this test spec, documentation and relationships between services can be derived. Example: tourApp/tour/logic_test.go
