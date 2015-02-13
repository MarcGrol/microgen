package gambler;

// Generated automatically: do not edit manually

import (
    "github.com/xebia/microgen/events"
)

type EventHandler interface {
     OnTourCreated ( event events.TourCreated ) ([]*events.Envelope,error)
    
}

type EventApplier interface {
     ApplyGamblerTeamCreated ( event events.GamblerTeamCreated ) error
     ApplyGamblerCreated ( event events.GamblerCreated ) error
     ApplyTourCreated ( event events.TourCreated ) error
    
}

