// Simple cache wrapper, used to abstract Redis/Memcache/etc behind a reusable
// API for simple use cases.
//
// The idea is that Redis could be swapped for another cache and the client wouldn't
// need to update another (except perhaps calls to New to provide different connection
// parameters).
package cache

import (
	"encoding/json"
	"github.com/KyleBanks/go-kit/log"
	"github.com/garyburd/redigo/redis"
)

// Cacher defines a mockable Cache interface that can store values in a key-value cache.
type Cacher interface {
	PutString(key string, value string) (interface{}, error)
	GetString(key string) (string, error)
	Delete(key string) error

	PutMarshaled(key string, value interface{}) (interface{}, error)
	GetMarshaled(key string, v interface{}) error
}

type Cache struct {
	pool *redis.Pool
}

// New instantiates and returns a new Cache.
func New(host string) *Cache {
	log.Info("Initializing Cache...")

	return &Cache{
		pool: &redis.Pool{
			MaxIdle:   5,
			MaxActive: 100,
			Dial: func() (redis.Conn, error) {
				return redis.Dial("tcp", host)
			},
		},
	}
}

// PutString stores a simple key-value pair in the cache.
func (c Cache) PutString(key string, value string) (interface{}, error) {
	r := c.pool.Get()
	defer r.Close()

	return r.Do("set", key, value)
}

// GetString returns the string value stored with the given key.
//
// If the key doesn't exist, an error is returned.
func (c Cache) GetString(key string) (string, error) {
	r := c.pool.Get()
	defer r.Close()

	return redis.String(r.Do("get", key))
}

// PutMarshaled stores a json marshalled value with the given key.
func (c Cache) PutMarshaled(key string, value interface{}) (interface{}, error) {
	// Marshal to JSON
	bytes, err := json.Marshal(value)
	if err != nil {
		return nil, err
	}

	// Store in the cache
	return c.PutString(key, string(bytes[:]))
}

// GetMarshaled retrieves an item from the cache with the specified key,
// and un-marshals it from JSON to the value provided.
//
// If they key doesn't exist, an error is returned.
func (c Cache) GetMarshaled(key string, v interface{}) error {
	cached, err := c.GetString(key)
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

// Delete removes an item from the cache by it's key.
func (c Cache) Delete(key string) error {
	r := c.pool.Get()
	defer r.Close()

	if _, err := r.Do("del", key); err != nil {
		return err
	}

	return nil
}
