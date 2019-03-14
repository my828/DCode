package handlers

import (
	"DCode/server/gateway/sessions"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

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
	// u, err := url.Parse(r.URL.Path)
	// u.Path = path.Join(u.Path, string([]byte(sessionID)))
	// log.Print("NEW PATH: ", u.Path)
	// responds with sessionID
	w.Header().Add(HeaderSessionID, string(sessionID))
	// if err := json.NewEncoder(w).Encode(sessionID); err != nil {
	// 	http.Error(w, fmt.Sprintf("errr"), http.StatusInternalServerError)
	// 	return
	// }
	log.Println("Header: ", w.Header().Get(HeaderSessionID))
	w.Write([]byte(sessionID))
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
