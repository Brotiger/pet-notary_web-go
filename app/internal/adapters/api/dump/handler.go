package dump

import (
	"net/http"

	"github.com/gorilla/mux"
)

const (
	dumpURL  = "/dump"
	dumpsURL = "/dumps"
)

type handler struct {
	dumpService IService
}

func NewHandler(service IService) *handler {
	return &handler{dumpService: service}
}

func (h *handler) Register(router *mux.Router) {
	router.HandleFunc(dumpsURL, h.GetAllDumps).Methods("GET")
}

func (h *handler) GetAllDumps(w http.ResponseWriter, r *http.Request) {
	//dumps := h.dumpService.GetAllDumps(10, 0)
	w.Write([]byte("hello world"))
}
