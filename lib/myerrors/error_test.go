package myerrors

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
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
		err := NewInternalErrorf("my %s error", "internal")
		assert.True(t, IsInternalError(err))
		assert.False(t, IsInvalidInputError(err))
		assert.False(t, IsNotFoundError(err))
		assert.False(t, IsNotAuthorizedError(err))
		assert.Equal(t, "my internal error", err.Error())
	}
	{
		err := NewInvalidInputErrorf("my %s error", "invalid input")
		assert.False(t, IsInternalError(err))
		assert.True(t, IsInvalidInputError(err))
		assert.False(t, IsNotFoundError(err))
		assert.False(t, IsNotAuthorizedError(err))
		assert.Equal(t, "my invalid input error", err.Error())
	}
	{
		err := NewNotFoundErrorf("my %s error", "not found")
		assert.False(t, IsInternalError(err))
		assert.False(t, IsInvalidInputError(err))
		assert.True(t, IsNotFoundError(err))
		assert.False(t, IsNotAuthorizedError(err))
		assert.Equal(t, "my not found error", err.Error())
	}
	{
		err := NewNotAuthorizedErrorf("my %s error", "not authorized")
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
