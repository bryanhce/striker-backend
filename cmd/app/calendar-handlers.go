package app

import (
	"backend/cmd/models"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
)

func (a *App) GetMonthlyTasks(w http.ResponseWriter, r *http.Request) {
	userId := mux.Vars(r)["userId"]
	dateStr := r.URL.Query().Get("year-month") 
	//dateStr validation
	dateArr := strings.Split(dateStr, "-")
	year := dateArr[0]
	month := dateArr[1]
	if !validateYear(year) {
		respondWithError(w, http.StatusNotFound, "year invalid")
		return
	}
	if !validateMonth(month) {
		respondWithError(w, http.StatusNotFound, "month invalid")
		return
	}

	// startDate := fmt.Sprint(dateStr, "-01")
	startDate := dateStr + "-01"
	var endDate string
	if (month == "01" || month == "03" || month == "05" || month == "07" ||
	month == "08" || month == "10" || month == "12") {
		endDate = dateStr + "-31"
	} else if (month == "04" || month == "06" || month == "09" || month == "11") {
		endDate = dateStr + "-30"
	} else {
		endDate = dateStr + "-28"
	}

	monthTaskList, err := models.GetMonthlyTasks(a.DB, userId, startDate, endDate)
	if err != nil {
		respondWithError(w, http.StatusNotFound, err.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, monthTaskList)
}