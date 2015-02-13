package tour;

// Generated automatically: do not edit manually

import (
    "github.com/xebia/microgen/events"
)

type EventHandler interface {
    
}

type EventApplier interface {
     ApplyCyclistCreated ( event events.CyclistCreated ) error
     ApplyEtappeCreated ( event events.EtappeCreated ) error
     ApplyTourCreated ( event events.TourCreated ) error
    
}

