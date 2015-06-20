package server

import (
	"github.com/stretchr/testify/assert"
	"github.com/zubairhamed/betwixt/tests"
	"testing"
)

func TestHandleHttpHome(t *testing.T) {
	server := tests.NewMockServer()

	fn := handleHttpHome(server)

	assert.NotNil(t, fn)
	response := fn(nil)

	assert.NotNil(t, response)
}

func TestHandleHttpDelete(t *testing.T) {
	server := tests.NewMockServer()

	fn := handleHttpDeleteClient(server)

	assert.NotNil(t, fn)
	response := fn(nil)

	assert.NotNil(t, response)
}

func TestHandleHttpViewClient(t *testing.T) {
	server := tests.NewMockServer()

	fn := handleHttpViewClient(server)

	assert.NotNil(t, fn)
	response := fn(nil)

	assert.NotNil(t, response)
}
