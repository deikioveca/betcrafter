package web_app

import (
	"log"
	"net/http"

	"github.com/deikioveca/betcrafter/internal/config"
	"github.com/deikioveca/betcrafter/internal/ticket/handler"
	"github.com/deikioveca/betcrafter/internal/ticket/service"
	"gorm.io/gorm"
)

type WebApp struct {
	DB *gorm.DB
	
	TicketService	service.TicketService
	TicketHandler	*handler.TicketHandler
}

func NewWebApp() *WebApp {
	cfg 	:= config.LoadConfig()
	db 		:= config.InitDB(cfg)

	ticketService := service.NewTicketService(db)
	ticketHandler := handler.NewTicketHandler(ticketService)

	return &WebApp{
		DB: 			db,
		TicketService: 	ticketService,
		TicketHandler: 	ticketHandler,
	}
}

func (a *WebApp) Run(addr string) {
	mux := http.NewServeMux()

	mux.HandleFunc("POST /ticket/create", 						a.TicketHandler.CreateTicket)
	mux.HandleFunc("GET /ticket/getPending", 					a.TicketHandler.GetPendingTickets)
	mux.HandleFunc("PATCH /ticket/updatePending", 				a.TicketHandler.UpdatePendingTicket)
	mux.HandleFunc("PATCH /ticket/updateDate", 					a.TicketHandler.UpdateTicketDate)

	mux.HandleFunc("GET /ticket/getById/{ticketID}", 			a.TicketHandler.GetTicketByID)
	mux.HandleFunc("GET /ticket/getByStatus/{status}", 			a.TicketHandler.GetTicketByStatus)

	mux.HandleFunc("GET /ticket/getStats", 						a.TicketHandler.GetTicketStats)
	mux.HandleFunc("GET /ticket/getPickedOutcomeStats", 		a.TicketHandler.GetPickedOutcomeStats)
	mux.HandleFunc("GET /ticket/getPickedOutcomeOddRangeStats", a.TicketHandler.GetPickedOutcomeOddRangeStats)	
	mux.HandleFunc("GET /ticket/getMostProfitablePick", 		a.TicketHandler.GetMostProfitablePickTypes)

	if err := http.ListenAndServe(addr, mux); err != nil {
		log.Fatalf("failed to start server: %v", err)
	}
}