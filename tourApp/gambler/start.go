package gambler

import (
	"errors"
	"fmt"
	"github.com/MarcGrol/microgen/myerrors"
	"github.com/MarcGrol/microgen/tourApp/events"
	"github.com/MarcGrol/microgen/tourApp/http"
	"github.com/MarcGrol/microgen/tourApp/infra"
	"github.com/gin-gonic/gin"
	"os"
	"strconv"
)

func Start(listenPort int, busAddress string, baseDir string) error {
	store, err := startStore(baseDir)
	if err != nil {
		return err
	}
	bus := startBus(busAddress, NewGamblerEventHandler(store))
	if bus == nil {
		return errors.New("Error starting bus")
	}
	startHttp(listenPort, NewGamblerCommandHandler(bus, store))
	return nil
}

func startStore(baseDir string) (*infra.EventStore, error) {
	dataDir := baseDir + "/" + "data"

	// create dir if not exists
	err := os.MkdirAll(dataDir, 0777)
	if err != nil {
		return nil, err
	}
	store := infra.NewEventStore(dataDir, "gambler.db")
	err = store.Open()
	if err != nil {
		return nil, err
	}
	return store, nil
}

func startBus(busAddress string, eventHandler EventHandler) *infra.EventBus {
	bus := infra.NewEventBus("tourApp", "gambler", busAddress)

	bus.Subscribe(events.TypeTourCreated, func(envelope *events.Envelope) error {
		return eventHandler.OnTourCreated(*envelope.TourCreated)
	})
	bus.Subscribe(events.TypeCyclistCreated, func(envelope *events.Envelope) error {
		return eventHandler.OnCyclistCreated(*envelope.CyclistCreated)
	})

	return bus
}

func startHttp(listenPort int, commandHandler CommandHandler) {
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
			err := commandHandler.HandleCreateGamblerCommand(command)
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
			err := commandHandler.HandleCreateGamblerTeamCommand(command)
			if err != nil {
				http.HandleError(c, err)
				return
			}
			c.JSON(200, *http.SuccessResponse())
		})
	}

	engine.Run(fmt.Sprintf(":%d", listenPort))
}
