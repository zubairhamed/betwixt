package main

import (
	"github.com/zubairhamed/betwixt/core/registry"
	. "github.com/zubairhamed/betwixt/examples/server"
)

func main() {
	s := NewDefaultServer(":8081")

	registry := registry.NewDefaultObjectRegistry()

	s.UseRegistry(registry)

	s.Start()
}


