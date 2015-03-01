package http

import (
	"github.com/MarcGrol/microgen/myerrors"
	"github.com/gin-gonic/gin"
)

type Response struct {
	Status bool             `json:"status"`
	Error  *ErrorDescriptor `json:"error"`
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

func HandleError(c *gin.Context, err error) {
	if err == nil {
		c.JSON(404, *ErrorResponse(99, "Error must be nil!!!!!!!!!"))
	}
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
