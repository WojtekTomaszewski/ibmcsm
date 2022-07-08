package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/WojtekTomaszewski/ibmctoken/ibmctoken"
	"net/http"
)

type httpclient interface {
	Do(req *http.Request) (*http.Response, error)
}

type secretsManager struct {
	endpoint string
	apikey   string
	client   httpclient
	token    *ibmctoken.Token
}

// NewSecretsManager creates new instance of SecretsManager for provided endpoint and access token.
func NewSecretsManager(endpoint string, apikey string) *secretsManager {
	return &secretsManager{
		endpoint: endpoint,
		client:   &http.Client{},
		token: &ibmctoken.Token{
			ApiKey: apikey,
		},
	}
}

// request is helper that makes HTTP request to SecretsManager instance.
func (sm *secretsManager) request(ctx context.Context, url string) (*http.Response, error) {

	if sm.token.Expired() {
		err := sm.token.RequestTokenWithContext(ctx)
		if err != nil {
			return nil, err
		}
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", sm.token.AccessToken))

	res, err := sm.client.Do(req)
	if err != nil {
		return nil, err
	}

	if res.StatusCode >= 400 {
		errResp := &secretError{}
		err = json.NewDecoder(res.Body).Decode(errResp)
		if err != nil {
			return nil, err
		}
		return nil, fmt.Errorf("reading secret failed %d %s", res.StatusCode, errResp.Errors[0].Message)
	}

	return res, nil
}

// ReadKeyValueSecret reads secrets of type Secrets from SecretsManager instance.
func (sm *secretsManager) ReadKeyValueSecret(id string) (*keyValueSecret, error) {
	return sm.ReadKeyValueSecretWithContext(context.Background(), id)
}

// ReadKeyValueSecretWithContext reads secrets of type Secrets from SecretsManager instance.
func (sm *secretsManager) ReadKeyValueSecretWithContext(ctx context.Context, id string) (*keyValueSecret, error) {
	urlPath := fmt.Sprintf("/api/v1/secrets/%s/%s", keyValueType, id)
	url := fmt.Sprintf("%s%s", sm.endpoint, urlPath)

	res, err := sm.request(ctx, url)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	secret := &keyValueSecret{}
	if err := json.NewDecoder(res.Body).Decode(secret); err != nil {
		fmt.Println(err)
		return nil, err
	}
	return secret, nil
}

// ReadUsernamePasswordSecret reads secrets of type username_password from SecretsManager instance.
func (sm *secretsManager) ReadUsernamePasswordSecret(id string) (*usernamePasswordSecret, error) {
	return sm.ReadUsernamePasswordSecretWithContext(context.Background(), id)
}

// ReadUsernamePasswordSecretWithContext reads secrets of type username_password from SecretsManager instance.
func (sm *secretsManager) ReadUsernamePasswordSecretWithContext(ctx context.Context, id string) (*usernamePasswordSecret, error) {
	urlPath := fmt.Sprintf("/api/v1/secrets/%s/%s", usernamePasswordType, id)
	url := fmt.Sprintf("%s%s", sm.endpoint, urlPath)

	res, err := sm.request(ctx, url)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	secret := &usernamePasswordSecret{}
	if err := json.NewDecoder(res.Body).Decode(secret); err != nil {
		fmt.Println(err)
		return nil, err
	}
	return secret, nil
}

// ReadArbitrarySecret reads secrets of type username_password from SecretsManager instance.
func (sm *secretsManager) ReadArbitrarySecret(id string) (*arbitrarySecret, error) {
	return sm.ReadArbitrarySecretWithContext(context.Background(), id)
}

// ReadArbitrarySecretWithContext reads secrets of type username_password from SecretsManager instance.
func (sm *secretsManager) ReadArbitrarySecretWithContext(ctx context.Context, id string) (*arbitrarySecret, error) {
	urlPath := fmt.Sprintf("/api/v1/secrets/%s/%s", arbitraryType, id)
	url := fmt.Sprintf("%s%s", sm.endpoint, urlPath)

	res, err := sm.request(ctx, url)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	secret := &arbitrarySecret{}
	if err := json.NewDecoder(res.Body).Decode(secret); err != nil {
		fmt.Println(err)
		return nil, err
	}
	return secret, nil
}
