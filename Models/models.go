package Models

import (
	"database/sql"

	"github.com/gorilla/mux"
)

type App struct {
	Router *mux.Router
	DB     *sql.DB
}

type Users struct {
	UserId string `json:"userId"`
	Dependencies []string `json:"dependencies"`
}

type SingleTask struct {
	Id string `json:"id"`
	DailyLogDate int `json:"dailyLogDate"`
	Type string `json:"type"`
	Description string `json:"description"`
	IsCompleted bool `json:"isCompleted"`
	Effort int `json:"effort"`
	UserId string `json:"userId"`
	ParentId string `json:"parentId"`
	Progress string `json:"progress"`
	Deadline int `json:"deadline"`
}

type TaskList struct {
	UserId string `json:"userId"`
	MultipleTasks []SingleTask `json:"multipleTasks"`
}

//type to use when json data is accepted from client side
//ID not generated as it is done by mysql
type DailyTaskPayLoad struct {
	DailyLogDate int `json:"dailyLogDate"`
	Type string `json:"type"`
	Description string `json:"description"`
	IsCompleted bool `json:"isCompleted"`
	Effort int `json:"effort"`
	UserId string `json:"userId"`
}

type MonthlyTaskPayLoad struct {
	DailyLogDate int `json:"dailyLogDate"`
	Type string `json:"type"`
	Description string `json:"description"`
	UserId string `json:"userId"`
	Progress string `json:"progress"`
	Deadline int `json:"deadline"`
}

type SubTaskPayLoad struct {
	DailyLogDate int `json:"dailyLogDate"`
	Type string `json:"type"`
	Description string `json:"description"`
	UserId string `json:"userId"`
	ParentId string `json:"parentId"`
	Progress string `json:"progress"`
	Deadline int `json:"deadline"`
}