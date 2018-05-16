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
	var json = `[{"Label":"%s", "URL": "%s"}]`
	expLabel := "LAbel"
	expURL := "a url"
	input := fmt.Sprintf(json, expLabel, expURL)

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
}

func Test_ReadClientsShouldReturnErrorOnLabelMissing(t *testing.T) {
	var json = `[{"Label":"", "URL": "string"}]`

	in := strings.NewReader(json)
	_, err := gowhisper.ReadClients(in)

	if err == nil {
		t.Errorf("expected err not to be nil")
	}

	hint := "Label"
	if !strings.Contains(err.Error(), "") {
		t.Errorf("expected error message to contain hint about '%s', got '%s'", hint, err.Error())
	}
}

func Test_ReadClientsShouldReturnErrorOnURLMissing(t *testing.T) {
	var json = `[{"Label":"string", "URL": "", "EmailAddress": "string"}]`

	in := strings.NewReader(json)
	_, err := gowhisper.ReadClients(in)

	if err == nil {
		t.Errorf("expected err not to be nil")
	}

	hint := "URL"
	if !strings.Contains(err.Error(), "") {
		t.Errorf("expected error message to contain hint about '%s', got '%s'", hint, err.Error())
	}
}
