package main

import (
	"bufio"
	"fmt"
	"os"
	"path"
	"strings"

	"github.com/chowder/duck-gen/duck"
)

func main() {
	token, err := readToken()
	if err != nil {
		token, err = getToken()
		if err != nil {
			fmt.Println("Could not get Duck token: ", err)
			os.Exit(1)
		}
	}

	address, err := duck.GenerateAddress(token)
	if err != nil {
		fmt.Println("Could not generate Private Duck Address: ", err)
		os.Exit(1)
	}

	fmt.Println(address)

	err = saveToken(token)
	if err != nil {
		fmt.Println("Could not save token: ", err)
	}
}

func getToken() (string, error) {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("Enter your Duck Address: ")
	if !scanner.Scan() {
		return "", fmt.Errorf("could not get duck address")
	}
	username := strings.TrimSuffix(strings.TrimSpace(scanner.Text()), "@duck.com")

	err := duck.GetLoginLink(username)
	if err != nil {
		return "", fmt.Errorf("could not trigger OTP login link: %w", err)
	}

	fmt.Print("Enter the one-time passphrase sent to your email: ")
	if !scanner.Scan() {
		return "", fmt.Errorf("could not get one-time passphrase")
	}
	otp := strings.TrimSpace(scanner.Text())

	loginResponse, err := duck.GetLogin(username, otp)
	if err != nil {
		return "", fmt.Errorf("could not login: %w", err)
	}

	dashboardResponse, err := duck.GetDashboard(loginResponse.Token)
	if err != nil {
		return "", fmt.Errorf("could not get access token: %w", err)
	}

	return dashboardResponse.User.AccessToken, nil
}

func getTokenFile() (string, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	return path.Join(home, ".duck_token"), nil
}

func readToken() (string, error) {
	tokenFile, err := getTokenFile()
	if err != nil {
		return "", err
	}

	token, err := os.ReadFile(tokenFile)
	if err != nil {
		return "", err
	}

	return strings.TrimSpace(string(token)), nil
}

func saveToken(token string) error {
	tokenFile, err := getTokenFile()
	if err != nil {
		return err
	}

	return os.WriteFile(tokenFile, []byte(token), 0600)
}
