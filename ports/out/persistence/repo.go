package persistence

type IRepoPort interface {
	Create(*interface{}) error
	Read(uint) (*interface{}, error)
	ReadAll() ([]*interface{}, error)
	Update(*interface{}) error
	Delete(uint) (*interface{}, error)
}