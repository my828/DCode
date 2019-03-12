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
	SessionID 	   SessionID  `jdon:"sessionId"`
	Editor         string     `json:"editor"`
}
// Thread-safe method for inserting a connection
func (s *SocketStore) InsertConnection(conn *websocket.Conn, ss *SessionState) int {
	s.Lock.Lock()
	// insert socket connection
	s.Connections[ss.SessionID] = append(s.Connections[ss.SessionID], conn)
	connID := len(s.Connctions[ss.SessionID])
	s.Lock.Unlock()
	return connID
}

// Thread-safe method for removing a connection
func (s *SocketStore) RemoveConnection(ss *SessionState, id int) {
	s.Lock.Lock()
	// insert socket connection
	s.Connections[ss.SessionID] := append(s.Connections[ss.SessionID][:id], s.Connections[ss.SessionID][id+1:]...)
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

func (s *SocketStore) WriteToConnections(message []byte, id sessions.SessionID)  {
	var writeError error;
	connList := s.Connections[id]
	for _, conn := range connList {
		writeError = conn.WriteMessage(websocket.TextMessage, message)
		if writeError != nil {
			s.RemoveConnection(id)
			conn.Close()
			return
		}
	}	
}