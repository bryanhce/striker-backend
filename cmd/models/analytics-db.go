package models

import (
	"context"
	"database/sql"
	"time"
)

//does not get get data from database about avergeDailyTaskCompleted
//but func handles the case and always returns 0
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
				AND isCompleted = 1
				UNION ALL
				SELECT SUM(effort) FROM alltasks
				WHERE userId = ?
				AND dailyLogDate >= ?
				AND dailyLogdate <= ?
				AND isCompleted = 1`


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
		var tempFloat sql.NullFloat64
		if count == 0 {
			scanErr = rows.Scan(
				&tempFloat,
			)	
			if tempFloat.Valid {
				breakdown.Assignments = tempFloat.Float64
			} else {
				breakdown.Assignments = 0
			}
		} else if count == 1 {
			scanErr = rows.Scan(
				&tempFloat,
			)
			if tempFloat.Valid {
				breakdown.Events = tempFloat.Float64
			} else {
				breakdown.Events = 0
			}
		} else if count == 2 {
			scanErr = rows.Scan(
				&tempFloat,
			)
			if tempFloat.Valid {
				breakdown.Notes = tempFloat.Float64
			} else {
				breakdown.Notes = 0
			}
		} else if count == 3 {
			scanErr = rows.Scan(
				&tempFloat,
			)
			if tempFloat.Valid {
				breakdown.TotalEffort = tempFloat.Float64
			} else {
				breakdown.TotalEffort = 0
			}
		} else if count == 4 {
			scanErr = rows.Scan(
				&tempFloat,
			)
			if tempFloat.Valid {
				breakdown.TotalCompletedEvents = tempFloat.Float64
			} else {
				breakdown.TotalCompletedEvents = 0
			}
		} else if count == 5 {
			scanErr = rows.Scan(
				&tempFloat,
			)
			if tempFloat.Valid {
				breakdown.TotalCompletedEffort = tempFloat.Float64
			} else {
				breakdown.TotalCompletedEffort = 0
			}
		} else if count == 6 {
			scanErr = rows.Scan(
				&tempFloat,
			)
			if tempFloat.Valid {
				breakdown.AverageDailyTaskCompleted = tempFloat.Float64
			} else {
				breakdown.AverageDailyTaskCompleted = 0
			}
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
			AND isCompleted = 1
			UNION ALL
			SELECT SUM(effort) FROM alltasks
			WHERE userId = ?
			AND isCompleted = 1
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
		var tempFloat sql.NullFloat64
		if count == 0 {
			scanErr = rows.Scan(
				&tempFloat,
			)	
			if tempFloat.Valid {
				breakdown.Assignments = tempFloat.Float64
			} else {
				breakdown.Assignments = 0
			}
		} else if count == 1 {
			scanErr = rows.Scan(
				&tempFloat,
			)
			if tempFloat.Valid {
				breakdown.Events = tempFloat.Float64
			} else {
				breakdown.Events = 0
			}
		} else if count == 2 {
			scanErr = rows.Scan(
				&tempFloat,
			)
			if tempFloat.Valid {
				breakdown.Notes = tempFloat.Float64
			} else {
				breakdown.Notes = 0
			}
		} else if count == 3 {
			scanErr = rows.Scan(
				&tempFloat,
			)
			if tempFloat.Valid {
				breakdown.TotalEffort = tempFloat.Float64
			} else {
				breakdown.TotalEffort = 0
			}
		} else if count == 4 {
			scanErr = rows.Scan(
				&tempFloat,
			)
			if tempFloat.Valid {
				breakdown.TotalCompletedEvents = tempFloat.Float64
			} else {
				breakdown.TotalCompletedEvents = 0
			}
		} else if count == 5 {
			scanErr = rows.Scan(
				&tempFloat,
			)
			if tempFloat.Valid {
				breakdown.TotalCompletedEffort = tempFloat.Float64
			} else {
				breakdown.TotalCompletedEffort = 0
			}
		} else if count == 6 {
			scanErr = rows.Scan(
				&tempFloat,
			)
			if tempFloat.Valid {
				breakdown.AverageDailyTaskCompleted = tempFloat.Float64
			} else {
				breakdown.AverageDailyTaskCompleted = 0
			}
		}
		if scanErr != nil {
			return nil, err
		}
		count++
	}
	
	return &breakdown, nil
}