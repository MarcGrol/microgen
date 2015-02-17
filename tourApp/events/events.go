package events

// Generated automatically by microgen: do not edit manually

import (
	"strconv"
	"time"
)

type TourCreated struct {
	Year int `json:"year"`
}

func (event *TourCreated) Wrap() *Envelope {
	envelope := new(Envelope)
	envelope.Type = TypeTourCreated
	envelope.TourCreated = event
	envelope.AggregateName = "tour"
	envelope.AggregateUid = strconv.Itoa(event.Year)
	envelope.SequenceNumber = 0 // TODO
	envelope.Timestamp = time.Now()
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
	envelope.AggregateName = "tour"
	envelope.AggregateUid = strconv.Itoa(event.Year)
	envelope.SequenceNumber = 0 // TODO
	envelope.Timestamp = time.Now()
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
	envelope.AggregateName = "tour"
	envelope.AggregateUid = strconv.Itoa(event.Year)
	envelope.SequenceNumber = 0 // TODO
	envelope.Timestamp = time.Now()
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
	envelope.AggregateName = "gambler"
	envelope.AggregateUid = event.GamblerUid
	envelope.SequenceNumber = 0 // TODO
	envelope.Timestamp = time.Now()
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
	envelope.AggregateName = "gambler"
	envelope.AggregateUid = event.GamblerUid
	envelope.SequenceNumber = 0 // TODO
	envelope.Timestamp = time.Now()
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
	envelope.AggregateName = "tour"
	envelope.AggregateUid = strconv.Itoa(event.Year)
	envelope.SequenceNumber = 0 // TODO
	envelope.Timestamp = time.Now()
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
	envelope.AggregateName = "tour"
	envelope.AggregateUid = strconv.Itoa(event.Year)
	envelope.SequenceNumber = 0 // TODO
	envelope.Timestamp = time.Now()
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
	envelope.AggregateName = "gambler"
	envelope.AggregateUid = event.GamblerUid
	envelope.SequenceNumber = 0 // TODO
	envelope.Timestamp = time.Now()
	return envelope
}

type Type int

const (
	TypeUnknown Type = iota
	TypeCyclistScoreCalculated
	TypeGamblerScoreCalculated
	TypeTourCreated
	TypeCyclistCreated
	TypeEtappeCreated
	TypeGamblerCreated
	TypeGamblerTeamCreated
	TypeEtappeResultsAvailable
)

func (t Type) String() string {
	switch t {
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
	case TypeGamblerCreated:
		return "GamblerCreated"
	case TypeGamblerTeamCreated:
		return "GamblerTeamCreated"
	case TypeEtappeResultsAvailable:
		return "EtappeResultsAvailable"

	}
	return "unknown"
}

type Envelope struct {
	SequenceNumber         uint64                  `json:"sequenceNumber"`
	AggregateName          string                  `json:"aggregateName"`
	AggregateUid           string                  `json:"aggregateUid"`
	Timestamp              time.Time               `json:"timestamp"`
	Type                   Type                    `json:"type"`
	TourCreated            *TourCreated            `json:"tourCreated"`
	CyclistCreated         *CyclistCreated         `json:"cyclistCreated"`
	EtappeCreated          *EtappeCreated          `json:"etappeCreated"`
	GamblerCreated         *GamblerCreated         `json:"gamblerCreated"`
	GamblerTeamCreated     *GamblerTeamCreated     `json:"gamblerTeamCreated"`
	EtappeResultsAvailable *EtappeResultsAvailable `json:"etappeResultsAvailable"`
	CyclistScoreCalculated *CyclistScoreCalculated `json:"cyclistScoreCalculated"`
	GamblerScoreCalculated *GamblerScoreCalculated `json:"gamblerScoreCalculated"`
}

type EventHandlerFunc func(Envelope *Envelope) error

type PublishSubscriber interface {
	Subscribe(eventType Type, f EventHandlerFunc) error
	Publish(Envelope *Envelope) error
}

type StoredItemHandlerFunc func(envelope *Envelope)

type Store interface {
	Store(envelope *Envelope) error
	Iterate(StoredItemHandlerFunc) error
}
