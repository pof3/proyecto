package models

import (
	"fmt"
	"log"
	"proyecto/db"
	"time"

	"github.com/surrealdb/surrealdb.go"
)

type Order struct {
	ID        string    `json:"id,omitempty"`
	Detail    string    `json:"detail,omitempty"`
	Price     uint      `json:"price,omitempty"`
	CreatedAt time.Time `json:"created_at"`
}

func PlaceOrder(userID string, detail string, price uint) (Order, error) {
	o := Order{
		Detail:    detail,
		Price:     price,
		CreatedAt: time.Now(),
	}
	data, err := db.DB.Create("order", o)
	if err != nil {
		return Order{}, err
	}
	order := make([]Order, 1)
	err = surrealdb.Unmarshal(data, &order)
	if err != nil {
		return Order{}, err
	}
	_, err = db.DB.Query("relate $userID ->ordered-> $orderID", map[string]string{
		"userID":  userID,
		"orderID": order[0].ID,
	})
	if err != nil {
		return Order{}, err
	}

	return order[0], nil
}

func GetOrder(userID string) ([]Order, error) {
	rawRes, err := db.DB.Query("select * from ordered where in = $userID", map[string]string{
		"userID": userID,
	})
	if err != nil {
		log.Fatal(err)
	}

	ordersIDs := make([]map[string]any, 0)
	_, err = surrealdb.UnmarshalRaw(rawRes, &ordersIDs)
	if err != nil {
		log.Fatal(err)
	}

	orders := make([]Order, 0)
	for _, orderID := range ordersIDs {
		o := new(Order)
		res, err := db.DB.Select(orderID["out"].(string))
		if err != nil {
			log.Fatal(err)
		}
		err = surrealdb.Unmarshal(res, &o)
		if err != nil {
			log.Fatal(err)
		}
		orders = append(orders, *o)
	}
	fmt.Printf("FROM HERE: %+v\n", orders)
	return orders, nil
}
