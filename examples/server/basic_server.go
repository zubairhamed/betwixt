package main

import (
	"github.com/zubairhamed/betwixt/core/registry"
	"github.com/zubairhamed/betwixt/server"
)

func main() {
	s := server.NewDefaultServer(":8081")

	registry := registry.NewDefaultObjectRegistry()

	s.UseRegistry(registry)

	s.Start()
}
