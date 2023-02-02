-- name: ListTokens :many
SELECT id, name, add_date
FROM token
WHERE id_user = $1;

-- name: AddToken :one
INSERT INTO token ("name", add_date, "value", id_user)
VALUES ($1, $2, $3, $4)
RETURNING id;

-- name: DeleteToken :exec
DELETE FROM token WHERE id = $1 AND id_user = $2;

-- name: GetToken :one
SELECT * FROM token WHERE id = $1 AND id_user = $2;

-- name: ListTokenVals :many
SELECT id_user, value FROM token;