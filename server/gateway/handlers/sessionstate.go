package handlers

import (
	"DCode/server/gateway/sessions"
)

// SessionState represents the collection of relevant data for our server
type SessionState struct {
	SessionID   sessions.SessionID `json:"sessionID"`
	Figures     string 			   `json:"figures"`
	Code        string			   `json:"code"`
}
