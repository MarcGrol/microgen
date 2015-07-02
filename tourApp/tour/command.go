package tour

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

type TourCommandHandler struct {
	bus     infra.PublishSubscriber
	store   infra.Store
	context *tourContext
}

func NewTourCommandHandler(bus infra.PublishSubscriber, store infra.Store, context *tourContext) CommandHandler {
	handler := new(TourCommandHandler)
	handler.bus = bus
	handler.store = store
	handler.context = context
	return handler
}

func (commandHandler *TourCommandHandler) Start(listenPort int) error {
	var err error
	engine := gin.Default()
	api := engine.Group("/api")
	{
		api.GET("/tour/:year", func(c *gin.Context) {
			year, err := strconv.Atoi(c.Params.ByName("year"))
			if err != nil {
				myhttp.HandleError(c, myerrors.NewInvalidInputError(err))
				return
			}
			tour, err := commandHandler.HandleGetTourQuery(year)
			if err != nil {
				myhttp.HandleError(c, err)
				return
			}
			c.JSON(200, *tour)
		})
		api.POST("/tour", func(c *gin.Context) {
			var command CreateTourCommand
			err = c.Bind(&command)
			if err != nil {
				myhttp.HandleError(c, myerrors.NewInvalidInputError(errors.New("Invalid tour-command")))
				return
			}
			err := commandHandler.HandleCreateTourCommand(&command)
			if err != nil {
				myhttp.HandleError(c, err)
				return
			}
			c.JSON(200, *myhttp.SuccessResponse())
		})
		api.POST("/tour/:year/etappe", func(c *gin.Context) {
			var command CreateEtappeCommand
			err = c.Bind(&command)
			if err != nil {
				myhttp.HandleError(c, myerrors.NewInvalidInputError(errors.New("Invalid etappe-command")))
				return
			}
			err = commandHandler.HandleCreateEtappeCommand(&command)
			if err != nil {
				myhttp.HandleError(c, err)
				return
			}
			c.JSON(200, *myhttp.SuccessResponse())
		})
		api.POST("/tour/:year/cyclist", func(c *gin.Context) {
			var command CreateCyclistCommand
			err = c.Bind(&command)
			if err != nil {
				myhttp.HandleError(c, myerrors.NewInvalidInputError(errors.New("Invalid cyclist-command")))
				return
			}
			err := commandHandler.HandleCreateCyclistCommand(&command)
			if err != nil {
				myhttp.HandleError(c, err)
				return
			}
			c.JSON(200, *myhttp.SuccessResponse())
		})
	}

	engine.Run(fmt.Sprintf(":%d", listenPort))

	return nil
}

func (ch *TourCommandHandler) validateCreateTourCommand(command *CreateTourCommand) error {
	v := validation.Validator{}
	v.GreaterThan("Year", 2010, command.Year)
	return v.Err
}

func (ch *TourCommandHandler) HandleCreateTourCommand(command *CreateTourCommand) error {
	// validate command
	err := ch.validateCreateTourCommand(command)
	if err != nil {
		return myerrors.NewInvalidInputError(err)
	}

	// get tour based on year
	_, found := ch.context.tours[command.Year]
	if found == true {
		return myerrors.NewInvalidInputErrorf("Tour %d already exists", command.Year)
	}

	// create event
	tourCreatedEvent := events.TourCreated{Year: command.Year}

	// store and emit resulting event
	err = ch.storeAndPublish([]*envelope.Envelope{tourCreatedEvent.Wrap()})
	if err != nil {
		return err
	}

	// add to local memory
	ch.context.ApplyTourCreated(&tourCreatedEvent)

	log.Printf("HandleCreateTourCommand completed:%+v -> %+v", command, tourCreatedEvent)

	return nil
}

