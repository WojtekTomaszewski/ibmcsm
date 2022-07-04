package ibmcsm

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/WojtekTomaszewski/ibmctoken"
	"net/http"
)

type secretsManager struct {
	endpoint string
	apikey   string
	token    *ibmctoken.Token
}

// NewSecretsManager creates new instance of SecretsManager for provided endpoint and access token
func NewSecretsManager(endpoint string, apikey string) *secretsManager {
	return &secretsManager{
		endpoint: endpoint,
		token: &ibmctoken.Token{
			ApiKey: apikey,
		},
	}
}

// Read reads secrets of type Secrets from SecretsManager instance
func (sm *secretsManager) ReadSecret(ctx context.Context, secret *secret) error {
	uri := fmt.Sprintf("/api/v1/secrets/%s/%s", secret.Resources[0].Type(), secret.Resources[0].Id())
	fullURI := fmt.Sprintf("%s%s", sm.endpoint, uri)

	if sm.token.Expired() {
		err := sm.token.RequestToken(ctx)
		if err != nil {
			return err
		}
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, fullURI, nil)
	if err != nil {
		return err
	}
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", sm.token.AccessToken))

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	if res.StatusCode >= 400 {
		return fmt.Errorf("reading secret failed with status code %d", res.StatusCode)
	}

	if err := json.NewDecoder(res.Body).Decode(secret); err != nil {
		return err
	}

	return nil

}
