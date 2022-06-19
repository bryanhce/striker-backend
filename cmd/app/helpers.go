package app

import (
	"encoding/json"
	"net/http"
	"strconv"
)

func respondWithError(w http.ResponseWriter, code int, message string) {
    respondWithJSON(w, code, map[string]string{"error": message})
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
    response, _ := json.Marshal(payload)

    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(code)
    w.Write(response)
}

func validateYear(year string) bool {
    yearInt, err := strconv.Atoi(year) 
    if err != nil {
        return false
    }
    if (yearInt <= 1900 || yearInt >= 2500 ) {
        return false
    }
    return true
}

func validateMonth(month string) bool {
    monthInt, err := strconv.Atoi(month) 
    if err != nil {
        return false
    }
    if (monthInt < 1 || monthInt > 12) {
        return false
    }
    return true
}

//todo for february need to check with year as well
func validateDay(day, month string) bool {
    dayInt, err := strconv.Atoi(day) 
    if err != nil {
        return false
    }
    if (month == "01" || month == "03" || month == "05" || month == "07" ||
	month == "08" || month == "10" || month == "12") {
		return dayInt >= 1 && dayInt <= 31
	} else if (month == "04" || month == "06" || month == "09" || month == "11") {
		return dayInt >= 1 && dayInt <= 30
	} else {
		return dayInt >= 1 && dayInt <= 28
	} 
}