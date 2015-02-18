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
Describe your "application" in terms of:
 - "service": one or more loosely couples "services"
 - "command": each service supports zero or more commands. Actors interact with the system via these commands (and queries)
 - "event" (with "attributes"): each commands emits zero or more events. An event is used to exchange of information between services in an async away

##Technical solution
- describe application in terns of "services", "commands" and "events". Example: application.go
- generate events and interfaces based on "description" of application. Example: tourApp/events/events.go and tourApp/tour/interface.go
- provide implementation for Bus (=exchange of events between services)
- provide implementation of append-ony store for persistence
