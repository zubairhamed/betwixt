package tests

import (
	"github.com/zubairhamed/betwixt"
)

func NewMockServer() betwixt.Server {
	return &MockServer{
		stats: &MockServerStatistics{},
	}
}

type MockServer struct {
	stats betwixt.ServerStatistics
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
