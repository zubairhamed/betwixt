package examples

import "flag"

type CliFlags struct {
	Name   string
	Server string
}

func StandardCommandLineFlags() *CliFlags {
	var name = flag.String("name", "betwixt", "Name for Node")
	var server = flag.String("server", "localhost:5683", "LWM2M Server")
	// var server = flag.String("server", "5.39.83.206:5683", "LWM2M Server")

	flag.Parse()

	return &CliFlags{
		Name:   *name,
		Server: *server,
	}
}
