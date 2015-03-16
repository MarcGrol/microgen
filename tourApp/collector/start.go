package collector

import (
	"errors"
	"fmt"
	"github.com/MarcGrol/microgen/infra"
	"github.com/MarcGrol/microgen/infra/bus"
	"github.com/MarcGrol/microgen/infra/myhttp"
	"github.com/MarcGrol/microgen/infra/store"
	"github.com/gin-gonic/gin"
	"os"
)

func Start(listenPort int, busAddress string, baseDir string) error {
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
	eventHandler := NewCollectorEventHandler(bus, store)
	err = eventHandler.Start()
	if err != nil {
		return err
	}

	// command-handler: start web-server: blocking call
	commandHandler := NewCollectorCommandHandler(bus, store)
	commandHandler.Start(listenPort)

	return nil
}

func (commandHandler *CollectorCommandHandler) Start(listenPort int) error {
	engine := gin.Default()
	api := engine.Group("/api")
	{
		api.GET("/events", func(c *gin.Context) {
			eventType := c.Params.ByName("eventType")
			aggregateType := c.Params.ByName("aggregateType")
			aggregateUid := c.Params.ByName("aggregateUid")
			results, err := commandHandler.HandleSearchQuery(eventType, aggregateType, aggregateUid)
			if err != nil {
				myhttp.HandleError(c, err)
				return
			}
			c.JSON(200, *results)
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
	st := store.NewEventStore(dataDir, "collector.db")
	err = st.Open()
	if err != nil {
		return nil, err
	}
	return st, nil
}

func createBus(busAddress string) (infra.PublishSubscriber, error) {
	bus := bus.NewEventBus("tourApp", "collector", busAddress)
	if bus == nil {
		return nil, errors.New("Error starting bus")
	}
	return bus, nil

}
