// Dosya: internal/api/handler_test.go

package api

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/admdnlr/GoCacheHub/internal/cache"
)

func TestSetHandler(t *testing.T) {
	cache := cache.NewCache()
	handler := NewCacheHandler(cache)

	server := httptest.NewServer(http.HandlerFunc(handler.SetHandler))
	defer server.Close()

	jsonData := `{"key":"testKey", "value":"testValue"}`
	resp, err := http.Post(server.URL, "application/json", strings.NewReader(jsonData))
	if err != nil {
		t.Fatalf("Failed to send POST request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected status code 200; got %v", resp.StatusCode)
	}
}

func TestGetHandler(t *testing.T) {
	cache := cache.NewCache()
	handler := NewCacheHandler(cache)
	cache.Set("testKey", "testValue")

	server := httptest.NewServer(http.HandlerFunc(handler.GetHandler))
	defer server.Close()

	resp, err := http.Get(server.URL + "?key=testKey")
	if err != nil {
		t.Fatalf("Failed to send GET request: %v", err)
	}
	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected status code 200; got %v", resp.StatusCode)
	}

	// Optionally, check response body
	// ...
}
