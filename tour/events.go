package tour;

// Generated automatically: do not edit manually

import (
    "github.com/xebia/microgen/events"
)

type EventHandler interface {
    
}

type EventApplier interface {
     ApplyTourCreated ( event events.TourCreated ) error
     ApplyCyclistCreated ( event events.CyclistCreated ) error
     ApplyEtappeCreated ( event events.EtappeCreated ) error
    
}

