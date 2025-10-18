package service

import (
	"errors"

	"github.com/deikioveca/betcrafter/internal/ticket/model"
	"github.com/deikioveca/betcrafter/internal/ticket/utils"
	"gorm.io/gorm"
)

type Creator interface {
	CreateTicket(ticketRequest *model.TicketRequest) (*model.Ticket, error)
}


func buildTicketMatches(matchRequests []model.MatchRequest) ([]model.TicketMatch, float64, error) {
	matches := make([]model.TicketMatch, 0, len(matchRequests))
	totalOdds := 1.0

	for _, match := range matchRequests {
		if err := ValidateLeague(match.League); err != nil {
			return nil, 0, err
		}

		if err := ValidatePickedOutcome(match.PickedOutcome); err != nil {
			return nil, 0, err
		}

		totalOdds *= match.Odd

		matches = append(matches, model.TicketMatch{
			League: 	   match.League,	
			HomeTeam:      match.HomeTeam,
			AwayTeam:      match.AwayTeam,
			PickedOutcome: match.PickedOutcome,
			Odd:           match.Odd,
			Arguments:     match.Arguments,
			Result:        utils.MatchPending,
		})
	}

	return matches, totalOdds, nil
}


func (t *ticketService) loadTicketWithMatches(ticketID uint) (*model.Ticket, error) {
	var ticket model.Ticket
	if err := t.db.Preload("Matches").First(&ticket, ticketID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, utils.ErrTicketNotFound
		}
		return nil, err
	}

	return &ticket, nil
}


func (t *ticketService) saveAndLoadTicket(ticket *model.Ticket) (*model.Ticket, error) {
	if err := t.db.Create(&ticket).Error; err != nil {
		return nil, err
	}

	return t.loadTicketWithMatches(ticket.ID)
}


func (t *ticketService) CreateTicket(ticketRequest *model.TicketRequest) (*model.Ticket, error) {
	if err := validateTicketRequest(ticketRequest); err != nil {
		return nil, err
	}

	matches, totalOdds, err := buildTicketMatches(ticketRequest.Matches)
	if err != nil {
		return nil, err
	}

	ticket := model.Ticket{
		Stake: 			ticketRequest.Stake,
		TotalOdds: 		totalOdds,
		PossibleWin: 	ticketRequest.Stake * totalOdds,
		ActualWin: 		0,
		CashOut: 		false,
		Status: 		utils.StatusPending,
		Matches: 		matches,	
	}

	return t.saveAndLoadTicket(&ticket)
}