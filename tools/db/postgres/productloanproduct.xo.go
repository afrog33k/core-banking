// Package postgres contains generated code from xo.
package postgres

// Code generated by xo. DO NOT EDIT.

import (
	"context"
	"fmt"

	"bnk.to/core/tools/db"
)

func (repo *Repository) InsertProductLoanProduct(ctx context.Context, plp *db.ProductLoanProduct) error {
	const stmt = `INSERT INTO product_loan_products
		(product_id, name, type, active, notes, allow_custom_repayment_allocation, category, currency_code, internal_controls, template_ids, settings, create_time, update_time)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13)
		RETURNING id`
	var id int
	row := repo.db.QueryRowContext(
		ctx, stmt,
		plp.ProductID, plp.Name, plp.Type, plp.Active, plp.Notes, plp.AllowCustomRepaymentAllocation, plp.Category, plp.CurrencyCode, plp.InternalControls, plp.TemplateIDs, plp.Settings, plp.CreateTime, plp.UpdateTime,
	)
	if err := row.Scan(&id); err != nil {
		return err
	}
	return nil
}

func (repo *Repository) ListProductLoanProducts(ctx context.Context, filter string, pageSize int32, orderBy string, after *db.ListPosition) (db.ListStat, []*db.ProductLoanProduct, *db.ListPosition, error) {
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
				quote_ident($1) AS ordered_idx, product_loan_products.id,

				product_loan_products.product_id,
				product_loan_products.name,
				product_loan_products.type,
				product_loan_products.active,
				product_loan_products.notes,
				product_loan_products.allow_custom_repayment_allocation,
				product_loan_products.category,
				product_loan_products.currency_code,
				product_loan_products.internal_controls,
				product_loan_products.template_ids,
				product_loan_products.settings,
				product_loan_products.create_time,
				product_loan_products.update_time
			FROM
				product_loan_products
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
			NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL
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
	rows, err := repo.db.QueryContext(ctx, query, "product_loan_products."+orderBy, after.Data, after.ID, pageSize)
	if err != nil {
		return db.ListStat{}, nil, nil, err
	}

	defer rows.Close()
	result := make([]*db.ProductLoanProduct, 0, pageSize)
	var nextPos db.ListPosition
	var listStat db.ListStat
	if !rows.Next() {
		return db.ListStat{}, nil, nil, rows.Err()
	}
	var x any
	if err := rows.Scan(
		&listStat.Total, &listStat.Remaining,
		&x, &x,
		&x, &x, &x, &x, &x, &x, &x, &x, &x, &x, &x, &x, &x,
	); err != nil {
		return db.ListStat{}, nil, nil, err
	}
	for rows.Next() {
		var next db.ProductLoanProduct
		if err := rows.Scan(
			&listStat.Total, &listStat.Remaining,
			&nextPos.Data, &nextPos.ID,
			&next.ProductID, &next.Name, &next.Type, &next.Active, &next.Notes, &next.AllowCustomRepaymentAllocation, &next.Category, &next.CurrencyCode, &next.InternalControls, &next.TemplateIDs, &next.Settings, &next.CreateTime, &next.UpdateTime,
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

func (repo *Repository) ProductLoanProductByID(ctx context.Context, id int32) (*db.ProductLoanProduct, error) {
	const stmt = `SELECT
		product_loan_products.id,

		product_loan_products.product_id,
		product_loan_products.name,
		product_loan_products.type,
		product_loan_products.active,
		product_loan_products.notes,
		product_loan_products.allow_custom_repayment_allocation,
		product_loan_products.category,
		product_loan_products.currency_code,
		product_loan_products.internal_controls,
		product_loan_products.template_ids,
		product_loan_products.settings,
		product_loan_products.create_time,
		product_loan_products.update_time
	FROM
		product_loan_products
	WHERE
		id = $1`

	var plp db.ProductLoanProduct
	row := repo.db.QueryRowContext(ctx, stmt, id)
	if err := row.Scan(
		&plp.ID,
		&plp.ProductID, &plp.Name, &plp.Type, &plp.Active, &plp.Notes, &plp.AllowCustomRepaymentAllocation, &plp.Category, &plp.CurrencyCode, &plp.InternalControls, &plp.TemplateIDs, &plp.Settings, &plp.CreateTime, &plp.UpdateTime,
	); err != nil {
		return nil, err
	}

	return &plp, nil
}

func (repo *Repository) UpdateProductLoanProductByID(ctx context.Context, plp *db.ProductLoanProduct) error {
	const stmt = `UPDATE product_loan_products
	SET product_id=$1,
		name=$2,
		type=$3,
		active=$4,
		notes=$5,
		allow_custom_repayment_allocation=$6,
		category=$7,
		currency_code=$8,
		internal_controls=$9,
		template_ids=$10,
		settings=$11,
		update_time=$12
	WHERE id = $13`
	_, err := repo.db.ExecContext(
		ctx, stmt,
		plp.ProductID, plp.Name, plp.Type, plp.Active, plp.Notes, plp.AllowCustomRepaymentAllocation, plp.Category, plp.CurrencyCode, plp.InternalControls, plp.TemplateIDs, plp.Settings, plp.UpdateTime, plp.ID,
	)
	return err
}

func (repo *Repository) DeleteProductLoanProductByID(ctx context.Context, id int32) error {
	const stmt = `DELETE FROM product_loan_products
	WHERE id = $1`
	_, err := repo.db.ExecContext(
		ctx, stmt,
		id,
	)
	return err
}

func (repo *Repository) ProductLoanProductByProductID(ctx context.Context, productID string) (*db.ProductLoanProduct, error) {
	const stmt = `SELECT
		product_loan_products.id,

		product_loan_products.product_id,
		product_loan_products.name,
		product_loan_products.type,
		product_loan_products.active,
		product_loan_products.notes,
		product_loan_products.allow_custom_repayment_allocation,
		product_loan_products.category,
		product_loan_products.currency_code,
		product_loan_products.internal_controls,
		product_loan_products.template_ids,
		product_loan_products.settings,
		product_loan_products.create_time,
		product_loan_products.update_time
	FROM
		product_loan_products
	WHERE
		product_id = $1`

	var plp db.ProductLoanProduct
	row := repo.db.QueryRowContext(ctx, stmt, productID)
	if err := row.Scan(
		&plp.ID,
		&plp.ProductID, &plp.Name, &plp.Type, &plp.Active, &plp.Notes, &plp.AllowCustomRepaymentAllocation, &plp.Category, &plp.CurrencyCode, &plp.InternalControls, &plp.TemplateIDs, &plp.Settings, &plp.CreateTime, &plp.UpdateTime,
	); err != nil {
		return nil, err
	}

	return &plp, nil
}

func (repo *Repository) UpdateProductLoanProductByProductID(ctx context.Context, plp *db.ProductLoanProduct) error {
	const stmt = `UPDATE product_loan_products
	SET product_id=$1,
		name=$2,
		type=$3,
		active=$4,
		notes=$5,
		allow_custom_repayment_allocation=$6,
		category=$7,
		currency_code=$8,
		internal_controls=$9,
		template_ids=$10,
		settings=$11,
		update_time=$12
	WHERE product_id = $13`
	_, err := repo.db.ExecContext(
		ctx, stmt,
		plp.ProductID, plp.Name, plp.Type, plp.Active, plp.Notes, plp.AllowCustomRepaymentAllocation, plp.Category, plp.CurrencyCode, plp.InternalControls, plp.TemplateIDs, plp.Settings, plp.UpdateTime, plp.ProductID,
	)
	return err
}

func (repo *Repository) DeleteProductLoanProductByProductID(ctx context.Context, productID string) error {
	const stmt = `DELETE FROM product_loan_products
	WHERE product_id = $1`
	_, err := repo.db.ExecContext(
		ctx, stmt,
		productID,
	)
	return err
}
