package utils

import "fmt"

type TicketError struct {
	Code    string	`json:"code"`
	Message string	`json:"message"`
}

func (t *TicketError) Error() string {
	return fmt.Sprintf("[%s]:  %s", t.Code, t.Message)
}

var (
	ErrInvalidStake 			= &TicketError{Code: "INVALID_STAKE", 			Message: "stake must be greater than 0"}
	ErrInvalidMatchesCount 		= &TicketError{Code: "INVALID_MATCH_COUNT", 	Message: "match count must be atleast 1"}
	ErrInvalidLeague			= &TicketError{Code: "INVALID_LEAUGE", 			Message: "picked league is invalid"}
	ErrInvalidOutcome			= &TicketError{Code: "INVALID_OUTCOME", 		Message: "picked outcome is invalid"}
	ErrInvalidStatus			= &TicketError{Code: "INVALID_STATUS", 			Message: "picked status is invalid"}
	ErrInvalidMatchResult		= &TicketError{Code: "INVALID_MATCH_RESULT", 	Message: "picked match result is invalid"}

	ErrTicketNotFound			= &TicketError{Code: "INVALID_TICKET_ID", 		Message: "ticket not found"}
	ErrTicketMatchNotFound		= &TicketError{Code: "INVALID_TICKET_MATCH_ID", Message: "ticket match not found"}

	ErrZeroTickets				= &TicketError{Code: "INVALID_TICKET_COUNT", 	Message: "0 tickets found"}
)