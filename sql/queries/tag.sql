-- name: ListTags :many
SELECT id, name, color FROM tag WHERE id_user = $1;

-- name: AddTag :one
INSERT INTO tag (name, color, id_user)
VALUES ($1, $2, $3)
RETURNING id;

-- name: DeleteTag :exec
DELETE FROM tag WHERE id = $1 AND id_user = $2;

-- name: AddTagAssoc :exec
INSERT INTO tag_link (id_bookmark, id_tag) VALUES ($1, $2);

-- name: DeleteTagAssoc :exec
DELETE FROM tag_link WHERE id_bookmark = $1 AND id_tag = $2;