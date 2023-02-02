// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.16.0

package database

import ()

type Bookmark struct {
	ID          int64  `json:"id"`
	Link        string `json:"link"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Save        bool   `json:"save"`
	AddDate     int64  `json:"addDate"`
	IDUser      int64  `json:"idUser"`
}

type Tag struct {
	ID     int64  `json:"id"`
	Name   string `json:"name"`
	Color  string `json:"color"`
	IDUser int64  `json:"idUser"`
}

type TagLink struct {
	IDBookmark int64 `json:"idBookmark"`
	IDTag      int64 `json:"idTag"`
}

type Token struct {
	ID      int64  `json:"id"`
	Name    string `json:"name"`
	AddDate int64  `json:"addDate"`
	Value   string `json:"value"`
	IDUser  int64  `json:"idUser"`
}

type User struct {
	ID       int64  `json:"id"`
	Username string `json:"username"`
	Name     string `json:"name"`
}
