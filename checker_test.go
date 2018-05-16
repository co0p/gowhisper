package gowhisper_test

import (
	"net/http"
	"testing"

	"github.com/co0p/gowhisper"
)

func Test_checker_CheckService_ShouldSetServiceToOffline_WhenServiceIsOffline(t *testing.T) {

	client := gowhisper.Client{
		URL: "http://localhost/test",
	}
	clients := []gowhisper.Client{client}
	checker := givenChecker(&clients)

	checker.CheckClient(&client)
	// TODO
}

func Test_checker_CheckService_ShouldSetServiceToOnline_WhenServiceIsOnline(t *testing.T) {

}

func givenChecker(clients *[]gowhisper.Client) gowhisper.Checker {
	return gowhisper.Checker{
		Clients:         clients,
		HTTPClient:      newHttpClient(),
		PollingInterval: 10,
	}
}

func newHttpClient() *http.Client {
	return http.DefaultClient
}
