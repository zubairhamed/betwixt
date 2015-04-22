package main

import (
    . "github.com/zubairhamed/goap"
    "log"
    . "github.com/zubairhamed/lwm2m"
)

func main() {

    repo := NewModelRepository()

    client := NewLwm2mClient(":0", "localhost:5683")

    client.AddResource(NewLWM2MResource(repo.GetModel(0), 0, 1, 2))
    client.AddResource(NewLWM2MResource(repo.GetModel(1), 1))
    client.AddResource(NewLWM2MResource(repo.GetModel(2), 0, 1, 2, 3, 4))
    client.AddResource(NewLWM2MResource(repo.GetModel(3), 0))
    client.AddResource(NewLWM2MResource(repo.GetModel(4), 0))
    client.AddResource(NewLWM2MResource(repo.GetModel(5), 0))
    client.AddResource(NewLWM2MResource(repo.GetModel(6)))
    client.AddResource(NewLWM2MResource(repo.GetModel(7)))

    client.OnStartup(func(evt *Event){
        client.Register("GOCLIENT")
    })

    client.OnRead(func(evt *Event){

    })

    client.OnRegistered(func(evt *Event){
        log.Println("Client is registered")
    })

    client.OnUnregistered(func(evt *Event){
        log.Println("Client is Unregistered")
    })


    client.OnExecute(func(evt *Event){

    })

    client.OnWrite(func(evt *Event){

    })

    client.Start()
}
