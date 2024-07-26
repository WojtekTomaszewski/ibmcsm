package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/WojtekTomaszewski/ibmcsm"
)

func main() {
	smEndpoint := os.Getenv("SM_URL")
	smId := os.Getenv("SM_SECRET_ID")
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
