package sessions

import (
	"path"
	"fmt"
	"net/http"
)

// HeaderSessionID is a custom header for transferring SessionID
const HeaderSessionID = "X-SessionID"

// SaveSession saves the `sessionState` to the store
func SaveSession(sessionID SessionID, sessionState interface{}, store Store) (SessionID, error) {
	err := store.Save(sessionID, sessionState)
	if err != nil {
		return InvalidSessionID, fmt.Errorf("internal error: %d", http.StatusInternalServerError)
	}
	return sessionID, nil
}

// GetSessionID gets the session ID from the client's request
func GetSessionID(r *http.Request) (SessionID, error) {
	sessionID := r.Header.Get(HeaderSessionID)
	if len(sessionID) == 0 {
		// return InvalidSessionID, fmt.Errorf("invalid sessionID: %d", http.StatusBadRequest)
		sessionID = path.Base(r.URL.Path)
	}
	return SessionID(sessionID), nil
}

// GetState extracts the SessionID from the request,
// gets the associated state from the provided store into
// the `sessionState` parameter, and returns the SessionID
func GetState(r *http.Request, sessionsStore Store, sessionState interface{}) (SessionID, error) {
	sessionID, err := GetSessionID(r)
	if err != nil {
		return InvalidSessionID, err
	}
	err = sessionsStore.Get(sessionID, sessionState)
	if err != nil {
		return InvalidSessionID, err
	}
	return sessionID, nil
}
