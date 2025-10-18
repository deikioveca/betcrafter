package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/deikioveca/betcrafter/internal/ticket/helper"
	"github.com/deikioveca/betcrafter/internal/ticket/model"
	"github.com/deikioveca/betcrafter/internal/ticket/service"
	"github.com/deikioveca/betcrafter/internal/ticket/utils"
)

type TicketHandler struct {
	service service.TicketService
}

func NewTicketHandler(s service.TicketService) *TicketHandler {
	return &TicketHandler{service: s}
}


func (t *TicketHandler) CreateTicket(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	var ticketRequest model.TicketRequest
	if err := json.NewDecoder(r.Body).Decode(&ticketRequest); err != nil {
		helper.WriteError(w, http.StatusBadRequest, "invalid request body")
		return
	}

	ticket, err := t.service.CreateTicket(&ticketRequest)
	if err != nil {
		switch err {
		case utils.ErrInvalidStake, utils.ErrInvalidMatchesCount, utils.ErrInvalidOutcome, utils.ErrInvalidLeague:
			helper.WriteError(w, http.StatusBadRequest, err.Error())
		default:
			helper.WriteError(w, http.StatusInternalServerError, "internal server error")
		}
		return
	}

	helper.WriteJSON(w, http.StatusCreated, ticket)
}


func (t *TicketHandler) GetPendingTickets(w http.ResponseWriter, r *http.Request) {
	tickets, err := t.service.GetPendingTickets()
	if err != nil {
		helper.WriteError(w, http.StatusInternalServerError, "internal server error")
		return
	}

	helper.WriteJSON(w, http.StatusOK, tickets)
}


func (t *TicketHandler) UpdatePendingTicket(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	var updateTicketRequest model.UpdateTicketRequest
	if err := json.NewDecoder(r.Body).Decode(&updateTicketRequest); err != nil {
		helper.WriteError(w, http.StatusBadRequest, "invalid request body")
		return
	}

	ticket, err := t.service.UpdatePendingTicket(&updateTicketRequest)
	if err != nil {
		switch err {
		case utils.ErrTicketNotFound, utils.ErrTicketMatchNotFound:
			helper.WriteError(w, http.StatusNotFound, err.Error())
		case utils.ErrInvalidStatus, utils.ErrInvalidMatchResult:
			helper.WriteError(w, http.StatusBadRequest, err.Error())
		default:
			helper.WriteError(w, http.StatusInternalServerError, "internal server error")
		}
		return
	}

	helper.WriteJSON(w, http.StatusOK, ticket)
}


func (t *TicketHandler) UpdateTicketDate(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	var updateTicketDateRequest model.UpdateTicketDateRequest
	if err := json.NewDecoder(r.Body).Decode(&updateTicketDateRequest); err != nil {
		helper.WriteError(w, http.StatusBadRequest, "invalid request body")
		return
	}

	ticket, err := t.service.UpdateTicketDate(&updateTicketDateRequest)
	if err != nil {
		switch err {
		case utils.ErrTicketNotFound:
			helper.WriteError(w, http.StatusNotFound, err.Error())
		default:
			helper.WriteError(w, http.StatusInternalServerError, "internal server error")
		}
		return
	}

	helper.WriteJSON(w, http.StatusOK, ticket)
}


func (t *TicketHandler) GetTicketByID(w http.ResponseWriter, r *http.Request) {
	pathValue := r.PathValue("ticketID")
	ticketID, err := strconv.Atoi(pathValue)
	if err != nil || ticketID < 0{
		helper.WriteError(w, http.StatusBadRequest, "invalid ticket id")
		return
	}

	ticket, err := t.service.GetTicketByID(uint(ticketID))
	if err != nil {
		switch err {
		case utils.ErrTicketNotFound:
			helper.WriteError(w, http.StatusNotFound, err.Error())
		default:
			helper.WriteError(w, http.StatusInternalServerError, "internal server error")
		}
		return
	}

	helper.WriteJSON(w, http.StatusOK, ticket)
}


func (t *TicketHandler) GetTicketByStatus(w http.ResponseWriter, r *http.Request) {
	ticketStatus := r.PathValue("status")

	tickets, err := t.service.GetTicketsByStatus(utils.TicketStatus(ticketStatus))
	if err != nil {
		switch err {
		case utils.ErrInvalidStatus:
			helper.WriteError(w, http.StatusBadRequest, err.Error())
		default:
			helper.WriteError(w, http.StatusInternalServerError, "internal server error")
		}
		return
	}

	helper.WriteJSON(w, http.StatusOK, tickets)
}


func (t *TicketHandler) GetTicketStats(w http.ResponseWriter, r *http.Request) {
	ticketStatsRequest, httpErr := helper.PrepareTicketStatsDate(r)
	if httpErr != nil {
		helper.WriteError(w, httpErr.Code, httpErr.Message)
		return
	}

	stats, err := t.service.GetTicketStats(ticketStatsRequest)
	if err != nil {
		helper.WriteError(w, http.StatusInternalServerError, "internal server error")
		return
	}

	helper.WriteJSON(w, http.StatusOK, stats)
}


func (t *TicketHandler) GetPickedOutcomeStats(w http.ResponseWriter, r *http.Request) {
	ticketStatsRequest, httpErr := helper.PrepareTicketStatsDate(r)
	if httpErr != nil {
		helper.WriteError(w, httpErr.Code, httpErr.Message)
		return
	}

	stats, err := t.service.GetPickedOutcomeStats(ticketStatsRequest)
	if err != nil {
		helper.WriteError(w, http.StatusInternalServerError, "internal server error")
		return
	}

	helper.WriteJSON(w, http.StatusOK, stats)
}


func (t *TicketHandler) GetPickedOutcomeOddRangeStats(w http.ResponseWriter, r *http.Request) {
	ticketStatsRequest, httpErr := helper.PrepareTicketStatsDate(r)
	if httpErr != nil {
		helper.WriteError(w, httpErr.Code, httpErr.Message)
		return
	}

	stats, err := t.service.GetPickedOutcomeOddRangeStats(ticketStatsRequest)
	if err != nil {
		helper.WriteError(w, http.StatusInternalServerError, "internal server error")
		return
	}

	helper.WriteJSON(w, http.StatusOK, stats)
}


func (t *TicketHandler) GetMostProfitablePickTypes(w http.ResponseWriter, r *http.Request) {
	ticketStatsRequest, httpErr := helper.PrepareTicketStatsDate(r)
	if httpErr != nil {
		helper.WriteError(w, httpErr.Code, httpErr.Message)
		return
	}

	stats, err := t.service.GetMostProfitablePickTypes(ticketStatsRequest)
	if err != nil {
		helper.WriteError(w, http.StatusInternalServerError, "internal server error")
		return
	}

	helper.WriteJSON(w, http.StatusOK, stats)
}