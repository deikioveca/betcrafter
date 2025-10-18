package service

import (
	"fmt"
	"sort"

	"github.com/deikioveca/betcrafter/internal/ticket/model"
	"github.com/deikioveca/betcrafter/internal/ticket/utils"
)

type Analyzer interface {
	GetTicketStats(ticketStatsRequest *model.TicketStatsRequest) (*model.TicketStats, error)

    GetPickedOutcomeStats(ticketStatsRequest *model.TicketStatsRequest) (map[string]*model.PickedOutcomeStats, error)

    GetPickedOutcomeOddRangeStats(ticketStatsRequest *model.TicketStatsRequest) (map[utils.PickedOutcome]map[string]*model.PickedOutcomeOddRangeStats, error)

    GetMostProfitablePickTypes(ticketStatsRequest *model.TicketStatsRequest) ([]*model.MostProfitablePickType, error)
}


func (t *ticketService) loadTicketsForStats(req *model.TicketStatsRequest) ([]model.Ticket, error) {
    var tickets []model.Ticket
    query := t.db

    if req.StartDate != nil && req.EndDate != nil {
        query = query.Where("created_at BETWEEN ? AND ?", req.StartDate.ToTime(), req.EndDate.ToTime())
    }

    if err := query.Preload("Matches").Find(&tickets).Error; err != nil {
        return nil, err
    }

    return tickets, nil
}


func calculateTicketStats(tickets []model.Ticket) *model.TicketStats {
    stats := &model.TicketStats{}

    for _, ticket := range tickets {
        stats.TotalTickets++
        stats.TotalStake += ticket.Stake
        switch ticket.Status {
        case utils.StatusWon:
            stats.WonCount++
            stats.TotalProfit += ticket.ActualWin - ticket.Stake
        case utils.StatusCashout:
            stats.CashOutCount++
            stats.TotalProfit += ticket.ActualWin - ticket.Stake
        case utils.StatusLost:
            stats.LostCount++
            stats.TotalProfit -= ticket.Stake
        case utils.StatusPending:
            stats.PendingCount++
        }
    }

    if stats.TotalStake > 0 {
        stats.ROI = (stats.TotalProfit / stats.TotalStake) * 100
    }

    settled := stats.WonCount + stats.CashOutCount + stats.LostCount
    if settled > 0 {
        stats.HitRate = (float64(stats.WonCount) / float64(settled)) * 100
    }

    if stats.TotalTickets > 0 {
        stats.AvgStake = stats.TotalStake / float64(stats.TotalTickets)
    }

    return stats
}


func (t *ticketService) GetTicketStats(ticketStatsRequest *model.TicketStatsRequest) (*model.TicketStats, error) {
    tickets, err := t.loadTicketsForStats(ticketStatsRequest)
    if err != nil {
        return nil, err
    }

    stats := calculateTicketStats(tickets)
    return stats, nil
}


func (t *ticketService) GetPickedOutcomeStats(ticketStatsRequest *model.TicketStatsRequest) (map[string]*model.PickedOutcomeStats, error) {
    tickets, err := t.loadTicketsForStats(ticketStatsRequest)
    if err != nil {
        return nil, err
    }

    stats := make(map[string]*model.PickedOutcomeStats)
    for _, outcome := range []utils.PickedOutcome{utils.HomeWin, utils.AwayWin, utils.Draw, utils.BTTS, utils.Over2_5Goals, utils.Over9_5Corners, utils.Under2_5Goals, utils.Under9_5Corners} {
        stats[string(outcome)] = &model.PickedOutcomeStats{}
    }

    for _, ticket := range tickets {
        for _, match := range ticket.Matches {
            s := stats[string(match.PickedOutcome)]

            switch match.Result {
            case utils.MatchCorrect:
                s.Wins++
            case utils.MatchWrong:
                s.Losses++
            }
        }
    }

    for _, s := range stats {
        total := s.Wins + s.Losses
        if total > 0 {
            s.WinRate = (float64(s.Wins) / float64(total)) * 100
        }
    }

    return stats, nil
}


type OddRange struct {
    Min float64
    Max float64
}

