package models

import (
	"context"
	"database/sql"
	"time"
)

func CreateUser(db *sql.DB, userId string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := "INSERT INTO `users` VALUES (?, null)"

	_, err := db.ExecContext(ctx, query, userId)
	if err != nil {
		return err
	}

	return nil
}