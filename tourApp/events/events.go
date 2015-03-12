package events

// Generated automatically by microgen: do not edit manually

import (
	"code.google.com/p/go-uuid/uuid"
	"encoding/json"
	"github.com/MarcGrol/microgen/lib/envelope"
	"log"
	"strconv"
	"time"
)

type Type int

const (
	TypeUnknown              Type = iota
	TypeGamblerTeamCreated        = 5
	TypeNewsItemCreated           = 10
	TypeTourCreated               = 1
	TypeCyclistCreated            = 2
	TypeEtappeCreated             = 3
	TypeEtappeResultsCreated      = 7
	TypeGamblerCreated            = 4
)

func GetAllEventTypes() []Type {
	return []Type{
		TypeEtappeCreated,
		TypeEtappeResultsCreated,
		TypeGamblerCreated,
		TypeGamblerTeamCreated,
		TypeNewsItemCreated,
		TypeTourCreated,
		TypeCyclistCreated,
	}
}

func GetTourEventTypes() []Type {
	return []Type{
		TypeTourCreated,
		TypeCyclistCreated,
		TypeEtappeCreated,
		TypeEtappeResultsCreated,
	}
}

func GetGamblerEventTypes() []Type {
	return []Type{
		TypeGamblerCreated,
		TypeGamblerTeamCreated,
	}
}

func GetNewsEventTypes() []Type {
	return []Type{
		TypeNewsItemCreated,
	}
}

func GetNotificationEventTypes() []Type {
	return []Type{}
}

func (t Type) String() string {
	switch t {
	case TypeTourCreated:
		return "TourCreated"
	case TypeCyclistCreated:
		return "CyclistCreated"
	case TypeEtappeCreated:
		return "EtappeCreated"
	case TypeEtappeResultsCreated:
		return "EtappeResultsCreated"
	case TypeGamblerCreated:
		return "GamblerCreated"
	case TypeGamblerTeamCreated:
		return "GamblerTeamCreated"
	case TypeNewsItemCreated:
		return "NewsItemCreated"

	}
	return "unknown"
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

type EtappeResultsCreated struct {
	Year                     int   `json:"year"`
	LastEtappeId             int   `json:"lastEtappeId"`
	BestDayCyclistIds        []int `json:"bestDayCyclistIds"`
	BestAllrondersCyclistIds []int `json:"bestAllrondersCyclistIds"`
	BestSprintersCyclistIds  []int `json:"bestSprintersCyclistIds"`
	BestClimberCyclistIds    []int `json:"bestClimberCyclistIds"`
}

func (event *EtappeResultsCreated) Wrap() *envelope.Envelope {
	var err error
	var t Type = TypeEtappeResultsCreated

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

func IsEtappeResultsCreated(envelope *envelope.Envelope) bool {
	var t Type = TypeEtappeResultsCreated
	return envelope.EventTypeName == t.String()
}

func GetIfIsEtappeResultsCreated(envelop *envelope.Envelope) (*EtappeResultsCreated, bool) {
	if IsEtappeResultsCreated(envelop) == false {
		return nil, false
	}
	event := UnWrapEtappeResultsCreated(envelop)
	return event, true
}

func UnWrapEtappeResultsCreated(envelop *envelope.Envelope) *EtappeResultsCreated {
	if IsEtappeResultsCreated(envelop) == false {
		return nil
	}
	var event EtappeResultsCreated
	err := json.Unmarshal([]byte(envelop.EventData), &event)
	if err != nil {
		log.Printf("Error unmarshalling event payload %+v", err)
		return nil
	}

	return &event
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

type NewsItemCreated struct {
	Uid       string    `json:"uid"`
	Year      int       `json:"year"`
	Timestamp time.Time `json:"timestamp"`
	Message   string    `json:"message"`
	Sender    string    `json:"sender"`
}

func (event *NewsItemCreated) Wrap() *envelope.Envelope {
	var err error
	var t Type = TypeNewsItemCreated

	envelope := new(envelope.Envelope)
	envelope.Uuid = uuid.New()
	envelope.SequenceNumber = 0 // Set later by event-store
	envelope.Timestamp = time.Now()
	envelope.AggregateName = "news"
	envelope.AggregateUid = event.Uid

	envelope.EventTypeName = t.String()
	blob, err := json.Marshal(event)
	if err != nil {
		log.Printf("Error marshalling event payload %+v", err)
		return nil //, err
	}
	envelope.EventData = string(blob)
	return envelope //, nil
}

func IsNewsItemCreated(envelope *envelope.Envelope) bool {
	var t Type = TypeNewsItemCreated
	return envelope.EventTypeName == t.String()
}

func GetIfIsNewsItemCreated(envelop *envelope.Envelope) (*NewsItemCreated, bool) {
	if IsNewsItemCreated(envelop) == false {
		return nil, false
	}
	event := UnWrapNewsItemCreated(envelop)
	return event, true
}

func UnWrapNewsItemCreated(envelop *envelope.Envelope) *NewsItemCreated {
	if IsNewsItemCreated(envelop) == false {
		return nil
	}
	var event NewsItemCreated
	err := json.Unmarshal([]byte(envelop.EventData), &event)
	if err != nil {
		log.Printf("Error unmarshalling event payload %+v", err)
		return nil
	}

	return &event
}
