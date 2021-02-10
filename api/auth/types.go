package auth

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"errors"
)

type AuthAPI struct {
	ApiKey string
	Host   string
}

type JWT string

type ChallengeGetResponse struct {
	ChallengeJWT JWT `json:"nuid.credential.challenge/jwt"`
}

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

type CredentialCreateResponse struct {
	NuID string                       `json:"nu/id"`
	Credential map[string]interface{} `json:"nuid/credential"`
}

type CredentialGetResponse struct {
	Credential map[string]interface{} `json:"nuid/credential"`
}
