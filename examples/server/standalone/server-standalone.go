package main

import (
	. "github.com/zubairhamed/betwixt/examples/server"
	"github.com/zubairhamed/betwixt"
)

func main() {
	s := NewDefaultServer(":8081")

	registry := betwixt.NewDefaultObjectRegistry()

	s.UseRegistry(registry)

	s.Start()
}


