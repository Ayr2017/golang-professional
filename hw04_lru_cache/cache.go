package hw04lrucache

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
		items:    make(map[Key]*ListItem, capacity),
	}
}

func (c *lruCache) Set(key Key, value interface{}) bool {
	var wasInCache bool
	if elem, ok := c.items[key]; ok {
		wasInCache = true
		c.queue.MoveToFront(elem)
		elem.Value.(*cacheItem).value = value
	} else {
		if c.queue.Len() > c.capacity {
			lastElem := c.queue.Back()
			c.queue.Remove(lastElem)
			delete(c.items, lastElem.Value.(*cacheItem).key)
		}

		elem = c.queue.PushFront(&cacheItem{key, value})
		c.items[key] = elem
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
	c.items = make(map[Key]*ListItem, c.capacity)
}

type cacheItem struct {
	key   Key
	value interface{}
}
