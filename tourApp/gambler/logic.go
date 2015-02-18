package gambler

import (
	"errors"
	"fmt"
	"github.com/MarcGrol/microgen/tourApp/events"
)

type GamblerCommandHandler struct {
	bus   events.PublishSubscriber
	store events.Store
}

func NewGamblerCommandHandler(bus events.PublishSubscriber, store events.Store) CommandHandler {
	handler := new(GamblerCommandHandler)
	handler.bus = bus
	handler.store = store
	return handler
}

func (gch *GamblerCommandHandler) HandleCreateGamblerCommand(command CreateGamblerCommand) error {
	// get gambler based on uid
	_, found := getGamblerOnUid(gch.store, command.GamblerUid)
	if found == true {
		return errors.New(fmt.Sprintf("gambler %s already exists", command.GamblerUid))
	}

	// apply business logic
	gambler := Gambler{}
	gamblerCreatedEvent := events.GamblerCreated{
		GamblerUid:   command.GamblerUid,
		GamblerName:  command.Name,
		GamblerEmail: command.Email}
	gambler.ApplyGamblerCreated(gamblerCreatedEvent)

	// store and emit resulting event
	return gch.publishAndStore([]*events.Envelope{gamblerCreatedEvent.Wrap()})
}

func (gch *GamblerCommandHandler) HandleCreateGamblerTeamCommand(command CreateGamblerTeamCommand) error {
	// get gambler based on uid
	gambler, found := getGamblerOnUid(gch.store, command.GamblerUid)
	if found == false {
		return errors.New(fmt.Sprintf("gambler %s does not exist", command.GamblerUid))
	}

	// apply business logic
	gamblerTeamCreatedEvent := events.GamblerTeamCreated{
		GamblerUid:      command.GamblerUid,
		Year:            command.Year,
		GamblerCyclists: command.CyclistIds}
	gambler.ApplyGamblerTeamCreated(gamblerTeamCreatedEvent)

	return gch.publishAndStore([]*events.Envelope{gamblerTeamCreatedEvent.Wrap()})
}

func (tch *GamblerCommandHandler) publishAndStore([]*events.Envelope) error {
	return errors.New("publishAndStore not implemented")
}

func getGamblerOnUid(store events.Store, uid string) (*Gambler, bool) {
	var gamblerCreatedEvent *events.GamblerCreated = nil

	calllback := func(envelope *events.Envelope) {
		if envelope.Type == events.TypeGamblerCreated && envelope.GamblerCreated != nil && envelope.GamblerCreated.GamblerUid == uid {
			gamblerCreatedEvent = envelope.GamblerCreated
		}
	}
	store.Iterate(calllback)

	if gamblerCreatedEvent == nil {
		return nil, false
	}

	gambler := Gambler{}
	gambler.ApplyGamblerCreated(*gamblerCreatedEvent)
	return &gambler, true
}

type Gambler struct {
	Uid      string
	Name     string
	Email    string
	Cyclists []Cyclist
}

type Cyclist struct {
	Id   int
	Name string
	Team string
}

func NewGambler(uid string, name string, email string) *Gambler {
	gambler := new(Gambler)
	gambler.Uid = uid
	gambler.Name = name
	gambler.Email = email
	return gambler
}

func (g *Gambler) ApplyGamblerCreated(event events.GamblerCreated) error {
	g.Uid = event.GamblerUid
	g.Name = event.GamblerName
	g.Email = event.GamblerEmail
	return nil
}

func (g *Gambler) ApplyTourCreated(event events.TourCreated) error {
	return nil
}

func (g *Gambler) ApplyGamblerTeamCreated(event events.GamblerTeamCreated) error {
	return nil
}
