package handlers

import (
	"DCode/server/gateway/sessions"
)

// SessionState represents the collection of relevant data for our server
type SessionState struct {
	SessionID   sessions.SessionID `json:"sessionID"`
	ActiveUsers []*string          `json:"activeUsers"`
	Figures     string
	Code        string
}
