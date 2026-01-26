Ticket de Soporte
package domain

type TicketType string
type TicketStatus string

const (
	Incident    TicketType = "INCIDENTE"
	Request     TicketType = "REQUERIMIENTO"

	Open        TicketStatus = "ABIERTO"
	InProgress  TicketStatus = "EN_PROCESO"
	Closed      TicketStatus = "CERRADO"
)

type Ticket struct {
	ID       string
	UserID   string
	Type     TicketType
	Status   TicketStatus
	Summary  string
}

func OpenTicket(id, userID string, tType TicketType, summary string) Ticket {
	return Ticket{
		ID:      id,
		UserID:  userID,
		Type:    tType,
		Status:  Open,
		Summary: summary,
	}
}
