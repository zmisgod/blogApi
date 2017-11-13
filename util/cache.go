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
