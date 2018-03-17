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
	var json = `[{"Label":"%s", "URL": "%s", "StatusCode": %d, "EmailAddress": "%s"}]`
	expLabel := "LAbel"
	expURL := "a url"
	expStatusCode := 100
	expEmailAddress := "notify me"
	input := fmt.Sprintf(json, expLabel, expURL, expStatusCode, expEmailAddress)

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
	if client.EmailAddress != expEmailAddress {
		t.Errorf("expected EmailAddress to be %s, got %s", expEmailAddress, client.EmailAddress)
	}
}

func Test_ReadClientsShouldReturnErrorOnLabelMissing(t *testing.T) {
	var json = `[{"Label":"", "URL": "string", "StatusCode": 200, "EmailAddress": "string"}]`

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
	var json = `[{"Label":"string", "URL": "", "StatusCode": 200, "EmailAddress": "string"}]`

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

func Test_ReadClientsShouldReturnErrorOnInvalidStatusCode(t *testing.T) {
	var json = `[{"Label":"string", "URL": "string", "StatusCode": 0, "EmailAddress": "string"}]`

	in := strings.NewReader(json)
	_, err := gowhisper.ReadClients(in)

	if err == nil {
		t.Errorf("expected err not to be nil")
	}

	hint := "StatusCode"
	if !strings.Contains(err.Error(), "") {
		t.Errorf("expected error message to contain hint about '%s', got '%s'", hint, err.Error())
	}
}

func Test_ReadClientsShouldReturnErrorOnEmailAddressMissing(t *testing.T) {
	var json = `[{"Label":"string", "URL": "string", "StatusCode": 200, "EmailAddress": ""}]`

	in := strings.NewReader(json)
	_, err := gowhisper.ReadClients(in)

	if err == nil {
		t.Errorf("expected err not to be nil")
	}

	hint := "EmailAddress"
	if !strings.Contains(err.Error(), "") {
		t.Errorf("expected error message to contain hint about '%s', got '%s'", hint, err.Error())
	}
}
