// Package postgres contains generated code from xo.
package postgres

// Code generated by xo. DO NOT EDIT.

import (
	"context"
	"fmt"

	"bnk.to/core/tools/db"
)

func (repo *Repository) InsertOrganization(ctx context.Context, o *db.Organization) error {
	const stmt = `WITH organizations_info AS (
		INSERT INTO contact_infos
			(full_name, addresses, telephones, emails, language)
			VALUES ($1, $2, $3, $4, $5)
			RETURNING id
	)
	INSERT INTO organizations
		(info, currency_code, date_format, date_time_format, timezone, create_time, update_time)
		SELECT organizations_info.id, $6, $7, $8, $9, $10, $11
		FROM organizations_info
		RETURNING id`
	var id int
	row := repo.db.QueryRowContext(
		ctx, stmt,
		o.Info.FullName, o.Info.Addresses, o.Info.Telephones, o.Info.Emails, o.Info.Language,
		o.CurrencyCode, o.DateFormat, o.DateTimeFormat, o.Timezone, o.CreateTime, o.UpdateTime,
	)
	if err := row.Scan(&id); err != nil {
		return err
	}
	return nil
}

func (repo *Repository) ListOrganizations(ctx context.Context, filter string, pageSize int32, orderBy string, after *db.ListPosition) (db.ListStat, []*db.Organization, *db.ListPosition, error) {
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
				quote_ident($1) AS ordered_idx, organizations.id,

				organizations_info_tbl.full_name,
				organizations_info_tbl.addresses,
				organizations_info_tbl.telephones,
				organizations_info_tbl.emails,
				organizations_info_tbl.language,

				organizations.currency_code,
				organizations.date_format,
				organizations.date_time_format,
				organizations.timezone,
				organizations.create_time,
				organizations.update_time
			FROM
				organizations
				JOIN contact_infos AS organizations_info_tbl ON organizations_info_tbl.id = organizations.info
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
			NULL, NULL, NULL, NULL, NULL,
			NULL, NULL, NULL, NULL, NULL, NULL
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
	rows, err := repo.db.QueryContext(ctx, query, "organizations."+orderBy, after.Data, after.ID, pageSize)
	if err != nil {
		return db.ListStat{}, nil, nil, err
	}

	defer rows.Close()
	result := make([]*db.Organization, 0, pageSize)
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
		&x, &x, &x, &x, &x, &x,
	); err != nil {
		return db.ListStat{}, nil, nil, err
	}
	for rows.Next() {
		var next db.Organization
		if err := rows.Scan(
			&listStat.Total, &listStat.Remaining,
			&nextPos.Data, &nextPos.ID,
			&next.Info.FullName, &next.Info.Addresses, &next.Info.Telephones, &next.Info.Emails, &next.Info.Language,
			&next.CurrencyCode, &next.DateFormat, &next.DateTimeFormat, &next.Timezone, &next.CreateTime, &next.UpdateTime,
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

func (repo *Repository) OrganizationByID(ctx context.Context, id int32) (*db.Organization, error) {
	const stmt = `SELECT
		organizations.id,

		organizations_info_tbl.full_name,
		organizations_info_tbl.addresses,
		organizations_info_tbl.telephones,
		organizations_info_tbl.emails,
		organizations_info_tbl.language,

		organizations.currency_code,
		organizations.date_format,
		organizations.date_time_format,
		organizations.timezone,
		organizations.create_time,
		organizations.update_time
	FROM
		organizations
		JOIN contact_infos AS organizations_info_tbl ON organizations_info_tbl.id = organizations.info
	WHERE
		id = $1`

	var o db.Organization
	row := repo.db.QueryRowContext(ctx, stmt, id)
	if err := row.Scan(
		&o.ID,
		&o.Info.FullName, &o.Info.Addresses, &o.Info.Telephones, &o.Info.Emails, &o.Info.Language,
		&o.CurrencyCode, &o.DateFormat, &o.DateTimeFormat, &o.Timezone, &o.CreateTime, &o.UpdateTime,
	); err != nil {
		return nil, err
	}

	return &o, nil
}

func (repo *Repository) UpdateOrganizationByID(ctx context.Context, o *db.Organization) error {
	const stmt = `WITH organizations_contact_infos AS (
		UPDATE contact_infos
		SET full_name=$1,
			addresses=$2,
			telephones=$3,
			emails=$4,
			language=$5
		FROM organizations
		WHERE organizations.info = contact_infos.id AND
			organizations.id = $6
	)
	UPDATE organizations
	SET currency_code=$7,
		date_format=$8,
		date_time_format=$9,
		timezone=$10,
		update_time=$11
	WHERE id = $12`
	_, err := repo.db.ExecContext(
		ctx, stmt,
		o.Info.FullName, o.Info.Addresses, o.Info.Telephones, o.Info.Emails, o.Info.Language, o.ID,
		o.CurrencyCode, o.DateFormat, o.DateTimeFormat, o.Timezone, o.UpdateTime, o.ID,
	)
	return err
}

func (repo *Repository) DeleteOrganizationByID(ctx context.Context, id int32) error {
	const stmt = `WITH organizations_contact_infos AS (
		DELETE FROM contact_infos
		USING organizations
		WHERE organizations.info = contact_infos.id AND
			organizations.id = $1
	)
	DELETE FROM organizations
	WHERE id = $1`
	_, err := repo.db.ExecContext(
		ctx, stmt,
		id,
	)
	return err
}
