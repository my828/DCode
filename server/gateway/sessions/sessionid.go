package sessions

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
)

// SessionID is used as a key for the Redis session store
type SessionID string

// InvalidSessionID represents an empty, invalid session ID
const InvalidSessionID SessionID = ""

const idLength = 6

// NewSessionID creates and returns a new digitally-signed session ID,
// using `signingKey` as the HMAC signing key. An error is returned only
// if there was an error generating random bytes for the session ID
func NewSessionID(signingKey string) (SessionID, error) {
	salt := make([]byte, idLength)
	_, err := rand.Read(salt)
	if err != nil {
		return "", fmt.Errorf("error while generating a session ID: %v", err)
	}
	sessionID := SessionID(base64.URLEncoding.EncodeToString(salt))
	return sessionID, nil
}

// String returns a string representation of the sessionID
func (sid SessionID) String() string {
	return string(sid)
}
