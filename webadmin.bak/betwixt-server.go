package main

import (
	"github.com/zubairhamed/betwixt"
	. "github.com/zubairhamed/betwixt/webadmin/webadmin"
)

func main() {
	s := NewDefaultServer("8081")

	registry := betwixt.NewDefaultObjectRegistry()

	s.UseRegistry(registry)

	s.Start()
}
