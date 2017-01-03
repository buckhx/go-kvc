package kvc

import (
	"sync"
	"time"
)

// MemKVC is an in memotry key value cache
type MemKVC struct {
	sync.RWMutex
	items map[Key]Value
}

// NewMem creates a new in memory key value cache
func NewMem() KVC {
	return &MemKVC{
		items: make(map[Key]Value),
	}
}

// Get fetch a value for the key
func (c *MemKVC) Get(k Key) Value {
	c.RLock()
	defer c.RUnlock()
	return c.UnsafeGet(k)
}

// Has returns true if the cache contains the key
func (c *MemKVC) Has(k Key) bool {
	c.RLock()
	defer c.RUnlock()
	return c.UnsafeHas(k)
}

// Set will set the value at the given key location
func (c *MemKVC) Set(k Key, v Value) {
	c.Lock()
	defer c.Unlock()
	c.UnsafeSet(k, v)
}

// SetTTL will set the value at key for the given amount of time to live
func (c *MemKVC) SetTTL(k Key, v Value, ttl time.Duration) {
	go func() {
		time.Sleep(ttl)
		c.Set(k, nil)
	}()
	c.Set(k, v)
}

// CompareAndSet sets the value if the cmp function returns true
// Only Unsafe* method's should be used in the cmp func since they do not acquire locks
func (c *MemKVC) CompareAndSet(k Key, v Value, cmp func() bool) bool {
	c.Lock()
	defer c.Unlock()
	ok := cmp()
	if ok {
		c.UnsafeSet(k, v)
	}
	return ok
}

// GetAndSet will apply a function to a value at the location of key in an atomic manner.
// Useful for things like atomic increments.
func (c *MemKVC) GetAndSet(k Key, fn func(cur Value) Value) {
	c.Lock()
	defer c.Unlock()
	cur := c.UnsafeGet(k)
	v := fn(cur)
	c.UnsafeSet(k, v)
}

// UnsafeGet gets a key value without acquiring a lock.
// CANNOT be access concurrently
func (c *MemKVC) UnsafeGet(k Key) Value {
	return c.items[k]
}

// UnsafeHas checks for key existence without acquiting a lock
// CANNOT be access concurrently
func (c *MemKVC) UnsafeHas(k Key) bool {
	_, ok := c.items[k]
	return ok
}

// UnsafeSet will set a value at the given value without acquiring a lock.
// CANNOT be access concurrently
func (c *MemKVC) UnsafeSet(k Key, v Value) {
	if v == nil {
		delete(c.items, k)
	} else {
		c.items[k] = v
	}
}
