package tour

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

	tour := NewTour()
	// no event-handler
	eventHandler := NewTourEventHandler(bus, store, tour)
	eventHandler.Start()

	// command-handler: start web-server: blocking call
	commandHandler := NewTourCommandHandler(bus, store, tour)
	commandHandler.Start(listenPort)

	return nil
}

func (commandHandler *TourCommandHandler) Start(listenPort int) error {
	var err error
	engine := gin.Default()
	api := engine.Group("/api")
	{
		api.GET("/tour/:year", func(c *gin.Context) {
			year, err := strconv.Atoi(c.Params.ByName("year"))
			if err != nil {
				myhttp.HandleError(c, myerrors.NewInvalidInputError(err))
				return
			}
			tour, err := commandHandler.HandleGetTourQuery(year)
			if err != nil {
				myhttp.HandleError(c, err)
				return
			}
			c.JSON(200, *tour)
		})
		api.POST("/tour", func(c *gin.Context) {
			var command CreateTourCommand
			err = c.Bind(&command)
			if err != nil {
				myhttp.HandleError(c, myerrors.NewInvalidInputError(errors.New("Invalid tour-command")))
				return
			}
			err := commandHandler.HandleCreateTourCommand(&command)
			if err != nil {
				myhttp.HandleError(c, err)
				return
			}
			c.JSON(200, *myhttp.SuccessResponse())
		})
		api.POST("/tour/:year/etappe", func(c *gin.Context) {
			var command CreateEtappeCommand
			err = c.Bind(&command)
			if err != nil {
				myhttp.HandleError(c, myerrors.NewInvalidInputError(errors.New("Invalid etappe-command")))
				return
			}
			err = commandHandler.HandleCreateEtappeCommand(&command)
			if err != nil {
				myhttp.HandleError(c, err)
				return
			}
			c.JSON(200, *myhttp.SuccessResponse())
		})
		api.POST("/tour/:year/cyclist", func(c *gin.Context) {
			var command CreateCyclistCommand
			err = c.Bind(&command)
			if err != nil {
				myhttp.HandleError(c, myerrors.NewInvalidInputError(errors.New("Invalid cyclist-command")))
				return
			}
			err := commandHandler.HandleCreateCyclistCommand(&command)
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
	st := store.NewEventStore(dataDir, "tour.db")
	err = st.Open()
	if err != nil {
		return nil, err
	}
	return st, nil
}

func createBus(busAddress string) (infra.PublishSubscriber, error) {
	bus := bus.NewEventBus("tourApp", "tour", busAddress)
	if bus == nil {
		return nil, errors.New("Error starting bus")
	}
	return bus, nil
}
