// Code generated by sqlc. DO NOT EDIT.
// source: storage.sql

package db

import (
	"context"
)

const createStorageItem = `-- name: CreateStorageItem :one
INSERT INTO storage (key, value, created_by) VALUES ($1, $2, $3) RETURNING id, key, created_by
`

type CreateStorageItemParams struct {
	Key       string `json:"key"`
	Value     string `json:"value"`
	CreatedBy int32  `json:"created_by"`
}

type CreateStorageItemRow struct {
	ID        int32  `json:"id"`
	Key       string `json:"key"`
	CreatedBy int32  `json:"created_by"`
}

func (q *Queries) CreateStorageItem(ctx context.Context, arg CreateStorageItemParams) (CreateStorageItemRow, error) {
	row := q.db.QueryRowContext(ctx, createStorageItem, arg.Key, arg.Value, arg.CreatedBy)
	var i CreateStorageItemRow
	err := row.Scan(&i.ID, &i.Key, &i.CreatedBy)
	return i, err
}

const deleteStorageItemByKey = `-- name: DeleteStorageItemByKey :exec
DELETE FROM storage WHERE key = $1
`

func (q *Queries) DeleteStorageItemByKey(ctx context.Context, key string) error {
	_, err := q.db.ExecContext(ctx, deleteStorageItemByKey, key)
	return err
}

const deleteStorageItemsByUserID = `-- name: DeleteStorageItemsByUserID :exec
DELETE FROM storage WHERE created_by = $1
`

func (q *Queries) DeleteStorageItemsByUserID(ctx context.Context, createdBy int32) error {
	_, err := q.db.ExecContext(ctx, deleteStorageItemsByUserID, createdBy)
	return err
}

const getStorageItemByKey = `-- name: GetStorageItemByKey :one
SELECT id, key, value, available, created_by, downloads, errors FROM storage
WHERE key = $1 LIMIT 1
`

func (q *Queries) GetStorageItemByKey(ctx context.Context, key string) (Storage, error) {
	row := q.db.QueryRowContext(ctx, getStorageItemByKey, key)
	var i Storage
	err := row.Scan(
		&i.ID,
		&i.Key,
		&i.Value,
		&i.Available,
		&i.CreatedBy,
		&i.Downloads,
		&i.Errors,
	)
	return i, err
}

const getStorageItemsByUserID = `-- name: GetStorageItemsByUserID :many
SELECT (id, key) FROM storage
WHERE created_by = $1
`

func (q *Queries) GetStorageItemsByUserID(ctx context.Context, createdBy int32) ([]interface{}, error) {
	rows, err := q.db.QueryContext(ctx, getStorageItemsByUserID, createdBy)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []interface{}{}
	for rows.Next() {
		var column_1 interface{}
		if err := rows.Scan(&column_1); err != nil {
			return nil, err
		}
		items = append(items, column_1)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getStorageItemsByUsername = `-- name: GetStorageItemsByUsername :many
SELECT (id, key) FROM storage
WHERE created_by = (SELECT id FROM users WHERE users.username = $1 LIMIT 1)
`

func (q *Queries) GetStorageItemsByUsername(ctx context.Context, username string) ([]interface{}, error) {
	rows, err := q.db.QueryContext(ctx, getStorageItemsByUsername, username)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []interface{}{}
	for rows.Next() {
		var column_1 interface{}
		if err := rows.Scan(&column_1); err != nil {
			return nil, err
		}
		items = append(items, column_1)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
