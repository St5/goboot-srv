// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: users.sql

package database

import (
	"context"
	"database/sql"

	"github.com/google/uuid"
)

const createUser = `-- name: CreateUser :one
INSERT INTO users (id, created_at, updated_at, email, hashed_password)
VALUES (gen_random_uuid(), now(), now(), $1, $2)
Returning id, email, created_at, updated_at, hashed_password, is_chirpy_red
`

type CreateUserParams struct {
	Email          string
	HashedPassword string
}

func (q *Queries) CreateUser(ctx context.Context, arg CreateUserParams) (User, error) {
	row := q.db.QueryRowContext(ctx, createUser, arg.Email, arg.HashedPassword)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Email,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.HashedPassword,
		&i.IsChirpyRed,
	)
	return i, err
}

const getUserByEmail = `-- name: GetUserByEmail :one
SELECT id, email, created_at, updated_at, hashed_password, is_chirpy_red FROM users WHERE email = $1
`

func (q *Queries) GetUserByEmail(ctx context.Context, email string) (User, error) {
	row := q.db.QueryRowContext(ctx, getUserByEmail, email)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Email,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.HashedPassword,
		&i.IsChirpyRed,
	)
	return i, err
}

const getUserByID = `-- name: GetUserByID :one
SELECT id, email, created_at, updated_at, hashed_password, is_chirpy_red FROM users WHERE id = $1
`

func (q *Queries) GetUserByID(ctx context.Context, id uuid.UUID) (User, error) {
	row := q.db.QueryRowContext(ctx, getUserByID, id)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Email,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.HashedPassword,
		&i.IsChirpyRed,
	)
	return i, err
}

const resetAllUsers = `-- name: ResetAllUsers :exec
DELETE FROM users
`

func (q *Queries) ResetAllUsers(ctx context.Context) error {
	_, err := q.db.ExecContext(ctx, resetAllUsers)
	return err
}

const updateChirpyRedByUserID = `-- name: UpdateChirpyRedByUserID :exec
UPDATE users SET is_chirpy_red = $1, updated_at = now()
WHERE id = $2
`

type UpdateChirpyRedByUserIDParams struct {
	IsChirpyRed sql.NullBool
	ID          uuid.UUID
}

func (q *Queries) UpdateChirpyRedByUserID(ctx context.Context, arg UpdateChirpyRedByUserIDParams) error {
	_, err := q.db.ExecContext(ctx, updateChirpyRedByUserID, arg.IsChirpyRed, arg.ID)
	return err
}

const updateUser = `-- name: UpdateUser :one
UPDATE users SET email = $1, hashed_password = $2, updated_at = now()
WHERE id = $3
RETURNING id, email, created_at, updated_at, hashed_password, is_chirpy_red
`

type UpdateUserParams struct {
	Email          string
	HashedPassword string
	ID             uuid.UUID
}

func (q *Queries) UpdateUser(ctx context.Context, arg UpdateUserParams) (User, error) {
	row := q.db.QueryRowContext(ctx, updateUser, arg.Email, arg.HashedPassword, arg.ID)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Email,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.HashedPassword,
		&i.IsChirpyRed,
	)
	return i, err
}
