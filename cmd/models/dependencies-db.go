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

func Get135(db *sql.DB, userId string) (*DependencyStruct, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	var is135 DependencyStruct

	query := `SELECT is135 FROM users WHERE id = ?`

	row := db.QueryRowContext(ctx, query, userId)

	err := row.Scan(
		&is135.Dependency,
	)
	if err != nil {
		return nil, err
	}

	return &is135, nil
}

func GetPomodoro(db *sql.DB, userId string) (*DependencyStruct, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	var isPomodoro DependencyStruct

	query := `SELECT isPomodoro FROM users WHERE id = ?`

	row := db.QueryRowContext(ctx, query, userId)

	err := row.Scan(
		&isPomodoro.Dependency,
	)
	if err != nil {
		return nil, err
	}

	return &isPomodoro, nil
}

func GetColourBlind(db *sql.DB, userId string) (*DependencyStruct, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	var isColourBlind DependencyStruct

	query := `SELECT isColourBlind FROM users WHERE id = ?`

	row := db.QueryRowContext(ctx, query, userId)

	err := row.Scan(
		&isColourBlind.Dependency,
	)
	if err != nil {
		return nil, err
	}

	return &isColourBlind, nil
}