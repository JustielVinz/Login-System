package secret

import "github.com/google/uuid"

func GenerateSessionKey() string {
	// Generate a UUID (session key)
	key := uuid.New()
	return key.String()
}
