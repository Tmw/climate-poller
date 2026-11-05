package client

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"os"

	"github.com/salex-org/ikea-dirigera-client/pkg/client"
)

type (
	Hub           = client.DirigeraHub
	Client        = client.Client
	Authorization = client.Authorization
)

func AuthFromString(token string) (Authorization, error) {
	var auth Authorization
	decoded, err := base64.StdEncoding.DecodeString(token)
	if err != nil {
		return auth, fmt.Errorf("failed to decode access-token: %w", err)
	}

	err = json.Unmarshal(decoded, &auth)
	return auth, err
}

func StringFromAuth(auth Authorization) (string, error) {
	var result string
	out, err := json.Marshal(auth)
	if err != nil {
		return result, fmt.Errorf("error marshalling JSON: %w", err)
	}

	result = base64.StdEncoding.EncodeToString(out)

	return result, nil
}

func GetAccessToken(hub Hub, clientName string) (*Authorization, error) {
	auth, err := client.Authorize(hub.Address, hub.Port, clientName, func() {
		fmt.Fprintf(os.Stderr, "Authorizing client: Press the button on the hub within 1 minute...")
	}, func() {
		fmt.Fprintf(os.Stderr, ".")
	})

	fmt.Fprintln(os.Stderr, "\n")

	if err != nil {
		return nil, fmt.Errorf("error authorizing: %w", err)
	}

	return &auth, nil
}
