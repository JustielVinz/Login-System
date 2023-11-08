package struct_test

import "time"

// SessionData represents user session data
type SessionData struct {
	Token          string
	CreationTime   time.Time
	ExpirationTime time.Time
}
