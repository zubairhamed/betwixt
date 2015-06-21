package tests

import (
	"github.com/zubairhamed/betwixt"
	"github.com/zubairhamed/go-commons/network"
)

func NewMockServer() betwixt.Server {
	return &MockServer{
		stats: 		&MockServerStatistics{},
		httpServer: network.NewDefaultHttpServer("8080"),
	}
}

type MockServer struct {
	stats betwixt.ServerStatistics
	httpServer *network.HttpServer
}

func (server *MockServer) Start() {

}

func (server *MockServer) UseRegistry(reg betwixt.Registry) {

}

func (server *MockServer) On(e betwixt.EventType, fn betwixt.FnEvent) {

}

func (server *MockServer) GetClients() map[string]betwixt.RegisteredClient {
	return make(map[string]betwixt.RegisteredClient)
}

func (server *MockServer) GetStats() betwixt.ServerStatistics {
	return server.stats
}

func (server *MockServer) GetHttpServer() (*network.HttpServer) {
	return nil
}

func (server *MockServer) GetClient(id string) betwixt.RegisteredClient {
	return nil
}
