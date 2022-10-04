package duck

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
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
	params := url.Values{"user": {user}}

	resp, err := http.Get(loginLinkUrl + "?" + params.Encode())
	if err != nil {
		return fmt.Errorf("getLoginLink failed: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return errors.New("getLoginLink returned " + resp.Status)
	}

	return nil
}

// GetLogin takes the passphrase from `getLoginLink`, and retrieves an one-time login token
func GetLogin(user string, passphrase string) (LoginResponse, error) {
	params := url.Values{
		"user": {user},
		"otp":  {passphrase},
	}
	httpResp, err := http.Get(loginUrl + "?" + params.Encode())
	if err != nil {
		return LoginResponse{}, fmt.Errorf("getLogin failed: %w", err)
	}
	defer httpResp.Body.Close()

	if httpResp.StatusCode != 200 {
		return LoginResponse{}, errors.New("getLogin returned " + httpResp.Status)
	}

	body, _ := io.ReadAll(httpResp.Body)

	resp := LoginResponse{}
	err = json.Unmarshal(body, &resp)
	if err != nil {
		return resp, fmt.Errorf("getLogin could not unmarshal response: %s, error: %w", body, err)
	}
	return resp, nil
}

// GetDashboard uses the one-time token from `getLogin` to retrieve a long-lasting access token
func GetDashboard(otpToken string) (DashboardResponse, error) {
	req, _ := http.NewRequest(http.MethodGet, dashboardUrl, nil)
	req.Header.Add("authorization", "Bearer "+otpToken)

	httpResp, err := http.DefaultClient.Do(req)
	if err != nil {
		return DashboardResponse{}, err
	}
	defer httpResp.Body.Close()

	if httpResp.StatusCode != 200 {
		return DashboardResponse{}, errors.New("getDashboard returned " + httpResp.Status)
	}

	body, _ := io.ReadAll(httpResp.Body)
	resp := DashboardResponse{}
	err = json.Unmarshal(body, &resp)
	if err != nil {
		return resp, fmt.Errorf("getDashboard failed to unrmashal response: %s, error: %w", body, err)
	}
	return resp, nil
}
