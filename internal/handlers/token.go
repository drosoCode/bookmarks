package handlers

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"github.com/drosocode/bookmarks/internal/auth"
	"github.com/drosocode/bookmarks/internal/config"
	"github.com/drosocode/bookmarks/internal/database"
	"github.com/drosocode/bookmarks/internal/utils"
	"github.com/go-chi/chi/v5"
)

type AddTokenData struct {
	Name string `json:"name"`
}

type ReturnToken struct {
	Token string `json:"token"`
	ID int64 `json:"id"`
}

// handle bookmarks
func SetupTokens(r *chi.Mux) {
	r2 := chi.NewRouter()
	r2.Use(auth.CheckUserMiddleware())
	r.Mount("/token", r2)
	r2.Get("/", ListTokens())
	r2.Post("/", AddToken())
	r2.Delete("/{id}", DeleteToken())
}

// GET token/
func ListTokens() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		userInfo := r.Context().Value(config.CtxUserKey).(auth.UserInfo)

		data, err := database.DB.ListTokens(context.Background(), userInfo.ID)
		if utils.IfError(w, r, err) {
			return
		}

		utils.JSON(w, r, 200, data)
	}
}

// POST token/
func AddToken() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		userInfo := r.Context().Value(config.CtxUserKey).(auth.UserInfo)

		settings := AddTagData{}
		err := json.NewDecoder(r.Body).Decode(&settings)
		if utils.IfError(w, r, err) {
			return
		}

		tok := auth.GenerateToken()
		id, err := database.DB.AddToken(context.Background(), database.AddTokenParams{Name: settings.Name, IDUser: userInfo.ID, AddDate: time.Now().Unix(), Value: tok})
		if utils.IfError(w, r, err) {
			return
		}
		config.Tokens[tok] = userInfo.ID

		utils.JSON(w, r, 200, ReturnToken{ID: id, Token: tok})
	}
}

// DELETE token/{id}
func DeleteToken() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		userInfo := r.Context().Value(config.CtxUserKey).(auth.UserInfo)
		id, err := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)
		if utils.IfError(w, r, err) {
			return
		}

		tok, err := database.DB.GetToken(context.Background(), database.GetTokenParams{IDUser: userInfo.ID, ID: id})
		if utils.IfError(w, r, err) {
			return
		}

		err = database.DB.DeleteToken(context.Background(), database.DeleteTokenParams{IDUser: userInfo.ID, ID: id})
		if utils.IfError(w, r, err) {
			return
		}

		auth.Logout(tok.Value)

		utils.JSON(w, r, 200, "")
	}
}
