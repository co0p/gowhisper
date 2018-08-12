// +build !unit

package gowhisper_test

import (
	"net/http"
	"testing"
	"time"

	"github.com/co0p/gowhisper/pkg"
)

func Test_checker_CheckService_ShouldSetServiceToOffline_WhenServiceIsOffline(t *testing.T) {

	client := gowhisper.Client{
		URL:    "http://hopefully.notawebsite.de",
		Online: true,
	}
	clients := []gowhisper.Client{client}
	checker := givenChecker(&clients)
	go checker.StartPolling()

	time.Sleep(time.Second * 3)
	if clients[0].Online {
		t.Errorf("expected server %s to be down, got %v", clients[0].URL, clients[0].Online)
	}

}

func Test_checker_CheckService_ShouldSetServiceToOnline_WhenServiceIsOnline(t *testing.T) {
	client := gowhisper.Client{
		URL:    "http://www.google.de",
		Online: false,
	}
	clients := []gowhisper.Client{client}
	checker := givenChecker(&clients)
	go checker.StartPolling()

	time.Sleep(time.Second * 3)
	if !clients[0].Online {
		t.Errorf("expected server %s to be up, got %v", clients[0].URL, clients[0].Online)
	}
}

func givenChecker(clients *[]gowhisper.Client) gowhisper.Checker {
	return gowhisper.Checker{
		Clients:         clients,
		HTTPClient:      newHttpClient(),
		PollingInterval: 2,
	}
}

func newHttpClient() *http.Client {
	return http.DefaultClient
}
