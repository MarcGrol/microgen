package events

// Generated automatically by microgen: do not edit manually

import (
	"code.google.com/p/go-uuid/uuid"
	"encoding/json"
	"github.com/MarcGrol/microgen/envelope"
	"log"
	"strconv"
	"time"
)

type Type int

const (
	TypeUnknown                Type = iota
	TypeGamblerCreated              = 4
	TypeGamblerTeamCreated          = 5
	TypeEtappeResultsAvailable      = 7
	TypeCyclistScoreCalculated      = 8
	TypeGamblerScoreCalculated      = 9
	TypeTourCreated                 = 1
	TypeCyclistCreated              = 2
	TypeEtappeCreated               = 3
)

func GetAllEventsTypes() []Type {
	return []Type{
		TypeGamblerTeamCreated,
		TypeEtappeResultsAvailable,
		TypeCyclistScoreCalculated,
		TypeGamblerScoreCalculated,
		TypeTourCreated,
		TypeCyclistCreated,
		TypeEtappeCreated,
		TypeGamblerCreated,
	}
}

func (t Type) String() string {
	switch t {
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
	case TypeCyclistScoreCalculated:
		return "CyclistScoreCalculated"
	case TypeGamblerScoreCalculated:
		return "GamblerScoreCalculated"
	case TypeTourCreated:
		return "TourCreated"

	}
	return "unknown"
}

type GamblerCreated struct {
	GamblerUid       string `json:"gamblerUid"`
	GamblerName      string `json:"gamblerName"`
	GamblerEmail     string `json:"gamblerEmail"`
	GamblerImageIUrl string `json:"gamblerImageIUrl"`
}

func (event *GamblerCreated) Wrap() *envelope.Envelope {
	var err error
	var t Type = TypeGamblerCreated

	envelope := new(envelope.Envelope)
	envelope.Uuid = uuid.New()
	envelope.SequenceNumber = 0 // Set later by event-store
	envelope.Timestamp = time.Now()
	envelope.AggregateName = "gambler"
	envelope.AggregateUid = event.GamblerUid

	envelope.EventTypeName = t.String()
	blob, err := json.Marshal(event)
	if err != nil {
		log.Printf("Error marshalling event payload %+v", err)
		return nil //, err
	}
	envelope.EventData = string(blob)
	return envelope //, nil
}

func IsGamblerCreated(envelope *envelope.Envelope) bool {
	var t Type = TypeGamblerCreated
	return envelope.EventTypeName == t.String()
}

func GetIfIsGamblerCreated(envelop *envelope.Envelope) (*GamblerCreated, bool) {
	if IsGamblerCreated(envelop) == false {
		return nil, false
	}
	event := UnWrapGamblerCreated(envelop)
	return event, true
}

func UnWrapGamblerCreated(envelop *envelope.Envelope) *GamblerCreated {
	if IsGamblerCreated(envelop) == false {
		return nil
	}
	var event GamblerCreated
	err := json.Unmarshal([]byte(envelop.EventData), &event)
	if err != nil {
		log.Printf("Error unmarshalling event payload %+v", err)
		return nil
	}

	return &event
}

type GamblerTeamCreated struct {
	GamblerUid      string `json:"gamblerUid"`
	Year            int    `json:"year"`
	GamblerCyclists []int  `json:"gamblerCyclists"`
}

func (event *GamblerTeamCreated) Wrap() *envelope.Envelope {
	var err error
	var t Type = TypeGamblerTeamCreated

	envelope := new(envelope.Envelope)
	envelope.Uuid = uuid.New()
	envelope.SequenceNumber = 0 // Set later by event-store
	envelope.Timestamp = time.Now()
	envelope.AggregateName = "gambler"
	envelope.AggregateUid = event.GamblerUid

	envelope.EventTypeName = t.String()
	blob, err := json.Marshal(event)
	if err != nil {
		log.Printf("Error marshalling event payload %+v", err)
		return nil //, err
	}
	envelope.EventData = string(blob)
	return envelope //, nil
}

