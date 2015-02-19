package tour

// Generated automatically by microgen: do not edit manually

import (
    "time"
    "github.com/MarcGrol/microgen/tourApp/events"
    "github.com/MarcGrol/microgen/myerrors"
)

// commands


type CreateTourCommand struct {
 Year  int `json:"year"`

}

type CreateCyclistCommand struct {
 Year  int `json:"year"`
 Id  int `json:"id"`
 Name  string `json:"name"`
 Team  string `json:"team"`

}

type CreateEtappeCommand struct {
 Year  int `json:"year"`
 Id  int `json:"id"`
 Date  time.Time `json:"date"`
 StartLocation  string `json:"startLocation"`
 FinishLocation  string `json:"finishLocation"`
 Length  int `json:"length"`
 Kind  int `json:"kind"`

}


type CommandHandler interface {
    
        HandleCreateTourCommand ( command CreateTourCommand ) *myerrors.Error
    
        HandleCreateCyclistCommand ( command CreateCyclistCommand ) *myerrors.Error
    
        HandleCreateEtappeCommand ( command CreateEtappeCommand ) *myerrors.Error
    
}

// events

type EventHandler interface {
    
}

type EventApplier interface {
     ApplyEtappeCreated ( event events.EtappeCreated ) *myerrors.Error
     ApplyTourCreated ( event events.TourCreated ) *myerrors.Error
     ApplyCyclistCreated ( event events.CyclistCreated ) *myerrors.Error
    
}

