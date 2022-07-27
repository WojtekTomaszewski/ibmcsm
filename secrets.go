package ibmcsm

import (
	"encoding/json"
	"time"
)

const (
	usernamePasswordType = "username_password"
	keyValueType         = "kv"
	arbitraryType        = "arbitrary"
)

type secretError struct {
	Metadata struct {
		CollectionType  string `json:"collection_type"`
		CollectionTotal int    `json:"collection_total"`
	} `json:"metadata"`
	Resources []struct {
		ErrorMessage string `json:"error_message"`
	} `json:"resources"`
	Errors []struct {
		Code    string `json:"code"`
		Message string `json:"message"`
	} `json:"errors"`
}

type keyValueSecret struct {
	Metadata  metadata                 `json:"metadata"`
	Resources []keyValueSecretResource `json:"resources"`
}

type keyValueSecretResource struct {
	CreatedBy      string    `json:"created_by,omitempty"`
	CreationDate   time.Time `json:"creation_date,omitempty"`
	Crn            string    `json:"crn,omitempty"`
	Downloaded     bool      `json:"downloaded,omitempty"`
	ID             string    `json:"id,omitempty"`
	Labels         []string  `json:"labels,omitempty"`
	LastUpdateDate time.Time `json:"last_update_date,omitempty"`
	LocksTotal     int       `json:"locks_total,omitempty"`
	Name           string    `json:"name,omitempty"`
	SecretData     struct {
		Payload *json.RawMessage `json:"payload,omitempty"`
	} `json:"secret_data,omitempty"`
	SecretGroupID    string    `json:"secret_group_id,omitempty"`
	SecretType       string    `json:"secret_type,omitempty"`
	State            int       `json:"state,omitempty"`
	StateDescription string    `json:"state_description,omitempty"`
	Versions         []version `json:"versions,omitempty"`
	VersionsTotal    int       `json:"versions_total,omitempty"`
}

type usernamePasswordSecret struct {
	Metadata  metadata                         `json:"metadata"`
	Resources []usernamePasswordSecretResource `json:"resources"`
}

type usernamePasswordSecretResource struct {
	CreatedBy      string    `json:"created_by,omitempty"`
	CreationDate   time.Time `json:"creation_date,omitempty"`
	Crn            string    `json:"crn,omitempty"`
	Downloaded     bool      `json:"downloaded,omitempty"`
	ID             string    `json:"id,omitempty"`
	Labels         []string  `json:"labels,omitempty"`
	LastUpdateDate time.Time `json:"last_update_date,omitempty"`
	LocksTotal     int       `json:"locks_total,omitempty"`
	Name           string    `json:"name,omitempty"`
	SecretData     struct {
		Username string `json:"username,omitempty"`
		Password string `json:"password,omitempty"`
	} `json:"secret_data,omitempty"`
	SecretGroupID    string    `json:"secret_group_id,omitempty"`
	SecretType       string    `json:"secret_type,omitempty"`
	State            int       `json:"state,omitempty"`
	StateDescription string    `json:"state_description,omitempty"`
	Versions         []version `json:"versions,omitempty"`
	VersionsTotal    int       `json:"versions_total,omitempty"`
}

type arbitrarySecret struct {
	Metadata  metadata                  `json:"metadata"`
	Resources []arbitrarySecretResource `json:"resources"`
}

type arbitrarySecretResource struct {
	CreatedBy      string    `json:"created_by,omitempty"`
	CreationDate   time.Time `json:"creation_date,omitempty"`
	Crn            string    `json:"crn,omitempty"`
	Downloaded     bool      `json:"downloaded,omitempty"`
	ID             string    `json:"id,omitempty"`
	Labels         []string  `json:"labels,omitempty"`
	LastUpdateDate time.Time `json:"last_update_date,omitempty"`
	LocksTotal     int       `json:"locks_total,omitempty"`
	Name           string    `json:"name,omitempty"`
	SecretData     struct {
		Payload string `json:"payload,omitempty"`
	} `json:"secret_data,omitempty"`
	SecretGroupID    string    `json:"secret_group_id,omitempty"`
	SecretType       string    `json:"secret_type,omitempty"`
	State            int       `json:"state,omitempty"`
	StateDescription string    `json:"state_description,omitempty"`
	Versions         []version `json:"versions,omitempty"`
	VersionsTotal    int       `json:"versions_total,omitempty"`
}

type metadata struct {
	CollectionType  string `json:"collection_type,omitempty"`
	CollectionTotal int    `json:"collection_total,omitempty"`
}

type version struct {
	CreatedBy        string    `json:"created_by,omitempty"`
	CreationDate     time.Time `json:"creation_date,omitempty"`
	Downloaded       bool      `json:"downloaded,omitempty"`
	ID               string    `json:"id,omitempty"`
	PayloadAvailable bool      `json:"payload_available,omitempty"`
}
