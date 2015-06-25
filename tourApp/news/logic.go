package news

//go:generate gen

import (
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/MarcGrol/microgen/infra"
	"github.com/MarcGrol/microgen/lib/envelope"
	"github.com/MarcGrol/microgen/lib/myerrors"
	"github.com/MarcGrol/microgen/lib/validation"
	"github.com/MarcGrol/microgen/tourApp/events"
)

type NewsEventHandler struct {
	bus   infra.PublishSubscriber
	store infra.Store
}

func NewNewsEventHandler(bus infra.PublishSubscriber, store infra.Store) *NewsEventHandler {
	handler := new(NewsEventHandler)
	handler.bus = bus
	handler.store = store
	return handler
}

func (eventHandler *NewsEventHandler) Start() error {
	for _, eventType := range events.GetTourEventTypes() {
		err := eventHandler.bus.Subscribe(eventType.String(), func(envelope *envelope.Envelope) error {
			return eventHandler.OnEvent(envelope)
		})
		if err != nil {
			return err
		}
	}

	return nil
}

func (eh *NewsEventHandler) OnEvent(envelop *envelope.Envelope) error {
	log.Printf(">>> Got event %+v", envelop)
	return doStore(eh.store, []*envelope.Envelope{envelop})
}

type NewsCommandHandler struct {
	bus   infra.PublishSubscriber
	store infra.Store
}

func NewNewsCommandHandler(bus infra.PublishSubscriber, store infra.Store) *NewsCommandHandler {
	handler := new(NewsCommandHandler)
	handler.bus = bus
	handler.store = store
	return handler
}

func (ch *NewsCommandHandler) validateCreateNewsItemCommand(command *CreateNewsItemCommand) error {
	v := validation.Validator{}
	v.GreaterThan("Year", 2014, command.Year)
	v.After("Timestamp", "2015-07-01T00:00:00Z", command.Timestamp)
	v.NotEmpty("Message", command.Message)
	v.NotEmpty("Sender", command.Sender)
	return v.Err
}

func (ch *NewsCommandHandler) HandleCreateNewsItemCommand(command *CreateNewsItemCommand) error {
	err := ch.validateCreateNewsItemCommand(command)
	if err != nil {
		return myerrors.NewInvalidInputError(err)
	}

	// apply business logic
	newsItemEvent := events.NewsItemCreated{
		Year:      command.Year,
		Timestamp: command.Timestamp,
		Message:   command.Message,
		Sender:    command.Sender}

	// store and emit resulting event
	return doStoreAndPublish(ch.store, ch.bus, []*envelope.Envelope{newsItemEvent.Wrap()})
}

func (ch *NewsCommandHandler) HandleGetNewsQuery(year int) (*News, error) {
	news, err := getNewsForYear(ch.store, year)
	if err != nil {
		return nil, myerrors.NewInternalError(err)
	}

	return news.years[year], nil
}

func doStore(store infra.Store, envelopes []*envelope.Envelope) error {
	for _, env := range envelopes {
		err := store.Store(env)
		if err != nil {
			log.Printf("Error storing event: %+v", err)
			return err
		}
		log.Printf("Successfully stored event: %+v", env)
	}
	return nil
}

func doStoreAndPublish(store infra.Store, bus infra.PublishSubscriber, envelopes []*envelope.Envelope) error {
	err := doStore(store, envelopes)
	if err != nil {
		return myerrors.NewInternalError(err)
	}
	for _, envelop := range envelopes {
		err = bus.Publish(envelop)
		if err != nil {
			return myerrors.NewInternalError(err)
		}
	}
	return nil
}

type NewsContext struct {
	years map[int]*News
}

