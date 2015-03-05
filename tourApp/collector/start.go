package collector

import (
	"errors"
	"fmt"
	"github.com/MarcGrol/microgen/infra"
	"github.com/MarcGrol/microgen/infra/bus"
	"github.com/MarcGrol/microgen/infra/http"
	"github.com/MarcGrol/microgen/infra/store"
	"github.com/MarcGrol/microgen/lib/envelope"
	"github.com/MarcGrol/microgen/tourApp/events"
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

func startStore(baseDir string) (infra.Store, error) {
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

func startBus(busAddress string, eventHandler EventHandler) infra.PublishSubscriber {
	bus := bus.NewEventBus("tourApp", "collector", busAddress)

	for _, eventType := range events.GetAllEventsTypes() {
		bus.Subscribe(eventType.String(), func(envelope *envelope.Envelope) error {
			return eventHandler.OnAnyEvent(envelope)
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
