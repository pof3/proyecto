package models

import (
	"errors"
	"github.com/surrealdb/surrealdb.go"
	"golang.org/x/crypto/bcrypt"
	"proyecto/db"
	"time"
)

type Person struct {
	ID        string    `json:"id,omitempty"`
	Name      string    `json:"name,omitempty"`
	Password  string    `json:"password,omitempty"`
	CreatedAt time.Time `json:"created_at,omitempty"`
}

func CreatePerson(name string, password string) (Person, error) {
	// Check in database if user exists with that name
	result, err := db.DB.Query("SELECT * FROM person WHERE name = $name;", map[string]string{
		"name": name,
	})
	if err != nil {
		return Person{}, err
	}
	p := make([]Person, 1)
	_, err = surrealdb.UnmarshalRaw(result, &p)
	if err != nil {
		return Person{}, err
	}
	if p[0].ID != "" {
		// if it's not a new record, check password
		err = bcrypt.CompareHashAndPassword([]byte(p[0].Password), []byte(password))
		if err != nil {
			return Person{}, errors.New("invalid password")
		}
		return p[0], nil
	}

	// User is new, proceed
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)

	person := Person{
		Name:      name,
		Password:  string(bytes),
		CreatedAt: time.Now(),
	}
	newRes, err := db.DB.Create("person", person)
	if err != nil {
		return Person{}, err
	}

	newPerson := make([]Person, 1)

	err = surrealdb.Unmarshal(newRes, &newPerson)
	if err != nil {
		return Person{}, err
	}

	return newPerson[0], nil
}
