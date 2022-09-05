// Package postgres contains generated code from xo.
package postgres

// Code generated by xo. DO NOT EDIT.

import (
	"context"
	"fmt"

	"bnk.to/core/tools/db"
)

func (repo *Repository) InsertDepositAccount(ctx context.Context, da *db.DepositAccount) error {
	const stmt = `INSERT INTO deposit_accounts
		(account_id, name, notes, holder_id, holder_type, state, type, accrued_amounts, assigned_branch_id, assigned_centre_id, assigned_user_id, balances, revolving_account_id, currency_code, interest_settings, internal_controls, settlement_account_ids, migration_event_id, overdraft_rate_settings, overdraft_settings, product_id, withholding_tax_source_id, create_time, approve_time, activate_time, update_time, close_time, appraise_time, interest_calculate_time, interest_store_time, overdraft_interest_review_time, arrears_set_time, lock_time, mature_time)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18, $19, $20, $21, $22, $23, $24, $25, $26, $27, $28, $29, $30, $31, $32, $33, $34)
		RETURNING id`
	var id int
	row := repo.db.QueryRowContext(
		ctx, stmt,
		da.AccountID, da.Name, da.Notes, da.HolderID, da.HolderType, da.State, da.Type, da.AccruedAmounts, da.AssignedBranchID, da.AssignedCentreID, da.AssignedUserID, da.Balances, da.RevolvingAccountID, da.CurrencyCode, da.InterestSettings, da.InternalControls, da.SettlementAccountIDs, da.MigrationEventID, da.OverdraftRateSettings, da.OverdraftSettings, da.ProductID, da.WithholdingTaxSourceID, da.CreateTime, da.ApproveTime, da.ActivateTime, da.UpdateTime, da.CloseTime, da.AppraiseTime, da.InterestCalculateTime, da.InterestStoreTime, da.OverdraftInterestReviewTime, da.ArrearsSetTime, da.LockTime, da.MatureTime,
	)
	if err := row.Scan(&id); err != nil {
		return err
	}
	return nil
}

