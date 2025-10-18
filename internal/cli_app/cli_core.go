package cli_app

import (
	"fmt"
	"strconv"

	"github.com/deikioveca/betcrafter/internal/ticket/model"
	"github.com/deikioveca/betcrafter/internal/ticket/service"
	"github.com/deikioveca/betcrafter/internal/ticket/utils"
)

func (c *CliClient) Core() {
	fmt.Println("\n==== Core menu ====")

	fmt.Println("0 - Back")
	fmt.Println("1 - Create ticket")
	fmt.Println("2 - Get pending tickets")
	fmt.Println("3 - Update pending ticket")
	fmt.Println("4 - Get tickets by status")
	fmt.Println("5 - Update ticket date")
	fmt.Println("6 - Get ticket by id")

	choice := c.readInput("Enter your choice: ")

	switch choice {
	case "0":
		return
	case "1":
		if err := c.CreateTicket(); err != nil {
			fmt.Println("Error:", err)
		}
	case "2":
		if err := c.GetPendingTickets(); err != nil {
			fmt.Println("Error:", err)
		}
	case "3":
		if err := c.UpdatePendingTicket(); err != nil {
			fmt.Println("Error:", err)
		}
	case "4":
		if err := c.GetTicketsByStatus(); err != nil {
			fmt.Println("Error:", err)
		}
	case "5":
		if err := c.UpdateTicketDate(); err != nil {
			fmt.Println("Error:", err)
		}
	case "6":
		if err := c.GetTicketById(); err != nil {
			fmt.Println("Error:", err)
		}
	default:
		fmt.Println("Invalid choice, try again.")
	}
}


func (c *CliClient) CreateTicket() error {
	fmt.Println("\n---- Create ticket ----")

	var stake float64
	for {
		stakeStr := c.readInput("Enter stake (must be bigger than 0): ")
		s, err := strconv.ParseFloat(stakeStr, 64)
		if err != nil || s <= 0 {
			fmt.Println("Invalid stake, try again.")
			continue
		}
		stake = s
		break
	}

	var numberOfMatches int
	for {
		numStr := c.readInput("Enter number of matches (atleast 1): ")
		n, err := strconv.Atoi(numStr)
		if err != nil || n <= 0 {
			fmt.Println("Invalid number of matches, try again.")
			continue
		}
		numberOfMatches = n
		break
	}

	matches := make([]model.MatchRequest, 0, numberOfMatches)
	for i := 1; i <= numberOfMatches; i++ {
		fmt.Printf("\n---- Match %d ----\n", i)

		var matchLeague utils.League
		for {
			league := c.readInput("League: ")
			if err := service.ValidateLeague(utils.League(league)); err != nil {
				fmt.Println(err)
				continue
			}
			matchLeague = utils.League(league)
			break
		}
		
		homeTeam := c.readInput("Home team: ")

		awayTeam := c.readInput("Away team: ")

		var outcome utils.PickedOutcome
		for {
			o := c.readInput("Outcome: ")
			if err := service.ValidatePickedOutcome(utils.PickedOutcome(o)); err != nil {
				fmt.Println(err)
				continue
			}
			outcome = utils.PickedOutcome(o)
			break
		}

		var odd float64
		for {
			oddStr := c.readInput("Odd: ")
			o, err := strconv.ParseFloat(oddStr, 64)
			if err != nil || o <= 0 {
				fmt.Println("Invalid odd, try again.")
				continue
			}
			odd = o
			break
		}

		arguments := c.readInput("Arguments: ")

		matches = append(matches, model.MatchRequest{
			League: 		matchLeague,
			HomeTeam: 		homeTeam,
			AwayTeam: 		awayTeam,
			PickedOutcome: 	outcome,
			Odd: 			odd,
			Arguments: 		arguments,
		})
	}

	ticketReq := &model.TicketRequest{
		Stake: 		stake,
		Matches: 	matches,
	}

	ticket, err := c.service.CreateTicket(ticketReq)
	if err != nil {
		return err
	}

	fmt.Printf("\nTicket created successfully! ID=%d | Stake=%.2f | PossibleWin=%.2f\n",
        ticket.ID, ticket.Stake, ticket.PossibleWin)

	return nil
}


