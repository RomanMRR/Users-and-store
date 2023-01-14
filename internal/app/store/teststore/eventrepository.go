package teststore

import (
	"http-rest-api/internal/app/model"
	"time"
)

// EventRepository ...
type EventRepository struct {
	store  *Store
	events map[time.Time]map[int]*model.User
}

// Create ...
// func (r *EventRepository) Create(e *model.User) error {
// 	if err := e.Validate(); err != nil {
// 		return err
// 	}

// 	r.events[e.Date][e.ID] = e
// 	e.ID = len(r.events)

// 	return nil
// }

// FindFprDay ...
// func (r *EventRepository) FindForDay(date time.Time, user_id int) ([]model.User, error) {
// 	var events []model.Event
// 	for key := range r.events {
// 		for key_id := range r.events[key] {
// 			if key == date && key_id == user_id {
// 				events = append(events, *r.events[key][key_id])
// 			}
// 		}
// 	}

// 	if len(events) == 0 {
// 		return nil, store.ErrRecordNotFound
// 	}

// 	return events, nil
// }
