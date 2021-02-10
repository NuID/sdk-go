package auth

import (
	"encoding/json"
	"os"
	"os/exec"
	"testing"
)

var api *AuthAPI

func TestRoundTrip(t *testing.T) {
	secret := "test secret"
	api = &AuthAPI{
		ApiKey: os.Getenv("NUID_AUTH_API_KEY"),
		Host: os.Getenv("NUID_AUTH_API_HOST"),
	}

	verifiedCredential, err := zk("verifiableFromSecret", secret)
	if err != nil {
		t.Fatalf("Failed to get verifiableFromSecret(%v) => %v", secret, err)
	}

	resp, credentialCreateBody, err := api.CredentialCreate(verifiedCredential)
	if err != nil {
		t.Fatalf("CredentialCreate() failed\nstatus=%d, body=%v, err=%v", resp.StatusCode, credentialCreateBody, err)
	}

	resp, credentialGetBody, err := api.CredentialGet(credentialCreateBody.NuID)
	if err != nil {
		t.Fatalf("CredentialGet() failed\nstatus=%d, nuid=%v, body=%v, err=%v", resp.StatusCode, credentialCreateBody.NuID, credentialGetBody, err)
	}

	resp, challengeGetBody, err := api.ChallengeGet(credentialGetBody.Credential)
	if err != nil {
		t.Fatalf("ChallengeGet()\nstatus=%d, credential=%v, body=%v, err=%v", resp.StatusCode, credentialGetBody.Credential, challengeGetBody, err)
	}

	claims, err := challengeGetBody.ChallengeJWT.Claims()
	if err != nil {
		t.Fatalf("Failed to decode JWT claims: %v", err)
	}
	proof, err := zk("proofFromSecretAndChallenge", secret, claims)
	if err != nil {
		t.Fatalf("Failed to get proofFromSecretAndChallenge(%v, %v) => %v", secret, challengeGetBody.ChallengeJWT, err)
	}

	resp, err = api.ChallengeVerify(challengeGetBody.ChallengeJWT, proof)
	if err != nil {
		t.Fatalf("ChallengeVerify()\nstatus=%d, challengeJWT=%v, proof=%v, err=%v", resp.StatusCode, challengeGetBody.ChallengeJWT, proof, err)
	}
}

func zk(zkCommand string, args ...interface{}) (data map[string]interface{}, err error) {
	cmdArgs, err := json.Marshal(args)
	if err != nil {
		return nil, err
	}

	out, err := exec.Command("nuid-cli", "zk", zkCommand, string(cmdArgs[:])).Output()
	if err != nil {
		return
	}

	err = json.Unmarshal(out, &data)
	return
}
