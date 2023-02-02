package auth

import (
	"context"
	"net/http"

	"github.com/drosocode/bookmarks/internal/config"
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
