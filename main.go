package main

import "github.com/NuID/sdk-go/api/auth"

// The NuID Go SDK provides access to simplified API interactions with NuID
// APIs. In the future we will also provide our crypto functions for generating
// and managing credentials. For now, please use the npm package `@nuid/zk` or
// `@nuid/cli` for generating crypto materials.

// Get a new Auth API Client struct to interact with the Auth API.
// See api/auth for more subpackage documentation.
func NewAuthAPIClient(apiKey string) *auth.APIClient {
	return auth.NewAPIClient(apiKey)
}

func main() {}
