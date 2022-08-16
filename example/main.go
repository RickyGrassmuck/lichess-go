package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/RickyGrassmuck/lichess-go/lichess"
)


func main() {
	token, ok := os.LookupEnv("LICHESS_TOKEN")
	if !ok {
		fmt.Println("Please set LICHESS_TOKEN env variable with your personal access token...")
		os.Exit(1)
	}
	c := clientWithResponses(token)

	PrintMyAccountInfo(c)
	PrintUserInfo(c, "joyrida")
}
type transport struct {
	headers map[string]string
	base    http.RoundTripper
}

func (t *transport) RoundTrip(req *http.Request) (*http.Response, error) {
	for k, v := range t.headers {
		req.Header.Add(k, v)
	}
	base := t.base
	if base == nil {
		base = http.DefaultTransport
	}
	return base.RoundTrip(req)
}

func clientWithResponses(apiToken string) *lichess.ClientWithResponses {
	client := &http.Client{
		Transport: &transport{
			headers: map[string]string{
				"Authorization": fmt.Sprintf("Bearer %s", apiToken),
			},
		},
	}

	return &lichess.ClientWithResponses{
		ClientInterface: &lichess.Client{
			Server: "https://lichess.org",
			Client: client,
		},
	}
}

func PrintMyAccountInfo(c *lichess.ClientWithResponses) {
	me, err := c.AccountMeWithResponse(context.Background())
	if err != nil {
		fmt.Printf("%v\n", err)
		os.Exit(1)
	}
	meJSON, err := json.MarshalIndent(me.JSON200, "", "  ")
	fmt.Printf("Error: %v\nStatus: %s\n\nReponse: %s\n\n", err, me.Status(), meJSON)
}

func PrintUserInfo(c *lichess.ClientWithResponses, username string) {
	user, _ := c.ApiUserWithResponse(context.Background(), "joyrida")
	userJSON, err := json.MarshalIndent(user.JSON200, "", "  ")
	fmt.Printf("Error: %v\n\nReponse: %s\n\n", err, userJSON)
}

