//author: ademdinler

package main

import (
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/admdnlr/GoCacheHub/internal/api"
	"github.com/admdnlr/GoCacheHub/internal/cache"
	"github.com/admdnlr/GoCacheHub/internal/storage"
	"github.com/gorilla/mux"
)

func main() {
	// Cache nesnesini oluştur
	c := cache.NewCache()

	// Storage nesnesini oluştur ve veri dosyasının yolu
	storagePath := "./data/cache_data.json"
	st := storage.NewStorage(storagePath)

	// Uygulama başladığında verileri yükle
	if err := st.Load(&c.Store); err != nil {
		log.Printf("Failed to load data from storage: %v", err)
	} else {
		log.Println("Data loaded successfully from storage.")
	}

	// CacheHandler nesnesini oluştur
	handler := api.NewCacheHandler(c)

	// Gorilla mux router'ını kullanarak HTTP handler'larını tanımla
	r := mux.NewRouter()
	r.HandleFunc("/set", handler.SetHandler).Methods("POST")
	r.HandleFunc("/get", handler.GetHandler).Methods("GET")

	// Graceful shutdown için sinyal yakalama
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		<-sigChan
		log.Println("Shutting down server...")

		// Uygulama kapanırken verileri kaydet
		if err := st.Save(c.Store); err != nil {
			log.Fatalf("Failed to save data to storage: %v", err)
		} else {
			log.Println("Data saved successfully to storage.")
		}

		os.Exit(0)
	}()

	// HTTP sunucusunu başlat ve router'ı kullan
	log.Println("Starting server on :9080")
	if err := http.ListenAndServe(":9080", r); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
