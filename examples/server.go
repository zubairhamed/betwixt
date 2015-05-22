package main
import (
    "github.com/zubairhamed/go-lwm2m"
    "log"
)

func main() {
    server := lwm2m.NewLWM2MServer()

    log.Println(server.GetModel(0))
    log.Println(server.GetModel(1))
    log.Println(server.GetModel(2))
    log.Println(server.GetModel(3))
    log.Println(server.GetModel(4))
}
