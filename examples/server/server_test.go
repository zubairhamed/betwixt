package server

import (
	"github.com/stretchr/testify/assert"
	. "github.com/zubairhamed/betwixt"
	"github.com/zubairhamed/go-commons/network"
	"testing"
)

func TestCoapRequests(t *testing.T) {
	server := &DefaultServer{
		coapServer: NewDefaultCoapServer(),
		httpServer: network.NewDefaultHttpServer(":8081"),
		clients:    make(map[string]RegisteredClient),
		stats:      &DefaultServerStatistics{},
	}

	SetupCoapRoutes(server)

	assert.Equal(t, 0, len(server.clients))
	server.register("betwixt-1", "127.0.0.1", nil)
	assert.Equal(t, 1, len(server.clients))
	server.register("betwixt-2", "127.0.0.1", nil)
	assert.Equal(t, 2, len(server.clients))
}

func TestHttpRequests(t *testing.T) {

}

func TestHandleHttpHome(t *testing.T) {
	server := NewMockServer()

	fn := handleHttpHome(server)

	assert.NotNil(t, fn)
	response := fn(nil)

	assert.NotNil(t, response)
}

//func TestHandleHttpDelete(t *testing.T) {
//	server := NewMockServer()
//
//	fn := handleHttpDeleteClient(server)
//
//	assert.NotNil(t, fn)
//	response := fn(nil)
//
//	assert.NotNil(t, response)
//}

func TestHandleHttpViewClient(t *testing.T) {
	server := NewMockServer()

	fn := handleHttpViewClient(server)

	assert.NotNil(t, fn)
	response := fn(nil)

	assert.NotNil(t, response)
}

func TestSetupRoutes(t *testing.T) {
	server := NewMockServer()

	SetupHttpRoutes(server)
}
