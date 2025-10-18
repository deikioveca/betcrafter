package model

import (
	"github.com/deikioveca/betcrafter/internal/ticket/utils"
	"gorm.io/gorm"
)

type Ticket struct {
	gorm.Model
	Stake			float64					`gorm:"not null"`
	TotalOdds		float64					`gorm:"not null"`
	PossibleWin		float64					`gorm:"not null"`
	ActualWin		float64					`gorm:"default:0"`
	CashOut			bool					`gorm:"default:false"`
	Status			utils.TicketStatus		`gorm:"type:varchar(20);default:'pending';index"`
	Matches			[]TicketMatch			`gorm:"foreignKey:TicketID;constraint:OnDelete:CASCADE"`
}


type TicketMatch struct {
	gorm.Model
	TicketID		uint				`gorm:"index;not null"`
	League			utils.League		`gorm:"type:varchar(64);not null;index"`
	HomeTeam		string				`gorm:"type:varchar(64);not null"`
	AwayTeam		string				`gorm:"type:varchar(64);not null"`
	PickedOutcome	utils.PickedOutcome	`gorm:"type:varchar(64);not null"`
	Odd				float64				`gorm:"not null"`
	Arguments		string				`gorm:"not null"`
	Result			utils.MatchResult	`gorm:"type:varchar(20);default:'pending'"`
}