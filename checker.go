package gowhisper

import (
	"log"
	"net/http"
	"time"
)

type Checker struct {
	Clients         *[]Client
	HTTPClient      *http.Client
	PollingInterval int
}

func (c *Checker) StartPolling() {

	pollingInterval := time.Duration(c.PollingInterval) * time.Second
	tick := time.Tick(pollingInterval)
	for {
		select {
		case <-tick:
			for i, _ := range *c.Clients {
				i := i // W000t
				go checkClient(c.HTTPClient, &(*c.Clients)[i])
			}
		}
	}
}

func checkClient(webClient *http.Client, service *Client) {
	resp, err := webClient.Get(service.URL)
	if err != nil {
		log.Printf("failed fetching '%s'", service.URL)
		return
	}

	if resp.StatusCode < 200 || resp.StatusCode > 399 {
		log.Printf("'%s' is down!", service.URL)
		service.Online = false
		return
	}

	// still online or recover case
	log.Printf("'%s' is up!", service.URL)
	service.Online = true
}
