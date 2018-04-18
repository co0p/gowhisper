package gowhisper

import (
	"log"
	"net/http"
	"sync"
	"time"
)

type Checker struct {
	Clients         *[]Client
	HTTPClient      *http.Client
	PollingInterval int
	mu              *sync.Mutex
}

func (c *Checker) StartPolling() {

	c.mu = &sync.Mutex{}
	pollingInterval := time.Duration(c.PollingInterval) * time.Second
	tick := time.Tick(pollingInterval)
	for {
		select {
		case <-tick:
			for i, _ := range *c.Clients {
				i := i // W000t
				go checkClient(c.mu, c.HTTPClient, &(*c.Clients)[i])
			}
		}
	}
}

func checkClient(mu *sync.Mutex, webClient *http.Client, service *Client) {
	resp, err := webClient.Get(service.URL)
	if err != nil {
		log.Printf("failed fetching '%s'", service.URL)
		return
	}
	online := false

	if resp.StatusCode < 200 || resp.StatusCode > 399 {
		log.Printf("'%s' is down!", service.URL)
		online = false
	} else {
		log.Printf("'%s' is up!", service.URL)
		online = true
	}

	mu.Lock()
	service.Online = online
	mu.Unlock()
}
