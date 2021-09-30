package micache

import (
	"container/list"
	"sync"
	"time"
)

var UTCNow = func() time.Time { return time.Now().UTC() }

type Options struct {
	Capacity int
	Policy   EvictionPolicy
}

type LRUCache struct {
	ItemRemoved func(*Item)
	Cap         int
	Policy      EvictionPolicy
	Items       map[string]*list.Element
	Lru         *list.List
	mu          *sync.Mutex
}

type Item struct {
	Key     string
	Value   interface{}
	Expires time.Time
}

func NewCache(o Options) *LRUCache {
	var cap int
	if o.Capacity > 0 {
		cap = o.Capacity
	} else {
		cap = 100
	}

	var pol EvictionPolicy
	if o.Policy != nil {
		pol = o.Policy
	} else {
		pol = NewNoEvictionPolicy()
	}

	return &LRUCache{
		ItemRemoved: func(*Item) {},
		Cap:         cap,
		Policy:      pol,
		Items:       map[string]*list.Element{},
		Lru:         list.New(),
		mu:          &sync.Mutex{},
	}
}

type EvictionPolicy interface {
	Apply(*Item) error
}

type NoEvictionPolicy struct{}

func NewNoEvictionPolicy() *NoEvictionPolicy {
	return new(NoEvictionPolicy)
}

func (p NoEvictionPolicy) Apply(*Item) error {
	return nil
}
