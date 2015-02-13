package gambler

// Generated automatically: do not edit manually

import (
	"github.com/xebia/microgen/events"
)

type CreateGamblerCommand struct {
	GamblerUid string `json:"gamblerUid"`
	Name       string `json:"name"`
	Email      string `json:"email"`
}

type CreateGamblerTeamCommand struct {
	GamblerUid string `json:"gamblerUid"`
	Year       int    `json:"year"`
	CyclistIds []int  `json:"cyclistIds"`
}

type CommandHandler interface {
	HandleCreateGamblerCommand(command CreateGamblerCommand) ([]*events.Envelope, error)

	HandleCreateGamblerTeamCommand(command CreateGamblerTeamCommand) ([]*events.Envelope, error)
}
