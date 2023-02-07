package main

import (
	"net/http"
	"time"

	"github.com/Brotiger/pet-notary_web-go/internal/composites"
	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	dumpComposite := composites.NewDumpComposite()
	dumpComposite.Handler.Register(router)

	server := &http.Server{
		Handler:      router,
		Addr:         "127.0.0.1:8000",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	server.ListenAndServe()
}
