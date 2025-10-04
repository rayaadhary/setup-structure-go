-- name: GetUserByUsername :one
SELECT id, username, password, created_at, updated_at
FROM users
WHERE username = $1;