func IsGamblerTeamCreated(envelope *envelope.Envelope) bool {
	var t Type = TypeGamblerTeamCreated
	return envelope.EventTypeName == t.String()
}

func GetIfIsGamblerTeamCreated(envelop *envelope.Envelope) (*GamblerTeamCreated, bool) {
	if IsGamblerTeamCreated(envelop) == false {
		return nil, false
	}
	event := UnWrapGamblerTeamCreated(envelop)
	return event, true
}

func UnWrapGamblerTeamCreated(envelop *envelope.Envelope) *GamblerTeamCreated {
	if IsGamblerTeamCreated(envelop) == false {
		return nil
	}
	var event GamblerTeamCreated
	err := json.Unmarshal([]byte(envelop.EventData), &event)
	if err != nil {
		log.Printf("Error unmarshalling event payload %+v", err)
		return nil
	}

	return &event
}

type EtappeResultsAvailable struct {
	Year                     int   `json:"year"`
	LastEtappeId             int   `json:"lastEtappeId"`
	BestDayCyclistIds        []int `json:"bestDayCyclistIds"`
	BestAllrondersCyclistIds []int `json:"bestAllrondersCyclistIds"`
	BestSprintersCyclistIds  []int `json:"bestSprintersCyclistIds"`
	BestClimberCyclistIds    []int `json:"bestClimberCyclistIds"`
}

func (event *EtappeResultsAvailable) Wrap() *envelope.Envelope {
	var err error
	var t Type = TypeEtappeResultsAvailable

	envelope := new(envelope.Envelope)
	envelope.Uuid = uuid.New()
	envelope.SequenceNumber = 0 // Set later by event-store
	envelope.Timestamp = time.Now()
	envelope.AggregateName = "tour"
	envelope.AggregateUid = strconv.Itoa(event.Year)
	envelope.EventTypeName = t.String()
	blob, err := json.Marshal(event)
	if err != nil {
		log.Printf("Error marshalling event payload %+v", err)
		return nil //, err
	}
	envelope.EventData = string(blob)
	return envelope //, nil
}

func IsEtappeResultsAvailable(envelope *envelope.Envelope) bool {
	var t Type = TypeEtappeResultsAvailable
	return envelope.EventTypeName == t.String()
}

func GetIfIsEtappeResultsAvailable(envelop *envelope.Envelope) (*EtappeResultsAvailable, bool) {
	if IsEtappeResultsAvailable(envelop) == false {
		return nil, false
	}
	event := UnWrapEtappeResultsAvailable(envelop)
	return event, true
}

func UnWrapEtappeResultsAvailable(envelop *envelope.Envelope) *EtappeResultsAvailable {
	if IsEtappeResultsAvailable(envelop) == false {
		return nil
	}
	var event EtappeResultsAvailable
	err := json.Unmarshal([]byte(envelop.EventData), &event)
	if err != nil {
		log.Printf("Error unmarshalling event payload %+v", err)
		return nil
	}

	return &event
}

type CyclistScoreCalculated struct {
	Year         int `json:"year"`
	CyclistId    int `json:"cyclistId"`
	LastEtappeId int `json:"lastEtappeId"`
	NewScore     int `json:"newScore"`
}

func (event *CyclistScoreCalculated) Wrap() *envelope.Envelope {
	var err error
	var t Type = TypeCyclistScoreCalculated

	envelope := new(envelope.Envelope)
	envelope.Uuid = uuid.New()
	envelope.SequenceNumber = 0 // Set later by event-store
	envelope.Timestamp = time.Now()
	envelope.AggregateName = "tour"
	envelope.AggregateUid = strconv.Itoa(event.Year)
	envelope.EventTypeName = t.String()
	blob, err := json.Marshal(event)
	if err != nil {
		log.Printf("Error marshalling event payload %+v", err)
		return nil //, err
	}
	envelope.EventData = string(blob)
	return envelope //, nil
}

