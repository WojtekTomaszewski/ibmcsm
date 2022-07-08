package ibmcsm

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/WojtekTomaszewski/ibmctoken"
	"net/http"
)

type httpclient interface {
	Do(req *http.Request) (*http.Response, error)
}

type SecretsManager struct {
	endpoint string
	apikey   string
	client   httpclient
}

// NewSecretsManager creates new instance of SecretsManager for provided endpoint and access token.
func NewSecretsManager(endpoint string, apikey string) *SecretsManager {

	return &SecretsManager{
		endpoint: endpoint,
		apikey:   apikey,
		client:   &http.Client{},
	}
}

// request is helper that makes HTTP request to SecretsManager instance.
func (sm *SecretsManager) request(ctx context.Context, url string) (*http.Response, error) {

	token := &ibmctoken.Token{
		ApiKey: sm.apikey,
		Client: sm.client,
	}

	if err := token.RequestToken(); err != nil {
		return nil, err
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", token.AccessToken))

	res, err := http.DefaultClient.Do(req)
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
func (sm *SecretsManager) ReadKeyValueSecret(id string) (*keyValueSecret, error) {
	return sm.ReadKeyValueSecretWithContext(context.Background(), id)
}

// ReadKeyValueSecretWithContext reads secrets of type Secrets from SecretsManager instance.
func (sm *SecretsManager) ReadKeyValueSecretWithContext(ctx context.Context, id string) (*keyValueSecret, error) {
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
func (sm *SecretsManager) ReadUsernamePasswordSecret(id string) (*usernamePasswordSecret, error) {
	return sm.ReadUsernamePasswordSecretWithContext(context.Background(), id)
}

// ReadUsernamePasswordSecretWithContext reads secrets of type username_password from SecretsManager instance.
func (sm *SecretsManager) ReadUsernamePasswordSecretWithContext(ctx context.Context, id string) (*usernamePasswordSecret, error) {
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
func (sm *SecretsManager) ReadArbitrarySecret(id string) (*arbitrarySecret, error) {
	return sm.ReadArbitrarySecretWithContext(context.Background(), id)
}

// ReadArbitrarySecretWithContext reads secrets of type username_password from SecretsManager instance.
func (sm *SecretsManager) ReadArbitrarySecretWithContext(ctx context.Context, id string) (*arbitrarySecret, error) {
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
