package tour

import (
	"errors"
	"fmt"
	"github.com/MarcGrol/microgen/myerrors"
	"github.com/MarcGrol/microgen/tourApp/events"
	"github.com/gin-gonic/gin"
	"strconv"
)

type Response struct {
	Status bool             `json:"status"`
	Error  *ErrorDescriptor `json:"errorDescriptor"`
}

type ErrorDescriptor struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func SuccessResponse() *Response {
	resp := new(Response)
	resp.Status = true
	resp.Error = nil
	return resp
}

func ErrorResponse(code int, message string) *Response {
	resp := new(Response)
	resp.Status = false
	resp.Error = new(ErrorDescriptor)
	resp.Error.Code = code
	resp.Error.Message = message
	return resp
}

func handleError(c *gin.Context, err error) {
	if myerrors.IsNotFoundError(err) {
		c.JSON(404, *ErrorResponse(1, err.Error()))
	} else if myerrors.IsInternalError(err) {
		c.JSON(500, *ErrorResponse(2, err.Error()))
	} else if myerrors.IsInvalidInputError(err) {
		c.JSON(400, *ErrorResponse(3, err.Error()))
	} else if myerrors.IsNotAuthorizedError(err) {
		c.JSON(403, *ErrorResponse(4, err.Error()))
	} else {
		c.JSON(500, *ErrorResponse(5, err.Error()))
	}
}

func Start(listenPort int, busAddress string) error {
	store, err := startStore()
	if err != nil {
		return err
	}
	bus := startBus(busAddress)
	if bus == nil {
		return errors.New("Error starting bus")
	}
	startHttp(listenPort, NewTourCommandHandler(bus, store), NewTourQueryHandler(bus, store))
	return nil
}

func startStore() (*events.EventStore, error) {
	store := events.NewEventStore()
	err := store.Open("tour.db")
	if err != nil {
		return nil, err
	}
	return store, nil
}

func startBus(busAddress string) *events.EventBus {
	return events.NewEventBus("tourApp", "tour", busAddress)
}

func startHttp(listenPort int, commandHandler CommandHandler, queryHandler *TourQueryHandler) {
	engine := gin.Default()
	api := engine.Group("/api")
	{
		api.GET("/tour/:year", func(c *gin.Context) {
			year, err := strconv.Atoi(c.Params.ByName("year"))
			if err != nil {
				handleError(c, myerrors.NewInvalidInputError(err))
				return
			}
			tour, err := queryHandler.GetTour(year)
			if err != nil {
				handleError(c, err)
			}
			c.JSON(200, *tour)
		})
		api.POST("/tour", func(c *gin.Context) {
			var command CreateTourCommand
			c.Bind(&command)
			err := commandHandler.HandleCreateTourCommand(command)
			if err != nil {
				handleError(c, err)
			}
			c.JSON(200, *SuccessResponse())
		})
		api.POST("/tour/:year/etappe", func(c *gin.Context) {
			var command CreateEtappeCommand
			c.Bind(&command)
			err := commandHandler.HandleCreateEtappeCommand(command)
			if err != nil {
				handleError(c, err)
			}
			c.JSON(200, *SuccessResponse())
		})
		api.POST("/tour/:year/cylist", func(c *gin.Context) {
			var command CreateCyclistCommand
			c.Bind(&command)
			err := commandHandler.HandleCreateCyclistCommand(command)
			if err != nil {
				handleError(c, err)
			}
			c.JSON(200, *SuccessResponse())
		})
	}

	engine.Run(fmt.Sprintf(":%d", listenPort))
}
