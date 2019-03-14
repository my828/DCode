package handlers

import (
	"log"
	"net/http"

	"DCode/server/gateway/sessions"

	"github.com/gorilla/websocket"
)

const clientDomain = "https://catsfordays.me"

// Upgrader checks the orgin and specs for websockets
var Upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		// if r.Header.Get("Origin") != clientDomain {
		// 	log.Print("Connection Refused", 403)
		// 	return false
		// }
		return true
	},
}

// WebSocket manages the websockets connection
type WebSocket struct {
	SS       *SocketStore
	CTX      *HandlerContext
	Upgrader websocket.Upgrader
}

// for testing
type msg struct {
	editorbody string
}

// NewWebSocket creates and populates websocket structs
func NewWebSocket(ss *SocketStore, ctx *HandlerContext) *WebSocket {
	return &WebSocket{
		ss,
		ctx,
		Upgrader,
	}
}

// WebSocketConnectionHandler manages new connections and data coming in and out of them
func (ws *WebSocket) WebSocketConnectionHandler(w http.ResponseWriter, r *http.Request) {
	sessionState := &SessionState{}
	_, err := sessions.GetState(r, ws.CTX.SessionsStore, sessionState)
	if err != nil {
		log.Println("session does not exist")
		http.Error(w, "session does not exist", http.StatusUnauthorized)
		return
	}

	// upgrade to websocket connection
	conn, err := ws.Upgrader.Upgrade(w, r, nil)
	if err != nil {
		http.Error(w, "Failed to open websocket connection", 401)
		return
	}
	// inserting connection to connection store
	ws.SS.InsertConnection(conn, sessionState)
}

// for testing editor body and websockets
// func echo(conn *websocket.Conn) {
// 	for { // infinite loop
// 		m := msg{}

// 		err := conn.ReadJSON(&m)
// 		if err != nil {
// 			fmt.Println("Error reading json.", err)
// 			conn.Close()
// 			break
// 		}

// 		fmt.Printf("Got message: %#v\n", m)

// 		if err = conn.WriteJSON(m); err != nil {
// 			fmt.Println(err)
// 		}
// 	}
// 	// cleanup
// }
