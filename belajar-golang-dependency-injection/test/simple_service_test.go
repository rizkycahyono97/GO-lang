package test

import (
	"github.com/stretchr/testify/assert"
	"programmerzamannow/belajar-golang-restful-api/simple"
	"testing"
)

func TestSimpleServiceError(t *testing.T) {
	simpleService, err := simple.InitializeService(true)
	assert.NotNil(t, err)
	assert.Nil(t, simpleService)
}

func TestSimpleService(t *testing.T) {
	simpleService, err := simple.InitializeService(false)
	assert.Nil(t, err)
	assert.NotNil(t, simpleService)
}
