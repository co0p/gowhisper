package main

import (
	"log"
	"os"

	"github.com/co0p/gowhisper"
)

func main() {
	flags, err := gowhisper.ParseFlags(os.Args[1:])
	if err != nil {
		log.Fatalf("%s", err)
	}

	in, err := os.Open(flags.ConfigurationFile)
	if err != nil {
		log.Fatalf("%s", err)
	}

	clients, err := gowhisper.ReadClients(in)
	if err != nil {
		log.Fatalf("%s", err)
	}
	log.Printf("loaded %d clients to poll ...", len(clients))
}
