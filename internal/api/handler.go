// Dosya: internal/api/handler.go

package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/admdnlr/GoCacheHub/internal/cache"
)

// KeyValue bir anahtar-değer çiftini temsil eder.
type KeyValue struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

// CacheHandler HTTP isteklerini işlemek için cache nesnesine erişim sağlar.
type CacheHandler struct {
	Cache *cache.Cache
}

// NewCacheHandler bir CacheHandler nesnesi oluşturur ve döndürür.
func NewCacheHandler(cache *cache.Cache) *CacheHandler {
	return &CacheHandler{
		Cache: cache,
	}
}

// SetHandler bir anahtar-değer çiftini cache'e ekler.
// İstek gövdesinden JSON olarak alınan anahtar ve değeri kullanır.
func (h *CacheHandler) SetHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	var kv KeyValue
	err := json.NewDecoder(r.Body).Decode(&kv)
	if err != nil {
		http.Error(w, "Error parsing JSON request body", http.StatusBadRequest)
		return
	}

	if kv.Key == "" || kv.Value == "" {
		http.Error(w, "Please provide both 'key' and 'value' in request body.", http.StatusBadRequest)
		return
	}

	h.Cache.Set(kv.Key, kv.Value)
	fmt.Fprintf(w, "Key '%s' set to '%s'.\n", kv.Key, kv.Value)
}

// GetHandler bir anahtara karşılık gelen değeri döndürür
func (h *CacheHandler) GetHandler(w http.ResponseWriter, r *http.Request) {
	key := r.URL.Query().Get("key")

	if key == "" {
		http.Error(w, "Please provide a 'key' query parameter.", http.StatusBadRequest)
		return
	}

	value, ok := h.Cache.Get(key)
	if !ok {
		http.Error(w, "Key not found.", http.StatusNotFound)
		return
	}

	fmt.Fprintf(w, "Value for key '%s': '%s'.\n", key, value)
}
