package cli_app

import (
	"bufio"
	"fmt"
	"os"
	"github.com/deikioveca/betcrafter/internal/ticket/service"
)

type CliClient struct {
	service		service.TicketService
	Reader		*bufio.Reader
}

func NewCliClient(ticketService service.TicketService) *CliClient {
	return &CliClient{
		service: 	ticketService,
		Reader: 	bufio.NewReader(os.Stdin),
	}
}


func (c *CliClient) Run() {
	for {
		fmt.Println("\n==== betcrafter-cli ====")

		c.printHelpfulTicketInfo()

		fmt.Println("0 - Exit")
		fmt.Println("1 - Core")
		fmt.Println("2 - Analyzer")
	
		choiceStr := c.readInput("Enter your choice: ")

		switch choiceStr {
		case "0":
			fmt.Println("Exiting...")
			return
		case "1":
			c.Core()
		case "2":
			c.Analyzer()
		default:
			fmt.Println("Invalid choice, try again.")
		}
	}
}