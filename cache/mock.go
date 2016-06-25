package cache

import (
	"errors"
	"github.com/KyleBanks/go-kit/log"
	"encoding/json"
)

// Mock provides a mocked Cache implementation for testing.
type Mock struct {
	cache map[string]string
}

// NewMock instantiates and returns a new Mock cache.
func NewMock() *Mock {
	return &Mock{
		cache: make(map[string]string),
	}
}

// PutString stores a simple key-value pair in the mock.
func (m Mock) PutString(key string, value string) (interface{}, error) {
	log.Info("Mock PutString:", key, value)

	m.cache[key] = value
	return nil, nil
}

// GetString returns the string value stored with the given key.
//
// If the key doesn't exist, an error is returned.
func (m Mock) GetString(key string) (string, error) {
	log.Info("Mock GetString:", key)

	val, ok := m.cache[key]
	if !ok {
		return "", errors.New("Key not found")
	}

	return val, nil
}

// PutMarshaled stores a json marshalled value with the given key.
func (m Mock) PutMarshaled(key string, value interface{}) (interface{}, error) {
	// Marshal to JSON
	bytes, err := json.Marshal(value)
	if err != nil {
		return nil, err
	}

	// Store in the cache
	return m.PutString(key, string(bytes[:]))
}

// GetMarshaled retrieves an item from the cache with the specified key,
// and un-marshals it from JSON to the value provided.
//
// If they key doesn't exist, an error is returned.
func (m Mock) GetMarshaled(key string, v interface{}) error {
	cached, err := m.GetString(key)
	if err != nil {
		return err
	}

	if len(cached) > 0 {
		if err := json.Unmarshal([]byte(cached), v); err != nil {
			return err
		}
	}

	return nil
}

// Delete removes an item from the mock by it's key.
func (m Mock) Delete(key string) error {
	log.Info("Mock Delete:", key)

	delete(m.cache, key)
	return nil
}
