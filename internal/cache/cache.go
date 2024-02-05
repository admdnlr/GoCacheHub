package cache

import (
	"sync"
)

// Cache anahtar-değer çiftlerini saklayan temel yapı.
type Cache struct {
	mu    sync.Mutex
	Store map[string]string // Store alanını dışarıdan erişilebilir yap
}

// NewCache yeni bir Cache nesnesi oluşturur ve başlatır.
func NewCache() *Cache {
	return &Cache{
		Store: make(map[string]string),
	}
}

// Set bir anahtar-değer çiftini cache'e ekler.
func (c *Cache) Set(key, value string) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.Store[key] = value
}

// Get bir anahtara karşılık gelen değeri döndürür.
func (c *Cache) Get(key string) (string, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()
	value, ok := c.Store[key]
	return value, ok
}

// Delete bir anahtar-değer çiftini cache'den siler.
func (c *Cache) Delete(key string) {
	c.mu.Lock()
	defer c.mu.Unlock()
	delete(c.Store, key)
}
