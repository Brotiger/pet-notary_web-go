package dump

import "github.com/Brotiger/pet-notary_web-go/internal/domain/dump"

type repository struct {
}

func NewRepository() *repository {
	return &repository{}
}

func (dr *repository) Create(dump *dump.Dump) *dump.Dump {
	return nil
}

func (dr *repository) GetOne(uuid string) *dump.Dump {
	return nil
}

func (dr *repository) GetAll(uuid string) []*dump.Dump {
	return nil
}

func (dr *repository) Delete(uuid string) error {
	return nil
}
