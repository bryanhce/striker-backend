package app

import (
	"backend/cmd/models"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func (a *App) Update135(w http.ResponseWriter, r *http.Request) {
	var uid string = mux.Vars(r)["userId"]

	struct135 := struct{
		Is135 bool `json:"is135"`
	}{}

	err := json.NewDecoder(r.Body).Decode(&struct135)
	if (err != nil) {
		fmt.Println(err.Error())
		respondWithError(w, http.StatusBadRequest, "invalid request payload")
		return
	}
	defer r.Body.Close()

	err = models.Update135(a.DB, uid, struct135.Is135)
	if err != nil {
		respondWithError(w, http.StatusNotFound, err.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, map[string]string{"result": "update successful"})
}

func (a *App) UpdatePomodoro(w http.ResponseWriter, r *http.Request) {
	var uid string = mux.Vars(r)["userId"]

	structPomo := struct{
		IsPomodoro bool `json:"isPomodoro"`
	}{}

	err := json.NewDecoder(r.Body).Decode(&structPomo)
	if (err != nil) {
		fmt.Println(err.Error())
		respondWithError(w, http.StatusBadRequest, "invalid request payload")
		return
	}
	defer r.Body.Close()

	err = models.UpdatePomodoro(a.DB, uid, structPomo.IsPomodoro)
	if err != nil {
		respondWithError(w, http.StatusNotFound, err.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, map[string]string{"result": "update successful"})
}

func (a *App) UpdateDarkMode(w http.ResponseWriter, r *http.Request) {
	var uid string = mux.Vars(r)["userId"]

	structDarkMode := struct{
		IsDarkMode bool `json:"isDarkMode"`
	}{}

	err := json.NewDecoder(r.Body).Decode(&structDarkMode)
	if (err != nil) {
		fmt.Println(err.Error())
		respondWithError(w, http.StatusBadRequest, "invalid request payload")
		return
	}
	defer r.Body.Close()

	err = models.UpdateDarkMode(a.DB, uid, structDarkMode.IsDarkMode)
	if err != nil {
		respondWithError(w, http.StatusNotFound, err.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, map[string]string{"result": "update successful"})
}

func (a *App) UpdateColourBlind(w http.ResponseWriter, r *http.Request) {
	var uid string = mux.Vars(r)["userId"]

	structColourBlind := struct{
		IsColourBlind bool `json:"isColourBlind"`
	}{}

	err := json.NewDecoder(r.Body).Decode(&structColourBlind)
	if (err != nil) {
		fmt.Println(err.Error())
		respondWithError(w, http.StatusBadRequest, "invalid request payload")
		return
	}
	defer r.Body.Close()

	err = models.UpdateColourBlind(a.DB, uid, structColourBlind.IsColourBlind)
	if err != nil {
		respondWithError(w, http.StatusNotFound, err.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, map[string]string{"result": "update successful"})
}