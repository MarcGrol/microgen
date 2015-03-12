package news

import (
	"github.com/MarcGrol/microgen/infra"
	"github.com/MarcGrol/microgen/lib/envelope"
	"github.com/MarcGrol/microgen/lib/myerrors"
	"github.com/MarcGrol/microgen/tourApp/events"
	"log"
	"strconv"
	"time"
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
	// TODO
	return nil
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

func (ch *NewsCommandHandler) HandleGetNewsQuery(command *CreateNewsItemCommand) (*News, error) {
	err := ch.validateCreateNewsItemCommand(command)
	if err != nil {
		return nil, myerrors.NewInvalidInputError(err)
	}
	news, err := getNews(ch.store, command.Year)
	if err != nil {
		return nil, myerrors.NewInternalError(err)
	}

	return news, nil
}

func doStore(store infra.Store, envelopes []*envelope.Envelope) error {
	for _, env := range envelopes {
		err := store.Store(env)
		if err != nil {
			log.Printf("Error storing event: %+v", err)
			return err
		}
		//log.Printf("Successfully stored event: %+v", env)
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

type News struct {
	Year     int
	NewItems []NewsItem
}

type NewsItem struct {
	Message   string
	Sender    string
	Timestamp time.Time
}

func getNews(store infra.Store, year int) (*News, error) {
	newsRelatedEvents, err := store.Get("new", strconv.Itoa(year))
	if err != nil || len(newsRelatedEvents) == 0 {
		return nil, err
	}

	news := new(News)
	applyEvents(newsRelatedEvents, news)
	return news, nil
}

func (news *News) ApplyTourCreated(event *events.TourCreated) {
	log.Fatal("news.%s not implemented", "ApplyTourCreated")
	// TODO create news item out of this
}

func (news *News) ApplyCyclistCreated(event *events.CyclistCreated) {
	log.Fatal("news.%s not implemented", "ApplyCyclistCreated")
	// TODO create news item out of this
}

func (news *News) ApplyEtappeCreated(event *events.EtappeCreated) {
	log.Fatal("news.%s not implemented", "ApplyEtappeCreated")
	// TODO create news item out of this
}

func (news *News) ApplyEtappeResultsCreated(event *events.EtappeResultsCreated) {
	log.Fatal("news.%s not implemented", "ApplyEtappeResultsCreated")
	// TODO create news item out of this
}

func (news *News) ApplyNewsItemCreated(event *events.NewsItemCreated) {
	log.Fatal("news.%s not implemented", "ApplyNewsItemCreated")
	// TODO create news item out of this
}
