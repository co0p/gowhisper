package main

import (
	"log"
	"os"

	"github.com/co0p/gowhisper"
)

func main() {
	flags, err := gowhisper.ParseFlags(os.Args[1:])
	if err != nil {
		log.Fatalf("failed startup: %s", err)
	}
	log.Printf("%v", flags)
}
