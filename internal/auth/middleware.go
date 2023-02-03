package auth

import (
	"context"
	"net/http"
	"regexp"
	"strconv"

	"github.com/drosocode/bookmarks/internal/config"
	"github.com/drosocode/bookmarks/internal/database"
	"github.com/drosocode/bookmarks/internal/utils"
)

type UserInfo struct {
	Token string
	ID    int64
}

// check user authentication and authorization and add user id/token to the request context
func CheckUserMiddleware() func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		fn := func(w http.ResponseWriter, r *http.Request) {
			// get user token from request
			token, err := GetToken(r)
			if utils.IfError(w, r, err) {
				return
			}
			// get user id from token
			uid, err := GetUserID(token)
			if utils.IfError(w, r, err) {
				return
			}
			// if all ok, add token and userid to context and continue
			r = r.WithContext(context.WithValue(r.Context(), config.CtxUserKey, UserInfo{Token: token, ID: uid}))
			next.ServeHTTP(w, r)
		}
		return http.HandlerFunc(fn)
	}
}

func CheckPathMiddleware() func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		fn := func(w http.ResponseWriter, r *http.Request) {
			rg := regexp.MustCompile(`^\/cache\/(\d*)\/`)
			if data := rg.FindStringSubmatch(r.URL.String()); len(data) > 1 {
				id, err := strconv.ParseInt(data[1], 10, 64)
				if utils.IfError(w, r, err) {
					return
				}
				userInfo := r.Context().Value(config.CtxUserKey).(UserInfo)
				_, err = database.DB.GetBookmark(context.Background(), database.GetBookmarkParams{IDUser: userInfo.ID, ID: id})
				if err != nil {
					utils.Error(w, r, 403, "unauthorized")
					return
				}
			}
			next.ServeHTTP(w, r)
		}
		return http.HandlerFunc(fn)
	}
}
