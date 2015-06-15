package server

import (
	"github.com/stretchr/testify/assert"
	"github.com/zubairhamed/betwixt"
	"github.com/zubairhamed/go-commons/network"
	"testing"
)

func TestCoapRequests(t *testing.T) {
	server := &DefaultServer{
		coapServer: NewDefaultCoapServer(),
		httpServer: network.NewDefaultHttpServer(":8081"),
		clients:    make(map[string]betwixt.RegisteredClient),
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
