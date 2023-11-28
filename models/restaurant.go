package models

import (
	"github.com/surrealdb/surrealdb.go"
	"proyecto/db"
)

type Restaurant struct {
	ID           string   `json:"id,omitempty"`
	Name         string   `json:"name,omitempty"`
	Dishes       []string `json:"dishes,omitempty"`
	DishesFields []Dish   `json:"dishes_fields,omitempty"`
}

func GetRestaurants() ([]Restaurant, error) {
	result, err := db.DB.Select("restaurant")
	if err != nil {
		return nil, err
	}
	restaurants := make([]Restaurant, 0)
	final := make([]Restaurant, 0)
	err = surrealdb.Unmarshal(result, &restaurants)
	if err != nil {
		return nil, err
	}
	for _, restaurant := range restaurants {
		for _, restaurantDish := range restaurant.Dishes {
			d := Dish{Additions: make([]string, 0)}
			dr, err := db.DB.Select(restaurantDish)
			if err != nil {
				return nil, err
			}
			err = surrealdb.Unmarshal(dr, &d)
			if err != nil {
				return nil, err
			}
			restaurant.DishesFields = append(restaurant.DishesFields, d)
		}
		final = append(final, restaurant)
	}
	return final, nil
}
