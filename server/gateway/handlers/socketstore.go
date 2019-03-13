package handlers

import (
	"fmt"
	"sync"
	"github.com/gorilla/websocket"
	"encoding/json"
	"github.com/streadway/amqp"
	"github.com/huibrm/DCode/server/gateway/sessions"


)

type SocketStore struct {
	Connections map[sessions.SessionID][]*websocket.Conn
	Lock sync.Mutex
}
// NewSocketStore makes a new map and mutex to safely store websockets
func NewSocketStore() *SocketStore{
	return &SocketStore{ 
		make(map[sessions.SessionID][]*websocket.Conn),
		sync.Mutex{},
	}
}
type MsgObj struct {
	SessionID 	   sessions.SessionID  `jdon:"sessionId"`
	Editor         string     `json:"editor"`
}
// Thread-safe method for inserting a connection
func (s *SocketStore) InsertConnection(conn *websocket.Conn, ss *SessionState) {
	s.Lock.Lock()
	// insert socket connection
	s.Connections[ss.SessionID] = append(s.Connections[ss.SessionID], conn)
	s.Lock.Unlock()

}

// Thread-safe method for removing a connection
func (s *SocketStore) RemoveConnection(sessionID sessions.SessionID, conn *websocket.Conn) {
	s.Lock.Lock()
	// insert socket connection
	var index int
	connections := s.Connections[sessionID]
	for i, connection := range connections {
		if conn == connection {
			index = i
		}
	}
	connections = append(connections[:index], connections[index+1:]...)

	s.Lock.Unlock()
}
 // fix after Harshitha writes microservice
func (s *SocketStore) Notify(msgs <-chan amqp.Delivery) error{
	for msg := range msgs {
		s.Lock.Lock()
		// this could be anything
		newMsg := &MsgObj{}
		if err := json.Unmarshal(msg.Body, newMsg); err != nil {
			return fmt.Errorf("Error unmarshalling userIDs %v", err)
		}

		s.WriteToConnections(msg.Body, newMsg.SessionID)
		s.Lock.Unlock()
	}
	return nil
}

func (s *SocketStore) WriteToConnections(message []byte, sessionID sessions.SessionID)  {
	var writeError error;
	connList := s.Connections[sessionID]
	for _, conn := range connList {
		writeError = conn.WriteMessage(websocket.TextMessage, message)
		if writeError != nil {
			s.RemoveConnection(sessionID, conn)
			conn.Close()
			return
		}
	}	
}