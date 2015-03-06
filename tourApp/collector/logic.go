package collector

import (
	"errors"
	"github.com/MarcGrol/microgen/infra"
	"github.com/MarcGrol/microgen/lib/envelope"
	"github.com/MarcGrol/microgen/lib/myerrors"
	"log"
)

func (eh *CollectorEventHandler) OnAnyEvent(envelop *envelope.Envelope) error {
	log.Printf("OnEvent: envelope: %+v", envelop)
	return doStore(eh.store, []*envelope.Envelope{envelop})
}

func doStore(store infra.Store, envelopes []*envelope.Envelope) error {
	for _, env := range envelopes {
		err := store.Store(env)
		if err != nil {
			log.Printf("Error storing event: %+v", err)
			return err
		}
		log.Printf("Successfully stored event: %+v", env)
	}
	return nil
}

func (ch *CollectorCommandHandler) HandleSearchQuery(eventType string, aggregateType string, aggregateUid string) (*SearchResults, error) {
	var err *myerrors.Error = nil

	results := new(SearchResults)
	if len(aggregateType) > 0 && len(aggregateUid) > 0 {
		results.Events, _ = ch.store.Get(aggregateType, aggregateUid)
	} else {
		err = myerrors.NewInternalError(errors.New("Search on event-type not supported"))
	}

	return nil, err
}

type SearchResults struct {
	Events []envelope.Envelope `json:"events"`
}
