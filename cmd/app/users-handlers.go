package app

import (
	"backend/cmd/models"
	"encoding/json"
	"fmt"
	"net/http"
)

func (a *App) CreateUser(w http.ResponseWriter, r *http.Request) {
	userId := struct{
		UserId string `json:"userId"`
	}{}

	err := json.NewDecoder(r.Body).Decode(&userId)
	if (err != nil) {
		fmt.Println(err.Error())
		respondWithError(w, http.StatusBadRequest, "invalid request payload")
		return
	}
	defer r.Body.Close()

	err = models.CreateUser(a.DB, userId.UserId)
	if err != nil {
		respondWithError(w, http.StatusNotFound, err.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, userId)
}