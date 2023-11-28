package models

import "time"

type Ingredient struct {
	ID        string    `json:"id,omitempty"`
	Name      string    `json:"name,omitempty"`
	Price     uint      `json:"price,omitempty"`
	CreatedAt time.Time `json:"created_at"`
}
