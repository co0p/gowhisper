package main

import (
	"log"
	"net"
	"net/http"
	"os"
	"time"

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
	log.Printf("loaded %d clients to check ...", len(clients))

	client := newHttpClient()
	notifier := gowhisper.MailNotifier{ApiURL: flags.NotifyURL, Client: client}
	notifier.Send(gowhisper.Message{})
}

func newHttpClient() *http.Client {
	transport := &http.Transport{
		Dial: (&net.Dialer{
			Timeout: 5 * time.Second,
		}).Dial,
		TLSHandshakeTimeout: 5 * time.Second,
	}
	return &http.Client{
		Timeout:   time.Second * 10,
		Transport: transport,
	}
}
