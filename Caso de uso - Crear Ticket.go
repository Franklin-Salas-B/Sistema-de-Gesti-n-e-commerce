Caso de uso - Crear Ticket
package application

import "ecommerce/internal/domain"

type TicketRepository interface {
	Save(ticket domain.Ticket) error
}

type CreateTicketUseCase struct {
	repo TicketRepository
}

func NewCreateTicketUseCase(r TicketRepository) CreateTicketUseCase {
	return CreateTicketUseCase{repo: r}
}

func (uc CreateTicketUseCase) Execute(
	id string,
	userID string,
	tType domain.TicketType,
	summary string,
) error {

	ticket := domain.OpenTicket(id, userID, tType, summary)
	return uc.repo.Save(ticket)
}
