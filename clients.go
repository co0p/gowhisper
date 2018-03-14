package gowhisper

import (
	"encoding/json"
	"io"
)

type Client struct {
	Label      string
	URL        string
	StatusCode int
	Notify     string
}

func ReadClients(in io.Reader) ([]Client, error) {
	var clients []Client
	err := json.NewDecoder(in).Decode(&clients)
	return clients, err
}
