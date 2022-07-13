package app

import (
	"backend/cmd/models"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func (a *App) CreateUser(w http.ResponseWriter, r *http.Request) {
	newUser := struct{
		UserId string `json:"userId"`
		Email string `json:"email"`
	}{}

	err := json.NewDecoder(r.Body).Decode(&newUser)
	if (err != nil) {
		fmt.Println(err.Error())
		respondWithError(w, http.StatusBadRequest, "invalid request payload")
		return
	}
	defer r.Body.Close()

	err = models.CreateUser(a.DB, newUser.UserId, newUser.Email)
	if err != nil {
		respondWithError(w, http.StatusNotFound, err.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, newUser)
}

func (a *App) GetLastLogin(w http.ResponseWriter, r *http.Request) {
	var uid string = mux.Vars(r)["userId"]

	lastLogin, err := models.GetLastLogin(a.DB, uid)
	if err != nil {
		respondWithError(w, http.StatusNotFound, err.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, lastLogin)
}

func (a *App) UpdateLastLogin(w http.ResponseWriter, r *http.Request) {
	var uid string = mux.Vars(r)["userId"]

	err := models.UpdateLastLogin(a.DB, uid)
	if err != nil {
		respondWithError(w, http.StatusNotFound, err.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, map[string]string{"result": "update successful"})
}

//deletes both from the users and alltasks table
func (a *App) DeleteUser(w http.ResponseWriter, r *http.Request) {
	var uid string = mux.Vars(r)["userId"]

	//the function order must be DeleteAllTasks before DeleteUser because of how
	//the database is designed with PK and FK
	err := models.DeleteAllTasks(a.DB, uid)
	if err != nil {
		respondWithError(w, http.StatusNotFound, err.Error())
		return
	}

	err = models.DeleteUser(a.DB, uid)
	if err != nil {
		respondWithError(w, http.StatusNotFound, err.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, map[string]string{"result": "delete successful"})
}