// Dosya: internal/api/handler.go

package api

import (
	"encoding/json"
	"fmt"
	"log"
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
func (h *CacheHandler) SetHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("SetHandler called") // Log mesajı

	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		log.Printf("Invalid request method: %v\n", r.Method) // Log mesajı
		return
	}

	var kv KeyValue
	err := json.NewDecoder(r.Body).Decode(&kv)
	if err != nil {
		http.Error(w, "Error parsing JSON request body", http.StatusBadRequest)
		log.Printf("Error parsing JSON: %v\n", err) // Log mesajı
		return
	}

	if kv.Key == "" || kv.Value == "" {
		http.Error(w, "Please provide both 'key' and 'value' in request body.", http.StatusBadRequest)
		log.Println("Empty key or value provided") // Log mesajı
		return
	}

	h.Cache.Set(kv.Key, kv.Value)
	log.Printf("Key '%s' set to '%s'.\n", kv.Key, kv.Value) // Log mesajı
	fmt.Fprintf(w, "Key '%s' set to '%s'.\n", kv.Key, kv.Value)
}

// GetHandler bir anahtara karşılık gelen değeri döndürür
func (h *CacheHandler) GetHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("GetHandler called") // Log mesajı

	key := r.URL.Query().Get("key")

	if key == "" {
		http.Error(w, "Please provide a 'key' query parameter.", http.StatusBadRequest)
		log.Println("Key query parameter is missing") // Log mesajı
		return
	}

	value, ok := h.Cache.Get(key)
	if !ok {
		http.Error(w, "Key not found.", http.StatusNotFound)
		log.Printf("Key '%s' not found.\n", key) // Log mesajı
		return
	}

	log.Printf("Value for key '%s': '%s'.\n", key, value) // Log mesajı
	fmt.Fprintf(w, "Value for key '%s': '%s'.\n", key, value)
}