func (ch *TourCommandHandler) validateCreateCyclistCommand(command *CreateCyclistCommand) error {
	v := validation.Validator{}
	v.GreaterThan("Year", 2010, command.Year)
	v.GreaterThan("Id", 0, command.Id)
	v.NotEmpty("Name", command.Name)
	v.NotEmpty("Team", command.Team)
	return v.Err
}

func (ch *TourCommandHandler) HandleCreateCyclistCommand(command *CreateCyclistCommand) error {
	// validate command
	err := ch.validateCreateCyclistCommand(command)
	if err != nil {
		return myerrors.NewInvalidInputError(err)
	}

	// get tour based on year
	tour, found := ch.context.tours[command.Year]
	if found == false {
		return myerrors.NewNotFoundErrorf("Tour %d does not exist", command.Year)
	}

	// verify if cyclist already exists
	if tour.hasCyclist(command.Id) {
		return myerrors.NewInvalidInputErrorf("Cyclist with id %d already exists", command.Id)
	}

	// create event
	cyclistCreatedEvent := events.CyclistCreated{Year: command.Year,
		CyclistId:   command.Id,
		CyclistName: command.Name,
		CyclistTeam: command.Team}

	// store and emit resulting event
	err = ch.storeAndPublish([]*envelope.Envelope{cyclistCreatedEvent.Wrap()})
	if err != nil {
		return err
	}

	ch.context.ApplyCyclistCreated(&cyclistCreatedEvent)

	log.Printf("HandleCreateCyclistCommand completed:%+v -> %+v", command, cyclistCreatedEvent)

	return nil
}

func (ch *TourCommandHandler) validateCreateEtappeCommand(command *CreateEtappeCommand) error {
	v := validation.Validator{}
	v.GreaterThan("Year", 2010, command.Year)
	v.GreaterThan("Id", 0, command.Id)
	v.NotEmpty("StartLocation", command.StartLocation)
	v.NotEmpty("FinishLocation", command.FinishLocation)
	v.GreaterThan("Length", 0, command.Length)
	v.GreaterThan("Kind", -1, command.Kind)
	v.After("Date", "2010-07-01T00:00:00Z", command.Date)

	return v.Err
}

func (ch *TourCommandHandler) HandleCreateEtappeCommand(command *CreateEtappeCommand) error {
	// validate command
	err := ch.validateCreateEtappeCommand(command)
	if err != nil {
		return myerrors.NewInvalidInputError(err)
	}

	// get tour based on year
	tour, found := ch.context.tours[command.Year]
	if found == false {
		return myerrors.NewNotFoundErrorf("Tour %d does not exist", command.Year)
	}

	// verify if etappe already exists
	if tour.hasEtappe(command.Id) {
		return myerrors.NewInvalidInputErrorf("Etappe with id %d already exists", command.Id)
	}

	// create event
	etappeCreatedEvent := events.EtappeCreated{
		Year:                 command.Year,
		EtappeId:             command.Id,
		EtappeDate:           command.Date,
		EtappeStartLocation:  command.StartLocation,
		EtappeFinishLocation: command.FinishLocation,
		EtappeLength:         command.Length,
		EtappeKind:           command.Kind}

	// store and emit resulting event
	err = ch.storeAndPublish([]*envelope.Envelope{etappeCreatedEvent.Wrap()})
	if err != nil {
		return err
	}

	ch.context.ApplyEtappeCreated(&etappeCreatedEvent)

	log.Printf("HandleCreateEtappeCommand completed:%+v -> %+v", command, etappeCreatedEvent)

	return nil
}

