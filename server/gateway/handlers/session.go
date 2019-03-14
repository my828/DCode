package handlers

import (
	"DCode/server/gateway/sessions"
	"net/http"

	"github.com/gorilla/mux"
	// "github.com/gorilla/websocket"
)

const clientDomain = "https://catsfordays.me"

// Upgrader checks the orgin and specs for websockets

// HeaderSessionID is a custom header for transferring SessionID
const HeaderSessionID = "X-SessionID"

// NewSessionHandler handles https requests to the `/dcode/v1/new` resource
func (hc *HandlerContext) NewSessionHandler(w http.ResponseWriter, r *http.Request) {
	sessionID, err := sessions.NewSessionID(hc.SigningKey)
	if err != nil {
		http.Error(w, "error creating a new session", http.StatusInternalServerError)
		return
	}
	sessionState := &SessionState{
		SessionID: sessionID,
	}
	_, err = sessions.SaveSession(sessionID, sessionState, hc.SessionsStore)
	if err != nil {
		http.Error(w, "error creating a new session", http.StatusInternalServerError)
		return
	}
	w.Header().Add(HeaderSessionID, string(sessionID))
	w.Write([]byte(sessionID))
}

// GetPageHandler handles requests to `/dcode/v1/{sessionID}`
func (hc *HandlerContext) GetPageHandler(w http.ResponseWriter, r *http.Request) {
	sessionState := &SessionState{}
	_, err := sessions.GetState(r, hc.SessionsStore, sessionState)
	if err != nil {
		http.Error(w, "error getting session", http.StatusInternalServerError)
		return
	}
	remoteAddress := IPAddress(r.RemoteAddr)
	val := hc.SocketStore.IPConnections[remoteAddress]
	if val == nil {
		connection, err := Upgrader.Upgrade(w, r, nil)
		if err != nil {
			http.Error(w, "error getting session", http.StatusInternalServerError)
			return
		}
		hc.SocketStore.IPConnections[remoteAddress] = connection
		hc.SocketStore.InsertConnection(connection, sessionState)
	} else {
		hc.SocketStore.InsertConnection(val, sessionState)
	}
}

// SessionExtensionHandler extends the session validity by another 48 hours
func (hc *HandlerContext) SessionExtensionHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	sessionID := vars["pageID"]
	err := hc.SessionsStore.Extend(sessions.SessionID(sessionID))
	if err != nil {
		http.Error(w, "error extending session", http.StatusInternalServerError)
		return
	}
	// responds with sessionID
	w.Header().Add(HeaderSessionID, string(sessionID))

	w.Write([]byte("session extended"))
}

// TestHandler tests a specific page -- should expire after a set time
// TODO: delete later
func (hc *HandlerContext) TestHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	sessionID := vars["pageID"]
	sessionState := &SessionState{}
	err := hc.SessionsStore.Get(sessions.SessionID(sessionID), sessionState)
	if err != nil {
		http.Error(w, "invalid session", http.StatusBadRequest)
		return
	}
	w.Write([]byte("valid dcode page id!"))
}
