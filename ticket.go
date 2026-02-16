package models

import (
	"errors"
	"fmt"
)

type Ticket struct {
	ID     int
	Title  string
	Status string
	UserID int
}

func (t *Ticket) UpdateStatus(newStatus string) error {
	if t.Status == "Cerrado" {
		return errors.New("no se puede modificar un ticket cerrado")
	}
	t.Status = newStatus
	return nil
}

func (t Ticket) GetInfo() string {
	return fmt.Sprintf("Ticket: %s | Status: %s", t.Title, t.Status)
}
