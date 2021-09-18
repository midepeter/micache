package micache

import (
	"fmt"
	"sync"

	"github.com/midepeter/micache/store"
)

//This is an interface to implement the primary functions in the cache
type Cache interface {
	Get(key string) ([]byte, error)
	Put(key string, value []byte) error
	Delete(key string) error
}

type cache struct {
	mu sync.Mutex
	s  *store.MemStore
}

func (c *cache) Get(key string) ([]byte, error) {
	c.mu.Lock()
	defer c.mu.Unlock()
	v, ok := c.s.Get(key)
	if !ok {
		return fmt.Errorf("Was unable to fetch that file")
	}

	return v, nil
}

func (c *cache) Put(key string, value []byte) error {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.s.Put(key, value)
	return nil
}

func (c *cache) Delete(key string) error {
	c.mu.Lock()
	defer c.mu.Unlock()
	err := c.s.Delete(key)
	return err
}

//Purge() removes all data from the memstore
func (c *cache) Purge() {
	c.Purge()
}

//Defining the cache struct so it does not return nil values
func Newcache(size int, opts ...func(*cache)) *cache {
	r := &cache{
		l: store.New(size),
	}
	for _, opt := range opts {
		opt(r)
	}
	return r
}
