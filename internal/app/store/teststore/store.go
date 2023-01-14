package teststore

// Store ...
type Store struct {
	eventRepository *EventRepository
}

// New ...
func New() *Store {
	return &Store{}
}

// Event ...
// func (s *Store) Event() store.Repository {
// 	if s.eventRepository != nil {
// 		return s.eventRepository
// 	}

// 	s.eventRepository = &EventRepository{
// 		store:  s,
// 		events: make(map[time.Time]*model.Event),
// 	}

// 	return s.eventRepository
// }
