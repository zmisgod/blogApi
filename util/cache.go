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
	itemList *list.List
	itemMap  map[string]*list.Element
	maxSize  int
	lock     sync.Mutex
}

func (cache *MyCache) Get(key string) interface{} {
	ele, ok := cache.itemMap[key]
	if !ok {
		return nil
	}
	cache.itemList.MoveToFront(ele)
	kv := ele.Value.(*keyValue)
	return kv.value
}

func (cache *MyCache) IsExist(key string) bool {
	if _, ok := cache.itemMap[key]; ok {
		return true
	}
	return false
}

func (cache *MyCache) Put(key string, value interface{}, timeout int16) error {
	cache.lock.Lock()
	defer cache.lock.Unlock()

	ele, ok := cache.itemMap[key]

	if ok {
		cache.itemList.MoveToFront(ele)
		kv := ele.Value.(*keyValue)
		kv.value = value
	} else {
		ele := cache.itemList.PushFront(&keyValue{key: key, value: value})
		cache.itemMap[key] = ele

		if cache.itemList.Len() > cache.maxSize {
			delElem := cache.itemList.Back()
			kv := delElem.Value.(*keyValue)
			cache.itemList.Remove(delElem)
			delete(cache.itemMap, kv.key)
		}
	}

	return nil
}

func (cache *MyCache) Delete(key string) error {
	cache.lock.Lock()
	defer cache.lock.Unlock()

	ele, ok := cache.itemMap[key]
	if ok {
		cache.itemList.Remove(ele)
		delete(cache.itemMap, key)
	}
	return nil
}

func (cache *MyCache) ClearAll() error {
	cache.lock.Lock()

	defer cache.lock.Unlock()

	for k, e := range cache.itemMap {
		cache.itemList.Remove(e)
		delete(cache.itemMap, k)
	}

	return nil
}

func NewMyCache(maxSize int) *MyCache {
	return &MyCache{
		itemList: list.New(),
		itemMap:  make(map[string]*list.Element),
		maxSize:  maxSize,
	}
}

func (cache *MyCache) Len() int {
	return cache.itemList.Len()
}
