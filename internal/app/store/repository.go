package store

type Data interface {
	Validate() error
}

// Repository ...
type Repository interface {
	Create(Data) error
	Update(Data) error
	Delete(id int, tableName string) error
	Find(whatLookingFor string, tableName string) ([]Data, error)
	// FindForMonth(month, year string, user_id int) ([]model.User, error)
	// FindForWeek(week, year string, user_id int) ([]model.User, error)
}
