package api_auth

import (
	"encoding/json"
	"net/http"
	"io/ioutil"
)

func (auth *AuthAPI) ChallengeGet(credential *Credential) (resp *http.Response, body *ChallengeGetResponse, err error) {
	resp, err = auth.post("/challenge", map[string]interface{}{
		"nuid/credential": credential,
	})
	defer resp.Body.Close()
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	err = json.Unmarshal(bodyBytes, body)
	if err != nil {
		return resp, nil, err
	}
	return resp, body, nil
}

func (auth *AuthAPI) ChallengeVerify(challengeJWT string, proof map[string]interface{}) (resp *http.Response, err error) {
	return auth.post("/challenge/verify", map[string]interface{}{
		"nuid.credential.challenge/jwt": challengeJWT,
		"nuid.credential/proof": proof,
	})
}

func (auth *AuthAPI) CredentialCreate(verifiedCredential *VerifiedCredential) (resp *http.Response, body *CredentialCreateResponse, err error) {
	resp, err = auth.post("/credential", map[string]interface{}{
		"nuid.credential/verified": verifiedCredential,
	})
	defer resp.Body.Close()
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	err = json.Unmarshal(bodyBytes, body)
	if err != nil {
		return resp, nil, err
	}
	return resp, body, nil
}

func (auth *AuthAPI) CredentialGet(nuid string) (resp *http.Response, body *CredentialGetResponse, err error) {
	resp, err = auth.get("/credential/" + nuid)
	if err != nil {
		return resp, nil, err
	}
	defer resp.Body.Close()
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	err = json.Unmarshal(bodyBytes, body)
	if err != nil {
		return resp, nil, err
	}
	return resp, body, nil
}
