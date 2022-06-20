package models

import "database/sql"

type Users struct {
	UserId string `json:"userId"`
	Dependencies []string `json:"dependencies"`
}

type SingleTask struct {
	Id string `json:"id"`
	DailyLogDate string `json:"dailyLogDate"`
	Type int `json:"taskType"`
	Description string `json:"description"`
	IsCompleted sql.NullBool `json:"isCompleted"`
	Effort sql.NullInt64 `json:"effort"`
	Priority sql.NullInt64 `json:"priority"`
	UserId string `json:"userId"`
	ParentId sql.NullString `json:"parentId"`
	Progress sql.NullInt64 `json:"progress"`
	Deadline sql.NullString `json:"deadline"`
}

//ID not generated as it is done by mysql
type SingleTaskPayLoad struct {
	DailyLogDate string `json:"dailyLogDate"`
	Type int `json:"taskType"`
	Description string `json:"description"`
	IsCompleted sql.NullBool `json:"isCompleted"`
	Effort sql.NullInt64 `json:"effort"`
	Priority sql.NullInt64 `json:"priority"`
	UserId string `json:"userId"`
	ParentId sql.NullString `json:"parentId"`
	Progress sql.NullInt64 `json:"progress"`
	Deadline sql.NullString `json:"deadline"`
} 

type AnalyticsBreakdown struct {
	Assignments int `json:"assignments"`
	Events int `json:"events"`
	Notes int `json:"notes"`
	TotalEffort int `json:"totalEffort"`
	TotalCompletedEvents int `json:"totalCompletedevents"`
	TotalCompletedEffort int `json:"TotalCompletedEffort"`
}