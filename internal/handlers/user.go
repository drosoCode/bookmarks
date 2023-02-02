package handlers

import (
	"net/http"

	"github.com/drosocode/bookmarks/internal/auth"
	"github.com/drosocode/bookmarks/internal/utils"
	"github.com/go-chi/chi/v5"
)

// handle users
func SetupUsers(r *chi.Mux) {
	r2 := chi.NewRouter()
	r.Mount("/user", r2)
	r2.Get("/login", LoginUser())
	r2.Get("/logout", LogoutUser())
}

type UserLogin struct {
	Token string `json:"token"`
}

// GET user/login
func LoginUser() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		uid, err := auth.Authenticate(r)
		if utils.IfError(w, r, err) {
			return
		}
		token := auth.Login(uid)

		utils.JSON(w, r, 200, UserLogin{Token: token})
	}
}

// GET user/logout
func LogoutUser() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		token, err := auth.GetToken(r)
		if utils.IfError(w, r, err) {
			return
		}
		auth.Logout(token)

		utils.JSON(w, r, 200, "ok")
	}
}
