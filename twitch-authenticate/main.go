package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"

	"golang.org/x/oauth2/clientcredentials"
	"golang.org/x/oauth2/twitch"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}

	oauth2Config := &clientcredentials.Config{
		ClientID:     os.Getenv("CLIENT_ID"),
		ClientSecret: os.Getenv("CLIENT_SECRET"),
		TokenURL:     twitch.Endpoint.TokenURL,
	}

	token, err := oauth2Config.Token(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Access token: %s\n", token.AccessToken)
	fmt.Printf("Refresh token: %s\n", token.RefreshToken)
	fmt.Printf("Token type: %s\n", token.TokenType)
	fmt.Printf("Expiry: %s\n", token.Expiry)
}