func (c *CliClient) GetPendingTickets() error {
	fmt.Println("\n---- Pending tickets ----")

	tickets, err := c.service.GetPendingTickets()
	if err != nil {
		return err
	}

	if len(tickets) == 0 {
		fmt.Println("No pending tickets")
		return nil
	}

	for _, t := range tickets {
		fmt.Printf("\nTicketID: %d | Stake: %.2f | Total odds: %.2f | Possible win: %.2f\n", t.ID, t.Stake, t.TotalOdds, t.PossibleWin)
		fmt.Println("----------------------------------")

		for i, m := range t.Matches {
			fmt.Printf("Match %d\n", i+1)
			fmt.Printf("  ID: %d\n", m.ID)
			fmt.Printf("  League: %s\n", m.League)
			fmt.Printf("  %s vs %s\n", m.HomeTeam, m.AwayTeam)
			fmt.Printf("  Picked outcome: %s | Odd: %.2f\n", m.PickedOutcome, m.Odd)
			fmt.Printf("  Arguments: %s\n", m.Arguments)
			fmt.Printf("  Result: %s\n", m.Result)
		}
		fmt.Println("Status:", t.Status)
		fmt.Println("==================================")
	}

	return nil
}


func (c *CliClient) UpdatePendingTicket() error {
	fmt.Println("\n---- Update pending ticket ----")

	var ticketID uint
	for {
		idStr := c.readInput("Enter ticket ID: ")
		id, err := strconv.Atoi(idStr)
		if err != nil || id <= 0 {
			fmt.Println("Invalid ticket id, try again.")
			continue
		}
		_, err = c.service.GetTicketByID(uint(id))
		if err != nil {
			fmt.Println("Error:", err)
			continue
		}
		ticketID = uint(id)
		break
	}

	var actualWin float64
	for {
		aWinInput := c.readInput("Actual win: ")
		actWin, err := strconv.ParseFloat(aWinInput, 64)
		if err != nil || actWin < 0 {
			fmt.Println("Invalid actual win input, try again.")
			continue
		}
		actualWin = actWin
		break
	}

	var cashout bool
	for {
		cashOutInput := c.readInput("Cashout(true || false): ")
		switch cashOutInput {
		case "true":
			cashout = true
		case "false":
			cashout = false
		default:
			fmt.Println("Invalid cashout input, try again.")
			continue
		}
		break
	}

	var status utils.TicketStatus
	for {
		statusInput := c.readInput("Status: ")
		if err := service.ValidateTicketStatus(utils.TicketStatus(statusInput)); err != nil {
			fmt.Println(err)
			continue
		}
		status = utils.TicketStatus(statusInput)
		break
	}

	var numberOfMatches int
	var matches []model.UpdateMatchResult
	for {
		numMatchesInput := c.readInput("Number of matches: ")
		numMatches, err := strconv.Atoi(numMatchesInput)
		if err != nil || numMatches <= 0 {
			fmt.Println("Invalid number of matches input, try again")
			continue
		}
		numberOfMatches = numMatches
		matches = make([]model.UpdateMatchResult, 0, numberOfMatches)

		for i := 1; i <= numberOfMatches; i++ {
			fmt.Printf("\n---- Match %d ----\n", i)

			var matchID uint
			for {
				matchIdInput := c.readInput("Match ID: ")
				matchiD, err := strconv.Atoi(matchIdInput)
				if err != nil || matchiD <= 0 {
					fmt.Println("Invalid match id, try again.")
					continue
				}
				matchID = uint(matchiD)
				break
			}

			var result utils.MatchResult
			for {
				matchResultInput := c.readInput("Match result: ")
				if err := service.ValidateMatchResult(utils.MatchResult(matchResultInput)); err != nil {
					fmt.Println(err)
					continue
				}
				result = utils.MatchResult(matchResultInput)
				break
			}

			matches = append(matches, model.UpdateMatchResult{
				MatchID: 	matchID,
				Result: 	result,
			})
		}
		break
	}

	updateReq := model.UpdateTicketRequest{
		TicketID: 		ticketID,
		ActualWin: 		actualWin,
		CashOut: 		cashout,
		Status: 		status,
		Matches: 		matches,
	}

	updatedTicket, err := c.service.UpdatePendingTicket(&updateReq)
	if err != nil {
		return fmt.Errorf("update pending ticket failed: %w", err)
	}

	fmt.Printf("\n✅ Ticket #%d updated successfully! Status: %s | Actual Win: %.2f | CashOut: %t\n",
		updatedTicket.ID, updatedTicket.Status, updatedTicket.ActualWin, updatedTicket.CashOut)

	return nil
}


