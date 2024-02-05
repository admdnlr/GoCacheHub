// Dosya: internal/cache/cache_test.go

package cache

import (
	"testing"
)

func TestCacheSetAndGet(t *testing.T) {
	cache := NewCache()
	key := "testKey"
	value := "testValue"

	cache.Set(key, value)
	gotValue, exists := cache.Get(key)
	if !exists || gotValue != value {
		t.Errorf("Get(%q) = %q, %v; want %q, true", key, gotValue, exists, value)
	}
}

func TestCacheDelete(t *testing.T) {
	cache := NewCache()
	key := "testKey"
	value := "testValue"

	cache.Set(key, value)
	cache.Delete(key)
	_, exists := cache.Get(key)
	if exists {
		t.Errorf("Delete(%q) failed; key %q still exists", key, key)
	}
}
