package util

import (
	"container/list"
	"sync"
)

type keyValue struct {
	key   string
	value interface{}
}

type MyCache struct {
	maxSize  int
	lock     sync.Mutex
	itemMap  map[string]*list.Element
	itemList *list.List
}

func (cache *MyCache) Get(key string) interface{} {
	elem, ok := cache.itemMap[key]
	if !ok {
		return nil
	}
	cache.itemList.MoveToFront(elem)
	kv := elem.Value.(*keyValue)
	return kv.value
}

func (cache *MyCache) Len() int {
	return cache.itemList.Len()
}

func NewMyCache(maxSize int) *MyCache {
	return &MyCache{
		itemList: list.New(),
		itemMap:  make(map[string]*list.Element),
		maxSize:  maxSize,
	}
}

func (cache *MyCache) ClearAll() error {
	cache.lock.Lock()
	defer cache.lock.Unlock()

	for k, v := range cache.itemMap {
		cache.itemList.Remove(v)
		delete(cache.itemMap, k)
	}
	return nil
}
