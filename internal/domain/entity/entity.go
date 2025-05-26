package entity

import "errors"

type User struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func NewUser(id, name string, age int) (User, error) {
	if name == "" {
		return User{}, errors.New("name is required")
	}

	if age <= 0 {
		return User{}, errors.New("age must be greater than zero")
	} else if age > 120 {
		return User{}, errors.New("age must be less than or equal to 120")
	}

	return User{
		ID:   id,
		Name: name,
		Age:  age,
	}, nil
}
