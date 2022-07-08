package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/WojtekTomaszewski/ibmcsm"
)

func main() {
	smEndpoint := "https://b1a31726-787e-4cdb-9534-99f4e2b5c560.eu-de.secrets-manager.appdomain.cloud"
	smId := "91bea799-80e2-d255-5da7-a8a20349f422"
	apiKey := os.Getenv("SM_API_KEY")

	sm := ibmcsm.NewSecretsManager(smEndpoint, apiKey)

	secret, err := sm.ReadKeyValueSecret(smId)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	sBytes, _ := json.Marshal(secret.Resources[0].SecretData.Payload)
	fmt.Printf(string(sBytes))

}
