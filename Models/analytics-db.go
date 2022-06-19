package models

import (
	"context"
	"database/sql"
	"time"
)

//todo figure out if there is a faster way to do the queries, now taking more
//than 3 sec which is too long for an api
func GetAnalytics(db *sql.DB, userId, startDate, endDate string) (*AnalyticsBreakdown, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	assignmentQuery := `SELECT COUNT(*) FROM alltasks
				WHERE userId = ? 
				AND dailyLogDate >= ? 
				AND dailyLogdate <= ?
				AND taskType = 0;`
	assignmentRow := db.QueryRowContext(ctx, assignmentQuery, userId, startDate, endDate)
	var assignments int
	err := assignmentRow.Scan(&assignments)
	if err != nil {
		return nil, err
	}

	eventQuery := `SELECT COUNT(*) FROM alltasks
				WHERE userId = ? 
				AND dailyLogDate >= ? 
				AND dailyLogdate <= ?
				AND taskType = 1;`
	eventRow := db.QueryRowContext(ctx, eventQuery, userId, startDate, endDate)
	var events int
	err = eventRow.Scan(&events)
	if err != nil {
		return nil, err
	}

	noteQuery := `SELECT COUNT(*) FROM alltasks
				WHERE userId = ? 
				AND dailyLogDate >= ? 
				AND dailyLogdate <= ?
				AND taskType = 2;`
	noteRow := db.QueryRowContext(ctx, noteQuery, userId, startDate, endDate)
	var notes int
	err = noteRow.Scan(&notes)
	if err != nil {
		return nil, err
	}

	effortQuery := `SELECT SUM(effort) FROM alltasks
				WHERE userId = ? 
				AND dailyLogDate >= ? 
				AND dailyLogdate <= ?;`
	effortRow := db.QueryRowContext(ctx, effortQuery, userId, startDate, endDate)
	var effort int
	err = effortRow.Scan(&effort)
	if err != nil {
		return nil, err
	}

	completedEventsQuery := `SELECT COUNT(*) FROM alltasks
				WHERE userId = ? 
				AND dailyLogDate >= ? 
				AND dailyLogdate <= ?
				AND isCompleted = 0
				AND (taskType = 0 OR taskType = 1);`
	completedEventsRow := db.QueryRowContext(ctx, completedEventsQuery, userId, startDate, endDate)
	var completedEvents int
	err = completedEventsRow.Scan(&completedEvents)
	if err != nil {
		return nil, err
	}

	completedEffortQuery := `SELECT SUM(effort) FROM alltasks
				WHERE userId = ? 
				AND dailyLogDate >= ? 
				AND dailyLogdate <= ?
				AND isCompleted = 0;`
	completedEffortRow := db.QueryRowContext(ctx, completedEffortQuery, userId, startDate, endDate)
	var completedEffort int
	err = completedEffortRow.Scan(&completedEffort)
	if err != nil {
		return nil, err
	}

	breakdown := AnalyticsBreakdown{
		Assignments: assignments,
		Events: events,
		Notes: notes,
		TotalEffort: effort,
		TotalCompletedEvents: completedEvents,
		TotalCompletedEffort: completedEffort,
	}

	return &breakdown, nil
}