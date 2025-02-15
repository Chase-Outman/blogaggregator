// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.28.0
// source: getuserbyid.sql

package database

import (
	"context"

	"github.com/google/uuid"
)

const getUserById = `-- name: GetUserById :one
SELECT id, created_at, updated_at, name FROM users
WHERE id = $1
`

func (q *Queries) GetUserById(ctx context.Context, id uuid.UUID) (User, error) {
	row := q.db.QueryRowContext(ctx, getUserById, id)
	var i User
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.Name,
	)
	return i, err
}
