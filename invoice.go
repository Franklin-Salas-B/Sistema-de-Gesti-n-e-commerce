package models

import (
	"errors"
	"fmt"
)

type Invoice struct {
	ID     int
	UserID int
	Amount float64
	Paid   bool
}

func (i *Invoice) Pay() error {
	if i.Paid {
		return errors.New("la factura ya fue pagada")
	}
	i.Paid = true
	return nil
}

func (i Invoice) GetInfo() string {
	return fmt.Sprintf("Invoice ID: %d | Amount: %.2f | Paid: %t",
		i.ID, i.Amount, i.Paid)
}
