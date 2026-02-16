package models

import (
	"errors"
	"fmt"
)

type User struct {
	ID       int
	Name     string
	Email    string
	password string
}

func NewUser(id int, name, email, password string) (*User, error) {
	if email == "" {
		return nil, errors.New("email no puede estar vacío")
	}

	return &User{
		ID:       id,
		Name:     name,
		Email:    email,
		password: password,
	}, nil
}

func (u User) Login(email, password string) error {
	if u.Email != email || u.password != password {
		return errors.New("credenciales incorrectas")
	}
	return nil
}

func (u User) GetInfo() string {
	return fmt.Sprintf("User: %s | Email: %s", u.Name, u.Email)
}
