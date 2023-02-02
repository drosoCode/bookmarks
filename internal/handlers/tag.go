package handlers

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/drosocode/bookmarks/internal/auth"
	"github.com/drosocode/bookmarks/internal/config"
	"github.com/drosocode/bookmarks/internal/database"
	"github.com/drosocode/bookmarks/internal/utils"
	"github.com/go-chi/chi/v5"
)

type AddTagData struct {
	Name string `json:"name"`
}

// handle bookmarks
func SetupTags(r *chi.Mux) {
	r2 := chi.NewRouter()
	r2.Use(auth.CheckUserMiddleware())
	r.Mount("/tag", r2)
	r2.Get("/", ListTags())
	r2.Post("/", AddTag())
	r2.Delete("/{id}", DeleteTag())
}

// GET tag/
func ListTags() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		userInfo := r.Context().Value(config.CtxUserKey).(auth.UserInfo)

		data, err := database.DB.ListTags(context.Background(), userInfo.ID)
		if utils.IfError(w, r, err) {
			return
		}

		utils.JSON(w, r, 200, data)
	}
}

// POST tag/
func AddTag() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		userInfo := r.Context().Value(config.CtxUserKey).(auth.UserInfo)

		settings := AddTagData{}
		err := json.NewDecoder(r.Body).Decode(&settings)
		if utils.IfError(w, r, err) {
			return
		}

		id, err := database.DB.AddTag(context.Background(), database.AddTagParams{Name: settings.Name, IDUser: userInfo.ID})
		if utils.IfError(w, r, err) {
			return
		}

		utils.JSON(w, r, 200, ReturnID{ID: id})
	}
}

// DELETE tag/{id}
func DeleteTag() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		userInfo := r.Context().Value(config.CtxUserKey).(auth.UserInfo)
		id, err := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)
		if utils.IfError(w, r, err) {
			return
		}

		err = database.DB.DeleteTag(context.Background(), database.DeleteTagParams{IDUser: userInfo.ID, ID: id})
		if utils.IfError(w, r, err) {
			return
		}

		utils.JSON(w, r, 200, "")
	}
}
