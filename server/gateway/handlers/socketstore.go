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
type MsgObj struct {
	SessionID 	   sessions.SessionID  `jdon:"sessionId"`
	Editor         string     `json:"editor"`
}
// Thread-safe method for inserting a connection
func (s *SocketStore) InsertConnection(conn *websocket.Conn, ss *SessionState) int {
	s.Lock.Lock()
	// insert socket connection
	s.Connections[ss.SessionID] = append(s.Connections[ss.SessionID], conn)
	connID := len(s.Connections[ss.SessionID])
	s.Lock.Unlock()
	return connID
}

// Thread-safe method for removing a connection
func (s *SocketStore) RemoveConnection(sessionID sessions.SessionID, id int) {
	s.Lock.Lock()
	// insert socket connection
	s.Connections[sessionID] = append(s.Connections[sessionID][:id], s.Connections[sessionID][id+1:]...)
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

func (s *SocketStore) WriteToConnections(message []byte, sessionID sessions.SessionID, connID int)  {
	var writeError error;
	connList := s.Connections[id]
	for _, conn := range connList {
		writeError = conn.WriteMessage(websocket.TextMessage, message)
		if writeError != nil {
			s.RemoveConnection(sessionID, connID)
			conn.Close()
			return
		}
	}	
}