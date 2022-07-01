package models

import (
	"context"
	"database/sql"
	"time"
)

func GetReminderEmails(db *sql.DB) (*[]ReminderEmail, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := `SELECT users.email, GROUP_CONCAT(description SEPARATOR ', ') AS description
				FROM users
				INNER JOIN alltasks ON users.id = alltasks.userId
				WHERE alltasks.deadline = CURRENT_DATE() 
				AND alltasks.description != ""
				AND (alltasks.progress = 0 OR alltasks.progress = 1)
				GROUP BY users.email`

	rows, err := db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var reminderEmailList []ReminderEmail
	for rows.Next() {
		var re ReminderEmail
		err := rows.Scan(
			&re.Email,
			&re.Description,
		)
		if err != nil {
			return nil, err
		}
		reminderEmailList = append(reminderEmailList, re)
	}

	return &reminderEmailList, err
}