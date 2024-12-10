package db

import (
	"context"
	"database/sql"
	"fmt"
	"time"
)

func (i impl) CreateTransaction(ctx context.Context, transaction *Transaction) (int, error) {
	query := `INSERT INTO transactions (amount, type, status, gateway_id, country_id, user_id, created_at) 
			  VALUES ($1, $2, $3, $4, $5, $6, $7) RETURNING id`

	err := i.db.QueryRowContext(ctx, query, transaction.Amount, transaction.Type, transaction.Status, transaction.GatewayID, transaction.CountryID, transaction.UserID, time.Now()).Scan(&transaction.ID)
	if err != nil {
		return 0, fmt.Errorf("failed to insert transaction: %v", err)
	}
	return transaction.ID, nil
}

func (i impl) GetTransactions(ctx context.Context) ([]Transaction, error) {
	rows, err := i.db.Query(`SELECT id, amount, type, status, user_id, gateway_id, country_id, created_at FROM transactions`)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch transactions: %v", err)
	}
	defer rows.Close()

	var transactions []Transaction
	for rows.Next() {
		var transaction Transaction
		if err := rows.Scan(&transaction.ID, &transaction.Amount, &transaction.Type, &transaction.Status, &transaction.UserID, &transaction.GatewayID, &transaction.CountryID, &transaction.CreatedAt, &transaction.UpdatedAt); err != nil {
			return nil, fmt.Errorf("failed to scan transaction: %v", err)
		}
		transactions = append(transactions, transaction)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return transactions, nil
}

func (i impl) GetTransactionByID(ctx context.Context, id int) (*Transaction, error) {
	query := `SELECT * FROM transactions WHERE id = $1`

	var transaction Transaction

	err := i.db.QueryRowContext(ctx, query, id).Scan(
		&transaction.ID, &transaction.Amount, &transaction.Type, &transaction.Status,
		&transaction.UserID, &transaction.GatewayID, &transaction.CountryID, &transaction.CreatedAt, &transaction.UpdatedAt,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, fmt.Errorf("failed to fetch transaction with ID %s: %w", id, err)
	}

	return &transaction, nil
}

func (i impl) UpdateTransactionByID(ctx context.Context, id int, transaction *Transaction) error {
	query := `UPDATE transactions 
				SET amount = $1, type = $2, status = $3, gateway_id = $4, country_id = $5, updated_at = $6, user_id = $7 
				WHERE id = $8`

	_, err := i.db.ExecContext(ctx, query, transaction.Amount, transaction.Type, transaction.Status,
		transaction.GatewayID, transaction.CountryID, time.Now().UTC(), transaction.UserID, id)
	if err != nil {
		return fmt.Errorf("failed to update transaction with ID %s: %w", id, err)
	}

	return nil
}
