package cache_node

type CacheNode interface {
	Add(int, any)
	Get(int) (any, error)
}

type InMemoryCache struct {
	data map[int]any
}
