// Package postgres contains generated code from xo.
package postgres

// Code generated by xo. DO NOT EDIT.

import (
	"context"
	"fmt"

	"bnk.to/core/tools/db"
)

func (repo *Repository) InsertTransactionLoanTransaction(ctx context.Context, tlt *db.TransactionLoanTransaction) error {
	const stmt = `WITH transaction_loan_transactions_card_transaction AS (
		INSERT INTO card_transactions
			(transaction_id, advice, amount, acceptor, card_id, authorization_id, transaction_time)
			VALUES ($1, $2, $3, $4, $5, $6, $7)
			RETURNING id
	)
	INSERT INTO transaction_loan_transactions
		(transaction_id, user_id, type, account_balances, adjustment_transaction_id, affected_amounts, amount, card_transaction, branch_id, centre_id, custom_payment_amounts, fees, installment_id, migration_event_id, notes, original_amount, original_transaction_id, parent_account_id, parent_loan_transaction_id, recalculation_method, taxes, terms, till_id, channel_id, transfer_details, create_time, book_time, value_time)
		SELECT $8, $9, $10, $11, $12, $13, $14, transaction_loan_transactions_card_transaction.id, $15, $16, $17, $18, $19, $20, $21, $22, $23, $24, $25, $26, $27, $28, $29, $30, $31, $32, $33, $34
		FROM transaction_loan_transactions_card_transaction
		RETURNING id`
	var id int
	row := repo.db.QueryRowContext(
		ctx, stmt,
		tlt.CardTransaction.TransactionID, tlt.CardTransaction.Advice, tlt.CardTransaction.Amount, tlt.CardTransaction.Acceptor, tlt.CardTransaction.CardID, tlt.CardTransaction.AuthorizationID, tlt.CardTransaction.TransactionTime,
		tlt.TransactionID, tlt.UserID, tlt.Type, tlt.AccountBalances, tlt.AdjustmentTransactionID, tlt.AffectedAmounts, tlt.Amount, tlt.BranchID, tlt.CentreID, tlt.CustomPaymentAmounts, tlt.Fees, tlt.InstallmentID, tlt.MigrationEventID, tlt.Notes, tlt.OriginalAmount, tlt.OriginalTransactionID, tlt.ParentAccountID, tlt.ParentLoanTransactionID, tlt.RecalculationMethod, tlt.Taxes, tlt.Terms, tlt.TillID, tlt.ChannelID, tlt.TransferDetails, tlt.CreateTime, tlt.BookTime, tlt.ValueTime,
	)
	if err := row.Scan(&id); err != nil {
		return err
	}
	return nil
}

