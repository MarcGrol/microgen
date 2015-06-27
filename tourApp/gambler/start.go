package gambler

import (
	"errors"
	"fmt"
	"os"
	"strconv"

	"github.com/MarcGrol/microgen/infra"
	"github.com/MarcGrol/microgen/infra/bus"
	"github.com/MarcGrol/microgen/infra/myhttp"
	"github.com/MarcGrol/microgen/infra/store"
	"github.com/MarcGrol/microgen/lib/myerrors"
	"github.com/gin-gonic/gin"
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

func (commandHandler *GamblerCommandHandler) Start(listenPort int) error {
	var err error
	engine := gin.Default()
	api := engine.Group("/api")
	{
		api.GET("/results/:year", func(c *gin.Context) {
			year, err := strconv.Atoi(c.Params.ByName("year"))
			if err != nil {
				myhttp.HandleError(c, myerrors.NewInvalidInputError(err))
				return
			}
			results, err := commandHandler.HandleGetResultsQuery(year)
			if err != nil {
				myhttp.HandleError(c, err)
				return
			}
			c.JSON(200, *results)
		})
		api.GET("/gambler/:gamblerUid/year/:year", func(c *gin.Context) {
			gamblerUid := c.Params.ByName("gamblerUid")
			year, err := strconv.Atoi(c.Params.ByName("year"))
			if err != nil {
				myhttp.HandleError(c, myerrors.NewInvalidInputError(err))
				return
			}
			gambler, err := commandHandler.HandleGetGamblerQuery(gamblerUid, year)
			if err != nil {
				myhttp.HandleError(c, err)
				return
			}
			c.JSON(200, *gambler)
		})
		api.POST("/gambler", func(c *gin.Context) {
			var command CreateGamblerCommand
			err = c.Bind(&command)
			if err != nil {
				myhttp.HandleError(c, myerrors.NewInvalidInputError(errors.New("Invalid create-gambler-command")))
				return
			}
			err := commandHandler.HandleCreateGamblerCommand(&command)
			if err != nil {
				myhttp.HandleError(c, err)
				return
			}
			c.JSON(200, *myhttp.SuccessResponse())
		})
		api.POST("gambler/:gamblerUid/year/:year/team", func(c *gin.Context) {
			var command CreateGamblerTeamCommand
			err = c.Bind(&command)
			if err != nil {
				myhttp.HandleError(c, myerrors.NewInvalidInputError(errors.New("Invalid create-gambler-team-command")))
				return
			}
			err := commandHandler.HandleCreateGamblerTeamCommand(&command)
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
