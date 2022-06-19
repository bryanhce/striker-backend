package models

import (
	"context"
	"database/sql"
	"time"
)

func GetMonthlyTasks(db *sql.DB, userId, startDate, endDate string) (*[]SingleTask, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := `SELECT * FROM alltasks
				WHERE userId = ? 
				AND dailyLogDate >= ? 
				AND dailyLogdate <= ?;`

	rows, err := db.QueryContext(ctx, query, userId, startDate, endDate)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var taskList []SingleTask
	for rows.Next() {
		var st SingleTask
		err := rows.Scan(
			&st.Id,
			&st.DailyLogDate,
			&st.Type,
			&st.Description,
			&st.IsCompleted,
			&st.Effort,
			&st.Priority,
			&st.UserId,
			&st.ParentId,
			&st.Progress,
			&st.Deadline,
		)
		if err != nil {
			return nil, err
		}
		taskList = append(taskList, st)
	}

	return &taskList, err
	
}