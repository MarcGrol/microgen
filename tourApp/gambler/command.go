package gambler

import (
	"errors"
	"fmt"
	"log"
	"strconv"

	"github.com/MarcGrol/microgen/infra"
	"github.com/MarcGrol/microgen/infra/myhttp"
	"github.com/MarcGrol/microgen/lib/envelope"
	"github.com/MarcGrol/microgen/lib/myerrors"
	"github.com/MarcGrol/microgen/lib/validation"
	"github.com/MarcGrol/microgen/tourApp/events"
	"github.com/gin-gonic/gin"
)

type GamblerCommandHandler struct {
	bus     infra.PublishSubscriber
	store   infra.Store
	context *GamblingContext
}

func NewGamblerCommandHandler(bus infra.PublishSubscriber, store infra.Store, context *GamblingContext) *GamblerCommandHandler {
	handler := new(GamblerCommandHandler)
	handler.bus = bus
	handler.store = store
	handler.context = context
	return handler
}

func (commandHandler *GamblerCommandHandler) Start(listenPort int) error {
	var err error
	engine := gin.Default()
	api := engine.Group("/api")
	{
		api.POST("/gambler", func(c *gin.Context) {
			var command CreateGamblerCommand
			err = c.Bind(&command)
			if err != nil {
				myhttp.HandleError(c, myerrors.NewInvalidInputError(errors.New("Invalid create-gambler-command")))
				return
			}
			err := commandHandler.HandleCreateGamblerCommand(&command)
			if err != nil {
				myhttp.HandleError(c, err)
				return
			}
			c.JSON(200, *myhttp.SuccessResponse())
		})
		api.GET("/gambler", func(c *gin.Context) {
			gamblers, err := commandHandler.HandleGetGamblersQuery()
			if err != nil {
				myhttp.HandleError(c, err)
				return
			}
			c.JSON(200, gamblers)
		})
		api.POST("gambler/:gamblerUid/year/:year/team", func(c *gin.Context) {
			var command CreateGamblerTeamCommand
			err = c.Bind(&command)
			if err != nil {
				myhttp.HandleError(c, myerrors.NewInvalidInputError(errors.New("Invalid create-gambler-team-command")))
				return
			}
			err := commandHandler.HandleCreateGamblerTeamCommand(&command)
			if err != nil {
				myhttp.HandleError(c, err)
				return
			}
			c.JSON(200, *myhttp.SuccessResponse())
		})
		api.GET("/gambler/:gamblerUid/year/:year", func(c *gin.Context) {
			gamblerUid := c.Params.ByName("gamblerUid")
			year, err := strconv.Atoi(c.Params.ByName("year"))
			if err != nil {
				myhttp.HandleError(c, myerrors.NewInvalidInputError(err))
				return
			}
			gambler, err := commandHandler.HandleGetGamblerQuery(gamblerUid, year)
			if err != nil {
				myhttp.HandleError(c, err)
				return
			}
			c.JSON(200, *gambler)
		})
		api.GET("/results/:year", func(c *gin.Context) {
			year, err := strconv.Atoi(c.Params.ByName("year"))
			if err != nil {
				myhttp.HandleError(c, myerrors.NewInvalidInputError(err))
				return
			}
			results, err := commandHandler.HandleGetResultsQuery(year)
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

func (ch *GamblerCommandHandler) validateCreateGamblerCommand(command *CreateGamblerCommand) error {
	v := validation.Validator{}
	v.NotEmpty("GamblerUid", command.GamblerUid)
	v.NotEmpty("Name", command.Name)
	v.NotEmpty("Email", command.Email)

	return v.Err
}

func (ch *GamblerCommandHandler) HandleCreateGamblerCommand(command *CreateGamblerCommand) error {
	err := ch.validateCreateGamblerCommand(command)
	if err != nil {
		return myerrors.NewInvalidInputError(err)
	}

	// apply business logic
	gamblerCreatedEvent := events.GamblerCreated{
		GamblerUid:   command.GamblerUid,
		GamblerName:  command.Name,
		GamblerEmail: command.Email}

	// store and emit resulting event
	err = doStoreAndPublish(ch.store, ch.bus, []*envelope.Envelope{gamblerCreatedEvent.Wrap()})
	if err != nil {
		return err
	}

	ch.context.ApplyGamblerCreated(&gamblerCreatedEvent)

	return nil
}

func (ch *GamblerCommandHandler) validateCreateGamblerTeamCommand(command *CreateGamblerTeamCommand) error {
	v := validation.Validator{}
	v.GreaterThan("Year", 2010, command.Year)
	v.NotEmpty("GamblerUid", command.GamblerUid)
	v.MinSliceLength("CyclistIds", 10, command.CyclistIds)
	v.NoDuplicates("CyclistIds", command.CyclistIds)
	return v.Err
}

func (ch *GamblerCommandHandler) HandleCreateGamblerTeamCommand(command *CreateGamblerTeamCommand) error {
	err := ch.validateCreateGamblerTeamCommand(command)
	if err != nil {
		return myerrors.NewInvalidInputError(err)
	}

	gamblingYear, exists := ch.context.years[command.Year]
	if exists == false {
		return myerrors.NewNotFoundErrorf("Tour %d not found", command.Year)
	}
	_, exists = ch.context.gamblers[command.GamblerUid]
	if exists == false {
		return myerrors.NewNotFoundErrorf("Gambler %s not found", command.GamblerUid)
	}

	err = cyclistsExist(gamblingYear.cyclistsForTour, command.CyclistIds)
	if err != nil {
		return myerrors.NewNotFoundError(err)
	}

	// apply business logic
	gamblerTeamCreatedEvent := events.GamblerTeamCreated{
		GamblerUid:      command.GamblerUid,
		Year:            command.Year,
		GamblerCyclists: command.CyclistIds}

	err = doStoreAndPublish(ch.store, ch.bus, []*envelope.Envelope{gamblerTeamCreatedEvent.Wrap()})
	if err != nil {
		return err
	}

	ch.context.ApplyGamblerTeamCreated(&gamblerTeamCreatedEvent)

	return nil
}

func cyclistsExist(allCyclists map[int]*Cyclist, cyclistIds []int) error {
	for _, id := range cyclistIds {
		_, exists := allCyclists[id]
		if exists == false {
			return fmt.Errorf("Cyclist %d does not exist", id)
		}
	}
	return nil
}

func doStore(store infra.Store, envelopes []*envelope.Envelope) error {
	for _, env := range envelopes {
		err := store.Store(env)
		if err != nil {
			log.Printf("Error storing event: %+v", err)
			return err
		}
		//log.Printf("Successfully stored event: %+v", env)
	}
	return nil
}

func doStoreAndPublish(store infra.Store, bus infra.PublishSubscriber, envelopes []*envelope.Envelope) error {
	err := doStore(store, envelopes)
	if err != nil {
		return myerrors.NewInternalError(err)
	}
	for _, envelop := range envelopes {
		err = bus.Publish(envelop)
		if err != nil {
			return myerrors.NewInternalError(err)
		}
	}
	return nil
}

func (ch *GamblerCommandHandler) HandleGetGamblerQuery(gamblerUid string, year int) (*Gambler, error) {

	log.Printf("************* HandleGetGamblerQuery:%+v", ch)
	gamblingYear, exists := ch.context.years[year]
	if exists == false {
		gambler, exists := ch.context.gamblers[gamblerUid]
		if exists == false {
			return nil, myerrors.NewNotFoundErrorf("Gambler %s not found", gamblerUid)
		}
		return &gambler, nil
	}
	gambler, found := gamblingYear.gamblersForTour[gamblerUid]
	if found == false {
		return nil, myerrors.NewNotFoundErrorf("Gambler with uid %s not found", gamblerUid)
	}

	log.Printf("HandleGetGamblerQuery.Gambler:%+v", gambler)

	return gambler, nil
}

func (ch *GamblerCommandHandler) HandleGetGamblersQuery() ([]Gambler, error) {
	gamblers := make([]Gambler, 0, 10)

	for _, gambler := range ch.context.gamblers {
		gamblers = append(gamblers, gambler)
	}
	return gamblers, nil
}

func (ch *GamblerCommandHandler) HandleGetResultsQuery(year int) (*Results, error) {
	return nil, errors.New("HandleGetResultsQuery not implemented")
}
