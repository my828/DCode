package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"sync"

	"DCode/server/gateway/sessions"

	"github.com/gorilla/websocket"
	"github.com/streadway/amqp"
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
	RabbitStore   *RabbitStore
	RedisStore  *sessions.RedisStore
}


// NewSocketStore makes a new map and mutex to safely store websockets
func NewSocketStore(rabbit *RabbitStore, redis *sessions.RedisStore) *SocketStore {
	return &SocketStore{
		Connections:   SessionConnections{},
		IPConnections: IPConnections{},
		Lock:          sync.Mutex{},
		RabbitStore: rabbit,
		RedisStore: redis,
	}
}

// Message is a struct to capture the
type Message struct {
	SessionID sessions.SessionID `json:"sessionID"`
	Figures   string             `json:"figures"`
	Code      string             `json:"code"`
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

// Listen receives messages from the websocket connection and populates the message queue
func (s *SocketStore) Listen(sessionState *SessionState) error {
	for {
		connections := s.Connections[sessionState.SessionID]
		for index, connection := range connections {
			m := &Message{}
			err := connection.ReadJSON(m)
			if err != nil {
				s.RemoveConnection(sessionState.SessionID, connection, index)
			} else {
				// save to redis
				newSS := &SessionState{
					m.SessionID,
					m.Figures,
					m.Code,
				}
				s.RedisStore.Save(m.SessionID, newSS)
				// save message to message queue
				if err := s.RabbitStore.Publish(m); err != nil {
					return fmt.Errorf("Error unmarshalling userIDs %v", err)	
				}

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


// Notify reads from Rabbit mq and writes to connections
func (s *SocketStore) Notify(msgs <-chan amqp.Delivery) error {
	for msg := range msgs {
		s.Lock.Lock()
		// this could be anything
		message := &Message{}
		if err := json.Unmarshal(msg.Body, message); err != nil {
			return fmt.Errorf("Error unmarshalling userIDs %v", err)
		}
		s.WriteToConnections(msg.Body, message.SessionID)
		s.Lock.Unlock()
	}
	return nil
}

// WriteToConnections writes to all the active connections for the session id
func (s *SocketStore) WriteToConnections(message []byte, sessionID sessions.SessionID) {
	var writeError error
	connList := s.Connections[sessionID]
	for index, conn := range connList {
		writeError = conn.WriteMessage(websocket.TextMessage, message)
		if writeError != nil {
			s.RemoveConnection(sessionID, conn, index)
			conn.Close()
			return
		}
	}
}