var OddRanges = []OddRange{
    {1.25, 1.50},
    {1.51, 1.76},
    {1.77, 2.00},
    {2.01, 2.26},
    {2.27, 2.52},
    {2.53, 2.75},
    {2.76, 3.01},
    {3.02, 3.27},
    {3.28, 3.53},
    {3.54, 3.79},
    {3.80, 4.05},
    {4.06, 4.31},
    {4.32, 4.57},
    {4.58, 4.83},
    {4.84, 5.00},
}


func formatOddRange(r OddRange) string {
    return fmt.Sprintf("%.2f-%.2f", r.Min, r.Max)
}

func (t *ticketService) GetPickedOutcomeOddRangeStats(ticketStatsRequest *model.TicketStatsRequest) (map[utils.PickedOutcome]map[string]*model.PickedOutcomeOddRangeStats, error) {
    tickets, err := t.loadTicketsForStats(ticketStatsRequest)
    if err != nil {
        return nil, err
    }

    result := make(map[utils.PickedOutcome]map[string]*model.PickedOutcomeOddRangeStats)
    outcomes := []utils.PickedOutcome{utils.HomeWin, utils.AwayWin, utils.Draw,utils.BTTS, utils.Over2_5Goals, utils.Over9_5Corners, utils.Under2_5Goals, utils.Under9_5Corners}

    for _, outcome := range outcomes {
        result[outcome] = make(map[string]*model.PickedOutcomeOddRangeStats)
        for _, r := range OddRanges {
            key := formatOddRange(r)
            result[outcome][key] = &model.PickedOutcomeOddRangeStats{}
        }
    }

    for _, ticket := range tickets {
        for _, match := range ticket.Matches {

            outcome := match.PickedOutcome
            for _, r := range OddRanges {
                if match.Odd >= r.Min && match.Odd <= r.Max {
                    key := formatOddRange(r)
                    s := result[outcome][key]
                    
                    s.Total++
                    s.TotalOdds += match.Odd

                    switch match.Result {
                    case utils.MatchCorrect:
                        s.Wins++
                    case utils.MatchWrong:
                        s.Losses++
                    }

                    if s.Total > 0 {
                        s.WinRate = (float64(s.Wins) / float64(s.Total)) * 100
                        s.AvgOdd = s.TotalOdds / float64(s.Total)
                    }
                    break
                }
            }
        }
    }

    return result, nil
}


func (t *ticketService) GetMostProfitablePickTypes(ticketStatsRequest *model.TicketStatsRequest) ([]*model.MostProfitablePickType, error) {
    tickets, err := t.loadTicketsForStats(ticketStatsRequest)
    if err != nil {
        return nil, err
    }

    type tempStats struct {
        Wins, Losses int64
        Profit       float64
    }

    statsMap := make(map[utils.PickedOutcome]*tempStats)

    for _, ticket := range tickets {
        if len(ticket.Matches) == 0 {
            continue
        }

        perMatchStake := ticket.Stake / float64(len(ticket.Matches))

        for _, match := range ticket.Matches {
            s, exists := statsMap[match.PickedOutcome]
            if !exists {
                s = &tempStats{}
                statsMap[match.PickedOutcome] = s
            }

            switch match.Result {
            case utils.MatchCorrect:
                s.Wins++
                s.Profit += (match.Odd - 1) * perMatchStake
            case utils.MatchWrong:
                s.Losses++
                s.Profit -= perMatchStake
            }
        }
    }

    var results []*model.MostProfitablePickType
    for outcome, s := range statsMap {
        total := s.Wins + s.Losses
        winRate := 0.0
        if total > 0 {
            winRate = (float64(s.Wins) / float64(total)) * 100
        }

        results = append(results, &model.MostProfitablePickType{
            PickedOutcome: string(outcome),
            TotalProfit:   s.Profit,
            Wins:          s.Wins,
            Losses:        s.Losses,
            Total:         total,
            WinRate:       winRate,
        })
    }

    sort.Slice(results, func(i, j int) bool {
        return results[i].TotalProfit > results[j].TotalProfit
    })

    if len(results) > 3 {
        results = results[:3]
    }

    return results, nil
}