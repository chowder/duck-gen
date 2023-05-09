package duck

import (
	"errors"
	"fmt"
)

const (
	loginLinkUrl = "https://quack.duckduckgo.com/api/auth/loginlink"
	loginUrl     = "https://quack.duckduckgo.com/api/auth/login"
	dashboardUrl = "https://quack.duckduckgo.com/api/email/dashboard"
)

type LoginResponse struct {
	Status string `json:"status"`
	Token  string `json:"token"`
	User   string `json:"user"`
}

type DashboardResponse struct {
	User struct {
		AccessToken string `json:"access_token"`
	} `json:"user"`
}

// GetLoginLink sends the OTP link/passphrase to the user's email
func GetLoginLink(user string) error {
	client := GetClient()

	r, err := client.R().
		SetQueryParam("user", user).
		Get(loginLinkUrl)

	if err != nil {
		return fmt.Errorf("getLoginLink failed: %w", err)
	}

	if r.StatusCode() != 200 {
		return errors.New("getLoginLink returned " + r.Status())
	}

	return nil
}

// GetLogin takes the passphrase from `getLoginLink`, and retrieves a one-time login token
func GetLogin(user string, passphrase string) (response LoginResponse, err error) {
	client := GetClient()

	r, err := client.R().
		SetQueryParam("user", user).
		SetQueryParam("otp", passphrase).
		SetResult(&response).
		Get(loginUrl)

	if err != nil {
		return response, fmt.Errorf("getLogin failed: %w", err)
	}

	if r.StatusCode() != 200 {
		return response, errors.New("getLogin returned " + r.Status())
	}

	return response, nil
}

// GetDashboard uses the one-time token from `getLogin` to retrieve a long-lasting access token
func GetDashboard(otpToken string) (response DashboardResponse, err error) {
	client := GetClient()

	r, err := client.R().
		SetHeader("authorization", "Bearer "+otpToken).
		SetResult(&response).
		Get(dashboardUrl)

	if r.StatusCode() != 200 {
		return response, errors.New("getDashboard returned " + r.Status())
	}

	return response, nil
}
