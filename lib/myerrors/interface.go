package myerrors

import ()

type InternalError interface {
	IsInternalError() bool
}

type InvalidInput interface {
	IsInvalidInputError() bool
}

type NotFound interface {
	IsNotFoundError() bool
}

type NotAuthorized interface {
	IsNotAuthorizedError() bool
}

func IsInternalError(err error) bool {
	if err != nil {
		if specificError, ok := err.(InternalError); ok {
			return specificError.IsInternalError()
		}
	}
	return false
}

func IsInvalidInputError(err error) bool {
	if err != nil {
		if specificError, ok := err.(InvalidInput); ok {
			return specificError.IsInvalidInputError()
		}
	}
	return false
}

func IsNotFoundError(err error) bool {
	if err != nil {
		if specificError, ok := err.(NotFound); ok {
			return specificError.IsNotFoundError()
		}
	}
	return false
}

func IsNotAuthorizedError(err error) bool {
	if err != nil {
		if specificError, ok := err.(NotAuthorized); ok {
			return specificError.IsNotAuthorizedError()
		}
	}
	return false
}
