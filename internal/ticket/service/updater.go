package service

import (
	"github.com/deikioveca/betcrafter/internal/ticket/model"
	"github.com/deikioveca/betcrafter/internal/ticket/utils"
)

type Updater interface {
	UpdatePendingTicket(updateTicketRequest *model.UpdateTicketRequest) (*model.Ticket, error)

	UpdateTicketDate(updateTicketDateRequest *model.UpdateTicketDateRequest) (*model.Ticket, error)
}


func (t *ticketService) updateTicketFields(ticket *model.Ticket) error {
	return t.db.Model(ticket).Updates(map[string]any{
		"actual_win": 	ticket.ActualWin,
		"cash_out": 	ticket.CashOut,
		"status": 		ticket.Status,
	}).Error
}


func (t *ticketService) updateTicketMatches(ticketID uint, matchUpdates []model.UpdateMatchResult) error {
    for _, match := range matchUpdates {
		if err := ValidateMatchResult(match.Result); err != nil {
			return err
		}

        err := t.db.Model(&model.TicketMatch{}).
            Where("id = ? AND ticket_id = ?", match.MatchID, ticketID).
            Update("result", string(match.Result))

		if err.Error != nil  {
			return err.Error
		}

		if err.RowsAffected == 0 {
			return utils.ErrTicketMatchNotFound
		}
		
    }
    return nil
}


func (t *ticketService) UpdatePendingTicket(updateTicketRequest *model.UpdateTicketRequest) (*model.Ticket, error) {
	ticket, err := t.loadTicketWithMatches(updateTicketRequest.TicketID)
	if err != nil {
		return nil, err
	}

	if err := ValidateTicketStatus(updateTicketRequest.Status); err != nil {
		return nil, err
	}

	ticket.ActualWin = 	updateTicketRequest.ActualWin
    ticket.CashOut 	 = 	updateTicketRequest.CashOut
    ticket.Status 	 = 	updateTicketRequest.Status

	if err := t.updateTicketFields(ticket); err != nil {
		return nil, err
	}

	if err := t.updateTicketMatches(ticket.ID, updateTicketRequest.Matches); err != nil {
		return nil, err
	}

	return t.loadTicketWithMatches(ticket.ID)
}


func (t *ticketService) UpdateTicketDate(updateTicketDateRequest *model.UpdateTicketDateRequest) (*model.Ticket, error) {
	ticket, err := t.loadTicketWithMatches(updateTicketDateRequest.TicketID)
	if err != nil {
		return nil, err
	}

	newDate := updateTicketDateRequest.NewDate.ToTime()
	ticket.CreatedAt = newDate

	if err := t.db.Save(&ticket).Error; err != nil {
		return nil, err
	}

	return ticket, nil
}