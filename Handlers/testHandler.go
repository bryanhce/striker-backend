package Handlers

import (
	"encoding/json"
	"net/http"
)

func (a *App) TestHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode( struct {
		Message string
	}{"test successful"})
}