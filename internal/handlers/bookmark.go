package handlers

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"
	"time"
	"path"
	"os"

	"github.com/drosocode/bookmarks/internal/auth"
	"github.com/drosocode/bookmarks/internal/config"
	"github.com/drosocode/bookmarks/internal/database"
	"github.com/drosocode/bookmarks/internal/processor"
	"github.com/drosocode/bookmarks/internal/utils"
	"github.com/go-chi/chi/v5"
)

type AddBookmarkData struct {
	Link string  `json:"link"`
	Tags []int64 `json:"tags"`
	Save bool    `json:"save"`
}

type ReturnID struct {
	ID int64 `json:"id"`
}

// handle bookmarks
func SetupBookmarks(r *chi.Mux) {
	r2 := chi.NewRouter()
	r2.Use(auth.CheckUserMiddleware())
	r.Mount("/bookmark", r2)
	r2.Get("/", ListBookmarks())
	r2.Post("/", AddBookmark())
	r2.Get("/{id}", GetBookmark())
	r2.Delete("/{id}", DeleteBookmark())
}

// GET bookmark/
func ListBookmarks() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		userInfo := r.Context().Value(config.CtxUserKey).(auth.UserInfo)

		data, err := database.DB.ListBookmarks(context.Background(), userInfo.ID)
		if utils.IfError(w, r, err) {
			return
		}

		utils.JSON(w, r, 200, data)
	}
}

// POST bookmark/
func AddBookmark() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		userInfo := r.Context().Value(config.CtxUserKey).(auth.UserInfo)

		settings := AddBookmarkData{}
		err := json.NewDecoder(r.Body).Decode(&settings)
		if utils.IfError(w, r, err) {
			return
		}

		id, err := database.DB.AddBookmark(context.Background(), database.AddBookmarkParams{Link: settings.Link, Save: settings.Save, AddDate: time.Now().Unix(), IDUser: userInfo.ID})
		if utils.IfError(w, r, err) {
			return
		}
		for _, t := range settings.Tags {
			database.DB.AddTagAssoc(context.Background(), database.AddTagAssocParams{IDTag: t, IDBookmark: id})
		}
		processor.AddBookmark(id, settings.Link, settings.Save)

		utils.JSON(w, r, 200, ReturnID{ID: id})
	}
}

// GET bookmark/{id}
func GetBookmark() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		userInfo := r.Context().Value(config.CtxUserKey).(auth.UserInfo)
		id, err := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)
		if utils.IfError(w, r, err) {
			return
		}

		data, err := database.DB.GetBookmark(context.Background(), database.GetBookmarkParams{IDUser: userInfo.ID, ID: id})
		if utils.IfError(w, r, err) {
			return
		}

		utils.JSON(w, r, 200, data)
	}
}

// DELETE bookmark/{id}
func DeleteBookmark() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		userInfo := r.Context().Value(config.CtxUserKey).(auth.UserInfo)
		id, err := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)
		if utils.IfError(w, r, err) {
			return
		}

		err = database.DB.DeleteBookmark(context.Background(), database.DeleteBookmarkParams{IDUser: userInfo.ID, ID: id})
		if utils.IfError(w, r, err) {
			return
		}
		os.RemoveAll(path.Join(config.Data.CachePath, string(id)))

		utils.JSON(w, r, 200, "")
	}
}
