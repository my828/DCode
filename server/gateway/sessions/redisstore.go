package sessions

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/go-redis/redis"
)

// RedisStore represents a session.Store backed by redis.
type RedisStore struct {
	// Redis client used to talk to redis server.
	Client *redis.Client
	// Used for key expiry time on redis.
	SessionDuration time.Duration
}

// NewRedisStore constructs a new RedisStore
func NewRedisStore(client *redis.Client, sessionDuration time.Duration) *RedisStore {
	return &RedisStore{
		Client:          client,
		SessionDuration: sessionDuration,
	}
}

// Save saves the provided `sessionState` and associated SessionID to the store
func (rs *RedisStore) Save(sessionID SessionID, sessionState interface{}) error {
	state, err := json.Marshal(sessionState)
	if err != nil {
		return err
	}
	redisKey := sessionID.getRedisKey()
	statusCmd := rs.Client.Set(redisKey, state, rs.SessionDuration)
	if statusCmd.Err() != nil {
		return fmt.Errorf("internal error: %v", statusCmd.Err())
	}
	return nil
}

// Get populates `sessionState` with the data previously saved
// for the given SessionID
func (rs *RedisStore) Get(sessionID SessionID, sessionState interface{}) error {
	redisKey := sessionID.getRedisKey()
	statusCmd := rs.Client.Get(redisKey)
	val, err := statusCmd.Result()
	if err != nil {
		return ErrStateNotFound
	}
	// unmarshal it back into the `sessionState` parameter
	err = json.Unmarshal([]byte(val), sessionState)
	return err
}

// Extend extends validity of session for another 48 hours
func (rs *RedisStore) Extend(sessionID SessionID) error {
	redisKey := sessionID.getRedisKey()
	res := rs.Client.Expire(redisKey, rs.SessionDuration)
	if res.Err() != nil {
		return fmt.Errorf("error extending the session")
	}
	return nil
}

// Delete deletes all state data associated with the SessionID from the store.
func (rs *RedisStore) Delete(sessionID SessionID) error {
	res := rs.Client.Del(sessionID.getRedisKey())
	if res == nil {
		return fmt.Errorf("session does not exist: %d", http.StatusBadRequest)
	}
	return nil
}

// getRedisKey() returns the redis key to use for the SessionID
func (sid SessionID) getRedisKey() string {
	// convert the SessionID to a string and add the prefix "sid:" to keep
	// SessionID keys separate from other keys that might end up in this
	// redis instance
	return "sid:" + sid.String()
}
