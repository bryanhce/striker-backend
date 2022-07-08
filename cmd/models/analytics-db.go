package models

import (
	"context"
	"database/sql"
	"time"
)

func GetAnalytics(db *sql.DB, userId, startDate, endDate string) (*AnalyticsBreakdown, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	query := `SELECT COUNT(id) FROM alltasks
				WHERE userId = ?
				AND dailyLogDate >= ?
				AND dailyLogdate <= ?
				GROUP BY taskType
				UNION ALL
				SELECT SUM(effort) FROM alltasks
				WHERE userId = ? 
				AND dailyLogDate >= ?
				AND dailyLogdate <= ?
				UNION ALL
				SELECT COUNT(id) FROM alltasks
				WHERE userId = ?
				AND dailyLogDate >= ?
				AND dailyLogdate <= ?
				AND isCompleted = 0
				UNION ALL
				SELECT SUM(effort) FROM alltasks
				WHERE userId = ?
				AND dailyLogDate >= ?
				AND dailyLogdate <= ?
				AND isCompleted = 0`


	rows, err := db.QueryContext(ctx, query, 
					userId, startDate, endDate,
					userId, startDate, endDate,
					userId, startDate, endDate,
					userId, startDate, endDate,
					)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	count := 0
	var breakdown AnalyticsBreakdown
	var scanErr error
	for rows.Next() {
		if count == 0 {
			scanErr = rows.Scan(
				&breakdown.Assignments,
			)	
		} else if count == 1 {
			scanErr = rows.Scan(
				&breakdown.Events,
			)
		} else if count == 2 {
			scanErr = rows.Scan(
				&breakdown.Notes,
			)
		} else if count == 3 {
			scanErr = rows.Scan(
				&breakdown.TotalEffort,
			)
		} else if count == 4 {
			scanErr = rows.Scan(
				&breakdown.TotalCompletedEvents,
			)
		} else if count == 5 {
			scanErr = rows.Scan(
				&breakdown.TotalCompletedEffort,
			)
		}
		if scanErr != nil {
			return nil, err
		}
		count++
	}
	
	return &breakdown, nil
}

func GetAllAnalytics(db *sql.DB, userId string) (*AnalyticsBreakdown, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	query := `SELECT COUNT(id) FROM alltasks
			WHERE userId = ?
			GROUP BY taskType
			UNION ALL
			SELECT SUM(effort) FROM alltasks
			WHERE userId = ?
			UNION ALL
			SELECT COUNT(id) FROM alltasks
			WHERE userId = ?
			AND isCompleted = 0
			UNION ALL
			SELECT SUM(effort) FROM alltasks
			WHERE userId = ?
			AND isCompleted = 0
			UNION ALL
			SELECT AVG(temp.taskCount) FROM (
			SELECT COUNT(id) as taskCount from alltasks 
			WHERE userId = ?
			AND effort != NULL AND priority != NULL
			AND isCompleted = 1
			GROUP BY dailyLogDate) as temp;`

	rows, err := db.QueryContext(ctx, query, userId, userId, userId, userId, userId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	count := 0
	var breakdown AnalyticsBreakdown
	var scanErr error
	for rows.Next() {
		if count == 0 {
			scanErr = rows.Scan(
				&breakdown.Assignments,
			)	
		} else if count == 1 {
			scanErr = rows.Scan(
				&breakdown.Events,
			)
		} else if count == 2 {
			scanErr = rows.Scan(
				&breakdown.Notes,
			)
		} else if count == 3 {
			scanErr = rows.Scan(
				&breakdown.TotalEffort,
			)
		} else if count == 4 {
			scanErr = rows.Scan(
				&breakdown.TotalCompletedEvents,
			)
		} else if count == 5 {
			scanErr = rows.Scan(
				&breakdown.TotalCompletedEffort,
			)
		}
		if scanErr != nil {
			return nil, err
		}
		count++
	}
	
	return &breakdown, nil
}