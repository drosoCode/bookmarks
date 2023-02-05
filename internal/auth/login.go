package auth

import (
	"context"
	"errors"
	"fmt"
	"net"
	"net/http"
	"strings"

	"github.com/drosocode/bookmarks/internal/config"
	"github.com/drosocode/bookmarks/internal/database"
	"github.com/drosocode/bookmarks/internal/utils"
)

// Function to authenticate the user, through trusted header auth
func Authenticate(r *http.Request) (int64, error) {
	source := net.ParseIP(strings.Split(r.RemoteAddr, ":")[0])
	proxies := config.Data.TrustedSources
	fmt.Println(source)
	fmt.Println(proxies)
	authorizedSource := false
	for _, p := range proxies {
		_, ipnet, err := net.ParseCIDR(p)
		if err == nil && ipnet.Contains(source) {
			authorizedSource = true
			break
		}
	}

	fmt.Printf("%+v\n", r.Header)
	if !authorizedSource {
		return -1, errors.New("source is not authorized to send authentication headers")
	}

	if len(config.Data.AllowedGroups) > 0 {
		ok := false
		groups := strings.Split(r.Header.Get(config.Data.RemoteGroupHeader), ",")
		for _, g := range config.Data.AllowedGroups {
			if utils.Contains(groups, g) {
				ok = true
				break
			}
		}
		if !ok {
			fmt.Println(config.Data.AllowedGroups)
			fmt.Println(groups)
			return -1, errors.New("unauthorized group")
		}
	}

	userData, err := database.DB.GetUser(context.Background(), r.Header.Get(config.Data.RemoteUserHeader))
	if err != nil {
		// auto register
		if config.Data.EnableRegistration {
			database.DB.AddUser(context.Background(), database.AddUserParams{Username: r.Header.Get(config.Data.RemoteUserHeader), Name: r.Header.Get(config.Data.RemoteNameHeader)})
		} else {
			return -1, errors.New("registration is disabled")
		}
	}
	return userData.ID, nil
}

// Function to grant access to a user, associates a User ID to a token and returns this token
func Login(uid int64) string {
	token := GenerateToken()
	config.Tokens[token] = uid
	return token
}

// Function to logout from a session, removes the association between a specific token and a User ID
func Logout(token string) {
	delete(config.Tokens, token)
}
