// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.16.0
// source: tag.sql

package database

import (
	"context"
)

const addTag = `-- name: AddTag :one
INSERT INTO tag (name, color, id_user)
VALUES ($1, $2, $3)
RETURNING id
`

type AddTagParams struct {
	Name   string `json:"name"`
	Color  string `json:"color"`
	IDUser int64  `json:"idUser"`
}

func (q *Queries) AddTag(ctx context.Context, arg AddTagParams) (int64, error) {
	row := q.db.QueryRowContext(ctx, addTag, arg.Name, arg.Color, arg.IDUser)
	var id int64
	err := row.Scan(&id)
	return id, err
}

const addTagAssoc = `-- name: AddTagAssoc :exec
INSERT INTO tag_link (id_bookmark, id_tag) VALUES ($1, $2)
`

type AddTagAssocParams struct {
	IDBookmark int64 `json:"idBookmark"`
	IDTag      int64 `json:"idTag"`
}

func (q *Queries) AddTagAssoc(ctx context.Context, arg AddTagAssocParams) error {
	_, err := q.db.ExecContext(ctx, addTagAssoc, arg.IDBookmark, arg.IDTag)
	return err
}

const deleteTag = `-- name: DeleteTag :exec
DELETE FROM tag WHERE id = $1 AND id_user = $2
`

type DeleteTagParams struct {
	ID     int64 `json:"id"`
	IDUser int64 `json:"idUser"`
}

func (q *Queries) DeleteTag(ctx context.Context, arg DeleteTagParams) error {
	_, err := q.db.ExecContext(ctx, deleteTag, arg.ID, arg.IDUser)
	return err
}

const deleteTagAssoc = `-- name: DeleteTagAssoc :exec
DELETE FROM tag_link WHERE id_bookmark = $1 AND id_tag = $2
`

type DeleteTagAssocParams struct {
	IDBookmark int64 `json:"idBookmark"`
	IDTag      int64 `json:"idTag"`
}

func (q *Queries) DeleteTagAssoc(ctx context.Context, arg DeleteTagAssocParams) error {
	_, err := q.db.ExecContext(ctx, deleteTagAssoc, arg.IDBookmark, arg.IDTag)
	return err
}

const listTags = `-- name: ListTags :many
SELECT id, name, color FROM tag WHERE id_user = $1
`

type ListTagsRow struct {
	ID    int64  `json:"id"`
	Name  string `json:"name"`
	Color string `json:"color"`
}

func (q *Queries) ListTags(ctx context.Context, idUser int64) ([]ListTagsRow, error) {
	rows, err := q.db.QueryContext(ctx, listTags, idUser)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []ListTagsRow
	for rows.Next() {
		var i ListTagsRow
		if err := rows.Scan(&i.ID, &i.Name, &i.Color); err != nil {
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
