package main
import (
    "github.com/zubairhamed/go-lwm2m/server"
    "github.com/zubairhamed/go-lwm2m/registry"
)

func main() {
    s := server.NewDefaultServer()

    registry := registry.NewDefaultObjectRegistry()

    s.UseRegistry(registry)

    s.Start()
}
