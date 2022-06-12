package Handlers

import (
	"encoding/json"
	"net/http"
)

func getTaskList(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	json.NewEncoder(w).Encode(task)
}