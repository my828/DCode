package handlers

import (
	"net/http"
)

// Middleware to handle cross origin requests

// CORSHandler enables all of the API calls to be callable
// cross-origin
type CORSHandler struct {
	handler http.Handler
}

// NewCORSHandler constructs and returns a new CORSHandler
func NewCORSHandler(handlerToWrap http.Handler) *CORSHandler {
	return &CORSHandler{
		handler: handlerToWrap,
	}
}

func (ch *CORSHandler) ServeHTTP(w http.ResponseWriter, r http.Request) {
	w.Header().Add("Access-Control-Allow-Origin", "*")
	w.Header().Add("Access-Control-Allow-Methods", "GET, PUT, POST, PATCH, DELETE")
	w.Header().Add("Access-Control-Allow-Headers", "Content-Type, Authorization")
	w.Header().Add("Access-Control-Expose-Headers", "Authorization")
	w.Header().Add("Access-Control-Max-Age", "600")
	// for preflight checks
	if r.Method == "Options" {
		w.WriteHeader(http.StatusOK)
		return
	}
	ch.ServeHTTP(w, r)
}