func (ch *TourCommandHandler) validateCreateEtappeResultsCommand(command *CreateEtappeResultsCommand) error {
	v := validation.Validator{}
	v.GreaterThan("Year", 2010, command.Year)
	v.GreaterThan("EtappeId", 0, command.EtappeId)

	v.MinSliceLength("BestDayCyclistIds", 10, command.BestDayCyclistIds)
	v.NoDuplicates("BestDayCyclistIds", command.BestDayCyclistIds)

	v.MinSliceLength("BestAllroundCyclistIds", 5, command.BestAllroundCyclistIds)
	v.NoDuplicates("BestAllroundCyclistIds", command.BestAllroundCyclistIds)

	v.MinSliceLength("BestClimbCyclistIds", 5, command.BestClimbCyclistIds)
	v.NoDuplicates("BestClimbCyclistIds", command.BestClimbCyclistIds)

	v.MinSliceLength("BestSprintCyclistIds", 5, command.BestSprintCyclistIds)
	v.NoDuplicates("BestSprintCyclistIds", command.BestSprintCyclistIds)

	return v.Err
}

func (ch *TourCommandHandler) HandleCreateEtappeResultsCommand(command *CreateEtappeResultsCommand) error {
	// validate command
	err := ch.validateCreateEtappeResultsCommand(command)
	if err != nil {
		return myerrors.NewInvalidInputError(err)
	}

	// get tour based on year
	tour, found := ch.context.tours[command.Year]
	if found == false {
		return myerrors.NewNotFoundErrorf("Tour %d does not exist", command.Year)
	}

	// verify that etappe already exists
	if tour.hasEtappe(command.EtappeId) == false {
		return myerrors.NewNotFoundErrorf("Etappe with id %d does not exist", command.EtappeId)
	}

	// verify that referenced cyclists already exists
	verify := fluentError{}
	verify.cyclistsExist("BestDayCyclistIds", tour, command.BestDayCyclistIds)
	verify.cyclistsExist("BestAllroundCyclistIds", tour, command.BestAllroundCyclistIds)
	verify.cyclistsExist("BestSprintCyclistIds", tour, command.BestSprintCyclistIds)
	verify.cyclistsExist("BestClimbCyclistIds", tour, command.BestClimbCyclistIds)
	if verify.err != nil {
		return verify.err
	}

	// compose event
	etappeResultCreatedEvent := events.EtappeResultsCreated{
		Year:                     command.Year,
		LastEtappeId:             command.EtappeId,
		BestDayCyclistIds:        command.BestDayCyclistIds,
		BestAllrounderCyclistIds: command.BestAllroundCyclistIds,
		BestSprinterCyclistIds:   command.BestSprintCyclistIds,
		BestClimberCyclistIds:    command.BestClimbCyclistIds}

	// store and emit resulting event
	err = ch.storeAndPublish([]*envelope.Envelope{etappeResultCreatedEvent.Wrap()})
	if err != nil {
		return err
	}

	ch.context.ApplyEtappeResultsCreated(&etappeResultCreatedEvent)

	log.Printf("HandleCreateEtappeResultsCommand completed:%+v -> %+v", command, etappeResultCreatedEvent)

	return nil
}

type fluentError struct {
	err error
}

func (v *fluentError) cyclistsExist(name string, tour *Tour, cyclistIds []int) error {
	if v.err == nil {
		for _, id := range cyclistIds {
			if tour.hasCyclist(id) == false {
				v.err = myerrors.NewNotFoundErrorf("%s: Cyclist with id %s does not exist",
					name, fmt.Sprintf("%d", id))
				break
			}
		}
	}
	return v.err
}

func (ch *TourCommandHandler) HandleGetTourQuery(year int) (*Tour, error) {
	// TODO validate input
	tour, found := ch.context.tours[year]
	if found == false {
		log.Printf("not found")
		return nil, myerrors.NewNotFoundErrorf("Tour %d not found", year)
	}
	log.Printf("GetTour:%+v", tour)

	return tour, nil
}

func (ch *TourCommandHandler) storeAndPublish(envelopes []*envelope.Envelope) error {
	for _, env := range envelopes {
		err := ch.store.Store(env)
		if err != nil {
			return myerrors.NewInternalError(err)
		}
		err = ch.bus.Publish(env)
		if err != nil {
			return myerrors.NewInternalError(err)
		}
	}
	return nil
}