func IsCyclistScoreCalculated(envelope *envelope.Envelope) bool {
	var t Type = TypeCyclistScoreCalculated
	return envelope.EventTypeName == t.String()
}

func GetIfIsCyclistScoreCalculated(envelop *envelope.Envelope) (*CyclistScoreCalculated, bool) {
	if IsCyclistScoreCalculated(envelop) == false {
		return nil, false
	}
	event := UnWrapCyclistScoreCalculated(envelop)
	return event, true
}

func UnWrapCyclistScoreCalculated(envelop *envelope.Envelope) *CyclistScoreCalculated {
	if IsCyclistScoreCalculated(envelop) == false {
		return nil
	}
	var event CyclistScoreCalculated
	err := json.Unmarshal([]byte(envelop.EventData), &event)
	if err != nil {
		log.Printf("Error unmarshalling event payload %+v", err)
		return nil
	}

	return &event
}

type GamblerScoreCalculated struct {
	Year         int    `json:"year"`
	GamblerUid   string `json:"gamblerUid"`
	LastEtappeId int    `json:"lastEtappeId"`
	NewScore     int    `json:"newScore"`
}

func (event *GamblerScoreCalculated) Wrap() *envelope.Envelope {
	var err error
	var t Type = TypeGamblerScoreCalculated

	envelope := new(envelope.Envelope)
	envelope.Uuid = uuid.New()
	envelope.SequenceNumber = 0 // Set later by event-store
	envelope.Timestamp = time.Now()
	envelope.AggregateName = "gambler"
	envelope.AggregateUid = event.GamblerUid

	envelope.EventTypeName = t.String()
	blob, err := json.Marshal(event)
	if err != nil {
		log.Printf("Error marshalling event payload %+v", err)
		return nil //, err
	}
	envelope.EventData = string(blob)
	return envelope //, nil
}

func IsGamblerScoreCalculated(envelope *envelope.Envelope) bool {
	var t Type = TypeGamblerScoreCalculated
	return envelope.EventTypeName == t.String()
}

func GetIfIsGamblerScoreCalculated(envelop *envelope.Envelope) (*GamblerScoreCalculated, bool) {
	if IsGamblerScoreCalculated(envelop) == false {
		return nil, false
	}
	event := UnWrapGamblerScoreCalculated(envelop)
	return event, true
}

func UnWrapGamblerScoreCalculated(envelop *envelope.Envelope) *GamblerScoreCalculated {
	if IsGamblerScoreCalculated(envelop) == false {
		return nil
	}
	var event GamblerScoreCalculated
	err := json.Unmarshal([]byte(envelop.EventData), &event)
	if err != nil {
		log.Printf("Error unmarshalling event payload %+v", err)
		return nil
	}

	return &event
}

type TourCreated struct {
	Year int `json:"year"`
}

func (event *TourCreated) Wrap() *envelope.Envelope {
	var err error
	var t Type = TypeTourCreated

	envelope := new(envelope.Envelope)
	envelope.Uuid = uuid.New()
	envelope.SequenceNumber = 0 // Set later by event-store
	envelope.Timestamp = time.Now()
	envelope.AggregateName = "tour"
	envelope.AggregateUid = strconv.Itoa(event.Year)
	envelope.EventTypeName = t.String()
	blob, err := json.Marshal(event)
	if err != nil {
		log.Printf("Error marshalling event payload %+v", err)
		return nil //, err
	}
	envelope.EventData = string(blob)
	return envelope //, nil
}

func IsTourCreated(envelope *envelope.Envelope) bool {
	var t Type = TypeTourCreated
	return envelope.EventTypeName == t.String()
}

func GetIfIsTourCreated(envelop *envelope.Envelope) (*TourCreated, bool) {
	if IsTourCreated(envelop) == false {
		return nil, false
	}
	event := UnWrapTourCreated(envelop)
	return event, true
}

