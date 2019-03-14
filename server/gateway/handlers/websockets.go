package handlers

import (
	"DCode/server/gateway/sessions"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

// Upgrader checks the orgin and specs for websockets
var Upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		// This function's purpose is to reject websocket upgrade requests if the
		// origin of the websockete handshake request is coming from unknown domains.
		// This prevents some random domain from opening up a socket with your server.
		// TODO: make sure you modify this for your HW to check if r.Origin is your host
		// if r.Header.Get("Origin") != "https://catsfordays.me" {
		// 	log.Print("Connection Refused", 403)
		// 	return false
		// }
		return true
	},
}

// WebSocket manages the websockets connection
type WebSocket struct {
	Ctx      *HandlerContext
	Upgrader websocket.Upgrader
}

// for testing
type msg struct {
	editorbody string
}

// NewWebSocket creates and populates websocket structs
func NewWebSocket(ctx *HandlerContext) *WebSocket {
	return &WebSocket{
		ctx,
		Upgrader,
	}
}

// // WebSocketConnectionHandler manages new connections and data coming in and out of them
// func (ws *WebSocket) WebSocketConnectionHandler(w http.ResponseWriter, r *http.Request) {
// 	log.Print("Inside Websocket Handler")
// 	// upgrading to websocket
// 	conn, err := ws.Upgrader.Upgrade(w, r, nil)
// 	if err != nil {
// 		http.Error(w, "Failed to open websocket connection", 401)
// 		return
// 	}
// 	remoteAddr := IPAddress(r.RemoteAddr)
// 	// inserting connection to connection store
// 	ws.Ctx.SocketStore.InsertConnection(remoteAddr, conn)

// 	go (func(remoteAddress IPAddress, conn *websocket.Conn) {
// 		defer conn.Close()
// 		defer ws.Ctx.SocketStore.RemoveIPConnection(remoteAddress)

// 		for {
// 			messageType, _, err := conn.ReadMessage()
// 			if messageType == websocket.TextMessage || messageType == websocket.BinaryMessage {
// 				log.Print("message sucessfully made")

// 			} else if messageType == websocket.CloseMessage {
// 				log.Print("Message was a closeMessage, try resending")
// 				break
// 			} else if err != nil {
// 				log.Print("Error reading message.")
// 				break
// 			}
// 		}

// 	})(remoteAddr, conn)
// }

// WebSocketConnectionHandler manages new connections and data coming in and out of them
func (ws *WebSocket) WebSocketConnectionHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Inside WebSocketHandler...")
	sessionState := &SessionState{}
	_, err := sessions.GetState(r, ws.Ctx.SessionsStore, sessionState)
	log.Println("After session was checked: ID = ", sessionState.SessionID)
	if err != nil {
		log.Println("Page does not exist")
		http.Error(w, "invalid page", http.StatusUnauthorized)
		return
	}

	// upgrading to websocket
	conn, err := ws.Upgrader.Upgrade(w, r, nil)
	if err != nil {
		http.Error(w, "Failed to open websocket connection", 401)
		return
	}
	// inserting connection to connection store
	ws.Ctx.SocketStore.InsertConnection(conn, sessionState)
	// ws.SS.InsertConnection(conn)
	go ws.Ctx.SocketStore.Listen(sessionState, conn)
}
