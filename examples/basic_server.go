package main
import "github.com/zubairhamed/go-lwm2m/server"

func main() {
    s := server.NewDefaultServer()

    s.Start()
}
