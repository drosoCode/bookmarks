package auth

import (
	"context"
	"crypto/rand"
	"encoding/hex"
	"errors"
	"net/http"

	"github.com/drosocode/bookmarks/internal/config"
	"github.com/drosocode/bookmarks/internal/database"
)

// Get the User ID from a token, returns an error if the association doesn't exists
func GetUserID(token string) (int64, error) {
	if uid, ok := config.Tokens[token]; ok {
		return uid, nil
	}
	return -1, errors.New("this token is not associated with a user")
}

// Function to retreive the user's token in a request, returns an error if not found
func GetToken(r *http.Request) (string, error) {
	var token string

	// check in the url ex: api/endpoint?token=xxxxx
	token = r.URL.Query().Get("token")
	if token != "" {
		return token, nil
	}

	// check in the headers ex: Authorization: Bearer xxxxx
	token = r.Header.Get("Authorization")
	if token != "" {
		return token[7:], nil
	}

	// check in the cookies
	tok, err := r.Cookie("bookmarktoken")
	if err == nil && tok.Value != "" {
		return tok.Value, nil
	}

	return "", errors.New("token not found")
}

// Generate a random token
func GenerateToken() string {
	bytes := make([]byte, 64)
	rand.Read(bytes)
	return hex.EncodeToString(bytes)
}

// load existing long lived tokens from the DB
func LoadTokens() error {
	data, err := database.DB.ListTokenVals(context.Background())
	if err != nil {
		return err
	}
	for _, x := range data {
		config.Tokens[x.Value] = x.IDUser
	}
	return nil
}
