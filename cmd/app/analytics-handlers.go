package app

import (
	"backend/models"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
)

func (a *App) GetAnalytics(w http.ResponseWriter, r *http.Request) {
	userId := mux.Vars(r)["userId"]
	startDate := r.URL.Query().Get("start-date") 
	endDate := r.URL.Query().Get("end-date")
	//date validation
	startDateArr := strings.Split(startDate, "-")
	endDateArr := strings.Split(endDate, "-")

	startYear := startDateArr[0]
	endYear := endDateArr[0]
	startMonth := startDateArr[1]
	endMonth := endDateArr[1]
	startDay := startDateArr[2]
	endDay := endDateArr[2]

	if !validateYear(startYear) {
		respondWithError(w, http.StatusNotFound, "start year invalid")
		return
	}

	if !validateYear(endYear) {
		respondWithError(w, http.StatusNotFound, "end year invalid")
		return
	}

	if !validateMonth(startMonth) {
		respondWithError(w, http.StatusNotFound, "start month invalid")
		return
	}

	if !validateMonth(endMonth) {
		respondWithError(w, http.StatusNotFound, "end month invalid")
		return
	}

	if !validateDay(startDay, startMonth) {
		respondWithError(w, http.StatusNotFound, "start day invalid")
		return
	}

	if !validateDay(endDay, endMonth) {
		respondWithError(w, http.StatusNotFound, "end day invalid")
		return
	}

	analyticsReport, err := models.GetAnalytics(a.DB, userId, startDate, endDate)
	if err != nil {
		respondWithError(w, http.StatusNotFound, err.Error())
		return
	}
	respondWithJSON(w, http.StatusOK, analyticsReport)
}