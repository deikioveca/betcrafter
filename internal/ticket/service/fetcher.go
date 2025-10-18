package service

import (
	"github.com/deikioveca/betcrafter/internal/ticket/model"
	"github.com/deikioveca/betcrafter/internal/ticket/utils"
)

type Fetcher interface {
	GetPendingTickets() ([]model.Ticket, error)

	GetTicketByID(ticketID uint) (*model.Ticket, error)
	
	GetTicketsByStatus(status utils.TicketStatus) ([]model.Ticket, error)
}

func (t *ticketService) GetTicketByID(ticketID uint) (*model.Ticket, error) {
	return t.loadTicketWithMatches(ticketID)
}


func (t *ticketService) GetPendingTickets() ([]model.Ticket, error) {
	var tickets []model.Ticket
	if err := t.db.Preload("Matches").Where("status = ?", utils.StatusPending).Find(&tickets).Error; err != nil {
		return nil, err
	}

	return tickets, nil
}


func (t *ticketService) GetTicketsByStatus(status utils.TicketStatus) ([]model.Ticket, error) {
	if err := ValidateTicketStatus(status); err != nil {
		return nil, err
	}

	var tickets []model.Ticket
	if err := t.db.Preload("Matches").Where("status = ?", status).Find(&tickets).Error; err != nil {
		return nil, err
	}

	return tickets, nil
}