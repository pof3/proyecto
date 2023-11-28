package models

import "time"

type Dish struct {
	ID              string       `json:"id,omitempty"`
	Name            string       `json:"name,omitempty"`
	Detail          string       `json:"detail,omitempty"`
	Price           uint         `json:"price,omitempty"`
	Additions       []string     `json:"additions,omitempty"`
	AdditionsStruct []Ingredient `json:"additions_struct,omitempty"`
	CreatedAt       time.Time    `json:"created_at"`
}
