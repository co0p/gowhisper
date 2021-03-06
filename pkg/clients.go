package gowhisper

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
)

type Client struct {
	Label  string
	URL    string
	Online bool
}

func ReadClients(in io.Reader) ([]Client, error) {
	var clients []Client
	err := json.NewDecoder(in).Decode(&clients)

	for k, v := range clients {
		if err := validateClientEntry(k, v); err != nil {
			return nil, err
		}
	}

	return clients, err
}

func validateClientEntry(idx int, client Client) error {

	if len(client.Label) < 1 {
		return errors.New(fmt.Sprintf("client entry #%d is missing Label", idx))
	}
	if len(client.URL) < 1 {
		return errors.New(fmt.Sprintf("client entry #%d is missing URL", idx))
	}

	return nil
}
