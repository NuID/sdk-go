package auth

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"io/ioutil"
)

func (auth *AuthAPI) ChallengeGet(credential map[string]interface{}) (resp *http.Response, body *ChallengeGetResponse, err error) {
	resp, err = auth.post("/challenge", map[string]interface{}{
		"nuid/credential": credential,
	})
	if err != nil {
		return
	}
	if resp.StatusCode != 201 {
		err = errors.New(fmt.Sprintf("Could not get challenge for credential, endpoint returned %s", resp.Status))
		return
	}

	defer resp.Body.Close()
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}
	if !bytes.Equal(bodyBytes, nil) {
		body = &ChallengeGetResponse{}
		err = json.Unmarshal(bodyBytes, body)
	}
	return
}

func (auth *AuthAPI) ChallengeVerify(challengeJWT JWT, proof map[string]interface{}) (resp *http.Response, err error) {
	resp, err = auth.post("/challenge/verify", map[string]interface{}{
		"nuid.credential.challenge/jwt": challengeJWT,
		"nuid.credential/proof": proof,
	})
	if err == nil && resp.StatusCode != 200 {
		err = errors.New(fmt.Sprintf("Could not verify challenge, endpoint returned %s", resp.Status))
	}
	return
}

func (auth *AuthAPI) CredentialCreate(verifiedCredential map[string]interface{}) (resp *http.Response, body *CredentialCreateResponse, err error) {
	resp, err = auth.post("/credential", map[string]interface{}{
		"nuid.credential/verified": verifiedCredential,
	})
	if err != nil {
		return
	}
	if resp.StatusCode != 201 {
		err = errors.New(fmt.Sprintf("Could not create credential, endpoint returned %s", resp.Status))
		return
	}

	defer resp.Body.Close()
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}
	if !bytes.Equal(bodyBytes, nil) {
		body = &CredentialCreateResponse{}
		err = json.Unmarshal(bodyBytes, body)
	}
	return
}

func (auth *AuthAPI) CredentialGet(nuid string) (resp *http.Response, body *CredentialGetResponse, err error) {
	resp, err = auth.get("/credential/" + nuid)
	if err != nil {
		return
	}
	if resp.StatusCode != 200 {
		err = errors.New(fmt.Sprintf("Could not get credential with nuid %s, endpoint returned %s", nuid, resp.Status))
		return
	}

	defer resp.Body.Close()
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil || bytes.Equal(bodyBytes, nil) {
		return
	}
	if !bytes.Equal(bodyBytes, nil) {
		body = &CredentialGetResponse{}
		err = json.Unmarshal(bodyBytes, body)
	}
	return
}
