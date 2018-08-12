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
			for clientIdx := range *c.Clients {
				client := &(*c.Clients)[clientIdx]
				go c.CheckClient(client)
			}
		}
	}
}

func (c *Checker) CheckClient(service *Client) {
	resp, err := c.HTTPClient.Get(service.URL)
	actualState := false
	if err != nil {
		log.Printf("failed fetching '%s'", service.URL)
	} else {
		actualState = resp.StatusCode > 199 && resp.StatusCode < 400
	}

	c.mu.Lock()
	service.Online = actualState
	log.Printf("'%s' is online=%v!", service.URL, service.Online)
	c.mu.Unlock()
}
