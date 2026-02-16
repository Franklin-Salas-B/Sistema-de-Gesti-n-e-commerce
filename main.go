package main

import (
	"fmt"

	"ecommerce/models"
	"ecommerce/process"
)

func main() {

	user, err := models.NewUser(1, "Franklin", "franklin@mail.com", "1234")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println(user.GetInfo())

	err = user.Login("franklin@mail.com", "1234")
	if err != nil {
		fmt.Println("Login fallido")
	} else {
		fmt.Println("Login exitoso")
	}

	ticket := models.Ticket{
		ID:     1,
		Title:  "Sin conexión",
		Status: "Abierto",
		UserID: user.ID,
	}

	fmt.Println(ticket.GetInfo())

	invoice, schedule, err :=
		process.ProcessTicketResolution(&ticket, 50.00)

	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println(ticket.GetInfo())
	fmt.Println(invoice.GetInfo())
	fmt.Println(schedule.GetInfo())

	err = invoice.Pay()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Factura pagada correctamente")
	fmt.Println(invoice.GetInfo())
}
