package results

// Generated automatically: do not edit manually

import (
	"github.com/xebia/microgen/events"
)

type CreateDayResultsCommand struct {
	Year                   int   `json:"year"`
	Id                     int   `json:"id"`
	BestDayCyclistIds      []int `json:"bestDayCyclistIds"`
	BestAllroundCyclistIds []int `json:"bestAllroundCyclistIds"`
	BestClimbCyclistIds    []int `json:"bestClimbCyclistIds"`
	BestSprintCyclistIds   []int `json:"bestSprintCyclistIds"`
}

type CommandHandler interface {
	HandleCreateDayResultsCommand(command CreateDayResultsCommand) ([]*events.Envelope, error)
}
