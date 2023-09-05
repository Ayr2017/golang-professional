package hw04lrucache

import "fmt"

type Key string

type Cache interface {
	Set(key Key, value interface{}) bool
	Get(key Key) (interface{}, bool)
	Clear()
}

type lruCache struct {
	capacity int
	queue    List
	items    map[Key]*ListItem
}

func NewCache(capacity int) Cache {
	return &lruCache{
		capacity: capacity,
		queue:    NewList(),
		items:    make(map[Key]*ListItem),
	}
}

func (c *lruCache) Set(key Key, value interface{}) bool {

	var wasInCache bool = false
	if elem, ok := c.items[key]; ok {
		wasInCache = true
		c.queue.MoveToFront(elem)
		elem.Value.(*cacheItem).value = value
	} else {
		elem = c.queue.PushFront(&cacheItem{key, value})
		c.items[key] = elem

		if c.queue.Len() > c.capacity {
			fmt.Println(c.queue.Len())
			lastElem := c.queue.Back()
			c.queue.Remove(lastElem)
			delete(c.items, lastElem.Value.(*cacheItem).key)
		}
	}

	return wasInCache
}

func (c *lruCache) Get(key Key) (interface{}, bool) {
	if elem, ok := c.items[key]; ok {
		c.queue.MoveToFront(elem)
		return elem.Value.(*cacheItem).value, true
	}

	return nil, false
}

func (c *lruCache) Clear() {
	c.queue = NewList()
	c.items = make(map[Key]*ListItem)
}

type cacheItem struct {
	key   Key
	value interface{}
}
