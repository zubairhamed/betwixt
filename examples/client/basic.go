package main

import (
	"github.com/zubairhamed/golwm2m"
    "github.com/zubairhamed/goap"
    "log"
    "time"
)

func main() {
	client := golwm2m.NewClient()

	client.Dial("udp", "localhost", 5683)

    client.Register("golwm2ma", func (msg *goap.Message){
        log.Println(goap.CoapCodeToString(msg.Code))
        log.Println(msg.GetOption(goap.OPTION_LOCATION_PATH))

        log.Println(msg.GetPath())

        time.Sleep(5 * time.Second)

        log.Println("Deregistering..")
        // client.Deregister()
    })
}

