package main

import (
	"fmt"
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
		log.Fatalf("failed to parse flags: %s", err)
	}

	in, err := os.Open(flags.ConfigurationFile)
	if err != nil {
		log.Fatalf("failed to read clients: %s", err)
	}

	clients, err := gowhisper.ReadClients(in)
	if err != nil {
		log.Fatalf("failed to parse clients: %s", err)
	}

	client := newHttpClient()
	notifier := gowhisper.MailNotifier{ApiURL: flags.NotifyURL, Client: client}
	notifier.Send(gowhisper.Message{})

	statusPage, err := gowhisper.NewStatusPage(&clients)
	if err != nil {
		log.Fatalf("failed to initialize status page: %s", err)
	}

	portStr := fmt.Sprintf(":%d", flags.Port)
	log.Printf("starting status page on port " + portStr)
	http.HandleFunc("/", statusPage.ServeHTTP)
	if err := http.ListenAndServe(portStr, nil); err != nil {
		log.Fatalf("failed to start statuspage listener: %s", err)
	}
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
