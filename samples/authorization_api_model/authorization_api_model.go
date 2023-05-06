package main

import (
	"context"
	"fmt"
	"net/http"
	"os"

	authorization "github.com/skeleton1231/amz-sp-api-gosdk/authorization-api-model"
)

func main() {
	// Set up the configuration for the Selling Partner API client
	cfg := authorization.NewConfiguration()
	cfg.BasePath = "https://sellingpartnerapi-na.amazon.com/"

	// Retrieve your LWA refresh token from a secure location
	refreshToken := "your-refresh-token-here"

	// Create a new client instance with the configured client and context
	apiClient := authorization.NewAPIClient(cfg)
	authApi := apiClient.AuthorizationApi

	// Set up the context with the LWA refresh token as a header
	ctx := context.WithValue(context.Background(), authorization.ContextAPIKey, &authorization.APIKey{Key: refreshToken})

	// Specify the parameters for the GetAuthorizationCode method

	SellingPartnerId := "your-selling-partner-id"
	DeveloperId := "your-developer-id"
	MwsAuthToken := "your-mws-auth-token"

	// Call the GetAuthorizationCode method on the AuthorizationApi object
	auth, resp, err := authApi.GetAuthorizationCode(ctx, SellingPartnerId, DeveloperId, MwsAuthToken)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error getting authorization code: %v\n", err)
		os.Exit(1)
	}

	// Access the response status code
	statusCode := resp.StatusCode

	// Check the status code and take action based on the result
	if statusCode == http.StatusOK {
		// Success! Do something with the authorization code
		authCode := auth.Payload.AuthorizationCode
		fmt.Printf("Authorization code: %s\n", authCode)
	} else if statusCode == http.StatusUnauthorized {
		// Authorization failed - handle the error
		fmt.Fprintf(os.Stderr, "Error getting authorization code: %v\n", err)
		os.Exit(1)
	} else {
		// Something else went wrong - handle the error
		fmt.Fprintf(os.Stderr, "Error getting authorization code: %v\n", err)
		os.Exit(1)
	}
}
