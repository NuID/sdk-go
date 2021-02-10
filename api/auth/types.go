package auth

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"errors"
)

// The AuthAPI struct contains the API Key and Host. Methods are callable from
// this type to interact with the API and parse responses.
type AuthAPI struct {
	// The ApiKey can be found in your free portal account at https://portal.nuid.io
	ApiKey string

	// The Host name of the auth API (you probably want https://auth.nuid.io)
	Host   string
}

type JWT string

// The response body for valid auth.ChallengeGet() responses.
type ChallengeGetResponse struct {
	// The Challenge JWT contains the challenge as its claims, to be
	// used when generating proofs in addition to the user's secret.
	ChallengeJWT JWT `json:"nuid.credential.challenge/jwt"`
}

// Get the claims map from a JWT
func (jwt JWT) Claims() (claims map[string]interface{}, err error) {
	parts := bytes.Split([]byte(jwt), []byte("."))
	encodedClaims := parts[1]
	if encodedClaims == nil {
		err = errors.New("Unable to decode JWT claims")
		return
	}
	claimsJSON, err := base64.RawStdEncoding.DecodeString(string(encodedClaims))
	if err != nil {
		return
	}
	err = json.Unmarshal(claimsJSON, &claims)
	if err != nil {
		return
	}
	return
}

// The response body for valid auth.CredentialCreate() responses.
type CredentialCreateResponse struct {
	// An encoded string unique to each credential. This value should be stored
	// alongside your user record during registration and referenced during
	// the login challenge and verify stages.
	NuID string                       `json:"nu/id"`

	// The user's credential. This value can be fetched from the API at any
	// time using auth.CredentialGet(nuid).
	Credential map[string]interface{} `json:"nuid/credential"`
}

// The response body for valid auth.CredentialGet() responses.
type CredentialGetResponse struct {
	// The user's credential.
	Credential map[string]interface{} `json:"nuid/credential"`
}
