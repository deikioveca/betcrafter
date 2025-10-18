package cli_app

import (
	"github.com/deikioveca/betcrafter/internal/config"
	"github.com/deikioveca/betcrafter/internal/ticket/service"
	"gorm.io/gorm"
)

type CliApp struct {
	DB *gorm.DB
	TicketService 	service.TicketService
	CliClient		*CliClient
}


func NewCliApp() *CliApp {
	cfg 	:= config.LoadConfig()
	db 		:= config.InitDB(cfg)

	ticketService 	:= 	service.NewTicketService(db)
	cliClient 		:=	NewCliClient(ticketService)

	return &CliApp{
		DB: 			db,
		TicketService: 	ticketService,
		CliClient: 		cliClient,
	}
}

func (c *CliApp) Run() {
	c.CliClient.Run()
}