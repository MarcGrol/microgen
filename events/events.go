package events

// Generated automatically by microgen: do not edit manually

import (
	"time"
)

type TourCreated struct {
	Year int `json:"year"`
}

func (event *TourCreated) Wrap() *Envelope {
	envelope := new(Envelope)
	envelope.Type = TypeTourCreated
	envelope.TourCreated = event
	return envelope
}

type CyclistCreated struct {
	Year        int    `json:"year"`
	CyclistId   int    `json:"cyclistId"`
	CyclistName string `json:"cyclistName"`
	CyclistTeam string `json:"cyclistTeam"`
}

func (event *CyclistCreated) Wrap() *Envelope {
	envelope := new(Envelope)
	envelope.Type = TypeCyclistCreated
	envelope.CyclistCreated = event
	return envelope
}

type EtappeCreated struct {
	Year                  int       `json:"year"`
	EtaopeId              int       `json:"etaopeId"`
	EtappeDate            time.Time `json:"etappeDate"`
	EtappeStartLocation   string    `json:"etappeStartLocation"`
	EtappeFinishtLocation string    `json:"etappeFinishtLocation"`
	EtappeLength          int       `json:"etappeLength"`
	EtappeKind            int       `json:"etappeKind"`
}

func (event *EtappeCreated) Wrap() *Envelope {
	envelope := new(Envelope)
	envelope.Type = TypeEtappeCreated
	envelope.EtappeCreated = event
	return envelope
}

type GamblerCreated struct {
	GamblerUid       string `json:"gamblerUid"`
	GamblerName      string `json:"gamblerName"`
	GamblerEmail     string `json:"gamblerEmail"`
	GamblerImageIUrl string `json:"gamblerImageIUrl"`
}

func (event *GamblerCreated) Wrap() *Envelope {
	envelope := new(Envelope)
	envelope.Type = TypeGamblerCreated
	envelope.GamblerCreated = event
	return envelope
}

type GamblerTeamCreated struct {
	GamblerUid      string `json:"gamblerUid"`
	Year            int    `json:"year"`
	GamblerCyclists []int  `json:"gamblerCyclists"`
}

func (event *GamblerTeamCreated) Wrap() *Envelope {
	envelope := new(Envelope)
	envelope.Type = TypeGamblerTeamCreated
	envelope.GamblerTeamCreated = event
	return envelope
}

type EtappeResultsAvailable struct {
	Year                     int   `json:"year"`
	LastEtappeId             int   `json:"lastEtappeId"`
	BestDayCyclistIds        []int `json:"bestDayCyclistIds"`
	BestAllrondersCyclistIds []int `json:"bestAllrondersCyclistIds"`
	BestSprintersCyclistIds  []int `json:"bestSprintersCyclistIds"`
	BestClimberCyclistIds    []int `json:"bestClimberCyclistIds"`
}

func (event *EtappeResultsAvailable) Wrap() *Envelope {
	envelope := new(Envelope)
	envelope.Type = TypeEtappeResultsAvailable
	envelope.EtappeResultsAvailable = event
	return envelope
}

type CyclistScoreCalculated struct {
	Year         int `json:"year"`
	CyclistId    int `json:"cyclistId"`
	LastEtappeId int `json:"lastEtappeId"`
	NewScore     int `json:"newScore"`
}

func (event *CyclistScoreCalculated) Wrap() *Envelope {
	envelope := new(Envelope)
	envelope.Type = TypeCyclistScoreCalculated
	envelope.CyclistScoreCalculated = event
	return envelope
}

type GamblerScoreCalculated struct {
	Year         int    `json:"year"`
	GamblerUid   string `json:"gamblerUid"`
	LastEtappeId int    `json:"lastEtappeId"`
	NewScore     int    `json:"newScore"`
}

func (event *GamblerScoreCalculated) Wrap() *Envelope {
	envelope := new(Envelope)
	envelope.Type = TypeGamblerScoreCalculated
	envelope.GamblerScoreCalculated = event
	return envelope
}

type Type int

const (
	TypeUnknown Type = iota
	TypeCyclistCreated
	TypeEtappeCreated
	TypeGamblerCreated
	TypeGamblerTeamCreated
	TypeEtappeResultsAvailable
	TypeCyclistScoreCalculated
	TypeGamblerScoreCalculated
	TypeTourCreated
)

func (t Type) String() string {
	switch t {
	case TypeGamblerCreated:
		return "GamblerCreated"
	case TypeGamblerTeamCreated:
		return "GamblerTeamCreated"
	case TypeEtappeResultsAvailable:
		return "EtappeResultsAvailable"
	case TypeCyclistScoreCalculated:
		return "CyclistScoreCalculated"
	case TypeGamblerScoreCalculated:
		return "GamblerScoreCalculated"
	case TypeTourCreated:
		return "TourCreated"
	case TypeCyclistCreated:
		return "CyclistCreated"
	case TypeEtappeCreated:
		return "EtappeCreated"

	}
	return "unknown"
}

type Envelope struct {
	Type                   Type                    `json:"type"`
	EtappeCreated          *EtappeCreated          `json:"etappeCreated"`
	GamblerCreated         *GamblerCreated         `json:"gamblerCreated"`
	GamblerTeamCreated     *GamblerTeamCreated     `json:"gamblerTeamCreated"`
	EtappeResultsAvailable *EtappeResultsAvailable `json:"etappeResultsAvailable"`
	CyclistScoreCalculated *CyclistScoreCalculated `json:"cyclistScoreCalculated"`
	GamblerScoreCalculated *GamblerScoreCalculated `json:"gamblerScoreCalculated"`
	TourCreated            *TourCreated            `json:"tourCreated"`
	CyclistCreated         *CyclistCreated         `json:"cyclistCreated"`
}

type EventHandlerFunc func(Envelope *Envelope) error

type PublishSubscriber interface {
	Subscribe(eventType Type, f EventHandlerFunc) error
	Publish(Envelope *Envelope) error
}

type StoredItemHandlerFunc func(envelope *Envelope) bool

type Store interface {
	Store(envelope *Envelope) error
	Iterate(StoredItemHandlerFunc) error
}
