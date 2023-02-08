package composites

import "github.com/Brotiger/pet-notary_web-go/internal/modules/dump"

type DumpComposite struct {
	Repository dump.IRepository
	Service    dump.IService
	Handler    dump.IHandler
}

func NewDumpComposite() *DumpComposite {
	repository := dump.NewRepository()
	service := dump.NewService(repository)
	handler := dump.NewHandler(service)

	return &DumpComposite{
		Repository: repository,
		Service:    service,
		Handler:    handler,
	}
}
