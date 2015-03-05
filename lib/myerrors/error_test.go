package myerrors

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestError(t *testing.T) {
	{
		err := makeNil()
		assert.Nil(t, err)
		assert.False(t, IsInternalError(err))
		assert.False(t, IsInvalidInputError(err))
		assert.False(t, IsNotFoundError(err))
		assert.False(t, IsNotAuthorizedError(err))
	}
	{
		err := errors.New("my unclassified error")
		assert.False(t, IsInternalError(err))
		assert.False(t, IsInvalidInputError(err))
		assert.False(t, IsNotFoundError(err))
		assert.False(t, IsNotAuthorizedError(err))
		assert.Equal(t, "my unclassified error", err.Error())
	}
	{
		err := makeInternalError()
		assert.True(t, IsInternalError(err))
		assert.False(t, IsInvalidInputError(err))
		assert.False(t, IsNotFoundError(err))
		assert.False(t, IsNotAuthorizedError(err))
		assert.Equal(t, "my internal error", err.Error())
	}
	{
		err := makeInvalidInputError()
		assert.False(t, IsInternalError(err))
		assert.True(t, IsInvalidInputError(err))
		assert.False(t, IsNotFoundError(err))
		assert.False(t, IsNotAuthorizedError(err))
		assert.Equal(t, "my invalid input error", err.Error())
	}
	{
		err := makeNotFoundError()
		assert.False(t, IsInternalError(err))
		assert.False(t, IsInvalidInputError(err))
		assert.True(t, IsNotFoundError(err))
		assert.False(t, IsNotAuthorizedError(err))
		assert.Equal(t, "my not found error", err.Error())
	}
	{
		err := makeAuthorisationError()
		assert.False(t, IsInternalError(err))
		assert.False(t, IsInvalidInputError(err))
		assert.False(t, IsNotFoundError(err))
		assert.True(t, IsNotAuthorizedError(err))
		assert.Equal(t, "my not authorized error", err.Error())
	}
}

func makeNil() error {
	return nil
}

func makeInternalError() error {
	return NewInternalError(errors.New("my internal error"))
}

func makeInvalidInputError() error {
	return NewInvalidInputError(errors.New("my invalid input error"))
}

func makeNotFoundError() error {
	return NewNotFoundError(errors.New("my not found error"))
}

func makeAuthorisationError() error {
	return NewNotAuthorizedError(errors.New("my not authorized error"))
}
