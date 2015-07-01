package tour

import (
	"errors"
	"os"

	"github.com/MarcGrol/microgen/infra"
	"github.com/MarcGrol/microgen/infra/bus"
	"github.com/MarcGrol/microgen/infra/store"
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

	// load events from store and populate in memory structure
	tour := NewTour()
	envelopes, err := store.GetAll()
	if err != nil {
		return err
	}
	tour.ApplyAll(envelopes)

	// has no events to listem for: start anyway
	eventHandler := NewTourEventHandler(bus, store, tour)
	eventHandler.Start()

	// command-handler: start web-server: blocking call
	commandHandler := NewTourCommandHandler(bus, store, tour)
	commandHandler.Start(listenPort)

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
