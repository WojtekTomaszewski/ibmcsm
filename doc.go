// Package ibmcsm provides API to read secrets from IBM Cloud Secrets Manager.
//
// Implemented
// - read username_password, kv and arbitrary secrets
//
// Not implemented
// - read iam_credentials,imported_cert,public_cert,private_cert secrets
//
//
// Example usage
//
//   // Initialize SecretsManager instance
//   sm := ibmcsm.NewSecretsManager(endpoint, apikey)
//   // Read kv secret with specific id
//   secret := ibmcsm.ReadKeyValueSecret(id)
//   // Read velue from kv secret for specific key
//   secret.Resources[0].SecretData.Payload["key"]
package ibmcsm
