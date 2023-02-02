-- name: AddUser :one
INSERT INTO users (username, name) VALUES ($1, $2) RETURNING id;

-- name: GetUser :one
SELECT id, name FROM users WHERE username = $1;

-- name: DeleteUser :exec
DELETE FROM users WHERE id = $1;