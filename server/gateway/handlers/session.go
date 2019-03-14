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

	// // add connection to SocketStore
	// remoteAddress := IPAddress(r.RemoteAddr)
	// connection := hc.SocketStore.IPConnections[remoteAddress]
	// sessionConnections := hc.SocketStore.Connections[sessionID]
	// sessionConnections = append(sessionConnections, connection)
	// hc.SocketStore.Connections[sessionID] = sessionConnections

	w.Header().Add(HeaderSessionID, string(sessionID))
	w.Write([]byte(sessionID))
}

// GetPageHandler handles requests to `/dcode/v1/{pageID}`
func (hc *HandlerContext) GetPageHandler(w http.ResponseWriter, r *http.Request) {
	sessionState := &SessionState{}
	_, err := sessions.GetState(r, hc.SessionsStore, sessionState)
	if err != nil {
		http.Error(w, "error getting session", http.StatusInternalServerError)
		return
	}

	// publish message to RabbitMQ
	message := &Message{
		SessionID: sessionState.SessionID,
		Figures:   sessionState.Figures,
		Code:      sessionState.Code,
	}
	hc.SocketStore.RabbitStore.Publish(message)

	w.Header().Add(HeaderSessionID, string(sessionState.SessionID))
	w.Write([]byte(sessionState.SessionID))
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
