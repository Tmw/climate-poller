package client

import (
	"encoding/json"
	"fmt"

	"github.com/salex-org/ikea-dirigera-client/pkg/client"
)

type (
	Hub           = client.DirigeraHub
	Client        = client.Client
	Authorization = client.Authorization
)

func AuthFromString(token string) (Authorization, error) {
	var auth Authorization
	err := json.Unmarshal([]byte(token), &auth)
	return auth, err
}

func StringFromAuth(auth Authorization) (string, error) {
	out, err := json.Marshal(auth)
	return string(out), err
}

func GetAccessToken(hub Hub, clientName string) (*Authorization, error) {
	auth, err := client.Authorize(hub.Address, hub.Port, clientName, func() {
		fmt.Printf("Authorizing client: Press the button on the hub within 1 minute...")
	}, func() {
		fmt.Printf(".")
	})

	if err != nil {
		return nil, fmt.Errorf("error authorizing: %w", err)
	}

	return &auth, nil
}
