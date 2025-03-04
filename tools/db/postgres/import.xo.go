// Package postgres contains generated code from xo.
package postgres

// Code generated by xo. DO NOT EDIT.

import (
	"context"
	"fmt"

	"bnk.to/core/tools/db"
)

func (repo *Repository) InsertImport(ctx context.Context, i *db.Import) error {
	const stmt = `INSERT INTO imports
		(import_id, importer_id, status, errors, progress)
		VALUES ($1, $2, $3, $4, $5)
		RETURNING id`
	var id int
	row := repo.db.QueryRowContext(
		ctx, stmt,
		i.ImportID, i.ImporterID, i.Status, i.Errors, i.Progress,
	)
	if err := row.Scan(&id); err != nil {
		return err
	}
	return nil
}

func (repo *Repository) ListImports(ctx context.Context, filter string, pageSize int32, orderBy string, after *db.ListPosition) (db.ListStat, []*db.Import, *db.ListPosition, error) {
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
				quote_ident($1) AS ordered_idx, imports.id,

				imports.import_id,
				imports.importer_id,
				imports.status,
				imports.errors,
				imports.progress
			FROM
				imports
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
			NULL, NULL, NULL, NULL, NULL
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
	rows, err := repo.db.QueryContext(ctx, query, "imports."+orderBy, after.Data, after.ID, pageSize)
	if err != nil {
		return db.ListStat{}, nil, nil, err
	}

	defer rows.Close()
	result := make([]*db.Import, 0, pageSize)
	var nextPos db.ListPosition
	var listStat db.ListStat
	if !rows.Next() {
		return db.ListStat{}, nil, nil, rows.Err()
	}
	var x any
	if err := rows.Scan(
		&listStat.Total, &listStat.Remaining,
		&x, &x,
		&x, &x, &x, &x, &x,
	); err != nil {
		return db.ListStat{}, nil, nil, err
	}
	for rows.Next() {
		var next db.Import
		if err := rows.Scan(
			&listStat.Total, &listStat.Remaining,
			&nextPos.Data, &nextPos.ID,
			&next.ImportID, &next.ImporterID, &next.Status, &next.Errors, &next.Progress,
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

func (repo *Repository) ImportByID(ctx context.Context, id int32) (*db.Import, error) {
	const stmt = `SELECT
		imports.id,

		imports.import_id,
		imports.importer_id,
		imports.status,
		imports.errors,
		imports.progress
	FROM
		imports
	WHERE
		id = $1`

	var i db.Import
	row := repo.db.QueryRowContext(ctx, stmt, id)
	if err := row.Scan(
		&i.ID,
		&i.ImportID, &i.ImporterID, &i.Status, &i.Errors, &i.Progress,
	); err != nil {
		return nil, err
	}

	return &i, nil
}

func (repo *Repository) UpdateImportByID(ctx context.Context, i *db.Import) error {
	const stmt = `UPDATE imports
	SET import_id=$1,
		importer_id=$2,
		status=$3,
		errors=$4,
		progress=$5
	WHERE id = $6`
	_, err := repo.db.ExecContext(
		ctx, stmt,
		i.ImportID, i.ImporterID, i.Status, i.Errors, i.Progress, i.ID,
	)
	return err
}

func (repo *Repository) DeleteImportByID(ctx context.Context, id int32) error {
	const stmt = `DELETE FROM imports
	WHERE id = $1`
	_, err := repo.db.ExecContext(
		ctx, stmt,
		id,
	)
	return err
}

func (repo *Repository) ImportByImportID(ctx context.Context, importID string) (*db.Import, error) {
	const stmt = `SELECT
		imports.id,

		imports.import_id,
		imports.importer_id,
		imports.status,
		imports.errors,
		imports.progress
	FROM
		imports
	WHERE
		import_id = $1`

	var i db.Import
	row := repo.db.QueryRowContext(ctx, stmt, importID)
	if err := row.Scan(
		&i.ID,
		&i.ImportID, &i.ImporterID, &i.Status, &i.Errors, &i.Progress,
	); err != nil {
		return nil, err
	}

	return &i, nil
}

func (repo *Repository) UpdateImportByImportID(ctx context.Context, i *db.Import) error {
	const stmt = `UPDATE imports
	SET import_id=$1,
		importer_id=$2,
		status=$3,
		errors=$4,
		progress=$5
	WHERE import_id = $6`
	_, err := repo.db.ExecContext(
		ctx, stmt,
		i.ImportID, i.ImporterID, i.Status, i.Errors, i.Progress, i.ImportID,
	)
	return err
}

func (repo *Repository) DeleteImportByImportID(ctx context.Context, importID string) error {
	const stmt = `DELETE FROM imports
	WHERE import_id = $1`
	_, err := repo.db.ExecContext(
		ctx, stmt,
		importID,
	)
	return err
}
