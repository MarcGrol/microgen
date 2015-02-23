package tour

import (
	"errors"
	"fmt"
	"github.com/MarcGrol/microgen/myerrors"
	"github.com/MarcGrol/microgen/tourApp/events"
	"github.com/MarcGrol/microgen/tourApp/http"
	"github.com/gin-gonic/gin"
	"os"
	"strconv"
)

func Start(listenPort int, busAddress string, baseDir string) error {
	store, err := startStore(baseDir)
	if err != nil {
		return err
	}
	bus := startBus(busAddress)
	if bus == nil {
		return errors.New("Error starting bus")
	}
	startHttp(listenPort, NewTourCommandHandler(bus, store))
	return nil
}

func startStore(baseDir string) (*events.EventStore, error) {
	dataDir := baseDir + "/" + "data"

	// create dir if not exists
	err := os.MkdirAll(dataDir, 0777)
	if err != nil {
		return nil, err
	}
	store := events.NewEventStore(dataDir, "tour.db")
	err = store.Open()
	if err != nil {
		return nil, err
	}
	return store, nil
}

func startBus(busAddress string) *events.EventBus {
	return events.NewEventBus("tourApp", "tour", busAddress)
}

func startHttp(listenPort int, commandHandler CommandHandler) {
	engine := gin.Default()
	api := engine.Group("/api")
	{
		api.GET("/tour/:year", func(c *gin.Context) {
			year, err := strconv.Atoi(c.Params.ByName("year"))
			if err != nil {
				http.HandleError(c, myerrors.NewInvalidInputError(err))
				return
			}
			tourOpaque, err := commandHandler.HandleGetTourQuery(year)
			tour, ok := tourOpaque.(*Tour)
			if err != nil || ok == false {
				http.HandleError(c, err)
			}
			c.JSON(200, *tour)
		})
		api.POST("/tour", func(c *gin.Context) {
			var command CreateTourCommand
			c.Bind(&command)
			err := commandHandler.HandleCreateTourCommand(command)
			if err != nil {
				http.HandleError(c, err)
			}
			c.JSON(200, *http.SuccessResponse())
		})
		api.POST("/tour/:year/etappe", func(c *gin.Context) {
			var command CreateEtappeCommand
			c.Bind(&command)
			err := commandHandler.HandleCreateEtappeCommand(command)
			if err != nil {
				http.HandleError(c, err)
			}
			c.JSON(200, *http.SuccessResponse())
		})
		api.POST("/tour/:year/cylist", func(c *gin.Context) {
			var command CreateCyclistCommand
			c.Bind(&command)
			err := commandHandler.HandleCreateCyclistCommand(command)
			if err != nil {
				http.HandleError(c, err)
			}
			c.JSON(200, *http.SuccessResponse())
		})
	}

	engine.Run(fmt.Sprintf(":%d", listenPort))
}
