package handlers

import (
	"log"
	"sync"

	"DCode/server/gateway/sessions"

	"github.com/gorilla/websocket"
)

// IPAddress represents the ip address obtaiend from request
type IPAddress string

// SessionConnections maps session ids to websocket connections
type SessionConnections map[sessions.SessionID][]*websocket.Conn

// IPConnections maps ip address to a web socket connection
type IPConnections map[IPAddress]*websocket.Conn

// SocketStore listens for updates to the message queue and notifies the
// client
type SocketStore struct {
	Connections   SessionConnections
	IPConnections IPConnections
	Lock          sync.Mutex
}

/// for test
// type SocketStore struct {
// 	Connections []*websocket.Conn
// 	Lock sync.Mutex
// }

// NewSocketStore makes a new map and mutex to safely store websockets
func NewSocketStore() *SocketStore {
	return &SocketStore{
		Connections:   SessionConnections{},
		IPConnections: IPConnections{},
		Lock:          sync.Mutex{},
	}
}

// Message is a struct to capture the
type Message struct {
	SessionID sessions.SessionID `json:"sessionID"`
	Code      string             `json:"code"`
	Figures   string             `json:"figures"`
}

// InsertConnection is a thread-safe method for inserting a connection
func (s *SocketStore) InsertConnection(connection *websocket.Conn, sessionState *SessionState) {
	// insert socket connection & unlock
	s.Lock.Lock()
	connections := s.Connections[sessions.SessionID(sessionState.SessionID)]
	connections = append(connections, connection)
	s.Connections[sessionState.SessionID] = connections
	s.Lock.Unlock()

	// process incoming messages from the client
	go s.Listen(sessionState)
}

// InsertConnection(connection *websocket.Conn, sessionState *SessionState)

// Listen receives messages from the websocket connection and populates the message queue
func (s *SocketStore) Listen(sessionState *SessionState) {
	for {
		connections := s.Connections[sessionState.SessionID]
		for index, connection := range connections {
			message := &Message{}
			err := connection.ReadJSON(message)
			if err != nil {
				s.RemoveConnection(sessionState.SessionID, connection, index)
			} else {
				// save message to message queue
				// rabbitstore.publish(message)
			}
		}
	}
}

// RemoveConnection thread-safe method for removing a connection
func (s *SocketStore) RemoveConnection(sessionID sessions.SessionID, connection *websocket.Conn, connectionIndex int) {
	s.Lock.Lock()
	err := connection.Close()
	if err != nil {
		log.Println("Error closing the connection")
	}
	connections := s.Connections[sessionID]
	connections = append(connections[:connectionIndex], connections[connectionIndex+1])
	s.Connections[sessionID] = connections
	s.Lock.Unlock()
}

// // FOR TESTING
// func (s *SocketStore) RemoveConnection(conn *websocket.Conn) {
// 	s.Lock.Lock()
// 	// insert socket connection
// 	var index int
// 	connections := s.Connections
// 	for i, connection := range connections {
// 		if conn == connection {
// 			index = i
// 		}
// 	}
// 	connections = append(connections[:index], connections[index+1:]...)

// // 	s.Lock.Unlock()
// // }
// // fix after Harshitha writes microservice
// func (s *SocketStore) Notify(msgs <-chan amqp.Delivery) error {
// 	for msg := range msgs {
// 		s.Lock.Lock()
// 		// this could be anything
// 		newMsg := &MsgObj{}
// 		if err := json.Unmarshal(msg.Body, newMsg); err != nil {
// 			return fmt.Errorf("Error unmarshalling userIDs %v", err)
// 		}

// 		s.WriteToConnections(msg.Body, newMsg.SessionID)
// 		s.Lock.Unlock()
// 	}
// 	return nil
// }

// func (s *SocketStore) WriteToConnections(message []byte, sessionID sessions.SessionID) {
// 	var writeError error
// 	connList := s.Connections[sessionID]
// 	for _, conn := range connList {
// 		writeError = conn.WriteMessage(websocket.TextMessage, message)
// 		if writeError != nil {
// 			s.RemoveConnection(sessionID, conn)
// 			conn.Close()
// 			return
// 		}
// 	}
// }
