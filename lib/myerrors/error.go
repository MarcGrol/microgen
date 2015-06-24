package myerrors

import "fmt"

type Error struct {
	underlyingError error
	errorType       errorType
}

type errorType int

const (
	errorTypeUnclassified errorType = iota
	errorTypeInternal
	errorTypeInvalidInput
	errorTypeNotFound
	errorTypeNotAuthorized
)

func NewInternalErrorf(format string, args interface{}) *Error {
	return NewInternalError(fmt.Errorf(format, args))
}

func NewInternalError(err error) *Error {
	newError := new(Error)
	newError.underlyingError = err
	newError.errorType = errorTypeInternal
	return newError
}

func NewInvalidInputErrorf(format string, args interface{}) *Error {
	return NewInvalidInputError(fmt.Errorf(format, args))
}

func NewInvalidInputError(err error) *Error {
	newError := new(Error)
	newError.underlyingError = err
	newError.errorType = errorTypeInvalidInput
	return newError
}

func NewNotFoundErrorf(format string, args interface{}) *Error {
	return NewNotFoundError(fmt.Errorf(format, args))
}

func NewNotFoundError(err error) *Error {
	newError := new(Error)
	newError.underlyingError = err
	newError.errorType = errorTypeNotFound
	return newError
}

func NewNotAuthorizedErrorf(format string, args interface{}) *Error {
	return NewNotAuthorizedError(fmt.Errorf(format, args))
}

func NewNotAuthorizedError(err error) *Error {
	newError := new(Error)
	newError.underlyingError = err
	newError.errorType = errorTypeNotAuthorized
	return newError
}

func (err Error) Error() string {
	return err.underlyingError.Error()
}

func (err Error) IsInternalError() bool {
	return err.errorType == errorTypeInternal
}

func (err Error) IsInvalidInputError() bool {
	return err.errorType == errorTypeInvalidInput
}

func (err Error) IsNotFoundError() bool {
	return err.errorType == errorTypeNotFound
}

func (err Error) IsNotAuthorizedError() bool {
	return err.errorType == errorTypeNotAuthorized
}
