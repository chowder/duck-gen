package duck

import (
	"errors"
	"fmt"
)

const generateAddressUrl = "https://quack.duckduckgo.com/api/email/addresses"

// GenerateAddress generates a new @duck.com email address
func GenerateAddress(token string) (string, error) {
	client := GetClient()

	type Response struct {
		Address string `json:"address"`
		Error   string `json:"error"`
	}

	var response = Response{}

	r, err := client.R().
		SetHeader("authorization", "Bearer "+token).
		SetResult(&response).
		Post(generateAddressUrl)

	if err != nil {
		return "", fmt.Errorf("generateAddress failed to unmarshal response: %s, error: %w", r.String(), err)
	}

	if len(response.Error) > 0 {
		return "", errors.New("generateAddress failed to generate email: " + response.Error)
	}

	return response.Address + "@duck.com", nil
}
