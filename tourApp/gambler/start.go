package gambler

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

	gamblingContext := NewGamblingContext()
	envelopes, err := store.GetAll()
	if err != nil {
		return err
	}
	gamblingContext.ApplyAll(envelopes)

	// event-handler
	eventHandler := NewGamblerEventHandler(bus, store, gamblingContext)
	err = eventHandler.Start()
	if err != nil {
		return err
	}

	// command-handler: start web-server: blocking call
	commandHandler := NewGamblerCommandHandler(bus, store, gamblingContext)
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
