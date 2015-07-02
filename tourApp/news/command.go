package news

import (
	"errors"
	"fmt"
	"log"
	"strconv"

	"github.com/MarcGrol/microgen/infra"
	"github.com/MarcGrol/microgen/infra/myhttp"
	"github.com/MarcGrol/microgen/lib/envelope"
	"github.com/MarcGrol/microgen/lib/myerrors"
	"github.com/MarcGrol/microgen/lib/validation"
	"github.com/MarcGrol/microgen/tourApp/events"
	"github.com/gin-gonic/gin"
)

type NewsCommandHandler struct {
	bus         infra.PublishSubscriber
	store       infra.Store
	newsContext *NewsContext
}

func NewNewsCommandHandler(bus infra.PublishSubscriber, store infra.Store, newsContext *NewsContext) *NewsCommandHandler {
	handler := new(NewsCommandHandler)
	handler.bus = bus
	handler.store = store
	handler.newsContext = newsContext
	return handler
}

func (commandHandler *NewsCommandHandler) Start(listenPort int) error {
	var err error
	engine := gin.Default()
	api := engine.Group("/api")
	{
		api.GET("/news/:year/news", func(c *gin.Context) {
			year, err := strconv.Atoi(c.Params.ByName("year"))
			if err != nil {
				myhttp.HandleError(c, myerrors.NewInvalidInputError(err))
				return
			}
			tour, err := commandHandler.HandleGetNewsQuery(year)
			if err != nil {
				myhttp.HandleError(c, err)
				return
			}
			c.JSON(200, *tour)
		})
		api.POST("/news/:year/news", func(c *gin.Context) {
			var command CreateNewsItemCommand
			err = c.Bind(&command)
			if err != nil {
				myhttp.HandleError(c, myerrors.NewInvalidInputError(errors.New("Invalid tour-command")))
				return
			}
			err = commandHandler.HandleCreateNewsItemCommand(&command)
			if err != nil {
				myhttp.HandleError(c, err)
				return
			}
			c.JSON(200, *myhttp.SuccessResponse())
		})
	}

	engine.Run(fmt.Sprintf(":%d", listenPort))

	return nil
}

func (ch *NewsCommandHandler) validateCreateNewsItemCommand(command *CreateNewsItemCommand) error {
	v := validation.Validator{}
	v.GreaterThan("Year", 2010, command.Year)
	v.After("Timestamp", "2010-07-01T00:00:00Z", command.Timestamp)
	v.NotEmpty("Message", command.Message)
	v.NotEmpty("Sender", command.Sender)
	return v.Err
}

func (ch *NewsCommandHandler) HandleCreateNewsItemCommand(command *CreateNewsItemCommand) error {
	log.Printf("******* HandleCreateNewsItemCommand for year:%d", command.Year)
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
	err = doStoreAndPublish(ch.store, ch.bus, []*envelope.Envelope{newsItemEvent.Wrap()})
	if err != nil {
		return err
	}

	// apply event in memory
	ch.newsContext.ApplyNewsItemCreated(&newsItemEvent)

	log.Printf("HandleCreateNewsItemCommand completed:%+v -> %+v", command, newsItemEvent)

	return nil
}

func (ch *NewsCommandHandler) HandleGetNewsQuery(year int) (*News, error) {
	log.Printf("******* HandleGetNewsQuery for year:%d", year)
	yearNews, exists := ch.newsContext.years[year]
	if exists == false {
		return nil, myerrors.NewNotFoundErrorf("No news for year %d found", year)
	}
	log.Printf("******* HandleGetNewsQuery for year:%+v", yearNews.NewsItems)

	return yearNews, nil
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
