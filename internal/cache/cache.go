package cache

const cacheCapacity = 100

var cache map[int]interface{}

func NewCache() {
	cache = make(map[int]interface{}, cacheCapacity)
}

func AddToCache(key int, item interface{}) {
	if len(cache) > cacheCapacity*0.9 {
		clear(cache)
	}
	cache[key] = item
}

func GetFromCache(key int) (interface{}, bool) {
	if v, ok := cache[key]; ok {
		return v, ok
	}
	return nil, false
}

func InvalidateCache(key int) {
	delete(cache, key)
}