func (repo *Repository) ListDepositAccounts(ctx context.Context, filter string, pageSize int32, orderBy string, after *db.ListPosition) (db.ListStat, []*db.DepositAccount, *db.ListPosition, error) {
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
				quote_ident($1) AS ordered_idx, deposit_accounts.id,

				deposit_accounts.account_id,
				deposit_accounts.name,
				deposit_accounts.notes,
				deposit_accounts.holder_id,
				deposit_accounts.holder_type,
				deposit_accounts.state,
				deposit_accounts.type,
				deposit_accounts.accrued_amounts,
				deposit_accounts.assigned_branch_id,
				deposit_accounts.assigned_centre_id,
				deposit_accounts.assigned_user_id,
				deposit_accounts.balances,
				deposit_accounts.revolving_account_id,
				deposit_accounts.currency_code,
				deposit_accounts.interest_settings,
				deposit_accounts.internal_controls,
				deposit_accounts.settlement_account_ids,
				deposit_accounts.migration_event_id,
				deposit_accounts.overdraft_rate_settings,
				deposit_accounts.overdraft_settings,
				deposit_accounts.product_id,
				deposit_accounts.withholding_tax_source_id,
				deposit_accounts.create_time,
				deposit_accounts.approve_time,
				deposit_accounts.activate_time,
				deposit_accounts.update_time,
				deposit_accounts.close_time,
				deposit_accounts.appraise_time,
				deposit_accounts.interest_calculate_time,
				deposit_accounts.interest_store_time,
				deposit_accounts.overdraft_interest_review_time,
				deposit_accounts.arrears_set_time,
				deposit_accounts.lock_time,
				deposit_accounts.mature_time
			FROM
				deposit_accounts
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
			NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL
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
	rows, err := repo.db.QueryContext(ctx, query, "deposit_accounts."+orderBy, after.Data, after.ID, pageSize)
	if err != nil {
		return db.ListStat{}, nil, nil, err
	}

	defer rows.Close()
	result := make([]*db.DepositAccount, 0, pageSize)
	var nextPos db.ListPosition
	var listStat db.ListStat
	if !rows.Next() {
		return db.ListStat{}, nil, nil, rows.Err()
	}
	var x any
	if err := rows.Scan(
		&listStat.Total, &listStat.Remaining,
		&x, &x,
		&x, &x, &x, &x, &x, &x, &x, &x, &x, &x, &x, &x, &x, &x, &x, &x, &x, &x, &x, &x, &x, &x, &x, &x, &x, &x, &x, &x, &x, &x, &x, &x, &x, &x,
	); err != nil {
		return db.ListStat{}, nil, nil, err
	}
	for rows.Next() {
		var next db.DepositAccount
		if err := rows.Scan(
			&listStat.Total, &listStat.Remaining,
			&nextPos.Data, &nextPos.ID,
			&next.AccountID, &next.Name, &next.Notes, &next.HolderID, &next.HolderType, &next.State, &next.Type, &next.AccruedAmounts, &next.AssignedBranchID, &next.AssignedCentreID, &next.AssignedUserID, &next.Balances, &next.RevolvingAccountID, &next.CurrencyCode, &next.InterestSettings, &next.InternalControls, &next.SettlementAccountIDs, &next.MigrationEventID, &next.OverdraftRateSettings, &next.OverdraftSettings, &next.ProductID, &next.WithholdingTaxSourceID, &next.CreateTime, &next.ApproveTime, &next.ActivateTime, &next.UpdateTime, &next.CloseTime, &next.AppraiseTime, &next.InterestCalculateTime, &next.InterestStoreTime, &next.OverdraftInterestReviewTime, &next.ArrearsSetTime, &next.LockTime, &next.MatureTime,
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

func (repo *Repository) DepositAccountByID(ctx context.Context, id int32) (*db.DepositAccount, error) {
	const stmt = `SELECT
		deposit_accounts.id,

		deposit_accounts.account_id,
		deposit_accounts.name,
		deposit_accounts.notes,
		deposit_accounts.holder_id,
		deposit_accounts.holder_type,
		deposit_accounts.state,
		deposit_accounts.type,
		deposit_accounts.accrued_amounts,
		deposit_accounts.assigned_branch_id,
		deposit_accounts.assigned_centre_id,
		deposit_accounts.assigned_user_id,
		deposit_accounts.balances,
		deposit_accounts.revolving_account_id,
		deposit_accounts.currency_code,
		deposit_accounts.interest_settings,
		deposit_accounts.internal_controls,
		deposit_accounts.settlement_account_ids,
		deposit_accounts.migration_event_id,
		deposit_accounts.overdraft_rate_settings,
		deposit_accounts.overdraft_settings,
		deposit_accounts.product_id,
		deposit_accounts.withholding_tax_source_id,
		deposit_accounts.create_time,
		deposit_accounts.approve_time,
		deposit_accounts.activate_time,
		deposit_accounts.update_time,
		deposit_accounts.close_time,
		deposit_accounts.appraise_time,
		deposit_accounts.interest_calculate_time,
		deposit_accounts.interest_store_time,
		deposit_accounts.overdraft_interest_review_time,
		deposit_accounts.arrears_set_time,
		deposit_accounts.lock_time,
		deposit_accounts.mature_time
	FROM
		deposit_accounts
	WHERE
		id = $1`

	var da db.DepositAccount
	row := repo.db.QueryRowContext(ctx, stmt, id)
	if err := row.Scan(
		&da.ID,
		&da.AccountID, &da.Name, &da.Notes, &da.HolderID, &da.HolderType, &da.State, &da.Type, &da.AccruedAmounts, &da.AssignedBranchID, &da.AssignedCentreID, &da.AssignedUserID, &da.Balances, &da.RevolvingAccountID, &da.CurrencyCode, &da.InterestSettings, &da.InternalControls, &da.SettlementAccountIDs, &da.MigrationEventID, &da.OverdraftRateSettings, &da.OverdraftSettings, &da.ProductID, &da.WithholdingTaxSourceID, &da.CreateTime, &da.ApproveTime, &da.ActivateTime, &da.UpdateTime, &da.CloseTime, &da.AppraiseTime, &da.InterestCalculateTime, &da.InterestStoreTime, &da.OverdraftInterestReviewTime, &da.ArrearsSetTime, &da.LockTime, &da.MatureTime,
	); err != nil {
		return nil, err
	}

	return &da, nil
}

func (repo *Repository) UpdateDepositAccountByID(ctx context.Context, da *db.DepositAccount) error {
	const stmt = `UPDATE deposit_accounts
	SET account_id=$1,
		name=$2,
		notes=$3,
		holder_id=$4,
		holder_type=$5,
		state=$6,
		type=$7,
		accrued_amounts=$8,
		assigned_branch_id=$9,
		assigned_centre_id=$10,
		assigned_user_id=$11,
		balances=$12,
		revolving_account_id=$13,
		currency_code=$14,
		interest_settings=$15,
		internal_controls=$16,
		settlement_account_ids=$17,
		migration_event_id=$18,
		overdraft_rate_settings=$19,
		overdraft_settings=$20,
		product_id=$21,
		withholding_tax_source_id=$22,
		approve_time=$23,
		activate_time=$24,
		update_time=$25,
		close_time=$26,
		appraise_time=$27,
		interest_calculate_time=$28,
		interest_store_time=$29,
		overdraft_interest_review_time=$30,
		arrears_set_time=$31,
		lock_time=$32,
		mature_time=$33
	WHERE id = $34`
	_, err := repo.db.ExecContext(
		ctx, stmt,
		da.AccountID, da.Name, da.Notes, da.HolderID, da.HolderType, da.State, da.Type, da.AccruedAmounts, da.AssignedBranchID, da.AssignedCentreID, da.AssignedUserID, da.Balances, da.RevolvingAccountID, da.CurrencyCode, da.InterestSettings, da.InternalControls, da.SettlementAccountIDs, da.MigrationEventID, da.OverdraftRateSettings, da.OverdraftSettings, da.ProductID, da.WithholdingTaxSourceID, da.ApproveTime, da.ActivateTime, da.UpdateTime, da.CloseTime, da.AppraiseTime, da.InterestCalculateTime, da.InterestStoreTime, da.OverdraftInterestReviewTime, da.ArrearsSetTime, da.LockTime, da.MatureTime, da.ID,
	)
	return err
}

func (repo *Repository) DeleteDepositAccountByID(ctx context.Context, id int32) error {
	const stmt = `DELETE FROM deposit_accounts
	WHERE id = $1`
	_, err := repo.db.ExecContext(
		ctx, stmt,
		id,
	)
	return err
}

func (repo *Repository) DepositAccountByAccountID(ctx context.Context, accountID string) (*db.DepositAccount, error) {
	const stmt = `SELECT
		deposit_accounts.id,

		deposit_accounts.account_id,
		deposit_accounts.name,
		deposit_accounts.notes,
		deposit_accounts.holder_id,
		deposit_accounts.holder_type,
		deposit_accounts.state,
		deposit_accounts.type,
		deposit_accounts.accrued_amounts,
		deposit_accounts.assigned_branch_id,
		deposit_accounts.assigned_centre_id,
		deposit_accounts.assigned_user_id,
		deposit_accounts.balances,
		deposit_accounts.revolving_account_id,
		deposit_accounts.currency_code,
		deposit_accounts.interest_settings,
		deposit_accounts.internal_controls,
		deposit_accounts.settlement_account_ids,
		deposit_accounts.migration_event_id,
		deposit_accounts.overdraft_rate_settings,
		deposit_accounts.overdraft_settings,
		deposit_accounts.product_id,
		deposit_accounts.withholding_tax_source_id,
		deposit_accounts.create_time,
		deposit_accounts.approve_time,
		deposit_accounts.activate_time,
		deposit_accounts.update_time,
		deposit_accounts.close_time,
		deposit_accounts.appraise_time,
		deposit_accounts.interest_calculate_time,
		deposit_accounts.interest_store_time,
		deposit_accounts.overdraft_interest_review_time,
		deposit_accounts.arrears_set_time,
		deposit_accounts.lock_time,
		deposit_accounts.mature_time
	FROM
		deposit_accounts
	WHERE
		account_id = $1`

	var da db.DepositAccount
	row := repo.db.QueryRowContext(ctx, stmt, accountID)
	if err := row.Scan(
		&da.ID,
		&da.AccountID, &da.Name, &da.Notes, &da.HolderID, &da.HolderType, &da.State, &da.Type, &da.AccruedAmounts, &da.AssignedBranchID, &da.AssignedCentreID, &da.AssignedUserID, &da.Balances, &da.RevolvingAccountID, &da.CurrencyCode, &da.InterestSettings, &da.InternalControls, &da.SettlementAccountIDs, &da.MigrationEventID, &da.OverdraftRateSettings, &da.OverdraftSettings, &da.ProductID, &da.WithholdingTaxSourceID, &da.CreateTime, &da.ApproveTime, &da.ActivateTime, &da.UpdateTime, &da.CloseTime, &da.AppraiseTime, &da.InterestCalculateTime, &da.InterestStoreTime, &da.OverdraftInterestReviewTime, &da.ArrearsSetTime, &da.LockTime, &da.MatureTime,
	); err != nil {
		return nil, err
	}

	return &da, nil
}

func (repo *Repository) UpdateDepositAccountByAccountID(ctx context.Context, da *db.DepositAccount) error {
	const stmt = `UPDATE deposit_accounts
	SET account_id=$1,
		name=$2,
		notes=$3,
		holder_id=$4,
		holder_type=$5,
		state=$6,
		type=$7,
		accrued_amounts=$8,
		assigned_branch_id=$9,
		assigned_centre_id=$10,
		assigned_user_id=$11,
		balances=$12,
		revolving_account_id=$13,
		currency_code=$14,
		interest_settings=$15,
		internal_controls=$16,
		settlement_account_ids=$17,
		migration_event_id=$18,
		overdraft_rate_settings=$19,
		overdraft_settings=$20,
		product_id=$21,
		withholding_tax_source_id=$22,
		approve_time=$23,
		activate_time=$24,
		update_time=$25,
		close_time=$26,
		appraise_time=$27,
		interest_calculate_time=$28,
		interest_store_time=$29,
		overdraft_interest_review_time=$30,
		arrears_set_time=$31,
		lock_time=$32,
		mature_time=$33
	WHERE account_id = $34`
	_, err := repo.db.ExecContext(
		ctx, stmt,
		da.AccountID, da.Name, da.Notes, da.HolderID, da.HolderType, da.State, da.Type, da.AccruedAmounts, da.AssignedBranchID, da.AssignedCentreID, da.AssignedUserID, da.Balances, da.RevolvingAccountID, da.CurrencyCode, da.InterestSettings, da.InternalControls, da.SettlementAccountIDs, da.MigrationEventID, da.OverdraftRateSettings, da.OverdraftSettings, da.ProductID, da.WithholdingTaxSourceID, da.ApproveTime, da.ActivateTime, da.UpdateTime, da.CloseTime, da.AppraiseTime, da.InterestCalculateTime, da.InterestStoreTime, da.OverdraftInterestReviewTime, da.ArrearsSetTime, da.LockTime, da.MatureTime, da.AccountID,
	)
	return err
}

func (repo *Repository) DeleteDepositAccountByAccountID(ctx context.Context, accountID string) error {
	const stmt = `DELETE FROM deposit_accounts
	WHERE account_id = $1`
	_, err := repo.db.ExecContext(
		ctx, stmt,
		accountID,
	)
	return err
}
