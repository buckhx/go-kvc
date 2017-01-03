package kvc

import "time"

// Key is the storage key value for lookup
type Key interface{}

// Value is the content of cache at the location of Key
type Value interface{}

// KVC is an interface for the cache
type KVC interface {
	Get(Key) Value
	Has(Key) bool
	Set(Key, Value)
	SetTTL(Key, Value, time.Duration)
	GetAndSet(Key, func(Value) Value)
	CompareAndSet(Key, Value, func() bool) bool
}