func (c *CliClient) GetTicketsByStatus() error {
	fmt.Println("\n---- Get Tickets By Status ----")

	for {
		status := c.readInput("Status: ")
		if err := service.ValidateTicketStatus(utils.TicketStatus(status)); err != nil {
			fmt.Println("Invalid status:", err)
			continue
		}

		tickets, err := c.service.GetTicketsByStatus(utils.TicketStatus(status))
		if err != nil {
			return err
		}

		if len(tickets) == 0 {
			fmt.Printf("No tickets founds with status '%s'.\n", status)
			return nil
		}

		fmt.Printf("\nFound %d %s tickets:\n", len(tickets), status)
		fmt.Println("======================================================")

		for _, t := range tickets {
			fmt.Printf("\nTicket ID: %d | Stake: %.2f | Total Odds: %.2f | Possible Win: %.2f\n", t.ID, t.Stake, t.TotalOdds, t.PossibleWin)
			fmt.Println("------------------------------------------------------")

			for i, m := range t.Matches {
				fmt.Printf("Match %d:\n", i+1)
				fmt.Printf("  MatchID: %d\n", m.ID)
				fmt.Printf("  League: %s\n", m.League)
				fmt.Printf("  %s vs %s\n", m.HomeTeam, m.AwayTeam)
				fmt.Printf("  Picked Outcome: %s | Odd: %.2f\n", m.PickedOutcome, m.Odd)
				fmt.Printf("  Arguments: %s\n", m.Arguments)
				fmt.Printf("  Result: %s\n", m.Result)
			}
			fmt.Println("Status:", t.Status)
			fmt.Println("======================================================")
		}
		return nil
	}
}


func (c *CliClient) UpdateTicketDate() error {
	fmt.Println("\n---- Update ticket date ----")

	var ticketID uint
	for {
		idStr := c.readInput("Enter ticket ID: ")
		id, err := strconv.Atoi(idStr)
		if err != nil || id <= 0 {
			fmt.Println("Invalid ticket id, try again.")
			continue
		}
		_, err = c.service.GetTicketByID(uint(id))
		if err != nil {
			fmt.Println("Error:", err)
			continue
		}
		ticketID = uint(id)
		break
	}

	var newDate utils.Date
	for {
		dateInput := c.readInput("Enter new date: ")
		if err := newDate.UnmarshalJSON([]byte(`"` + dateInput + `"`)); err != nil {
			fmt.Println("Invalid date format. Use Year-Month-Day.")
			continue
		}
		break
	}

	req := &model.UpdateTicketDateRequest{
		TicketID: 	ticketID,
		NewDate: 	newDate,
	}

	ticket, err := c.service.UpdateTicketDate(req)
	if err != nil {
		return err
	}

	fmt.Printf("Ticket with id: %d updated date successfully. New date: %v",  ticket.ID, ticket.CreatedAt)
	return nil
}


func (c *CliClient) GetTicketById() error {
	fmt.Println("\n---- Get ticket by id ----")

	var ticketId uint
	for {
		idStr := c.readInput("Enter ticket ID: ")
		id, err := strconv.Atoi(idStr)
		if err != nil || id <= 0 {
			fmt.Println("Invalid ticket id, try again.")
			continue
		}

		ticketId = uint(id)
		ticket, err := c.service.GetTicketByID(ticketId)
		if err != nil {
			return err
		}

		fmt.Printf("TicketID:%d\n", ticket.ID)
		fmt.Printf("Created:%s\n", ticket.CreatedAt.Format("2006-01-02 15:04:05"))
		fmt.Printf("Updated:%s\n", ticket.UpdatedAt.Format("2006-01-02 15:04:05"))
		fmt.Printf("Stake:%.2f\n", ticket.Stake)
		fmt.Printf("Total odds:%.2f\n", ticket.TotalOdds)
		fmt.Printf("Possible win:%.2f\n", ticket.PossibleWin)
		fmt.Printf("Actual win:%.2f\n", ticket.ActualWin)
		fmt.Printf("Cashout:%t\n", ticket.CashOut)
		fmt.Printf("Status:%s\n", ticket.Status)

		for i, match := range ticket.Matches {
			fmt.Printf("Match: %d\n", i+1)
			fmt.Printf("   MatchID:%d\n", match.ID)
			fmt.Printf("   League:%s\n", match.League)
			fmt.Printf("   Home team:%s\n", match.HomeTeam)
			fmt.Printf("   Away team:%s\n", match.AwayTeam)
			fmt.Printf("   Outcome:%s\n", match.PickedOutcome)
			fmt.Printf("   Odd:%.2f", match.Odd)
			fmt.Printf("   Arguments:%s\n", match.Arguments)
			fmt.Printf("   Result:%s\n", match.Result)
		}
		break
	}

	return nil
}