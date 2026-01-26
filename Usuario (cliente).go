Usuario (cliente)
package domain

type User struct {
	ID        string
	Name      string
	Email     string
	IsActive  bool
}

func NewUser(id, name, email string) User {
	return User{
		ID:       id,
		Name:     name,
		Email:    email,
		IsActive: true,
	}
}

