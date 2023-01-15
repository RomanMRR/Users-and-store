package store

type Data interface {
	Validate() error
}

// Repository  interface
type Repository interface {
	Create(Data) error
	Update(Data) error
	Delete(id int, tableName string) error
	Find(whatLookingFor string, tableName string) ([]Data, error)
}