type News struct {
	newsItems []*NewsItem
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

func getNewsForYear(store infra.Store, year int) (*NewsContext, error) {
	tourRelatedEvents, err := store.Get("tour", strconv.Itoa(year))
	if err != nil {
		return nil, err
	}

	newsRelatedEvents, err := store.Get("news", strconv.Itoa(year))
	if err != nil {
		return nil, err
	}

	newsContext := &NewsContext{
		years: make(map[int]*News),
	}
	applyEvents(append(tourRelatedEvents, newsRelatedEvents...), newsContext)
	return newsContext, nil
}

func getYearNews(newsContext *NewsContext, year int) *News {

	news, exists := newsContext.years[year]
	if exists == false {
		newsContext.years[year] = &News{
			newsItems: make([]*NewsItem, 0, 10),
			cyclists:  make(map[int]*cyclist),
			etappes:   make(map[int]*etappe),
		}
		news = newsContext.years[year]
	}

	return news
}

func (newsContext *NewsContext) ApplyTourCreated(event *events.TourCreated) {
	log.Printf(">>> ApplyTourCreated before:%+v -> %+v", event, newsContext)
	yearData := getYearNews(newsContext, event.Year)
	yearData.newsItems = append(yearData.newsItems,
		&NewsItem{
			Timestamp: time.Now(),
			Sender:    "admin",
			Message:   fmt.Sprintf("The tour of %d is about to start", event.Year),
		})

	log.Printf(">>> ApplyTourCreated after:%+v -> %+v", event, newsContext)
}

func (newsContext *NewsContext) ApplyCyclistCreated(event *events.CyclistCreated) {
	log.Printf(">>> ApplyCyclistCreated before:%+v -> %+v", event, newsContext)
	yearData := getYearNews(newsContext, event.Year)
	yearData.cyclists[event.CyclistId] =
		&cyclist{
			number: event.CyclistId,
			name:   event.CyclistName,
			team:   event.CyclistTeam}

	log.Printf(">>> ApplyCyclistCreated after:%+v -> %+v", event, newsContext)
}

func (newsContext *NewsContext) ApplyEtappeCreated(event *events.EtappeCreated) {
	log.Printf(">>> ApplyEtappeCreated before:%+v -> %+v", event, newsContext)
	yearData := getYearNews(newsContext, event.Year)
	yearData.etappes[event.EtappeId] =
		&etappe{
			id:             event.EtappeId,
			date:           event.EtappeDate,
			startLocation:  event.EtappeStartLocation,
			finishLocation: event.EtappeFinishLocation,
			length:         event.EtappeLength,
			kind:           event.EtappeKind}

	log.Printf(">>> ApplyEtappeCreated after:%+v -> %+v", event, newsContext)
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
		"Etappe %d from %s to %s has finished.\nEtappe result:\n-1- %s\n-2- %s\n-3- %s\n",
		etappe.id,
		etappe.startLocation,
		etappe.finishLocation,
		dayFirst.name,
		daySecond.name,
		dayThird.name), nil
}

func (newsContext *NewsContext) ApplyEtappeResultsCreated(event *events.EtappeResultsCreated) {
	log.Printf(">>> ApplyEtappeResultsCreated before:%+v -> %+v", event, newsContext)
	yearData := getYearNews(newsContext, event.Year)
	msg, err := composeMsgText(yearData, event)
	if err != nil {
		return
	}

	yearData.newsItems = append(yearData.newsItems,
		&NewsItem{
			Sender:    "admin",
			Timestamp: yearData.etappes[event.LastEtappeId].date,
			Message:   msg})

	log.Printf(">>> ApplyEtappeResultsCreated after:%+v -> %+v", event, newsContext)
}

func (newsContext *NewsContext) ApplyNewsItemCreated(event *events.NewsItemCreated) {
	log.Printf(">>> ApplyNewsItemCreated before:%+v -> %+v", event, newsContext)
	yearData := getYearNews(newsContext, event.Year)
	yearData.newsItems = append(yearData.newsItems,
		&NewsItem{
			Sender:    event.Sender,
			Timestamp: event.Timestamp,
			Message:   event.Message})
	log.Printf(">>> ApplyNewsItemCreated after:%+v -> %+v", event, newsContext)
}
