package gambler

// Generated automatically by microgen: do not edit manually

import (
    
    "github.com/MarcGrol/microgen/tourApp/events"
    "github.com/MarcGrol/microgen/myerrors"
)

// commands


   
	type CreateGamblerCommand struct {
	 GamblerUid  string `json:"gamblerUid" binding:"required"`
	 Name  string `json:"name" binding:"required"`
	 Email  string `json:"email" binding:"required"`
	
	}

    
  

   
	type CreateGamblerTeamCommand struct {
	 GamblerUid  string `json:"gamblerUid" binding:"required"`
	 Year  int `json:"year" binding:"required"`
	 CyclistIds [] int `json:"cyclistIds" `
	
	}

    
  

  


type CommandHandler interface {
    
    	
        HandleCreateGamblerCommand ( command CreateGamblerCommand ) *myerrors.Error
    	
    
    	
        HandleCreateGamblerTeamCommand ( command CreateGamblerTeamCommand ) *myerrors.Error
    	
    
    	 
        HandleGetGamblerQuery (  gamblerUid string,  year int,  ) (*Gambler, *myerrors.Error)
    	
    
}

// events

type EventHandler interface {
     OnTourCreated ( event events.TourCreated ) error
     OnCyclistCreated ( event events.CyclistCreated ) error
    
}

type EventApplier interface {
     ApplyTourCreated ( event events.TourCreated )
     ApplyGamblerCreated ( event events.GamblerCreated )
     ApplyCyclistCreated ( event events.CyclistCreated )
     ApplyGamblerTeamCreated ( event events.GamblerTeamCreated )
    
}

