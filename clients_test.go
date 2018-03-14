package gowhisper_test

import (
	"fmt"
	"strings"
	"testing"

	"github.com/co0p/gowhisper"
)

func Test_ReadClientsShouldThrowErrorOnInvalidJSON(t *testing.T) {
	invalid := "[{int a},{},{]"

	in := strings.NewReader(invalid)
	_, err := gowhisper.ReadClients(in)

	if err == nil {
		t.Errorf("expected err not to be nil, got %v", err)
	}
}

func Test_ReadClientsShouldReturnClientsFromJSON(t *testing.T) {
	var json = `[{"Label":"%s", "URL": "%s", "StatusCode": %d, "Notify": "%s"}]`
	expLabel := "LAbel"
	expURL := "a url"
	expStatusCode := 100
	expNotify := "notify me"
	input := fmt.Sprintf(json, expLabel, expURL, expStatusCode, expNotify)

	in := strings.NewReader(input)
	clients, err := gowhisper.ReadClients(in)

	if err != nil {
		t.Errorf("expected err to be nil, got %v", err)
	}

	client := clients[0]
	if client.Label != expLabel {
		t.Errorf("expected Label to be %s, got %s", expLabel, client.Label)
	}
	if client.URL != expURL {
		t.Errorf("expected URL to be %s, got %s", expURL, client.URL)
	}
	if client.StatusCode != expStatusCode {
		t.Errorf("expected StatusCode to be %d, got %d", expStatusCode, client.StatusCode)
	}
	if client.Notify != expNotify {
		t.Errorf("expected Notify to be %s, got %s", expNotify, client.Notify)
	}
}
