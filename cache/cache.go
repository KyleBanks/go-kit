// Simple cache wrapper, used to abstract Redis/Memcache/etc behind a reusable
// API for simple use cases.
//
// The idea is that Redis could be swapped for another cache and the client wouldn't
// need to update another (except perhaps calls to New to provide different connection
// parameters).
package cache

import (
	"github.com/garyburd/redigo/redis"
	"github.com/KyleBanks/go-kit/log"
)

type Cache struct {
	pool *redis.Pool
}

// New instantiates and returns a new Cache.
func New(host string) *Cache {
	log.Info("Initializing Cache...")

	return &Cache{
		pool: &redis.Pool{
			MaxIdle:   5,
			MaxActive: 10,
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

// Delete removes an item from the cache by it's key.
func (c Cache) Delete(key string) error {
	r := c.pool.Get()
	defer r.Close()

	if _, err := r.Do("del", key); err != nil {
		return err
	}

	return nil
}
