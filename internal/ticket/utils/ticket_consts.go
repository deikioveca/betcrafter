package utils

type TicketStatus string
const (
	StatusPending 	TicketStatus	= "pending"
	StatusWon	 	TicketStatus	= "won"
	StatusLost		TicketStatus	= "lost"
	StatusCashout	TicketStatus	= "cashout"
)


type MatchResult string
const (
	MatchPending 	MatchResult 	= "pending"
	MatchCorrect 	MatchResult 	= "correct"
	MatchWrong   	MatchResult 	= "wrong"
)


type PickedOutcome string
const (
	HomeWin			PickedOutcome	= "home_win"
	AwayWin			PickedOutcome	= "away_win"
	Draw			PickedOutcome	= "draw"
	BTTS			PickedOutcome	= "btts"
	Over2_5Goals	PickedOutcome	= "over_2_5_goals"
	Under2_5Goals   PickedOutcome   = "under_2_5_goals"
	Over9_5Corners	PickedOutcome	= "over_9_5_corners"
	Under9_5Corners PickedOutcome   = "under_9_5_corners"
)


type League string
const (
	PremierLeague 		League 	= 	"premier_league"
	LaLiga		  		League 	= 	"la_liga"
	SerieA        		League 	= 	"serie_a"
	Bundesliga    		League 	= 	"bundesliga"
	FrenchLeagueOne	  	League 	= 	"french_league_one"
	ChampionsLeague		League 	= 	"champions_league"
	EuropeLeague 		League 	= 	"europe_league"
	ConferenceLeague	League 	= 	"conference_league"
)