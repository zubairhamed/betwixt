package main

import (
    . "github.com/zubairhamed/goap"
    "log"
    . "github.com/zubairhamed/lwm2m"
)

func main() {

    client := NewLwm2mClient(":0", "localhost:5683")

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
