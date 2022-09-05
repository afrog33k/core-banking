// Package postgres contains generated code from xo.
package postgres

// Code generated by xo. DO NOT EDIT.

import (
	"context"
	"fmt"

	"bnk.to/core/tools/db"
)

func (repo *Repository) InsertCardTransaction(ctx context.Context, ct *db.CardTransaction) error {
	const stmt = `INSERT INTO card_transactions
		(transaction_id, advice, amount, acceptor, card_id, authorization_id, transaction_time)
		VALUES ($1, $2, $3, $4, $5, $6, $7)
		RETURNING id`
	var id int
	row := repo.db.QueryRowContext(
		ctx, stmt,
		ct.TransactionID, ct.Advice, ct.Amount, ct.Acceptor, ct.CardID, ct.AuthorizationID, ct.TransactionTime,
	)
	if err := row.Scan(&id); err != nil {
		return err
	}
	return nil
}

func (repo *Repository) ListCardTransactions(ctx context.Context, filter string, pageSize int32, orderBy string, after *db.ListPosition) (db.ListStat, []*db.CardTransaction, *db.ListPosition, error) {
	if filter != "" {
		return db.ListStat{}, nil, nil, fmt.Errorf("filter is unimplemented")
	}
	if orderBy == "" {
		orderBy = "id"
	}
	whereClause := `
		(
			(all_entries.ordered_idx > $2) OR
			(all_entries.ordered_idx = $2 AND all_entries.id > $3)
		)
	`
	if after == nil {
		// Use placeholder values but always evaluate to true.
		whereClause = "$2::INTEGER = $3"
		after = &db.ListPosition{
			Data: 0,
			ID:   0,
		}
	}
	const stmt = `WITH all_entries AS (
			SELECT
				quote_ident($1) AS ordered_idx, card_transactions.id,

				card_transactions.transaction_id,
				card_transactions.advice,
				card_transactions.amount,
				card_transactions.acceptor,
				card_transactions.card_id,
				card_transactions.authorization_id,
				card_transactions.transaction_time
			FROM
				card_transactions
			WHERE
				%s
		), all_count AS (
			SELECT
				COUNT(*) AS count
			FROM
				all_entries
		), filtered AS (
			SELECT
				*
			FROM
				all_entries
			WHERE %s
		), filtered_count AS (
			SELECT
				COUNT(*) AS count
			FROM
				filtered
		)
		SELECT
			all_count.count, filtered_count.count,
			NULL, NULL,
			NULL, NULL, NULL, NULL, NULL, NULL, NULL
		FROM
			all_count
			CROSS JOIN filtered_count
		UNION ALL
		(
			SELECT
				*
			FROM
				all_count
				CROSS JOIN filtered_count
				CROSS JOIN filtered
			ORDER BY
				quote_ident($1), filtered.id
			LIMIT
				$4
		)`

	filterSQL := "TRUE" // TODO
	query := fmt.Sprintf(stmt, filterSQL, whereClause)
	rows, err := repo.db.QueryContext(ctx, query, "card_transactions."+orderBy, after.Data, after.ID, pageSize)
	if err != nil {
		return db.ListStat{}, nil, nil, err
	}

	defer rows.Close()
	result := make([]*db.CardTransaction, 0, pageSize)
	var nextPos db.ListPosition
	var listStat db.ListStat
	if !rows.Next() {
		return db.ListStat{}, nil, nil, rows.Err()
	}
	var x any
	if err := rows.Scan(
		&listStat.Total, &listStat.Remaining,
		&x, &x,
		&x, &x, &x, &x, &x, &x, &x,
	); err != nil {
		return db.ListStat{}, nil, nil, err
	}
	for rows.Next() {
		var next db.CardTransaction
		if err := rows.Scan(
			&listStat.Total, &listStat.Remaining,
			&nextPos.Data, &nextPos.ID,
			&next.TransactionID, &next.Advice, &next.Amount, &next.Acceptor, &next.CardID, &next.AuthorizationID, &next.TransactionTime,
		); err != nil {
			return db.ListStat{}, nil, nil, err
		}
		result = append(result, &next)
	}
	if rows.Err() != nil {
		return db.ListStat{}, nil, nil, rows.Err()
	}
	return listStat, result, &nextPos, nil
}

func (repo *Repository) CardTransactionByID(ctx context.Context, id int32) (*db.CardTransaction, error) {
	const stmt = `SELECT
		card_transactions.id,

		card_transactions.transaction_id,
		card_transactions.advice,
		card_transactions.amount,
		card_transactions.acceptor,
		card_transactions.card_id,
		card_transactions.authorization_id,
		card_transactions.transaction_time
	FROM
		card_transactions
	WHERE
		id = $1`

	var ct db.CardTransaction
	row := repo.db.QueryRowContext(ctx, stmt, id)
	if err := row.Scan(
		&ct.ID,
		&ct.TransactionID, &ct.Advice, &ct.Amount, &ct.Acceptor, &ct.CardID, &ct.AuthorizationID, &ct.TransactionTime,
	); err != nil {
		return nil, err
	}

	return &ct, nil
}

func (repo *Repository) UpdateCardTransactionByID(ctx context.Context, ct *db.CardTransaction) error {
	const stmt = `UPDATE card_transactions
	SET transaction_id=$1,
		advice=$2,
		amount=$3,
		acceptor=$4,
		card_id=$5,
		authorization_id=$6,
		transaction_time=$7
	WHERE id = $8`
	_, err := repo.db.ExecContext(
		ctx, stmt,
		ct.TransactionID, ct.Advice, ct.Amount, ct.Acceptor, ct.CardID, ct.AuthorizationID, ct.TransactionTime, ct.ID,
	)
	return err
}

func (repo *Repository) DeleteCardTransactionByID(ctx context.Context, id int32) error {
	const stmt = `DELETE FROM card_transactions
	WHERE id = $1`
	_, err := repo.db.ExecContext(
		ctx, stmt,
		id,
	)
	return err
}
