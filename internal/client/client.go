package client

import (
	"fmt"

	"github.com/salex-org/ikea-dirigera-client/pkg/client"
)

func NewClient(hub *Hub, token string) (Client, error) {
	auth, err := AuthFromString(token)
	if err != nil {
		return nil, fmt.Errorf("error parsing auth token: %w", err)
	}

	return client.Connect(hub.Address, hub.Port, &auth), nil
}

// Get the first hub in the list.
func FirstHub() (*Hub, error) {
	hubs, err := client.Scan()
	if err != nil {
		return nil, fmt.Errorf("error scanning for hubs: %w", err)
	}

	if len(hubs) == 0 {
		return nil, fmt.Errorf("no hubs found")
	}

	hub := hubs[0]
	return &hub, nil
}
