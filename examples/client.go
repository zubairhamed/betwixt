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

        objSecurity1 := NewLwm2mObject(0, 0)
        objSecurity.putResource(0, value)

        client.AddObject(objSecurity)
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


    client.OnExecute(func(evt *Event, m *ObjectInstance, i *ResourceInstance) (*LWM2MResponse) {

    })

    client.OnWrite(func(evt *Event) (*LWM2MResponse) {

    })

    client.Start()
}


func setupResources (c *Client) {

    c.AddObjects(OBJECT_LWM2M_SECURITY,
        NewObjectInstance(OBJECT_LWM2M_SECURITY, 0),
        NewObjectInstance(OBJECT_LWM2M_SECURITY, 1),
        NewObjectInstance(OBJECT_LWM2M_SECURITY, 2),
    )

    c.AddObjects(OBJECT_LWM2M_SERVER,
        NewObjectInstance(OBJECT_LWM2M_SERVER, 1),
    )

    c.AddObjects(OBJECT_LWM2M_ACCESS_CONTROL,
        NewObjectInstance(OBJECT_LWM2M_ACCESS_CONTROL, 0),
        NewObjectInstance(OBJECT_LWM2M_ACCESS_CONTROL, 1),
        NewObjectInstance(OBJECT_LWM2M_ACCESS_CONTROL, 2),
        NewObjectInstance(OBJECT_LWM2M_ACCESS_CONTROL, 3),
    )

    c.AddObjects(OBJECT_LWM2M_DEVICE,
        NewObjectInstance(OBJECT_LWM2M_DEVICE, 0),
    )

    c.AddObjects(OBJECT_LWM2M_CONNECTIVITY_MONITORING,
        NewObjectInstance(OBJECT_LWM2M_CONNECTIVITY_MONITORING, 0),
    )

    c.AddObjects(OBJECT_LWM2M_FIRMWARE_UPDATE,
        NewObjectInstance(OBJECT_LWM2M_FIRMWARE_UPDATE, 0),
    )

    c.AddObjects(OBJECT_LWM2M_LOCATION)
    c.AddObjects(OBJECT_LWM2M_CONNECTIVITY_STATISTICS)
/*
    client
        ObjectModel
            ObjectInstances
                Resources
*/
}