package models

import "database/sql"

type Users struct {
	UserId string `json:"userId"`
	Dependencies []string `json:"dependencies"`
}

type SingleTask struct {
	Id string `json:"id"`
	DailyLogDate int `json:"dailyLogDate"`
	Type int `json:"taskType"`
	Description string `json:"description"`
	IsCompleted sql.NullBool `json:"isCompleted"`
	Effort sql.NullInt64 `json:"effort"`
	Priority sql.NullInt64 `json:"priority"`
	UserId string `json:"userId"`
	ParentId sql.NullString `json:"parentId"`
	Progress sql.NullInt64 `json:"progress"`
	Deadline sql.NullInt64 `json:"deadline"`
}

//todo check if this is being used
type TaskList struct {
	UserId string `json:"userId"`
	MultipleTasks []SingleTask `json:"multipleTasks"`
}

//ID not generated as it is done by mysql
type SingleTaskPayLoad struct {
	DailyLogDate int `json:"dailyLogDate"`
	Type int `json:"taskType"`
	Description string `json:"description"`
	IsCompleted sql.NullBool `json:"isCompleted"`
	Effort sql.NullInt64 `json:"effort"`
	Priority sql.NullInt64 `json:"priority"`
	UserId string `json:"userId"`
	ParentId sql.NullString `json:"parentId"`
	Progress sql.NullInt64 `json:"progress"`
	Deadline sql.NullInt64 `json:"deadline"`
} 
