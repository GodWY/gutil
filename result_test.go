package gutil

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRetSuccess(t *testing.T) {
	assert.Equal(t, RetSuccess("error", nil).Code, int32(0))
	assert.Equal(t, RetSuccess("error", nil).Detail, "error")

	assert.Equal(t, RetFail(100, "xxxx").Code, int32(100))
	assert.Equal(t, RetFail(100, "xxxx").Detail, "xxxx")

	result := errors.New(`{"code":100,"detail":"sm"}`)
	assert.Equal(t, RetError(result).Code, int32(100))
	assert.Equal(t, RetError(result).Detail, "sm")
}
