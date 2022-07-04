package ibmcsm

import "time"

type secretType string

const (
	UsernamePasswordType secretType = "username_password"
	KeyValueType         secretType = "kv"
	ArbitraryType        secretType = "arbitrary"
)

type secretresource interface {
	Id() string
	Type() secretType
}

// readSecret represents secret in Secrets Manager.
type secret struct {
	Metadata struct {
		CollectionType  string `json:"collection_type"`
		CollectionTotal int    `json:"collection_total"`
	} `json:"metadata"`
	Resources []secretresource `json:"resources"`
}

type readSecretResource struct {
	CreatedBy        string        `json:"created_by"`
	CreationDate     time.Time     `json:"creation_date"`
	Crn              string        `json:"crn"`
	Downloaded       bool          `json:"downloaded"`
	ID               string        `json:"id"`
	Labels           []interface{} `json:"labels"`
	LastUpdateDate   time.Time     `json:"last_update_date"`
	LocksTotal       int           `json:"locks_total"`
	Name             string        `json:"name"`
	SecretData       interface{}   `json:"secret_data"`
	SecretGroupID    string        `json:"secret_group_id"`
	SecretType       secretType    `json:"secret_type"`
	State            int           `json:"state"`
	StateDescription string        `json:"state_description"`
	Versions         []struct {
		CreatedBy        string    `json:"created_by"`
		CreationDate     time.Time `json:"creation_date"`
		Downloaded       bool      `json:"downloaded"`
		ID               string    `json:"id"`
		PayloadAvailable bool      `json:"payload_available"`
	} `json:"versions"`
	VersionsTotal int `json:"versions_total"`
}

func (s *readSecretResource) Id() string {
	return s.ID
}

func (s *readSecretResource) Type() secretType {
	return s.SecretType
}

type secretsDataUsernamePassword struct {
	Username string `json:"username,omitempty"`
	Password string `json:"password,omitempty"`
}

type secretsDataKeyValue struct {
	Payload map[string]string `json:"payload,omitempty"`
}

type secretsDataArbitrary struct {
	Payload string `json:"payload,omitempty"`
}

type SecretOptions struct {
	Id             string     `json:"id,omitempty"`
	SecretType     secretType `json:"secret_type,omitempty"`
	Name           string     `json:"name,omitempty"`
	Description    string     `json:"description,omitempty"`
	Labels         []string   `json:"labels,omitempty"`
	SecretGroupID  string     `json:"secret_group_id,omitempty"`
	ExpirationDate time.Time  `json:"expiration_date,omitempty"`
}

func NewSecret(opt *SecretOptions) *secret {
	secret := &secret{}
	secretResource := &readSecretResource{}

	switch opt.SecretType {
	case UsernamePasswordType:
		secretResource.SecretData = secretsDataUsernamePassword{}
	case KeyValueType:
		secretResource.SecretData = secretsDataKeyValue{}
	case ArbitraryType:
		secretResource.SecretData = secretsDataArbitrary{}
	}

	if len(opt.Id) > 0 {
		secretResource.ID = opt.Id
		secretResource.SecretType = opt.SecretType
		secret.Resources = append(secret.Resources, secretResource)
	}

	return secret
}
