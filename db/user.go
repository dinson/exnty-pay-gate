package db

import (
	"context"
	"fmt"
	"time"
)

func (i impl) CreateUser(ctx context.Context, user User) error {
	query := `INSERT INTO users (username, email, country_id, created_at, updated_at) 
			  VALUES ($1, $2, $3, $4, $5) RETURNING id`

	err := i.db.QueryRow(query, user.Username, user.Email, user.CountryID, time.Now(), time.Now()).Scan(&user.ID)
	if err != nil {
		return fmt.Errorf("failed to insert user: %v", err)
	}
	return nil
}

func (i impl) GetUsers(ctx context.Context) ([]User, error) {
	rows, err := i.db.Query(`SELECT id, username, email, country_id, created_at, updated_at FROM users`)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch users: %v", err)
	}
	defer rows.Close()

	var users []User
	for rows.Next() {
		var user User
		if err := rows.Scan(&user.ID, &user.Username, &user.Email, &user.CountryID, &user.CreatedAt, &user.UpdatedAt); err != nil {
			return nil, fmt.Errorf("failed to scan user: %v", err)
		}
		users = append(users, user)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return users, nil
}
