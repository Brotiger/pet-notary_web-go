package composites

import (
	"github.com/Brotiger/pet-notary_web-go/internal/adapters"
	api "github.com/Brotiger/pet-notary_web-go/internal/adapters/api/dump"
	db "github.com/Brotiger/pet-notary_web-go/internal/adapters/db/dump"
	domain "github.com/Brotiger/pet-notary_web-go/internal/domain/dump"
)

type DumpComposite struct {
	Repository domain.IRepository
	Service    api.IService
	Handler    adapters.IHandler
}

func NewDumpComposite() *DumpComposite {
	repository := db.NewRepository()
	service := domain.NewService(repository)
	handler := api.NewHandler(service)

	return &DumpComposite{
		Repository: repository,
		Service:    service,
		Handler:    handler,
	}
}
