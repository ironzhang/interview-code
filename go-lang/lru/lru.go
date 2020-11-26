package lru

type Cache struct {
	maxEntries int

	onEvicted func(key, value interface{})
}

// 构建一个 url cache
// maxEntries 指定能缓存的最大元素数量，<= 0 则不限制
// 如果设置了 onEvicted，当元素被从缓存中移除时，会回调 onEvicted
func New(maxEntries int, onEvicted func(key, value interface{})) *Cache {
	return nil
}

// Init 初始化缓存
func (c *Cache) Init(maxEntries int, onEvicted func(key, value interface{})) *Cache {
	return c
}

// Len 获取缓存的元素数量
func (c *Cache) Len() int {
	return 0
}

// Add 添加一个元素
func (c *Cache) Add(key, value interface{}) {
}

// Get 通过 key 获取一个元素
func (c *Cache) Get(key interface{}) (interface{}, bool) {
	return nil, false
}

// Remove 通过 key 移除一个元素
func (c *Cache) Remove(key interface{}) {
}

// RemoveOldest 移除最老的一个元素
func (c *Cache) RemoveOldest() {
}

// Clear 清除所有元素
func (c *Cache) Clear() {
}
