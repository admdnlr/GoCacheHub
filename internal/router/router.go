// Dosya: internal/router/router.go

package router

import (
	"github.com/admdnlr/GoCacheHub/internal/api"
	"github.com/gorilla/mux"
)

// NewRouter, uygulama için router'ı oluşturur ve yapılandırır.
func NewRouter(apiHandler *api.CacheHandler) *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/set", apiHandler.SetHandler).Methods("POST")
	r.HandleFunc("/get", apiHandler.GetHandler).Methods("GET")
	// Daha fazla route burada tanımlanabilir.
	return r
}
