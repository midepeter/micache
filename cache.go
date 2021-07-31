package micache

import (
	"fmt"
	"sync"
)

//This is an interface to implement the primary functions in the cache
type Cache interface {
	Get(key string) ([]byte, error)
	Set(key string, value []byte) error
	Delete(key string) error

type cache struct {
	mu sync.Mutex
	l  *lru.Cache
}

func (c *cache) Get(key string) ([]byte, error) {
	c.mu.Lock()
	defer c.mu.Unlock()
	v, ok := c.l.Get(key)
	if !ok {
		return fmt.Errorf("Was unable to fetch that file")
	}

	return v, nil
}

func (c *cache) Set(key string, value []byte) error {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.l.Set(key, value)
	return nil

}

func (c *cache) Delete(key string) error {
	c.mu.Lock()
	defer c.mu.Unlock()
	err := c.l.Delete(key)
	return err
}

//Defining the cache struct so it does not return nil values
func Newcache(size int, opts ...func(*cache)) *cache {
	r := &cache{
		l: lru.New(size),
	}
	for _, opt := range opts {
		opt(r)
	}
	return r
}
