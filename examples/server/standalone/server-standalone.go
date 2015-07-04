package main

import (
	"github.com/zubairhamed/betwixt"
	. "github.com/zubairhamed/betwixt/examples/server"
)

func main() {
	s := NewDefaultServer(":8081")

	registry := betwixt.NewDefaultObjectRegistry()

	s.UseRegistry(registry)

	s.Start()
}
