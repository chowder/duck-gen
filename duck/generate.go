package duck

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
)

// GenerateAddress generates a new @duck.com email address
func GenerateAddress(token string) (string, error) {
	const url = "https://quack.duckduckgo.com/api/email/addresses"

	req, _ := http.NewRequest(http.MethodPost, url, nil)
	req.Header.Add("authorization", "Bearer "+token)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", fmt.Errorf("generateAddress failed: %w", err)
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)

	type result struct {
		Address string `json:"address"`
		Error   string `json:"error"`
	}

	r := result{}
	err = json.Unmarshal(body, &r)
	if err != nil {
		return "", fmt.Errorf("generateAddress failed to unmarshal response: %s, error: %w", body, err)
	}

	if len(r.Error) > 0 {
		return "", errors.New("generateAddress failed to generate email: " + r.Error)
	}

	return r.Address + "@duck.com", nil
}
