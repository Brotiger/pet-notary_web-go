package dump

type service struct {
	dumpRepository IRepository
}

func NewService(repository IRepository) *service {
	return &service{dumpRepository: repository}
}

func (ds *service) CreateDump(CreateDumpDto *CreateDto) *Dump {
	return nil
}

func (ds *service) DeleteDump(uuid string) error {
	return nil
}

func (ds *service) GetDumpById(uuid string) *Dump {
	return nil
}

func (ds *service) GetAllDumps(limit, offset int) []*Dump {
	return nil
}
