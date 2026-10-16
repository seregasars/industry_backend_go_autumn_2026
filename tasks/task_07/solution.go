package main

import (
	"container/list"
	"sync"
)

type LRU[K comparable, V any] interface {
	Get(key K) (value V, ok bool)
	Set(key K, value V)
}
type entry[K comparable, V any] struct {
	key   K
	value V
}
type LRUCache[K comparable, V any] struct {
	capacity int
	mu       sync.Mutex
	ll       list.List
	items    map[K]*list.Element
}

func NewLRUCache[K comparable, V any](capacity int) *LRUCache[K, V] {
	cache := LRUCache[K, V]{}
	cache.capacity = capacity
	cache.items = make(map[K]*list.Element)
	return &cache	
}

func (c *LRUCache[K, V]) Get(key K) (value V, ok bool) {
	c.mu.Lock()
	defer c.mu.Unlock()
	e, ok := c.items[key]
	if !ok {
		return
	}
	value = e.Value.(entry[K, V]).value
	c.ll.MoveToFront(e)
	return value, ok
}

func (c *LRUCache[K, V]) Set(key K, value V) {
	c.mu.Lock()
	defer c.mu.Unlock()
	if c.capacity <= 0 {
		return
	}
	e, ok := c.items[key]
	if ok {
		e.Value = entry[K, V]{key: key,value: value}
		return
	}
	if c.ll.Len() >= c.capacity{
		back := c.ll.Back()
		keyOld := back.Value.(entry[K, V]).key
		delete(c.items, keyOld)
		c.ll.Remove(back)
	}
	newElement := c.ll.PushFront(entry[K,V]{key: key,value: value})
	c.items[key] = newElement
}
