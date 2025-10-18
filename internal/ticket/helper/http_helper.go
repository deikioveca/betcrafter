package helper

import (
	"encoding/json"
	"net/http"

	"github.com/deikioveca/betcrafter/internal/ticket/model"
	"github.com/deikioveca/betcrafter/internal/ticket/utils"
)

type HTTPError struct {
	Code 	int
	Message	string
}


func WriteJSON(w http.ResponseWriter, httpStatus int, data any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(httpStatus)
	json.NewEncoder(w).Encode(data)
}


func WriteError(w http.ResponseWriter, httpStatus int, msg string) {
	WriteJSON(w, httpStatus, map[string]string{"error": msg})
}


func PrepareTicketStatsDate(r *http.Request) (*model.TicketStatsRequest, *HTTPError) {
	startStr := r.URL.Query().Get("start")
	endStr := r.URL.Query().Get("end")

	if startStr == "" || endStr == "" {
		return nil, &HTTPError{Code: http.StatusBadRequest, Message: "start and end dates are required"}
	}

	startDate := utils.Date{}
	if err := startDate.UnmarshalJSON([]byte(`"` + startStr + `"`)); err != nil {
		return nil, &HTTPError{Code: http.StatusBadRequest, Message: "invalid start date format"}
	}

	endDate := utils.Date{}
	if err := endDate.UnmarshalJSON([]byte(`"` + endStr + `"`)); err != nil {
		return nil, &HTTPError{Code: http.StatusBadRequest, Message: "invalid end date format"}
	}

	return &model.TicketStatsRequest{StartDate: &startDate, EndDate: &endDate}, nil
}