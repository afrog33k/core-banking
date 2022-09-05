// Package postgres contains generated code from xo.
package postgres

// Code generated by xo. DO NOT EDIT.

import (
	"context"
	"fmt"

	"bnk.to/core/tools/db"
)

func (repo *Repository) InsertRoleClientRole(ctx context.Context, rcr *db.RoleClientRole) error {
	const stmt = `INSERT INTO role_client_roles
		(role_id, name, client_type, description, require_id, can_guarantee, can_open_accounts, use_default_address, create_time, update_time)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)
		RETURNING id`
	var id int
	row := repo.db.QueryRowContext(
		ctx, stmt,
		rcr.RoleID, rcr.Name, rcr.ClientType, rcr.Description, rcr.RequireID, rcr.CanGuarantee, rcr.CanOpenAccounts, rcr.UseDefaultAddress, rcr.CreateTime, rcr.UpdateTime,
	)
	if err := row.Scan(&id); err != nil {
		return err
	}
	return nil
}

func (repo *Repository) ListRoleClientRoles(ctx context.Context, filter string, pageSize int32, orderBy string, after *db.ListPosition) (db.ListStat, []*db.RoleClientRole, *db.ListPosition, error) {
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
				quote_ident($1) AS ordered_idx, role_client_roles.id,

				role_client_roles.role_id,
				role_client_roles.name,
				role_client_roles.client_type,
				role_client_roles.description,
				role_client_roles.require_id,
				role_client_roles.can_guarantee,
				role_client_roles.can_open_accounts,
				role_client_roles.use_default_address,
				role_client_roles.create_time,
				role_client_roles.update_time
			FROM
				role_client_roles
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
			NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL
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
	rows, err := repo.db.QueryContext(ctx, query, "role_client_roles."+orderBy, after.Data, after.ID, pageSize)
	if err != nil {
		return db.ListStat{}, nil, nil, err
	}

	defer rows.Close()
	result := make([]*db.RoleClientRole, 0, pageSize)
	var nextPos db.ListPosition
	var listStat db.ListStat
	if !rows.Next() {
		return db.ListStat{}, nil, nil, rows.Err()
	}
	var x any
	if err := rows.Scan(
		&listStat.Total, &listStat.Remaining,
		&x, &x,
		&x, &x, &x, &x, &x, &x, &x, &x, &x, &x,
	); err != nil {
		return db.ListStat{}, nil, nil, err
	}
	for rows.Next() {
		var next db.RoleClientRole
		if err := rows.Scan(
			&listStat.Total, &listStat.Remaining,
			&nextPos.Data, &nextPos.ID,
			&next.RoleID, &next.Name, &next.ClientType, &next.Description, &next.RequireID, &next.CanGuarantee, &next.CanOpenAccounts, &next.UseDefaultAddress, &next.CreateTime, &next.UpdateTime,
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

func (repo *Repository) RoleClientRoleByID(ctx context.Context, id int32) (*db.RoleClientRole, error) {
	const stmt = `SELECT
		role_client_roles.id,

		role_client_roles.role_id,
		role_client_roles.name,
		role_client_roles.client_type,
		role_client_roles.description,
		role_client_roles.require_id,
		role_client_roles.can_guarantee,
		role_client_roles.can_open_accounts,
		role_client_roles.use_default_address,
		role_client_roles.create_time,
		role_client_roles.update_time
	FROM
		role_client_roles
	WHERE
		id = $1`

	var rcr db.RoleClientRole
	row := repo.db.QueryRowContext(ctx, stmt, id)
	if err := row.Scan(
		&rcr.ID,
		&rcr.RoleID, &rcr.Name, &rcr.ClientType, &rcr.Description, &rcr.RequireID, &rcr.CanGuarantee, &rcr.CanOpenAccounts, &rcr.UseDefaultAddress, &rcr.CreateTime, &rcr.UpdateTime,
	); err != nil {
		return nil, err
	}

	return &rcr, nil
}

func (repo *Repository) UpdateRoleClientRoleByID(ctx context.Context, rcr *db.RoleClientRole) error {
	const stmt = `UPDATE role_client_roles
	SET role_id=$1,
		name=$2,
		client_type=$3,
		description=$4,
		require_id=$5,
		can_guarantee=$6,
		can_open_accounts=$7,
		use_default_address=$8,
		update_time=$9
	WHERE id = $10`
	_, err := repo.db.ExecContext(
		ctx, stmt,
		rcr.RoleID, rcr.Name, rcr.ClientType, rcr.Description, rcr.RequireID, rcr.CanGuarantee, rcr.CanOpenAccounts, rcr.UseDefaultAddress, rcr.UpdateTime, rcr.ID,
	)
	return err
}

func (repo *Repository) DeleteRoleClientRoleByID(ctx context.Context, id int32) error {
	const stmt = `DELETE FROM role_client_roles
	WHERE id = $1`
	_, err := repo.db.ExecContext(
		ctx, stmt,
		id,
	)
	return err
}

func (repo *Repository) RoleClientRoleByRoleID(ctx context.Context, roleID string) (*db.RoleClientRole, error) {
	const stmt = `SELECT
		role_client_roles.id,

		role_client_roles.role_id,
		role_client_roles.name,
		role_client_roles.client_type,
		role_client_roles.description,
		role_client_roles.require_id,
		role_client_roles.can_guarantee,
		role_client_roles.can_open_accounts,
		role_client_roles.use_default_address,
		role_client_roles.create_time,
		role_client_roles.update_time
	FROM
		role_client_roles
	WHERE
		role_id = $1`

	var rcr db.RoleClientRole
	row := repo.db.QueryRowContext(ctx, stmt, roleID)
	if err := row.Scan(
		&rcr.ID,
		&rcr.RoleID, &rcr.Name, &rcr.ClientType, &rcr.Description, &rcr.RequireID, &rcr.CanGuarantee, &rcr.CanOpenAccounts, &rcr.UseDefaultAddress, &rcr.CreateTime, &rcr.UpdateTime,
	); err != nil {
		return nil, err
	}

	return &rcr, nil
}

func (repo *Repository) UpdateRoleClientRoleByRoleID(ctx context.Context, rcr *db.RoleClientRole) error {
	const stmt = `UPDATE role_client_roles
	SET role_id=$1,
		name=$2,
		client_type=$3,
		description=$4,
		require_id=$5,
		can_guarantee=$6,
		can_open_accounts=$7,
		use_default_address=$8,
		update_time=$9
	WHERE role_id = $10`
	_, err := repo.db.ExecContext(
		ctx, stmt,
		rcr.RoleID, rcr.Name, rcr.ClientType, rcr.Description, rcr.RequireID, rcr.CanGuarantee, rcr.CanOpenAccounts, rcr.UseDefaultAddress, rcr.UpdateTime, rcr.RoleID,
	)
	return err
}

func (repo *Repository) DeleteRoleClientRoleByRoleID(ctx context.Context, roleID string) error {
	const stmt = `DELETE FROM role_client_roles
	WHERE role_id = $1`
	_, err := repo.db.ExecContext(
		ctx, stmt,
		roleID,
	)
	return err
}
