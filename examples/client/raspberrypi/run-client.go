package main

import (
	. "github.com/zubairhamed/betwixt"
)

func main() {
	cli := StandardCommandLineFlags()

	registry := NewDefaultObjectRegistry()
	c := NewDefaultClient(":0", cli.Server, registry)

	setupResources(c, registry)

	c.OnStartup(func() {
		c.Register(cli.Name)

		// TODO: Randomly fire change events for values changed
	})

	c.Start()
}

func setupResources(client LWM2MClient, reg Registry) {

}
