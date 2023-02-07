package dump

import "github.com/Brotiger/pet-notary_web-go/internal/domain/dump"

type IService interface {
	CreateDump(CreateDumpDto *dump.CreateDto) *dump.Dump
	DeleteDump(uuid string) error
	GetDumpById(uuid string) *dump.Dump
	GetAllDumps(limit, offset int) []*dump.Dump
}
