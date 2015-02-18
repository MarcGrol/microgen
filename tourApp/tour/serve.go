package tour

import (
	"errors"
	"fmt"
	"github.com/MarcGrol/microgen/tourApp/events"
	"github.com/gin-gonic/gin"
	"log"
	"strconv"
)

type Response struct {
	status bool
	error  *ErrorDescriptor
}

type ErrorDescriptor struct {
	code    int
	message string
}

func SuccessResponse() *Response {
	resp := new(Response)
	resp.status = true
	resp.error = nil
	return resp
}

func ErrorResponse(code int, message string) *Response {
	resp := new(Response)
	resp.status = false
	resp.error = new(ErrorDescriptor)
	resp.error.code = code
	resp.error.message = message
	return resp
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
				c.JSON(400, *ErrorResponse(4, err.Error()))
				return
			}
			tour, err := queryHandler.GetTour(year)
			if err != nil {
				c.JSON(404, *ErrorResponse(5, err.Error()))
				return
			}
			log.Printf("GetTour: %v", *tour)

			c.JSON(200, *tour)
		})
		api.POST("/tour", func(c *gin.Context) {
			var command CreateTourCommand
			status := c.Bind(&command)
			if status == false {
				c.JSON(400, *ErrorResponse(1, "Invalid input"))
				return
			}
			err := commandHandler.HandleCreateTourCommand(command)
			if err != nil {
				c.JSON(400, *ErrorResponse(2, err.Error()))
				return
			}
			c.JSON(200, *SuccessResponse())
		})
		api.POST("/tour/:year/etappe", func(c *gin.Context) {
			var command CreateEtappeCommand
			status := c.Bind(&command)
			if status == false {
				c.JSON(400, *ErrorResponse(1, "Invalid input"))
				return
			}
			err := commandHandler.HandleCreateEtappeCommand(command)
			if err != nil {
				c.JSON(400, *ErrorResponse(3, err.Error()))
				return
			}
			c.JSON(200, *SuccessResponse())
		})
		api.POST("/tour/:year/cylist", func(c *gin.Context) {
			var command CreateCyclistCommand
			status := c.Bind(&command)
			if status == false {
				c.JSON(400, *ErrorResponse(1, "Invalid input"))
				return
			}
			err := commandHandler.HandleCreateCyclistCommand(command)
			if err != nil {
				c.JSON(400, *ErrorResponse(3, err.Error()))
				return
			}
			c.JSON(200, *SuccessResponse())
		})
	}

	engine.Run(fmt.Sprintf(":%d", listenPort))
}
