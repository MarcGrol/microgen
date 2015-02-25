package collector

import (
	"errors"
	"fmt"
	"github.com/MarcGrol/microgen/tourApp/events"
	"github.com/MarcGrol/microgen/tourApp/http"
	"github.com/gin-gonic/gin"
	"os"
)

func Start(listenPort int, busAddress string, baseDir string) error {
	store, err := startStore(baseDir)
	if err != nil {
		return err
	}
	bus := startBus(busAddress, NewCollectorEventHandler(store))
	if bus == nil {
		return errors.New("Error starting bus")
	}
	startHttp(listenPort, NewCollectorCommandHandler(bus, store))
	return nil
}

func startStore(baseDir string) (*events.EventStore, error) {
	dataDir := baseDir + "/" + "data"

	// create dir if not exists
	err := os.MkdirAll(dataDir, 0777)
	if err != nil {
		return nil, err
	}
	store := events.NewEventStore(dataDir, "gambler.db")
	err = store.Open()
	if err != nil {
		return nil, err
	}
	return store, nil
}

func startBus(busAddress string, eventHandler EventHandler) *events.EventBus {
	bus := events.NewEventBus("tourApp", "collector", busAddress)

	for _, eventType := range events.GetAllEventsTypes() {
		bus.Subscribe(eventType, func(envelope *events.Envelope) error {
			return eventHandler.OnEvent(envelope)
		})
	}

	return bus
}

func startHttp(listenPort int, commandHandler CommandHandler) {
	engine := gin.Default()
	api := engine.Group("/api")
	{
		api.GET("/events", func(c *gin.Context) {
			eventType := c.Params.ByName("eventType")
			aggregateType := c.Params.ByName("aggregateType")
			aggregateUid := c.Params.ByName("aggregateUid")
			results, err := commandHandler.HandleSearchQuery(eventType, aggregateType, aggregateUid)
			if err != nil {
				http.HandleError(c, err)
				return
			}
			c.JSON(200, *results)
		})
	}

	engine.Run(fmt.Sprintf(":%d", listenPort))
}
