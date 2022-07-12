package models

import (
	"context"
	"database/sql"
	"time"
)

func GetTaskList(db *sql.DB, userId string, date string) (*[]SingleTask, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := "SELECT * FROM `alltasks` WHERE `userId` = ? AND `dailyLogDate` = ?"

	//QueryContext expects one or more rows
	rows, err := db.QueryContext(ctx, query, userId, date)
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
			&st.HasChildren,
		)
		if err != nil {
			return nil, err
		}
		taskList = append(taskList, st)
	}

	return &taskList, err
	
}

func (st *SingleTask) GetSingleTask(db *sql.DB) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := "SELECT * FROM `alltasks` WHERE `id` = ?"

	//QueryRowContext expect to return at most 1 row
	row := db.QueryRowContext(ctx, query, st.Id)

	err := row.Scan(
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
		&st.HasChildren,
	)

	if err != nil {
		return err
	}

	return nil 
}

func (st *SingleTask) CreateSingleTask(db *sql.DB) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	query := "INSERT INTO `alltasks` VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)"

	_, err := db.ExecContext(ctx, query,
		st.Id, 
		st.DailyLogDate,
		st.Type,
		st.Description,
		st.IsCompleted,
		st.Effort,
		st.Priority,
		st.UserId,
		st.ParentId,
		st.Progress,
		st.Deadline,
		st.HasChildren,
	)
	if err != nil {
		return err
	}

	return nil
}

func (st *SingleTask) UpdateTask(db *sql.DB) error { 
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	query := `UPDATE alltasks SET 
	dailyLogDate = ?, 
	taskType = ?, 
	description = ?, 
	isCompleted = ?,
	effort = ?, 
	priority = ?,
	parentId = ?, 
	progress = ?,
	deadline = ?,
	hasChildren = ? 
	WHERE id = ?;`

	_, err := db.ExecContext(ctx, query, 
		st.DailyLogDate,
		st.Type,
		st.Description,
		st.IsCompleted,
		st.Effort,
		st.Priority,
		st.ParentId,
		st.Progress,
		st.Deadline,
		st.HasChildren,
		st.Id,)
	if err != nil {
		return err
	}

	return nil
}

func (st *SingleTask) DeleteTask(db *sql.DB) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := "DELETE FROM `alltasks` WHERE `id` = ?"

	_, err := db.ExecContext(ctx, query, st.Id)
	if err != nil {
		return err
	}

	return nil
}

func DeleteAllTasks(db *sql.DB, uid string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := "DELETE FROM `alltasks` WHERE `userId` = ?"

	_, err := db.ExecContext(ctx, query, uid)
	if err != nil {
		return err
	}

	return nil
}
