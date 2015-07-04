package server

import (
	"github.com/zubairhamed/betwixt"
	"log"
	"os"
	"runtime/pprof"
	"strconv"
	"testing"
)

func BenchmarkServer(b *testing.B) {
	f, _ := os.Create("profiler.prof")
	pprof.StartCPUProfile(f)
	server := NewDefaultServer(":8181")
	reg := betwixt.NewDefaultObjectRegistry()
	server.UseRegistry(reg)

	cli := betwixt.NewDefaultClient(":0", "localhost:5683", reg)
	cli.OnStartup(func() {
		for i := 1; i <= 5000; i++ {
			name := "bet" + strconv.Itoa(i)
			log.Println("Calling ", name)
			cli.Register(name)
		}
		log.Println("Done!")
		pprof.StopCPUProfile()
		os.Exit(0)
	})

	server.On(betwixt.EVENT_START, func() {
		log.Println("EVent atart")
		cli.Start()
	})

	server.Start()
}
