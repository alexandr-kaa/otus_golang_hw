package hw04_lru_cache //nolint:golint,stylecheck

type Key string

type Cache interface {
	Set(key Key, value interface{}) bool
	Get(key Key) (interface{}, bool)
	Clear()
}

type lruCache struct {
	// Place your code here:
	capacity int
	queue    List
	items    map[Key]*listItem // - items
}

func (cache *lruCache) Set(key Key, value interface{}) bool {
	if valueMap, ok := cache.items[key]; ok {
		if cacheItemvalue, okTransform := valueMap.Value.(cacheItem); okTransform {
			cacheItemvalue.value = value
			valueMap.Value = cacheItemvalue
			cache.queue.MoveToFront(valueMap)
			return true
		}
		return false
	}
	newItem := cacheItem{value: value, key: key}
	cache.checkAndDelete()
	front := cache.queue.PushFront(newItem)
	cache.items[key] = front
	return false
}
func (cache *lruCache) checkAndDelete() {
	if cache == nil {
		return
	}
	if cache.queue.Len() < cache.capacity {
		return
	}
	last := cache.queue.Back()
	if cacheLastItem, ok := last.Value.(cacheItem); ok {
		delete(cache.items, cacheLastItem.key)
		cache.queue.Remove(last)
	}
}

/* Удаляем из map и queue лишний элемент с наименьшей частотой использования.
func (cache *lruCache) checkAndDelete(keyexternal Key) {
	values := make([]cacheItem, 0)
	for key, value := range cache.items {
		if item, ok := value.Value.(cacheItem); ok && (key != keyexternal) {
			values = append(values, item)
		}
	}
	sort.Slice(values, func(i, j int) bool {
		return values[i].count < values[j].count
	})
	if len(values) > cache.capacity-1 {
		cache.queue.Remove(cache.items[values[0].key])
		delete(cache.items, values[0].key)
	}
}*/

func (cache *lruCache) Get(key Key) (interface{}, bool) {
	if valueMap, ok := cache.items[key]; ok {
		cache.queue.MoveToFront(valueMap)
		// Нужен для mock теста
		if valueMap == nil {
			return nil, true
		}
		if cacheItemvalue, okTransform := valueMap.Value.(cacheItem); okTransform {
			valueMap.Value = cacheItemvalue
			return cacheItemvalue.value, true
		}
		panic("Transform datatype invalid!!!")
	}
	return nil, false
}

func (cache *lruCache) Clear() {
	for key := range cache.items {
		cache.queue.Remove(cache.items[key])
	}
	cache.queue = NewList()
}

type cacheItem struct {
	// Place your code here
	value interface{}
	key   Key
}

func NewCache(capa int) Cache {
	return &lruCache{capacity: capa, queue: NewList(), items: make(map[Key]*listItem)}
}

func NewCacheExistingList(cap int, inlist List) Cache {
	mapData := make(map[Key]*listItem)

	if _, ok := inlist.Front().Value.(cacheItem); !ok {
		panic("wrong data type")
	}

	for cap < inlist.Len() {
		inlist.Remove(inlist.Back())
	}

	for i := inlist.Front(); i != nil; i = i.Next {
		mapData[i.Value.(cacheItem).key] = i
	}

	return &lruCache{capacity: cap, queue: inlist, items: mapData}
}
