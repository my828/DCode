package handlers

import (
	"DCode/server/gateway/sessions"
)

// HandlerContext will be a collection of all the contextual variables that the
// http handlers will need access to
type HandlerContext struct {
	SigningKey    string
	SessionsStore sessions.Store
	SocketStore   *SocketStore

}

// NewHandlerContext constructs new handler context ensuring that the dependencies are
// valid values
func NewHandlerContext(signingKey string, sessionStore sessions.Store, socketStore *SocketStore) *HandlerContext {
	if sessionStore != nil {
		return &HandlerContext{
			SigningKey:    signingKey,
			SessionsStore: sessionStore,
			SocketStore:   socketStore,
		}
	}
	return nil
}
