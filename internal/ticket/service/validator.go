package service

import (
	"github.com/deikioveca/betcrafter/internal/ticket/model"
	"github.com/deikioveca/betcrafter/internal/ticket/utils"
)

func ValidateMatchResult(result utils.MatchResult) error {
	switch result {
	case utils.MatchCorrect, utils.MatchWrong, utils.MatchPending:
		return nil
	default:
		return utils.ErrInvalidMatchResult
	}
}


func ValidateLeague(league utils.League) error {
	switch league {
	case utils.PremierLeague, utils.LaLiga, utils.SerieA, utils.Bundesliga, utils.FrenchLeagueOne, utils.ChampionsLeague, utils.EuropeLeague, utils.ConferenceLeague:
		return nil
	default:
		return utils.ErrInvalidLeague
	}
}


func ValidatePickedOutcome(outcome utils.PickedOutcome) error {
    switch outcome {
	case utils.HomeWin, utils.AwayWin, utils.Draw, utils.BTTS, utils.Over2_5Goals, utils.Over9_5Corners, utils.Under2_5Goals, utils.Under9_5Corners:
        return nil
    default:
        return utils.ErrInvalidOutcome
    }
}


func ValidateTicketStatus(status utils.TicketStatus) error {
	switch status {
		case utils.StatusWon, utils.StatusLost, utils.StatusCashout, utils.StatusPending:
			return nil
		default:
			return utils.ErrInvalidStatus
	}
}


func validateTicketRequest(ticketRequest *model.TicketRequest) error {
	if ticketRequest.Stake <= 0 {
		return utils.ErrInvalidStake
	}

	if len(ticketRequest.Matches) == 0 {
		return utils.ErrInvalidMatchesCount
	}

	return nil
}