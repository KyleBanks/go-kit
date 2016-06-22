package cache

import (
	"github.com/rafaeljusto/redigomock"
	"github.com/garyburd/redigo/redis"
)

// Mock provides a mocked Cache implementation for testing.
type Mock struct {
	conn *redigomock.Conn
}

// NewMock instantiates and returns a new Mock cache.
func NewMock() *Cache {
	return &Mock{
		conn: redigomock.NewConn(),
	}
}

// PutString stores a simple key-value pair in the mock.
func (m Mock) PutString(key string, value string) (interface{}, error) {
	return m.conn.Do("set", key, value)
}

// GetString returns the string value stored with the given key.
//
// If the key doesn't exist, an error is returned.
func (m Mock) GetString(key string) (string, error) {
	return redis.String(m.conn.Do("get", key))
}

// Delete removes an item from the mock by it's key.
func (m Mock) Delete(key string) error {
	if _, err := m.conn.Do("del", key); err != nil {
		return err
	}

	return nil
}
