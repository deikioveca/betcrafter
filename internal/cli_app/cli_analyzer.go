package cli_app

import "fmt"

func (c *CliClient) Analyzer() {
	fmt.Println("\n==== Analyzer menu ====")

	fmt.Println("0 - Back")
	fmt.Println("1 - Get ticket stats")
	fmt.Println("2 - Get picked outcome stats")
	fmt.Println("3 - Get picked outcome odd range stats")
	fmt.Println("4 - Get most profitable pick types")

	choice := c.readInput("Enter your choice: ")

	switch choice {
	case "0":
		return
	case "1":
		if err := c.GetTicketStats(); err != nil { 
			fmt.Println("Error:", err) 
		}
	case "2":
		if err := c.GetPickedOutcomeStats(); err != nil { 
			fmt.Println("Error:", err) 
		}
	case "3":
		if err := c.GetPickedOutcomeOddRangeStats(); err != nil { 
			fmt.Println("Error:", err) 
		}
	case "4":
		if err := c.GetMostProfitablePickTypes(); err != nil {
			 fmt.Println("Error:", err) 
			}
	default:
		fmt.Println("Invalid choice, try again.")
	}
}


func (c *CliClient) GetTicketStats() error {
	fmt.Println("\n---- Get ticket stats ----")

	req, err := c.readTicketStatsDateRange()
	if err != nil {
		return err
	}

	stats, err := c.service.GetTicketStats(req)
	if err != nil {
		return err
	}

	fmt.Println("\n--- Stats ---")
	fmt.Printf("Total Tickets: 		%d\n", stats.TotalTickets)
	fmt.Printf("Won Tickets: 		%d\n", stats.WonCount)
	fmt.Printf("Cashout Tickets: 	%d\n", stats.CashOutCount)
	fmt.Printf("Lost Tickets: 		%d\n", stats.LostCount)
	fmt.Printf("Pending Tickets: 	%d\n", stats.PendingCount)
	fmt.Printf("Total Stake: 		%.2f\n", stats.TotalStake)
	fmt.Printf("Total Profit: 		%.2f\n", stats.TotalProfit)
	fmt.Printf("ROI: 			%.2f%%\n", stats.ROI)
	fmt.Printf("Hit Rate: 		%.2f%%\n", stats.HitRate)
	fmt.Printf("Average Stake: 		%.2f\n", stats.AvgStake)

	return nil
}


func (c *CliClient) GetPickedOutcomeStats() error {
	fmt.Println("\n---- Get picked outcome stats ----")

	req, err := c.readTicketStatsDateRange()
	if err != nil {
		return err
	}

	stats, err := c.service.GetPickedOutcomeStats(req)
	if err != nil {
		return err
	}

	fmt.Println("\n--- Stats ---")
	for outcome, s := range stats {
		fmt.Printf("%s: Wins=%d, Losses=%d, WinRate=%.2f%%\n", outcome, s.Wins, s.Losses, s.WinRate)
	}

	return nil
}


func (c *CliClient) GetPickedOutcomeOddRangeStats() error {
	fmt.Println("\n---- Get picked outcome odd range stats ----")

	req, err := c.readTicketStatsDateRange()
	if err != nil {
		return err
	}

	stats, err := c.service.GetPickedOutcomeOddRangeStats(req)
	if err != nil {
		return err
	}

	fmt.Println("\n--- Stats ---")
	for outcome, ranges := range stats {
		fmt.Printf("\n%s:\n", outcome)
		for oddRange, s := range ranges {
			fmt.Printf(" %s: Total=%d, Wins=%d, Losses=%d, WinRate=%.2f%%, AvgOdd=%.2f\n", oddRange, s.Total, s.Wins, s.Losses, s.WinRate, s.AvgOdd)
		}
	}

	return nil
}


func (c *CliClient) GetMostProfitablePickTypes() error {
	fmt.Println("\n---- Get most profitable pick types ----")

	req, err := c.readTicketStatsDateRange()
	if err != nil {
		return err
	}

	stats, err := c.service.GetMostProfitablePickTypes(req)
	if err != nil {
		return err
	}

	fmt.Println("\n--- Stats ---")
	for i, pickType := range stats {
		fmt.Printf("%d. %s: TotalProfit=%.2f, Wins=%d, Losses=%d, Total=%d, WinRate=%.2f%%\n",
				i+1, pickType.PickedOutcome, pickType.TotalProfit, pickType.Wins, pickType.Losses, pickType.Total, pickType.WinRate)
	}

	return nil
}