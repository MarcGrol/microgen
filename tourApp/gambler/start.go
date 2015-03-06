package gambler

import (
	"errors"
	"fmt"
	"github.com/MarcGrol/microgen/infra"
	"github.com/MarcGrol/microgen/infra/bus"
	"github.com/MarcGrol/microgen/infra/http"
	"github.com/MarcGrol/microgen/infra/store"
	//"github.com/MarcGrol/microgen/lib/envelope"
	"github.com/MarcGrol/microgen/lib/myerrors"
	//"github.com/MarcGrol/microgen/tourApp/events"
	"github.com/gin-gonic/gin"
	"os"
	"strconv"
)

func Start(listenPort int, busAddress string, baseDir string) error {

	//start store
	store, err := createStore(baseDir)
	if err != nil {
		return err
	}

	// create and start event bus
	bus, err := createBus(busAddress)
	if err != nil {
		return err
	}

	// event-handler
	eventHandler := NewGamblerEventHandler(bus, store)
	err = eventHandler.Start()
	if err != nil {
		return err
	}

	// command-handler: start web-server: blocking call
	commandHandler := NewGamblerCommandHandler(bus, store)
	commandHandler.Start(listenPort)

	return nil
}

func (eventHandler *GamblerEventHandler) Start() error {
	// {
	// 	var t events.Type = events.TypeTourCreated
	// 	topic := t.String()
	// 	eventHandler.bus.Subscribe(topic, func(envelop *envelope.Envelope) error {
	// 		event := events.UnWrapTourCreated(envelop)
	// 		return eventHandler.OnTourCreated(event)
	// 	})
	// }
	// {
	// 	var t events.Type = events.TypeCyclistCreated
	// 	topic := t.String()
	// 	eventHandler.bus.Subscribe(topic, func(envelop *envelope.Envelope) error {
	// 		event := events.UnWrapCyclistCreated(envelop)
	// 		return eventHandler.OnCyclistCreated(event)
	// 	})
	// }
	return nil
}

func (commandHandler *GamblerCommandHandler) Start(listenPort int) {
	engine := gin.Default()
	api := engine.Group("/api")
	{
		api.GET("/gambler/:gamblerUid/year/:year", func(c *gin.Context) {
			gamblerUid := c.Params.ByName("gamblerUid")
			year, err := strconv.Atoi(c.Params.ByName("year"))
			if err != nil {
				http.HandleError(c, myerrors.NewInvalidInputError(err))
				return
			}
			gambler, err := commandHandler.HandleGetGamblerQuery(gamblerUid, year)
			if err != nil {
				http.HandleError(c, err)
				return
			}
			c.JSON(200, *gambler)
		})
		api.POST("/gambler", func(c *gin.Context) {
			var command CreateGamblerCommand
			ok := c.Bind(&command)
			if ok == false {
				http.HandleError(c, myerrors.NewInvalidInputError(errors.New("Invalid create-gambler-command")))
				return
			}
			err := commandHandler.HandleCreateGamblerCommand(&command)
			if err != nil {
				http.HandleError(c, err)
				return
			}
			c.JSON(200, *http.SuccessResponse())
		})
		api.POST("gambler/:gamblerUid/year/:year/team", func(c *gin.Context) {
			var command CreateGamblerTeamCommand
			ok := c.Bind(&command)
			if ok == false {
				http.HandleError(c, myerrors.NewInvalidInputError(errors.New("Invalid create-gambler-team-command")))
				return
			}
			err := commandHandler.HandleCreateGamblerTeamCommand(&command)
			if err != nil {
				http.HandleError(c, err)
				return
			}
			c.JSON(200, *http.SuccessResponse())
		})
	}

	engine.Run(fmt.Sprintf(":%d", listenPort))
}

func createStore(baseDir string) (infra.Store, error) {
	dataDir := baseDir + "/" + "data"

	// create dir if not exists
	err := os.MkdirAll(dataDir, 0777)
	if err != nil {
		return nil, err
	}
	st := store.NewEventStore(dataDir, "gambler.db")
	err = st.Open()
	if err != nil {
		return nil, err
	}
	return st, nil
}

func createBus(busAddress string) (infra.PublishSubscriber, error) {
	bus := bus.NewEventBus("tourApp", "gambler", busAddress)
	if bus == nil {
		return nil, errors.New("Error starting bus")
	}
	return bus, nil
}
