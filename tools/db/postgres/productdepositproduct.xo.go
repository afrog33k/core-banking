// Package postgres contains generated code from xo.
package postgres

// Code generated by xo. DO NOT EDIT.

import (
	"context"
	"fmt"

	"bnk.to/core/tools/db"
)

func (repo *Repository) InsertProductDepositProduct(ctx context.Context, pdp *db.ProductDepositProduct) error {
	const stmt = `INSERT INTO product_deposit_products
		(product_id, name, notes, category, currencies, controls, settings, active, template_ids, type, create_time, update_time)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12)
		RETURNING id`
	var id int
	row := repo.db.QueryRowContext(
		ctx, stmt,
		pdp.ProductID, pdp.Name, pdp.Notes, pdp.Category, pdp.Currencies, pdp.Controls, pdp.Settings, pdp.Active, pdp.TemplateIDs, pdp.Type, pdp.CreateTime, pdp.UpdateTime,
	)
	if err := row.Scan(&id); err != nil {
		return err
	}
	return nil
}

func (repo *Repository) ListProductDepositProducts(ctx context.Context, filter string, pageSize int32, orderBy string, after *db.ListPosition) (db.ListStat, []*db.ProductDepositProduct, *db.ListPosition, error) {
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
				quote_ident($1) AS ordered_idx, product_deposit_products.id,

				product_deposit_products.product_id,
				product_deposit_products.name,
				product_deposit_products.notes,
				product_deposit_products.category,
				product_deposit_products.currencies,
				product_deposit_products.controls,
				product_deposit_products.settings,
				product_deposit_products.active,
				product_deposit_products.template_ids,
				product_deposit_products.type,
				product_deposit_products.create_time,
				product_deposit_products.update_time
			FROM
				product_deposit_products
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
			NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL
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
	rows, err := repo.db.QueryContext(ctx, query, "product_deposit_products."+orderBy, after.Data, after.ID, pageSize)
	if err != nil {
		return db.ListStat{}, nil, nil, err
	}

	defer rows.Close()
	result := make([]*db.ProductDepositProduct, 0, pageSize)
	var nextPos db.ListPosition
	var listStat db.ListStat
	if !rows.Next() {
		return db.ListStat{}, nil, nil, rows.Err()
	}
	var x any
	if err := rows.Scan(
		&listStat.Total, &listStat.Remaining,
		&x, &x,
		&x, &x, &x, &x, &x, &x, &x, &x, &x, &x, &x, &x,
	); err != nil {
		return db.ListStat{}, nil, nil, err
	}
	for rows.Next() {
		var next db.ProductDepositProduct
		if err := rows.Scan(
			&listStat.Total, &listStat.Remaining,
			&nextPos.Data, &nextPos.ID,
			&next.ProductID, &next.Name, &next.Notes, &next.Category, &next.Currencies, &next.Controls, &next.Settings, &next.Active, &next.TemplateIDs, &next.Type, &next.CreateTime, &next.UpdateTime,
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

func (repo *Repository) ProductDepositProductByID(ctx context.Context, id int32) (*db.ProductDepositProduct, error) {
	const stmt = `SELECT
		product_deposit_products.id,

		product_deposit_products.product_id,
		product_deposit_products.name,
		product_deposit_products.notes,
		product_deposit_products.category,
		product_deposit_products.currencies,
		product_deposit_products.controls,
		product_deposit_products.settings,
		product_deposit_products.active,
		product_deposit_products.template_ids,
		product_deposit_products.type,
		product_deposit_products.create_time,
		product_deposit_products.update_time
	FROM
		product_deposit_products
	WHERE
		id = $1`

	var pdp db.ProductDepositProduct
	row := repo.db.QueryRowContext(ctx, stmt, id)
	if err := row.Scan(
		&pdp.ID,
		&pdp.ProductID, &pdp.Name, &pdp.Notes, &pdp.Category, &pdp.Currencies, &pdp.Controls, &pdp.Settings, &pdp.Active, &pdp.TemplateIDs, &pdp.Type, &pdp.CreateTime, &pdp.UpdateTime,
	); err != nil {
		return nil, err
	}

	return &pdp, nil
}

func (repo *Repository) UpdateProductDepositProductByID(ctx context.Context, pdp *db.ProductDepositProduct) error {
	const stmt = `UPDATE product_deposit_products
	SET product_id=$1,
		name=$2,
		notes=$3,
		category=$4,
		currencies=$5,
		controls=$6,
		settings=$7,
		active=$8,
		template_ids=$9,
		type=$10,
		update_time=$11
	WHERE id = $12`
	_, err := repo.db.ExecContext(
		ctx, stmt,
		pdp.ProductID, pdp.Name, pdp.Notes, pdp.Category, pdp.Currencies, pdp.Controls, pdp.Settings, pdp.Active, pdp.TemplateIDs, pdp.Type, pdp.UpdateTime, pdp.ID,
	)
	return err
}

func (repo *Repository) DeleteProductDepositProductByID(ctx context.Context, id int32) error {
	const stmt = `DELETE FROM product_deposit_products
	WHERE id = $1`
	_, err := repo.db.ExecContext(
		ctx, stmt,
		id,
	)
	return err
}

func (repo *Repository) ProductDepositProductByProductID(ctx context.Context, productID string) (*db.ProductDepositProduct, error) {
	const stmt = `SELECT
		product_deposit_products.id,

		product_deposit_products.product_id,
		product_deposit_products.name,
		product_deposit_products.notes,
		product_deposit_products.category,
		product_deposit_products.currencies,
		product_deposit_products.controls,
		product_deposit_products.settings,
		product_deposit_products.active,
		product_deposit_products.template_ids,
		product_deposit_products.type,
		product_deposit_products.create_time,
		product_deposit_products.update_time
	FROM
		product_deposit_products
	WHERE
		product_id = $1`

	var pdp db.ProductDepositProduct
	row := repo.db.QueryRowContext(ctx, stmt, productID)
	if err := row.Scan(
		&pdp.ID,
		&pdp.ProductID, &pdp.Name, &pdp.Notes, &pdp.Category, &pdp.Currencies, &pdp.Controls, &pdp.Settings, &pdp.Active, &pdp.TemplateIDs, &pdp.Type, &pdp.CreateTime, &pdp.UpdateTime,
	); err != nil {
		return nil, err
	}

	return &pdp, nil
}

func (repo *Repository) UpdateProductDepositProductByProductID(ctx context.Context, pdp *db.ProductDepositProduct) error {
	const stmt = `UPDATE product_deposit_products
	SET product_id=$1,
		name=$2,
		notes=$3,
		category=$4,
		currencies=$5,
		controls=$6,
		settings=$7,
		active=$8,
		template_ids=$9,
		type=$10,
		update_time=$11
	WHERE product_id = $12`
	_, err := repo.db.ExecContext(
		ctx, stmt,
		pdp.ProductID, pdp.Name, pdp.Notes, pdp.Category, pdp.Currencies, pdp.Controls, pdp.Settings, pdp.Active, pdp.TemplateIDs, pdp.Type, pdp.UpdateTime, pdp.ProductID,
	)
	return err
}

func (repo *Repository) DeleteProductDepositProductByProductID(ctx context.Context, productID string) error {
	const stmt = `DELETE FROM product_deposit_products
	WHERE product_id = $1`
	_, err := repo.db.ExecContext(
		ctx, stmt,
		productID,
	)
	return err
}
