package dump

type repository struct {
}

func NewRepository() *repository {
	return &repository{}
}

func (dr *repository) Create(dump *Dump) *Dump {
	return nil
}

func (dr *repository) GetOneById(id string) *Dump {
	return nil
}

func (dr *repository) GetAll(id string) []*Dump {
	return nil
}

func (dr *repository) Delete(id string) error {
	return nil
}
