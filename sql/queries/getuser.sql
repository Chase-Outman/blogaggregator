-- name: GetUser :one
SELECT * from users
WHERE name = $1;