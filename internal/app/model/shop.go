package model

import (
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/pkg/errors"
)

const ShopTable = "shops"

type Shop struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	Address string `json:"address"`
	Working bool   `json:"working"`
	Owner   string `json:"owner"`
}

type UpdateShopInput struct {
	ID      *int    `json:"id"`
	Name    *string `json:"name"`
	Address *string `json:"address"`
	Working *bool   `json:"working"`
	Owner   *string `json:"owner"`
}

// Validate model shop
func (s *Shop) Validate() error {
	return validation.ValidateStruct(
		s,
		validation.Field(&s.Name, validation.Required, validation.Length(2, 100)),
		validation.Field(&s.Address, validation.Required, validation.Length(6, 100)),
		validation.Field(&s.Working, validation.Required),
	)
}

func (i *UpdateShopInput) Validate() error {
	if i.Address == nil && i.Name == nil && i.Owner == nil && i.Working == nil {
		return errors.New("Shop update structure has no values")
	}

	return nil
}
