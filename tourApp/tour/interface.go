package tour

// Generated automatically by microgen: do not edit manually

import (
    "time"
    "github.com/MarcGrol/microgen/tourApp/events"
)

// commands


   
	type CreateTourCommand struct {
	 Year  int `json:"year" binding:"required"`
	
	}

    
  

   
	type CreateCyclistCommand struct {
	 Year  int `json:"year" binding:"required"`
	 Id  int `json:"id" binding:"required"`
	 Name  string `json:"name" binding:"required"`
	 Team  string `json:"team" binding:"required"`
	
	}

    
  

   
	type CreateEtappeCommand struct {
	 Year  int `json:"year" binding:"required"`
	 Id  int `json:"id" binding:"required"`
	 Date  time.Time `json:"date" binding:"required"`
	 StartLocation  string `json:"startLocation" binding:"required"`
	 FinishLocation  string `json:"finishLocation" binding:"required"`
	 Length  int `json:"length" binding:"required"`
	 Kind  int `json:"kind" binding:"required"`
	
	}

    
  

  


type CommandHandler interface {
    
    	
        HandleCreateTourCommand ( command CreateTourCommand ) error
    	
    
    	
        HandleCreateCyclistCommand ( command CreateCyclistCommand ) error
    	
    
    	
        HandleCreateEtappeCommand ( command CreateEtappeCommand ) error
    	
    
    	 
        HandleGetTourQuery (  year int,  ) (*Tour, error)
    	
    
}

// events

type EventHandler interface {
    
}

type EventApplier interface {
     ApplyTourCreated ( event events.TourCreated )
     ApplyCyclistCreated ( event events.CyclistCreated )
     ApplyEtappeCreated ( event events.EtappeCreated )
    
}

