package handlers

import (
	"github.com/huibrm/DCode/server/gateway/sessions"
)

// HandlerContext will be a collection of all the contextual variables that the
// http handlers will need access to
type HandlerContext struct {
	SigningKey    string
	SessionsStore sessions.Store
}

// NewHandlerContext constructs new handler context ensuring that the dependencies are
// valid values
func NewHandlerContext(signingKey string, sessionStore sessions.Store) *HandlerContext {
	if sessionStore != nil {
		return &HandlerContext{
			SigningKey:    signingKey,
			SessionsStore: sessionStore,
		}
	}
	return nil
}
