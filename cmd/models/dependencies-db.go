package models

import (
	"context"
	"database/sql"
	"time"
)

func Update135(db *sql.DB, userId string, boolValue bool) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := `UPDATE users SET is135 = ? WHERE id = ?`

	_, err := db.ExecContext(ctx, query, boolValue, userId)
	if err != nil {
		return err
	}

	return nil
}

func UpdatePomodoro(db *sql.DB, userId string, boolValue bool) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := `UPDATE users SET isPomodoro = ? WHERE id = ?`

	_, err := db.ExecContext(ctx, query, boolValue, userId)
	if err != nil {
		return err
	}

	return nil
}

func UpdateDarkMode(db *sql.DB, userId string, boolValue bool) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := `UPDATE users SET isDarkMode = ? WHERE id = ?`

	_, err := db.ExecContext(ctx, query, boolValue, userId)
	if err != nil {
		return err
	}

	return nil
}

func UpdateColourBlind(db *sql.DB, userId string, boolValue bool) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := `UPDATE users SET isColourBlind = ? WHERE id = ?`

	_, err := db.ExecContext(ctx, query, boolValue, userId)
	if err != nil {
		return err
	}

	return nil
}