package main

import (
	"github.com/zubairhamed/betwixt/registry"
	"github.com/zubairhamed/betwixt/server"
)

func main() {
	s := server.NewDefaultServer()

	registry := registry.NewDefaultObjectRegistry()

	s.UseRegistry(registry)

	s.Start()
}
