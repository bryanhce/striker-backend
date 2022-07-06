package app

import (
	"backend/cmd/models"
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
)

func (a *App) GetTaskList(w http.ResponseWriter, r *http.Request) {
	userId := mux.Vars(r)["userId"] //path param
	dateStr := r.URL.Query().Get("date") //query param

	dateArr := strings.Split(dateStr, "-")
	year := dateArr[0]
	month := dateArr[1]
	day := dateArr[2]

	if !validateYear(year) {
		respondWithError(w, http.StatusNotFound, "year is invalid")
		return
	}

	if !validateMonth(month) {
		respondWithError(w, http.StatusNotFound, "month is invalid")
		return
	}

	if !validateDay(day, month) {
		respondWithError(w, http.StatusNotFound, "day is invalid")
		return
	}

	taskList, err := models.GetTaskList(a.DB, userId, dateStr)
	if err != nil {
		respondWithError(w, http.StatusNotFound, err.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, taskList)
}

func (a *App) GetSingleTask(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]

	singleTask := models.SingleTask{Id: id}
	err := singleTask.GetSingleTask(a.DB)
	if err != nil {
		switch err {
        case sql.ErrNoRows:
            respondWithError(w, http.StatusNotFound, "id is invalid")
        default:
            respondWithError(w, http.StatusInternalServerError, err.Error())
        }
        return
	}

	respondWithJSON(w, http.StatusOK, singleTask)
}

func (a *App) CreateSingleTask(w http.ResponseWriter, r *http.Request) {
	dateStr := r.URL.Query().Get("date") //query param
	dateArr := strings.Split(dateStr, "-")
	year := dateArr[0]
	month := dateArr[1]
	day := dateArr[2]

	if !validateYear(year) {
		respondWithError(w, http.StatusNotFound, "year is invalid")
		return
	}

	if !validateMonth(month) {
		respondWithError(w, http.StatusNotFound, "month is invalid")
		return
	}

	if !validateDay(day, month) {
		respondWithError(w, http.StatusNotFound, "day is invalid")
		return
	}

	newTask := models.SingleTask{DailyLogDate: dateStr}
	var holder map[string]interface{}

	err := json.NewDecoder(r.Body).Decode(&holder)
	if (err != nil) {
		fmt.Println(err.Error())
		respondWithError(w, http.StatusBadRequest, "invalid request payload")
		return
	}
	defer r.Body.Close()

	if (holder["description"] == nil) || (holder["taskType"] == nil) || 
	(holder["userId"] == nil) || (holder["id"] == nil) || (holder["hasChildren"] == nil) {
		respondWithError(w, http.StatusBadRequest, "invalid request payload")
		return
	}

	//type casting from interface to specific type
	newTask.Id = holder["id"].(string)
	newTask.Type = int(holder["taskType"].(float64))
	newTask.Description = holder["description"].(string)
	newTask.UserId = holder["userId"].(string)
	newTask.HasChildren = holder["hasChildren"].(bool)

	if (holder["isCompleted"] == nil) {
		newTask.IsCompleted.Scan(nil)
	} else {
		newTask.IsCompleted.Scan(holder["isCompleted"].(bool))
	}

	if (holder["effort"] == nil) {
		newTask.IsCompleted.Scan(nil)
	} else {
		newTask.Effort.Scan(int(holder["effort"].(float64)))
	}

	if (holder["priority"] == nil) {
		newTask.IsCompleted.Scan(nil)
	} else {
		newTask.Priority.Scan(int(holder["priority"].(float64)))
	}

	if (holder["parentId"] == nil) {
		newTask.IsCompleted.Scan(nil)
	} else {
		newTask.ParentId.Scan(holder["parentId"].(string))
	}

	if (holder["progress"] == nil) {
		newTask.IsCompleted.Scan(nil)
	} else {
		newTask.Progress.Scan(int(holder["progress"].(float64)))
	}

	if (holder["deadline"] == nil) {
		newTask.IsCompleted.Scan(nil)
	} else {
		newTask.Deadline.Scan(holder["deadline"].(string))
	}

	err = newTask.CreateSingleTask(a.DB)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return 
	}

	respondWithJSON(w, http.StatusCreated, newTask)
}

func (a *App) UpdateSingleTask(w http.ResponseWriter, r *http.Request) {
	var idParam string = mux.Vars(r)["id"]
	updatedTask := models.SingleTask{Id: idParam}

	err := updatedTask.GetSingleTask(a.DB)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "invalid task id")
		return 
	}

	var holder map[string]interface{}
	err = json.NewDecoder(r.Body).Decode(&holder)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return 
	}
	defer r.Body.Close()

	//no check for userId and id as they do not change
	if (holder["dailyLogDate"] != nil) {
		updatedTask.DailyLogDate = (holder["dailyLogDate"].(string))
	}

	if (holder["taskType"] != nil) {
		updatedTask.Type = int(holder["taskType"].(float64))
	}

	if (holder["description"] != nil) {
		updatedTask.Description = holder["description"].(string)
	}

	if (holder["isCompleted"] != nil) {
		updatedTask.IsCompleted.Scan(holder["isCompleted"].(bool))
	}

	if (holder["effort"] != nil) {
		updatedTask.Effort.Scan(int(holder["effort"].(float64)))
	}

	if (holder["priority"] != nil) {
		updatedTask.Priority.Scan(int(holder["priority"].(float64)))
	}

	if (holder["parentId"] != nil) {
		updatedTask.ParentId.Scan(holder["parentId"].(string))
	}

	if (holder["progress"] != nil) {
		updatedTask.Progress.Scan(int(holder["progress"].(float64)))
	}

	if (holder["deadline"] != nil) {
		updatedTask.Deadline.Scan(holder["deadline"].(string))
	}

	if (holder["hasChildren"] != nil) {
		updatedTask.HasChildren = holder["hasChildren"].(bool)
	}

	err = updatedTask.UpdateTask(a.DB)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, updatedTask)
}

func (a *App) DeleteSingleTask(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	task := models.SingleTask{Id: id}
	err := task.DeleteTask(a.DB)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJSON(w, http.StatusOK, map[string]string{"result": "delete successful"})
}