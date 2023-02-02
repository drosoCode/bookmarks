-- name: ListBookmarks :many
SELECT id,
    link,
    name,
    description,
    save,
    add_date,
    array_agg(id_tag)::bigint[] AS tags
FROM bookmark b JOIN tag_link t ON (b.id = t.id_bookmark)
WHERE id_user = $1
GROUP BY id;

-- name: GetBookmark :one
SELECT id,
    link,
    name,
    description,
    save,
    add_date,
    array_agg(id_tag)::bigint[] AS tags
FROM bookmark b JOIN tag_link t ON (b.id = t.id_bookmark)
WHERE id_user = $1
    AND id = $2
GROUP BY id;

-- name: AddBookmark :one
INSERT INTO bookmark (link, name, description, save, add_date, id_user)
VALUES ($1, '', '', $2, $3, $4)
RETURNING id;

-- name: SetBookmarkData :exec
UPDATE bookmark
SET name = $1, description = $2
WHERE id = $3;

-- name: DeleteBookmark :exec
DELETE FROM bookmark WHERE id = $1 AND id_user = $2;

