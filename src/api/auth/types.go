package api_auth

type Credential interface{}

type AuthAPI struct {
	apiKey string
	url    string
}

type VerifiedCredential interface{}

type ChallengeGetResponse struct {
	challengeJWT string `json:"nuid.credential.challenge/jwt"`
}

type CredentialCreateResponse struct {
	nuid string            `json:"nu/id"`
	credential *Credential `json:"nuid/credential"`
}

type CredentialGetResponse struct {
	credential *Credential `json:"nuid/credential"`
}
