package results;

// Generated automatically: do not edit manually

import (
    "github.com/xebia/microgen/events"
)

type EventHandler interface {
     OnTourCreated ( event events.TourCreated ) ([]*events.Envelope,error)
     OnEtappeCreated ( event events.EtappeCreated ) ([]*events.Envelope,error)
     OnCyclistCreated ( event events.CyclistCreated ) ([]*events.Envelope,error)
     OnGamblerCreated ( event events.GamblerCreated ) ([]*events.Envelope,error)
     OnGamblerTeamCreated ( event events.GamblerTeamCreated ) ([]*events.Envelope,error)
    
}

type EventApplier interface {
     ApplyCyclistCreated ( event events.CyclistCreated ) error
     ApplyGamblerCreated ( event events.GamblerCreated ) error
     ApplyGamblerTeamCreated ( event events.GamblerTeamCreated ) error
     ApplyEtappeResultsAvailable ( event events.EtappeResultsAvailable ) error
     ApplyCyclistScoreCalculated ( event events.CyclistScoreCalculated ) error
     ApplyGamblerScoreCalculated ( event events.GamblerScoreCalculated ) error
     ApplyTourCreated ( event events.TourCreated ) error
     ApplyEtappeCreated ( event events.EtappeCreated ) error
    
}

