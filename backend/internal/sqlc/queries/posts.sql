-- name: CreatePost :one
INSERT INTO posts (title, content)
VALUES ($1, $2)
RETURNING id, title, content, created_at, updated_at;

-- name: GetPost :one
SELECT id, title, content, created_at, updated_at
FROM posts
WHERE id = $1
LIMIT 1;

-- name: ListPosts :many
SELECT id, title, content, created_at, updated_at
FROM posts
ORDER BY created_at DESC
LIMIT $1 OFFSET $2;

-- name: UpdatePost :one
UPDATE posts
SET title = $1,
    content = $2,
    updated_at = now()
WHERE id = $3
RETURNING id, title, content, created_at, updated_at;

-- name: DeletePost :exec
DELETE FROM posts
WHERE id = $1;
