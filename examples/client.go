package main

import (
    . "github.com/zubairhamed/goap"
    "log"
    . "github.com/zubairhamed/lwm2m"
)

func main() {
    client := NewLwm2mClient(":0", "localhost:5683")


    // LWM2MRequest
        // Model
        // Resource

    // LWM2MResponse

    /*
        ResourceDefs

        ResourceInstances

        LWM2MResource myResource =

        client.OnRead(Resource, ResourceInstance) {

        }

        objSecurity := NewLwm2mObject(0)
        objSecurity.addResource
    */



    setupResources(client)


    client.OnStartup(func(evt *Event){
        client.Register("GOCLIENT")
    })

    client.OnRead(func(evt *Event, m *ObjectInstance, i *ResourceInstance) (*LWM2MResponse) {
        log.Println(evt.Data["objectModel"].(*ObjectModel))
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


func setupResources (c *Client) {

    instance1 ;=  newObjectInstance(object_lwm2m_security)
    instance.Set(0, val)

    c.AddObject(OBJECT_LWM2M_SECURITY, instance1, instance2, instance3)


/*
    client
        ObjectModel
            ObjectInstances
                Resources



    client.AddResource(0, 0, 1, 2)
    client.AddResource(1, 1)
    client.AddResource(2, 0, 1, 2, 3, 4)
    client.AddResource(3, 0)
    client.AddResource(4, 0)
    client.AddResource(5, 0)
    client.AddResource(6)
    client.AddResource(7)

    client.AddObject(0)
    client.AddResourceInstance(0, 0)
    client.AddResourceInstance(0, 1)
    client.AddResourceInstance(0, 2)
*/
}