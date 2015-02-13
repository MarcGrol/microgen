package gambler

// Generated automatically by microgen: do not edit manually

import (
    
    "github.com/xebia/microgen/events"
)

// commands


type CreateGamblerCommand struct {
 GamblerUid  string `json:"gamblerUid"`
 Name  string `json:"name"`
 Email  string `json:"email"`

}

type CreateGamblerTeamCommand struct {
 GamblerUid  string `json:"gamblerUid"`
 Year  int `json:"year"`
 CyclistIds [] int `json:"cyclistIds"`

}


type CommandHandler interface {
    
        HandleCreateGamblerCommand ( command CreateGamblerCommand ) ([]*events.Envelope,error)
    
        HandleCreateGamblerTeamCommand ( command CreateGamblerTeamCommand ) ([]*events.Envelope,error)
    
}

// events

type EventHandler interface {
     OnTourCreated ( event events.TourCreated ) ([]*events.Envelope,error)
    
}

type EventApplier interface {
     ApplyTourCreated ( event events.TourCreated ) error
     ApplyGamblerTeamCreated ( event events.GamblerTeamCreated ) error
     ApplyGamblerCreated ( event events.GamblerCreated ) error
    
}

