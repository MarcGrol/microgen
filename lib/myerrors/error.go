package myerrors

import "errors"

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

func NewInternalErrorf(msg string) *Error {
	return NewInternalError(errors.New(msg))
}

func NewInternalError(err error) *Error {
	newError := new(Error)
	newError.underlyingError = err
	newError.errorType = errorTypeInternal
	return newError
}

func NewInvalidInputErrorf(msg string) *Error {
	return NewInvalidInputError(errors.New(msg))
}

func NewInvalidInputError(err error) *Error {
	newError := new(Error)
	newError.underlyingError = err
	newError.errorType = errorTypeInvalidInput
	return newError
}

func NewNotFoundErrorf(msg string) *Error {
	return NewNotFoundError(errors.New(msg))
}

func NewNotFoundError(err error) *Error {
	newError := new(Error)
	newError.underlyingError = err
	newError.errorType = errorTypeNotFound
	return newError
}

func NewNotAuthorizedErrorf(msg string) *Error {
	return NewNotAuthorizedError(errors.New(msg))
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
