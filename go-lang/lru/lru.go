package lru

type Cache struct {
	maxEntries int

	onEvicted func(key, value interface{})
}

// 构建一个 url cache
// maxEntries 指定能缓存的最大元素数量
// 如果设置了 onEvicted，当元素被从缓存中移除时，会回调 onEvicted
func New(maxEntries int, onEvicted func(key, value interface{})) *Cache {
	return nil
}

func (c *Cache) Init(maxEntries int, onEvicted func(key, value interface{})) *Cache {
	return nil
}

func (c *Cache) Len() int {
	return 0
}

func (c *Cache) Add(key, value interface{}) {
}

func (c *Cache) Get(key interface{}) (interface{}, bool) {
	return nil, false
}

func (c *Cache) Remove(key interface{}) {
}

func (c *Cache) RemoveOldest() {
}

func (c *Cache) Clear() {
}