func (repo *Repository) ListTransactionLoanTransactions(ctx context.Context, filter string, pageSize int32, orderBy string, after *db.ListPosition) (db.ListStat, []*db.TransactionLoanTransaction, *db.ListPosition, error) {
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
				quote_ident($1) AS ordered_idx, transaction_loan_transactions.id,

				transaction_loan_transactions_card_transaction_tbl.transaction_id,
				transaction_loan_transactions_card_transaction_tbl.advice,
				transaction_loan_transactions_card_transaction_tbl.amount,
				transaction_loan_transactions_card_transaction_tbl.acceptor,
				transaction_loan_transactions_card_transaction_tbl.card_id,
				transaction_loan_transactions_card_transaction_tbl.authorization_id,
				transaction_loan_transactions_card_transaction_tbl.transaction_time,

				transaction_loan_transactions.transaction_id,
				transaction_loan_transactions.user_id,
				transaction_loan_transactions.type,
				transaction_loan_transactions.account_balances,
				transaction_loan_transactions.adjustment_transaction_id,
				transaction_loan_transactions.affected_amounts,
				transaction_loan_transactions.amount,
				transaction_loan_transactions.branch_id,
				transaction_loan_transactions.centre_id,
				transaction_loan_transactions.custom_payment_amounts,
				transaction_loan_transactions.fees,
				transaction_loan_transactions.installment_id,
				transaction_loan_transactions.migration_event_id,
				transaction_loan_transactions.notes,
				transaction_loan_transactions.original_amount,
				transaction_loan_transactions.original_transaction_id,
				transaction_loan_transactions.parent_account_id,
				transaction_loan_transactions.parent_loan_transaction_id,
				transaction_loan_transactions.recalculation_method,
				transaction_loan_transactions.taxes,
				transaction_loan_transactions.terms,
				transaction_loan_transactions.till_id,
				transaction_loan_transactions.channel_id,
				transaction_loan_transactions.transfer_details,
				transaction_loan_transactions.create_time,
				transaction_loan_transactions.book_time,
				transaction_loan_transactions.value_time
			FROM
				transaction_loan_transactions
				JOIN card_transactions AS transaction_loan_transactions_card_transaction_tbl ON transaction_loan_transactions_card_transaction_tbl.id = transaction_loan_transactions.card_transaction
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
			NULL, NULL, NULL, NULL, NULL, NULL, NULL,
			NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL
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
	rows, err := repo.db.QueryContext(ctx, query, "transaction_loan_transactions."+orderBy, after.Data, after.ID, pageSize)
	if err != nil {
		return db.ListStat{}, nil, nil, err
	}

	defer rows.Close()
	result := make([]*db.TransactionLoanTransaction, 0, pageSize)
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
		&x, &x, &x, &x, &x, &x, &x, &x, &x, &x, &x, &x, &x, &x, &x, &x, &x, &x, &x, &x, &x, &x, &x, &x, &x, &x, &x,
	); err != nil {
		return db.ListStat{}, nil, nil, err
	}
	for rows.Next() {
		var next db.TransactionLoanTransaction
		if err := rows.Scan(
			&listStat.Total, &listStat.Remaining,
			&nextPos.Data, &nextPos.ID,
			&next.CardTransaction.TransactionID, &next.CardTransaction.Advice, &next.CardTransaction.Amount, &next.CardTransaction.Acceptor, &next.CardTransaction.CardID, &next.CardTransaction.AuthorizationID, &next.CardTransaction.TransactionTime,
			&next.TransactionID, &next.UserID, &next.Type, &next.AccountBalances, &next.AdjustmentTransactionID, &next.AffectedAmounts, &next.Amount, &next.BranchID, &next.CentreID, &next.CustomPaymentAmounts, &next.Fees, &next.InstallmentID, &next.MigrationEventID, &next.Notes, &next.OriginalAmount, &next.OriginalTransactionID, &next.ParentAccountID, &next.ParentLoanTransactionID, &next.RecalculationMethod, &next.Taxes, &next.Terms, &next.TillID, &next.ChannelID, &next.TransferDetails, &next.CreateTime, &next.BookTime, &next.ValueTime,
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

func (repo *Repository) TransactionLoanTransactionByID(ctx context.Context, id int32) (*db.TransactionLoanTransaction, error) {
	const stmt = `SELECT
		transaction_loan_transactions.id,

		transaction_loan_transactions_card_transaction_tbl.transaction_id,
		transaction_loan_transactions_card_transaction_tbl.advice,
		transaction_loan_transactions_card_transaction_tbl.amount,
		transaction_loan_transactions_card_transaction_tbl.acceptor,
		transaction_loan_transactions_card_transaction_tbl.card_id,
		transaction_loan_transactions_card_transaction_tbl.authorization_id,
		transaction_loan_transactions_card_transaction_tbl.transaction_time,

		transaction_loan_transactions.transaction_id,
		transaction_loan_transactions.user_id,
		transaction_loan_transactions.type,
		transaction_loan_transactions.account_balances,
		transaction_loan_transactions.adjustment_transaction_id,
		transaction_loan_transactions.affected_amounts,
		transaction_loan_transactions.amount,
		transaction_loan_transactions.branch_id,
		transaction_loan_transactions.centre_id,
		transaction_loan_transactions.custom_payment_amounts,
		transaction_loan_transactions.fees,
		transaction_loan_transactions.installment_id,
		transaction_loan_transactions.migration_event_id,
		transaction_loan_transactions.notes,
		transaction_loan_transactions.original_amount,
		transaction_loan_transactions.original_transaction_id,
		transaction_loan_transactions.parent_account_id,
		transaction_loan_transactions.parent_loan_transaction_id,
		transaction_loan_transactions.recalculation_method,
		transaction_loan_transactions.taxes,
		transaction_loan_transactions.terms,
		transaction_loan_transactions.till_id,
		transaction_loan_transactions.channel_id,
		transaction_loan_transactions.transfer_details,
		transaction_loan_transactions.create_time,
		transaction_loan_transactions.book_time,
		transaction_loan_transactions.value_time
	FROM
		transaction_loan_transactions
		JOIN card_transactions AS transaction_loan_transactions_card_transaction_tbl ON transaction_loan_transactions_card_transaction_tbl.id = transaction_loan_transactions.card_transaction
	WHERE
		id = $1`

	var tlt db.TransactionLoanTransaction
	row := repo.db.QueryRowContext(ctx, stmt, id)
	if err := row.Scan(
		&tlt.ID,
		&tlt.CardTransaction.TransactionID, &tlt.CardTransaction.Advice, &tlt.CardTransaction.Amount, &tlt.CardTransaction.Acceptor, &tlt.CardTransaction.CardID, &tlt.CardTransaction.AuthorizationID, &tlt.CardTransaction.TransactionTime,
		&tlt.TransactionID, &tlt.UserID, &tlt.Type, &tlt.AccountBalances, &tlt.AdjustmentTransactionID, &tlt.AffectedAmounts, &tlt.Amount, &tlt.BranchID, &tlt.CentreID, &tlt.CustomPaymentAmounts, &tlt.Fees, &tlt.InstallmentID, &tlt.MigrationEventID, &tlt.Notes, &tlt.OriginalAmount, &tlt.OriginalTransactionID, &tlt.ParentAccountID, &tlt.ParentLoanTransactionID, &tlt.RecalculationMethod, &tlt.Taxes, &tlt.Terms, &tlt.TillID, &tlt.ChannelID, &tlt.TransferDetails, &tlt.CreateTime, &tlt.BookTime, &tlt.ValueTime,
	); err != nil {
		return nil, err
	}

	return &tlt, nil
}

func (repo *Repository) UpdateTransactionLoanTransactionByID(ctx context.Context, tlt *db.TransactionLoanTransaction) error {
	const stmt = `WITH transaction_loan_transactions_card_transactions AS (
		UPDATE card_transactions
		SET transaction_id=$1,
			advice=$2,
			amount=$3,
			acceptor=$4,
			card_id=$5,
			authorization_id=$6,
			transaction_time=$7
		FROM transaction_loan_transactions
		WHERE transaction_loan_transactions.card_transaction = card_transactions.id AND
			transaction_loan_transactions.id = $8
	)
	UPDATE transaction_loan_transactions
	SET transaction_id=$9,
		user_id=$10,
		type=$11,
		account_balances=$12,
		adjustment_transaction_id=$13,
		affected_amounts=$14,
		amount=$15,
		branch_id=$16,
		centre_id=$17,
		custom_payment_amounts=$18,
		fees=$19,
		installment_id=$20,
		migration_event_id=$21,
		notes=$22,
		original_amount=$23,
		original_transaction_id=$24,
		parent_account_id=$25,
		parent_loan_transaction_id=$26,
		recalculation_method=$27,
		taxes=$28,
		terms=$29,
		till_id=$30,
		channel_id=$31,
		transfer_details=$32,
		book_time=$33,
		value_time=$34
	WHERE id = $35`
	_, err := repo.db.ExecContext(
		ctx, stmt,
		tlt.CardTransaction.TransactionID, tlt.CardTransaction.Advice, tlt.CardTransaction.Amount, tlt.CardTransaction.Acceptor, tlt.CardTransaction.CardID, tlt.CardTransaction.AuthorizationID, tlt.CardTransaction.TransactionTime, tlt.ID,
		tlt.TransactionID, tlt.UserID, tlt.Type, tlt.AccountBalances, tlt.AdjustmentTransactionID, tlt.AffectedAmounts, tlt.Amount, tlt.BranchID, tlt.CentreID, tlt.CustomPaymentAmounts, tlt.Fees, tlt.InstallmentID, tlt.MigrationEventID, tlt.Notes, tlt.OriginalAmount, tlt.OriginalTransactionID, tlt.ParentAccountID, tlt.ParentLoanTransactionID, tlt.RecalculationMethod, tlt.Taxes, tlt.Terms, tlt.TillID, tlt.ChannelID, tlt.TransferDetails, tlt.BookTime, tlt.ValueTime, tlt.ID,
	)
	return err
}

func (repo *Repository) DeleteTransactionLoanTransactionByID(ctx context.Context, id int32) error {
	const stmt = `WITH transaction_loan_transactions_card_transactions AS (
		DELETE FROM card_transactions
		USING transaction_loan_transactions
		WHERE transaction_loan_transactions.card_transaction = card_transactions.id AND
			transaction_loan_transactions.id = $1
	)
	DELETE FROM transaction_loan_transactions
	WHERE id = $1`
	_, err := repo.db.ExecContext(
		ctx, stmt,
		id,
	)
	return err
}

func (repo *Repository) TransactionLoanTransactionByTransactionID(ctx context.Context, transactionID string) (*db.TransactionLoanTransaction, error) {
	const stmt = `SELECT
		transaction_loan_transactions.id,

		transaction_loan_transactions_card_transaction_tbl.transaction_id,
		transaction_loan_transactions_card_transaction_tbl.advice,
		transaction_loan_transactions_card_transaction_tbl.amount,
		transaction_loan_transactions_card_transaction_tbl.acceptor,
		transaction_loan_transactions_card_transaction_tbl.card_id,
		transaction_loan_transactions_card_transaction_tbl.authorization_id,
		transaction_loan_transactions_card_transaction_tbl.transaction_time,

		transaction_loan_transactions.transaction_id,
		transaction_loan_transactions.user_id,
		transaction_loan_transactions.type,
		transaction_loan_transactions.account_balances,
		transaction_loan_transactions.adjustment_transaction_id,
		transaction_loan_transactions.affected_amounts,
		transaction_loan_transactions.amount,
		transaction_loan_transactions.branch_id,
		transaction_loan_transactions.centre_id,
		transaction_loan_transactions.custom_payment_amounts,
		transaction_loan_transactions.fees,
		transaction_loan_transactions.installment_id,
		transaction_loan_transactions.migration_event_id,
		transaction_loan_transactions.notes,
		transaction_loan_transactions.original_amount,
		transaction_loan_transactions.original_transaction_id,
		transaction_loan_transactions.parent_account_id,
		transaction_loan_transactions.parent_loan_transaction_id,
		transaction_loan_transactions.recalculation_method,
		transaction_loan_transactions.taxes,
		transaction_loan_transactions.terms,
		transaction_loan_transactions.till_id,
		transaction_loan_transactions.channel_id,
		transaction_loan_transactions.transfer_details,
		transaction_loan_transactions.create_time,
		transaction_loan_transactions.book_time,
		transaction_loan_transactions.value_time
	FROM
		transaction_loan_transactions
		JOIN card_transactions AS transaction_loan_transactions_card_transaction_tbl ON transaction_loan_transactions_card_transaction_tbl.id = transaction_loan_transactions.card_transaction
	WHERE
		transaction_id = $1`

	var tlt db.TransactionLoanTransaction
	row := repo.db.QueryRowContext(ctx, stmt, transactionID)
	if err := row.Scan(
		&tlt.ID,
		&tlt.CardTransaction.TransactionID, &tlt.CardTransaction.Advice, &tlt.CardTransaction.Amount, &tlt.CardTransaction.Acceptor, &tlt.CardTransaction.CardID, &tlt.CardTransaction.AuthorizationID, &tlt.CardTransaction.TransactionTime,
		&tlt.TransactionID, &tlt.UserID, &tlt.Type, &tlt.AccountBalances, &tlt.AdjustmentTransactionID, &tlt.AffectedAmounts, &tlt.Amount, &tlt.BranchID, &tlt.CentreID, &tlt.CustomPaymentAmounts, &tlt.Fees, &tlt.InstallmentID, &tlt.MigrationEventID, &tlt.Notes, &tlt.OriginalAmount, &tlt.OriginalTransactionID, &tlt.ParentAccountID, &tlt.ParentLoanTransactionID, &tlt.RecalculationMethod, &tlt.Taxes, &tlt.Terms, &tlt.TillID, &tlt.ChannelID, &tlt.TransferDetails, &tlt.CreateTime, &tlt.BookTime, &tlt.ValueTime,
	); err != nil {
		return nil, err
	}

	return &tlt, nil
}

func (repo *Repository) UpdateTransactionLoanTransactionByTransactionID(ctx context.Context, tlt *db.TransactionLoanTransaction) error {
	const stmt = `WITH transaction_loan_transactions_card_transactions AS (
		UPDATE card_transactions
		SET transaction_id=$1,
			advice=$2,
			amount=$3,
			acceptor=$4,
			card_id=$5,
			authorization_id=$6,
			transaction_time=$7
		FROM transaction_loan_transactions
		WHERE transaction_loan_transactions.card_transaction = card_transactions.id AND
			transaction_loan_transactions.transaction_id = $8
	)
	UPDATE transaction_loan_transactions
	SET transaction_id=$9,
		user_id=$10,
		type=$11,
		account_balances=$12,
		adjustment_transaction_id=$13,
		affected_amounts=$14,
		amount=$15,
		branch_id=$16,
		centre_id=$17,
		custom_payment_amounts=$18,
		fees=$19,
		installment_id=$20,
		migration_event_id=$21,
		notes=$22,
		original_amount=$23,
		original_transaction_id=$24,
		parent_account_id=$25,
		parent_loan_transaction_id=$26,
		recalculation_method=$27,
		taxes=$28,
		terms=$29,
		till_id=$30,
		channel_id=$31,
		transfer_details=$32,
		book_time=$33,
		value_time=$34
	WHERE transaction_id = $35`
	_, err := repo.db.ExecContext(
		ctx, stmt,
		tlt.CardTransaction.TransactionID, tlt.CardTransaction.Advice, tlt.CardTransaction.Amount, tlt.CardTransaction.Acceptor, tlt.CardTransaction.CardID, tlt.CardTransaction.AuthorizationID, tlt.CardTransaction.TransactionTime, tlt.TransactionID,
		tlt.TransactionID, tlt.UserID, tlt.Type, tlt.AccountBalances, tlt.AdjustmentTransactionID, tlt.AffectedAmounts, tlt.Amount, tlt.BranchID, tlt.CentreID, tlt.CustomPaymentAmounts, tlt.Fees, tlt.InstallmentID, tlt.MigrationEventID, tlt.Notes, tlt.OriginalAmount, tlt.OriginalTransactionID, tlt.ParentAccountID, tlt.ParentLoanTransactionID, tlt.RecalculationMethod, tlt.Taxes, tlt.Terms, tlt.TillID, tlt.ChannelID, tlt.TransferDetails, tlt.BookTime, tlt.ValueTime, tlt.TransactionID,
	)
	return err
}

func (repo *Repository) DeleteTransactionLoanTransactionByTransactionID(ctx context.Context, transactionID string) error {
	const stmt = `WITH transaction_loan_transactions_card_transactions AS (
		DELETE FROM card_transactions
		USING transaction_loan_transactions
		WHERE transaction_loan_transactions.card_transaction = card_transactions.id AND
			transaction_loan_transactions.transaction_id = $1
	)
	DELETE FROM transaction_loan_transactions
	WHERE transaction_id = $1`
	_, err := repo.db.ExecContext(
		ctx, stmt,
		transactionID,
	)
	return err
}
