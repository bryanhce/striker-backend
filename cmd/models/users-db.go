package models

import (
	"context"
	"database/sql"
	"time"
)

func CreateUser(db *sql.DB, userId, email string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := "INSERT INTO `users` VALUES (?, ?, SUBSTRING_INDEX(CONVERT_TZ(CURRENT_DATE(),'+00:00','+08:00'), ' ', 1), false, false, false, false)"

	_, err := db.ExecContext(ctx, query, userId, email)
	if err != nil {
		return err
	}

	return nil
}

func GetLastLogin(db *sql.DB, userId string) (*LastLoginStruct, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	var loginStruct LastLoginStruct;

	query := `SELECT lastLogin FROM users WHERE id = ?`

	row := db.QueryRowContext(ctx, query, userId)

	err := row.Scan(
		&loginStruct.LastLogin,
	)
	if err != nil {
		return nil, err
	}

	return &loginStruct, nil
}

func UpdateLastLogin(db *sql.DB, userId string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := `UPDATE users SET lastLogin = CURRENT_DATE() WHERE id = ?`

	_, err := db.ExecContext(ctx, query, userId)
	if err != nil {
		return err
	}

	return nil
}

func DeleteUser(db *sql.DB, uid string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := `DELETE FROM users WHERE id = ?`

	_, err := db.ExecContext(ctx, query, uid)
	if err != nil {
		return err
	}

	return nil
}