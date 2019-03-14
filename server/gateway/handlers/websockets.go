package handlers

// import (
// 	"log"
// 	"net/http"

// 	"github.com/gorilla/websocket"
// 	"github.com/huibrm/DCode/server/gateway/sessions"
// )
// // Upgrader checks the orgin and specs for websockets
// var Upgrader = websocket.Upgrader{
// 	ReadBufferSize:  1024,
// 	WriteBufferSize: 1024,
// 	CheckOrigin: func(r *http.Request) bool {
// 		// This function's purpose is to reject websocket upgrade requests if the
// 		// origin of the websockete handshake request is coming from unknown domains.
// 		// This prevents some random domain from opening up a socket with your server.
// 		// TODO: make sure you modify this for your HW to check if r.Origin is your host
// 		// if r.Header.Get("Origin") != "https://catsfordays.me" {
// 		// 	log.Print("Connection Refused", 403)
// 		// 	return false
// 		// }
// 		return true
// 	},
// }
// // WebSocket manages the websockets connection
// type WebSocket struct {
// 	SS       *SocketStore
// 	CTX      *HandlerContext
// 	Upgrader websocket.Upgrader
// }
// // for testing
// type msg struct {
// 	editorbody string
// }
// // NewWebSocket creates and populates websocket structs
// func NewWebSocket(ss *SocketStore, ctx *HandlerContext) *WebSocket {
// 	return &WebSocket{
// 		ss,
// 		ctx,
// 		Upgrader,
// 	}
// }

// // WebSocketConnectionHandler manages new connections and data coming in and out of them
// func (ws *WebSocket) WebSocketConnectionHandler(w http.ResponseWriter, r *http.Request) {
// 	// handle the websocket handshake

// 	log.Print("Inside Websocket Handler")
// 	//getting auth token and adding to query string
// 	sessionState := &SessionState{}
// 	_, err := sessions.GetState(r, ws.CTX.SessionsStore, sessionState)
// 	log.Println("After session was checked: ID = ", sessionState)
// 	if err != nil {
// 		log.Println("Page does not exist")
// 		http.Error(w, "invalid page", http.StatusUnauthorized)
// 		return
// 	}

// 	// upgrading to websocket
// 	conn, err := ws.Upgrader.Upgrade(w, r, nil)
// 	if err != nil {
// 		http.Error(w, "Failed to open websocket connection", 401)
// 		return
// 	}
// 	// inserting connection to connection store
// 	ws.SS.InsertConnection(conn, sessionState)
// 	go ws.SS.Listen(conn,sessionState.SessionID)
// 	// ws.SS.InsertConnection(conn)

// 	// go (func(conn *websocket.Conn, ws *WebSocket) {
// 	go (func(conn *websocket.Conn, sessionState *SessionState, ws *WebSocket) {
// 		defer conn.Close()
// 		defer ws.SS.RemoveConnection(sessionState.SessionID, conn)
// 		// defer ws.SS.RemoveConnection(conn)

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
// 			// ignore ping and pong messages
// 		}

// 	})(conn, sessionState, ws)
// 	// })(conn, ws)

// }


// // for testing editor body and websockets
// // func echo(conn *websocket.Conn) {
// // 	for { // infinite loop
// // 		m := msg{}

// // // WebSocket manages the websockets connection
// // type WebSocket struct {
// // 	SS       *SocketStore
// // 	CTX      *HandlerContext
// // 	Upgrader websocket.Upgrader
// // }

// // // for testing
// // type msg struct {
// // 	editorbody string
// // }

// // // NewWebSocket creates and populates websocket structs
// // func NewWebSocket(ss *SocketStore, ctx *HandlerContext) *WebSocket {
// // 	return &WebSocket{
// // 		ss,
// // 		ctx,
// // 		Upgrader,
// // 	}
// // }

// // // WebSocketConnectionHandler manages new connections and data coming in and out of them
// // func (ws *WebSocket) WebSocketConnectionHandler(w http.ResponseWriter, r *http.Request) {
// // 	sessionState := &SessionState{}
// // 	_, err := sessions.GetState(r, ws.CTX.SessionsStore, sessionState)
// // 	if err != nil {
// // 		log.Println("session does not exist")
// // 		http.Error(w, "session does not exist", http.StatusUnauthorized)
// // 		return
// // 	}

// // 	// upgrade to websocket connection
// // 	conn, err := ws.Upgrader.Upgrade(w, r, nil)
// // 	if err != nil {
// // 		http.Error(w, "Failed to open websocket connection", 401)
// // 		return
// // 	}
// // 	// inserting connection to connection store
// // 	ws.SS.InsertConnection(conn, sessionState)
// // }

// // // for testing editor body and websockets
// // // func echo(conn *websocket.Conn) {
// // // 	for { // infinite loop
// // // 		m := msg{}

// // // 		err := conn.ReadJSON(&m)
// // // 		if err != nil {
// // // 			fmt.Println("Error reading json.", err)
// // // 			conn.Close()
// // // 			break
// // // 		}

// // // 		fmt.Printf("Got message: %#v\n", m)

// // // 		if err = conn.WriteJSON(m); err != nil {
// // // 			fmt.Println(err)
// // // 		}
// // // 	}
// // // 	// cleanup
// // // }
