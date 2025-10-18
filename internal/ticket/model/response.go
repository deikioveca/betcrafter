package model

type TicketStats struct {
	TotalTickets int64   `json:"total_tickets"`
	WonCount     int64   `json:"won_count"`
	LostCount    int64   `json:"lost_count"`
	PendingCount int64   `json:"pending_count"`
	CashOutCount int64   `json:"cashout_count"`
	TotalStake   float64 `json:"total_stake"`
	TotalProfit  float64 `json:"total_profit"`
	ROI          float64 `json:"roi"`
	HitRate      float64 `json:"hit_rate"`
	AvgStake     float64 `json:"average_stake"`
}

type PickedOutcomeStats struct {
	Wins    int64   `json:"wins"`
	Losses  int64   `json:"losses"`
	WinRate float64 `json:"win_rate"`
}

type PickedOutcomeOddRangeStats struct {
	Wins        int64       `json:"wins"`
	Losses      int64       `json:"losses"`
    Total       int64       `json:"total"`
    WinRate     float64     `json:"win_rate"`
    AvgOdd      float64     `json:"average_odd"`
    TotalOdds   float64     `json:"-"`
}

type MostProfitablePickType struct {
	PickedOutcome	string		`json:"picked_outcome"`
	TotalProfit		float64		`json:"total_profit"`
	Wins          	int64   	`json:"wins"`
    Losses        	int64   	`json:"losses"`
    Total         	int64   	`json:"total"`
    WinRate       	float64 	`json:"win_rate"`
}