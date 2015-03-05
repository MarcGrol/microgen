package tour

import (
	"errors"
	"fmt"
	"github.com/MarcGrol/microgen/infra"
	"github.com/MarcGrol/microgen/infra/bus"
	"github.com/MarcGrol/microgen/infra/http"
	"github.com/MarcGrol/microgen/infra/store"
	"github.com/MarcGrol/microgen/lib/myerrors"
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

func startStore(baseDir string) (infra.Store, error) {
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

func startBus(busAddress string) infra.PublishSubscriber {
	return bus.NewEventBus("tourApp", "tour", busAddress)
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
			tour, err := commandHandler.HandleGetTourQuery(year)
			if err != nil {
				http.HandleError(c, err)
				return
			}
			c.JSON(200, *tour)
		})
		api.POST("/tour", func(c *gin.Context) {
			var command CreateTourCommand
			ok := c.Bind(&command)
			if ok == false {
				http.HandleError(c, myerrors.NewInvalidInputError(errors.New("Invalid tour-command")))
				return
			}
			err := commandHandler.HandleCreateTourCommand(&command)
			if err != nil {
				http.HandleError(c, err)
				return
			}
			c.JSON(200, *http.SuccessResponse())
		})
		api.POST("/tour/:year/etappe", func(c *gin.Context) {
			var command CreateEtappeCommand
			ok := c.Bind(&command)
			if ok == false {
				http.HandleError(c, myerrors.NewInvalidInputError(errors.New("Invalid etappe-command")))
				return
			}
			err := commandHandler.HandleCreateEtappeCommand(&command)
			if err != nil {
				http.HandleError(c, err)
				return
			}
			c.JSON(200, *http.SuccessResponse())
		})
		api.POST("/tour/:year/cyclist", func(c *gin.Context) {
			var command CreateCyclistCommand
			ok := c.Bind(&command)
			if ok == false {
				http.HandleError(c, myerrors.NewInvalidInputError(errors.New("Invalid cyclist-command")))
				return
			}
			err := commandHandler.HandleCreateCyclistCommand(&command)
			if err != nil {
				http.HandleError(c, err)
				return
			}
			c.JSON(200, *http.SuccessResponse())
		})
	}

	engine.Run(fmt.Sprintf(":%d", listenPort))
}
