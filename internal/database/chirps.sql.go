// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: chirps.sql

package database

import (
	"context"

	"github.com/google/uuid"
)

const createChirp = `-- name: CreateChirp :one
INSERT INTO chirps (id, created_at, updated_at, user_id, body)
VALUES (gen_random_uuid(), now(), now(), $1, $2)
Returning id, created_at, updated_at, body, user_id
`

type CreateChirpParams struct {
	UserID uuid.UUID
	Body   string
}

func (q *Queries) CreateChirp(ctx context.Context, arg CreateChirpParams) (Chirp, error) {
	row := q.db.QueryRowContext(ctx, createChirp, arg.UserID, arg.Body)
	var i Chirp
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.Body,
		&i.UserID,
	)
	return i, err
}

const deleteChirpByID = `-- name: DeleteChirpByID :exec
DELETE FROM chirps WHERE id = $1
`

func (q *Queries) DeleteChirpByID(ctx context.Context, id uuid.UUID) error {
	_, err := q.db.ExecContext(ctx, deleteChirpByID, id)
	return err
}

const getAllChirps = `-- name: GetAllChirps :many
SELECT id, created_at, updated_at, body, user_id FROM chirps ORDER BY created_at
`

func (q *Queries) GetAllChirps(ctx context.Context) ([]Chirp, error) {
	rows, err := q.db.QueryContext(ctx, getAllChirps)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Chirp
	for rows.Next() {
		var i Chirp
		if err := rows.Scan(
			&i.ID,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.Body,
			&i.UserID,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getAllChirpsDesc = `-- name: GetAllChirpsDesc :many
SELECT id, created_at, updated_at, body, user_id FROM chirps ORDER BY created_at DESC
`

func (q *Queries) GetAllChirpsDesc(ctx context.Context) ([]Chirp, error) {
	rows, err := q.db.QueryContext(ctx, getAllChirpsDesc)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Chirp
	for rows.Next() {
		var i Chirp
		if err := rows.Scan(
			&i.ID,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.Body,
			&i.UserID,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getChirpByID = `-- name: GetChirpByID :one
SELECT id, created_at, updated_at, body, user_id FROM chirps WHERE id = $1
`

func (q *Queries) GetChirpByID(ctx context.Context, id uuid.UUID) (Chirp, error) {
	row := q.db.QueryRowContext(ctx, getChirpByID, id)
	var i Chirp
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.Body,
		&i.UserID,
	)
	return i, err
}

const getChirpsByUserID = `-- name: GetChirpsByUserID :many
SELECT id, created_at, updated_at, body, user_id FROM chirps WHERE user_id = $1 ORDER BY $2
`

type GetChirpsByUserIDParams struct {
	UserID  uuid.UUID
	Column2 interface{}
}

func (q *Queries) GetChirpsByUserID(ctx context.Context, arg GetChirpsByUserIDParams) ([]Chirp, error) {
	rows, err := q.db.QueryContext(ctx, getChirpsByUserID, arg.UserID, arg.Column2)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Chirp
	for rows.Next() {
		var i Chirp
		if err := rows.Scan(
			&i.ID,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.Body,
			&i.UserID,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const resetAllChirps = `-- name: ResetAllChirps :exec
DELETE FROM chirps
`

func (q *Queries) ResetAllChirps(ctx context.Context) error {
	_, err := q.db.ExecContext(ctx, resetAllChirps)
	return err
}
