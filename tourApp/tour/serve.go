package tour

import (
	"fmt"
	"github.com/gin-gonic/gin"
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

func StartHttp(listenPort int, commandHandler CommandHandler, eventHandler EventHandler) {
	engine := gin.Default()
	api := engine.Group("/api")
	{
		api.POST("/tour", func(c *gin.Context) {
			var command CreateTourCommand
			status := c.Bind(&command)
			if status == false {
				c.JSON(400, *ErrorResponse(1, "Invalid input"))
			} else {
				err := commandHandler.HandleCreateTourCommand(command)
				if err != nil {
					c.JSON(400, *ErrorResponse(2, err.Error()))
				} else {
					c.JSON(200, *SuccessResponse())
				}
			}
		})
		api.POST("/tour/:year/etappe", func(c *gin.Context) {
			var command CreateEtappeCommand
			status := c.Bind(&command)
			if status == false {
				c.JSON(400, *ErrorResponse(1, "Invalid input"))
			} else {
				err := commandHandler.HandleCreateEtappeCommand(command)
				if err != nil {
					c.JSON(400, *ErrorResponse(3, err.Error()))
				} else {
					c.JSON(200, *SuccessResponse())
				}
			}
		})
		api.POST("/tour/:year/cylist", func(c *gin.Context) {
			var command CreateCyclistCommand
			status := c.Bind(&command)
			if status == false {
				c.JSON(400, *ErrorResponse(1, "Invalid input"))
			} else {
				err := commandHandler.HandleCreateCyclistCommand(command)
				if err != nil {
					c.JSON(400, *ErrorResponse(3, err.Error()))
				} else {
					c.JSON(200, *SuccessResponse())
				}
			}
		})
	}

	engine.Run(fmt.Sprintf(":%d", listenPort))
}
