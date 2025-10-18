package cli_app

import (
	"fmt"
	"strings"

	"github.com/deikioveca/betcrafter/internal/ticket/model"
	"github.com/deikioveca/betcrafter/internal/ticket/utils"
)


func (c *CliClient) printHelpfulTicketInfo() {
	fmt.Println("Leagues:		premier_league, la_liga, serie_a, bundesliga, french_league_one, champions_league, europe_league, conference_league")

	fmt.Println("Outcome:		home_win, away_win, draw, btts, over_2_5_goals, over_9_5_corners, under_2_5_goals, under_9_5_corners")

	fmt.Println("Match result:		pending, correct, wrong")

	fmt.Println("Ticket status:		pending, won, lost, cashout")
}


func (c *CliClient) readInput(prompt string) string {
	fmt.Print(prompt)
	input, _ := c.Reader.ReadString('\n')
	return strings.TrimSpace(input)
}


func (c *CliClient) readTicketStatsDateRange() (*model.TicketStatsRequest, error) {
	var startDate utils.Date
	for {
		startDateInput := c.readInput("Enter start date: ")
		if err := startDate.UnmarshalJSON([]byte(`"` + startDateInput + `"`)); err != nil {
			fmt.Println("Invalid date format. Use Year-Month-Day.")
			continue
		}
		break
	}

	var endDate utils.Date
	for {
		endDateInput := c.readInput("Enter end date: ")
		if err := endDate.UnmarshalJSON([]byte(`"` + endDateInput + `"`)); err != nil {
			fmt.Println("Invalid date format. User Year-Month-Day.")
			continue
		}
		break
	}

	return &model.TicketStatsRequest{StartDate: &startDate, EndDate: &endDate}, nil
}