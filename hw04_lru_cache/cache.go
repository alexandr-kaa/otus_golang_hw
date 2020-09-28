package hw04_lru_cache //nolint:golint,stylecheck
import "sort"

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
			cacheItemvalue.count = cacheItemvalue.count + 1
			cacheItemvalue.value = value
			valueMap.Value = cacheItemvalue
			cache.queue.MoveToFront(valueMap)
			return true
		}
		return false
	}
	newItem := cacheItem{value: value, key: key, count: 0}
	front := cache.queue.PushFront(newItem)
	cache.items[key] = front
	//элемент key должен остаться
	cache.checkAndDelete(key)
	return false
}

//Удаляем из map и queu лишний элемент с наименьшей частотой использования
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
}

func (cache *lruCache) Get(key Key) (interface{}, bool) {
	if valueMap, ok := cache.items[key]; ok {
		cache.queue.MoveToFront(valueMap)
		if cacheItemvalue, okTransform := valueMap.Value.(cacheItem); okTransform {
			cacheItemvalue.count = cacheItemvalue.count + 1
			valueMap.Value = cacheItemvalue
			return cacheItemvalue.value, true
		}
	}
	return nil, false
}

func (cache *lruCache) Clear() {
	for key := range cache.items {
		cache.queue.Remove(cache.items[key])
	}

}

type cacheItem struct {
	// Place your code here
	value interface{}
	key   Key
	//для подсчета частоты ссылок
	count int
}
type listItemPointer *listItem

func NewCache(capa int) Cache {

	return &lruCache{capacity: capa, queue: NewList(), items: make(map[Key]*listItem)}
}
