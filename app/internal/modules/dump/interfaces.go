package dump

import "github.com/gorilla/mux"

type IRepository interface {
	Create(dump *Dump) *Dump
	GetOneById(id string) *Dump
	GetAll(id string) []*Dump
	Delete(id string) error
}

type IHandler interface {
	Register(router *mux.Router)
}

type IService interface {
	CreateDump(CreateDumpDto *CreateDto) *Dump
	DeleteDump(id string) error
	GetDumpById(id string) *Dump
	GetAllDumps(limit, offset int) []*Dump
}
