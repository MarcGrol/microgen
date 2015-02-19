package myerrors

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestError(t *testing.T) {
	{
		err := errors.New("my unclassified error")
		assert.False(t, IsInternalError(err))
		assert.False(t, IsInvalidInputError(err))
		assert.False(t, IsNotFoundError(err))
		assert.False(t, IsNotAuthorizedError(err))
		assert.Equal(t, "my unclassified error", err.Error())
	}
	{
		err := NewInternalError(errors.New("my internal error"))
		assert.True(t, IsInternalError(err))
		assert.False(t, IsInvalidInputError(err))
		assert.False(t, IsNotFoundError(err))
		assert.False(t, IsNotAuthorizedError(err))
		assert.Equal(t, "my internal error", err.Error())
	}
	{
		err := NewInvalidInputError(errors.New("my invalid input error"))
		assert.False(t, IsInternalError(err))
		assert.True(t, IsInvalidInputError(err))
		assert.False(t, IsNotFoundError(err))
		assert.False(t, IsNotAuthorizedError(err))
		assert.Equal(t, "my invalid input error", err.Error())
	}
	{
		err := NewNotFoundError(errors.New("my not found error"))
		assert.False(t, IsInternalError(err))
		assert.False(t, IsInvalidInputError(err))
		assert.True(t, IsNotFoundError(err))
		assert.False(t, IsNotAuthorizedError(err))
		assert.Equal(t, "my not found error", err.Error())
	}
	{
		err := NewNotAuthorizedError(errors.New("my not authorized error"))
		assert.False(t, IsInternalError(err))
		assert.False(t, IsInvalidInputError(err))
		assert.False(t, IsNotFoundError(err))
		assert.True(t, IsNotAuthorizedError(err))
		assert.Equal(t, "my not authorized error", err.Error())
	}
}