func UnWrapTourCreated(envelop *envelope.Envelope) *TourCreated {
	if IsTourCreated(envelop) == false {
		return nil
	}
	var event TourCreated
	err := json.Unmarshal([]byte(envelop.EventData), &event)
	if err != nil {
		log.Printf("Error unmarshalling event payload %+v", err)
		return nil
	}

	return &event
}

type CyclistCreated struct {
	Year        int    `json:"year"`
	CyclistId   int    `json:"cyclistId"`
	CyclistName string `json:"cyclistName"`
	CyclistTeam string `json:"cyclistTeam"`
}

func (event *CyclistCreated) Wrap() *envelope.Envelope {
	var err error
	var t Type = TypeCyclistCreated

	envelope := new(envelope.Envelope)
	envelope.Uuid = uuid.New()
	envelope.SequenceNumber = 0 // Set later by event-store
	envelope.Timestamp = time.Now()
	envelope.AggregateName = "tour"
	envelope.AggregateUid = strconv.Itoa(event.Year)
	envelope.EventTypeName = t.String()
	blob, err := json.Marshal(event)
	if err != nil {
		log.Printf("Error marshalling event payload %+v", err)
		return nil //, err
	}
	envelope.EventData = string(blob)
	return envelope //, nil
}

func IsCyclistCreated(envelope *envelope.Envelope) bool {
	var t Type = TypeCyclistCreated
	return envelope.EventTypeName == t.String()
}

func GetIfIsCyclistCreated(envelop *envelope.Envelope) (*CyclistCreated, bool) {
	if IsCyclistCreated(envelop) == false {
		return nil, false
	}
	event := UnWrapCyclistCreated(envelop)
	return event, true
}

func UnWrapCyclistCreated(envelop *envelope.Envelope) *CyclistCreated {
	if IsCyclistCreated(envelop) == false {
		return nil
	}
	var event CyclistCreated
	err := json.Unmarshal([]byte(envelop.EventData), &event)
	if err != nil {
		log.Printf("Error unmarshalling event payload %+v", err)
		return nil
	}

	return &event
}

type EtappeCreated struct {
	Year                 int       `json:"year"`
	EtappeId             int       `json:"etappeId"`
	EtappeDate           time.Time `json:"etappeDate"`
	EtappeStartLocation  string    `json:"etappeStartLocation"`
	EtappeFinishLocation string    `json:"etappeFinishLocation"`
	EtappeLength         int       `json:"etappeLength"`
	EtappeKind           int       `json:"etappeKind"`
}

func (event *EtappeCreated) Wrap() *envelope.Envelope {
	var err error
	var t Type = TypeEtappeCreated

	envelope := new(envelope.Envelope)
	envelope.Uuid = uuid.New()
	envelope.SequenceNumber = 0 // Set later by event-store
	envelope.Timestamp = time.Now()
	envelope.AggregateName = "tour"
	envelope.AggregateUid = strconv.Itoa(event.Year)
	envelope.EventTypeName = t.String()
	blob, err := json.Marshal(event)
	if err != nil {
		log.Printf("Error marshalling event payload %+v", err)
		return nil //, err
	}
	envelope.EventData = string(blob)
	return envelope //, nil
}

func IsEtappeCreated(envelope *envelope.Envelope) bool {
	var t Type = TypeEtappeCreated
	return envelope.EventTypeName == t.String()
}

func GetIfIsEtappeCreated(envelop *envelope.Envelope) (*EtappeCreated, bool) {
	if IsEtappeCreated(envelop) == false {
		return nil, false
	}
	event := UnWrapEtappeCreated(envelop)
	return event, true
}

func UnWrapEtappeCreated(envelop *envelope.Envelope) *EtappeCreated {
	if IsEtappeCreated(envelop) == false {
		return nil
	}
	var event EtappeCreated
	err := json.Unmarshal([]byte(envelop.EventData), &event)
	if err != nil {
		log.Printf("Error unmarshalling event payload %+v", err)
		return nil
	}

	return &event
}
