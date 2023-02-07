package dump

type IRepository interface {
	Create(dump *Dump) *Dump
	GetOne(uuid string) *Dump
	GetAll(uuid string) []*Dump
	Delete(uuid string) error
}
