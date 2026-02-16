package models

import (
	"fmt"
	"time"
)

type Schedule struct {
	ID       int
	TicketID int
	Date     time.Time
}

func (s Schedule) GetInfo() string {
	return fmt.Sprintf("Schedule for Ticket %d on %s",
		s.TicketID,
		s.Date.Format("2006-01-02 15:04"))
}
