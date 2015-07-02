package news

import (
	"fmt"
	"log"
	"time"

	"github.com/MarcGrol/microgen/lib/envelope"
	"github.com/MarcGrol/microgen/tourApp/events"
)

type NewsContext struct {
	years map[int]*News
}

type News struct {
	NewsItems []*NewsItem
	cyclists  map[int]*cyclist
	etappes   map[int]*etappe
}

// +gen slice:"SortBy,Where,Select[string],Any,First"
type NewsItem struct {
	Timestamp time.Time `json:"timestamp"`
	Sender    string    `json:"sender"`
	Message   string    `json:"message"`
}

type cyclist struct {
	number int
	name   string
	team   string
}

type etappe struct {
	id             int
	date           time.Time
	startLocation  string
	finishLocation string
	length         int
	kind           int
}

func NewNewsContext() *NewsContext {
	return &NewsContext{
		years: make(map[int]*News),
	}
}

func getYearNews(newsContext *NewsContext, year int) *News {

	news, exists := newsContext.years[year]
	if exists == false {
		newsContext.years[year] = &News{
			NewsItems: make([]*NewsItem, 0, 10),
			cyclists:  make(map[int]*cyclist),
			etappes:   make(map[int]*etappe),
		}
		news = newsContext.years[year]
	}

	return news
}

func (newsContext *NewsContext) ApplyAll(envelopes []envelope.Envelope) {
	applyEvents(envelopes, newsContext)
}

func (newsContext *NewsContext) ApplyTourCreated(event *events.TourCreated) {
	yearData := getYearNews(newsContext, event.Year)

	log.Printf(">>> ApplyTourCreated before:%+v -> %+v (%d)",
		event, newsContext, len(yearData.NewsItems))

	yearData.NewsItems = append(yearData.NewsItems,
		&NewsItem{
			Timestamp: time.Now(),
			Sender:    "admin",
			Message:   fmt.Sprintf("The tour of %d is about to start", event.Year),
		})

	log.Printf(">>> ApplyTourCreated after:%+v -> %+v (%d)", event, newsContext,
		len(yearData.NewsItems))
}

func (newsContext *NewsContext) ApplyCyclistCreated(event *events.CyclistCreated) {
	yearData := getYearNews(newsContext, event.Year)

	log.Printf(">>> ApplyCyclistCreated before:%+v -> %+v (%d)", event, newsContext, len(yearData.NewsItems))

	yearData.cyclists[event.CyclistId] =
		&cyclist{
			number: event.CyclistId,
			name:   event.CyclistName,
			team:   event.CyclistTeam}

	log.Printf(">>> ApplyCyclistCreated after:%+v -> %+v (%d)", event, newsContext, len(yearData.NewsItems))
}

func (newsContext *NewsContext) ApplyEtappeCreated(event *events.EtappeCreated) {
	yearData := getYearNews(newsContext, event.Year)

	log.Printf(">>> ApplyEtappeCreated before:%+v -> %+v (%d)",
		event, newsContext, len(yearData.NewsItems))

	yearData.etappes[event.EtappeId] =
		&etappe{
			id:             event.EtappeId,
			date:           event.EtappeDate,
			startLocation:  event.EtappeStartLocation,
			finishLocation: event.EtappeFinishLocation,
			length:         event.EtappeLength,
			kind:           event.EtappeKind}

	log.Printf(">>> ApplyEtappeCreated after:%+v -> %+v (%d)",
		event, newsContext, len(yearData.NewsItems))
}

func composeMsgText(yearData *News, event *events.EtappeResultsCreated) (string, error) {
	etappe, exists := yearData.etappes[event.LastEtappeId]
	if exists == false {
		return "", fmt.Errorf("Etappe %d not found", event.LastEtappeId)
	}

	dayFirst, exists := yearData.cyclists[event.BestDayCyclistIds[0]]
	if exists == false {
		return "", fmt.Errorf("Cyclist %d not found", event.LastEtappeId)
	}

	daySecond, exists := yearData.cyclists[event.BestDayCyclistIds[1]]
	if exists == false {
		return "", fmt.Errorf("Cyclist %d not found", event.LastEtappeId)
	}

	dayThird, exists := yearData.cyclists[event.BestDayCyclistIds[2]]
	if exists == false {
		return "", fmt.Errorf("Cyclist %d not found", event.LastEtappeId)
	}

	return fmt.Sprintf(
		"Results for etappe %d from %s to %s:\n-1- %s\n-2- %s\n-3- %s\n",
		etappe.id,
		etappe.startLocation,
		etappe.finishLocation,
		dayFirst.name,
		daySecond.name,
		dayThird.name), nil
}

func (newsContext *NewsContext) ApplyEtappeResultsCreated(event *events.EtappeResultsCreated) {
	yearData := getYearNews(newsContext, event.Year)

	log.Printf(">>> ApplyEtappeResultsCreated before:%+v -> %+v (%d)",
		event, newsContext, len(yearData.NewsItems))

	msg, err := composeMsgText(yearData, event)
	if err != nil {
		return
	}

	yearData.NewsItems = append(yearData.NewsItems,
		&NewsItem{
			Sender:    "admin",
			Timestamp: yearData.etappes[event.LastEtappeId].date,
			Message:   msg})

	log.Printf(">>> ApplyEtappeResultsCreated after:%+v -> %+v (%d)", event, newsContext,
		len(yearData.NewsItems))
}

func (newsContext *NewsContext) ApplyNewsItemCreated(event *events.NewsItemCreated) {
	log.Printf(">>> ApplyNewsItemCreated before:%+v -> %+v", event, newsContext)
	yearData := getYearNews(newsContext, event.Year)
	yearData.NewsItems = append(yearData.NewsItems,
		&NewsItem{
			Sender:    event.Sender,
			Timestamp: event.Timestamp,
			Message:   event.Message})
	log.Printf(">>> ApplyNewsItemCreated after:%+v -> %+v", event, newsContext)
}
