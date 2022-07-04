# Secrets Manager

sercretsmanager provides API to work with IBM Cloud Secrets Manager secrets

Implemented
- read secret

Not implemented
- create secret
- create group
- delete secret
- delete group
- list groups
- list secrets
- read group

Example usage

```go
// Initialize SecretsManager instance
sm := secretsmanager.NewSecretsManager(endpoint, token)
// Initialize secret of type 'kv' (refer to Secrete Manager docs for secrets types) with ID
secret := secretsmanager.NewKeyValueSecret(ID)
// Read secret from Secrets Manager
err := sm.Read(secret)
```

---
Readme created from Go doc with [goreadme](https://github.com/posener/goreadme)
