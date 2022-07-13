package models

import "database/sql"

type Users struct {
	UserId string `json:"userId"`
	Dependencies string `json:"dependencies"`
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
	HasChildren bool `json:"hasChildren"`
}

type AnalyticsBreakdown struct {
	Assignments float64 `json:"assignments"`
	Events float64 `json:"events"`
	Notes float64 `json:"notes"`
	TotalEffort float64 `json:"totalEffort"`
	TotalCompletedEvents float64 `json:"totalCompletedEvents"`
	TotalCompletedEffort float64 `json:"totalCompletedEffort"`
	AverageDailyTaskCompleted float64 `json:"averageDailyTaskCompleted"`
}

type LastLoginStruct struct {
	LastLogin string `json:"lastLogin"`
}

type ReminderEmail struct {
	Email string `json:"email"`
	Description string `json:"description"`
}

type DependencyStruct struct {
	Dependency bool `json:"dependency"`
}
