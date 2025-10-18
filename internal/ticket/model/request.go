package model

import "github.com/deikioveca/betcrafter/internal/ticket/utils"

type TicketRequest struct {
	Stake   float64        `json:"stake"`
	Matches []MatchRequest `json:"matches"`
}


type MatchRequest struct {
	League        utils.League        `json:"league"`
	HomeTeam      string              `json:"home_team"`
	AwayTeam      string              `json:"away_team"`
	PickedOutcome utils.PickedOutcome `json:"picked_outcome"`
	Odd           float64             `json:"odd"`
	Arguments     string              `json:"arguments"`
}


type UpdateMatchResult struct {
	MatchID uint              `json:"match_id"`
	Result  utils.MatchResult `json:"result"`
}


type UpdateTicketRequest struct {
	TicketID  uint                `json:"ticket_id"`
	ActualWin float64             `json:"actual_win"`
	CashOut   bool                `json:"cash_out"`
	Status    utils.TicketStatus  `json:"status"`
	Matches   []UpdateMatchResult `json:"matches"`
}


type TicketStatsRequest struct {
	StartDate *utils.Date `json:"start_date"`
	EndDate   *utils.Date `json:"end_date"`
}


type UpdateTicketDateRequest struct {
	TicketID	uint		`json:"ticket_id"`
	NewDate		utils.Date	`json:"new_date"`
}