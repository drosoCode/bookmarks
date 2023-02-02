package main

import (
	"embed"
	"fmt"
	"io/fs"
	"net/http"

	"github.com/drosocode/bookmarks/internal/auth"
	"github.com/drosocode/bookmarks/internal/config"
	"github.com/drosocode/bookmarks/internal/database"
	"github.com/drosocode/bookmarks/internal/processor"
	handler "github.com/drosocode/bookmarks/internal/handlers"
	_ "github.com/lib/pq"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
)

//go:embed api static
var embedFS embed.FS

// go run github.com/playwright-community/playwright-go/cmd/playwright install --with-deps
// go run .\main.go --db-host="10.10.2.1" --db-user="bookmarks" --db-password="secret" --db-name="bookmarks" --registration=true --serve="127.0.0.1:9000"

func main() {
	config.ParseConfig()
	database.ConfigDB(config.Data.DB)
	config.Tokens = map[string]int64{}
	config.CtxUserKey = config.CtxKey("userinfo")
	auth.LoadTokens()

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300,
	}))

	api := chi.NewRouter()

	r.Mount("/api", api)
	handler.SetupUsers(api)
	handler.SetupBookmarks(api)
	handler.SetupTags(api)
	handler.SetupTokens(api)

	staticFS := fs.FS(embedFS)

	apiDir, _ := fs.Sub(staticFS, "api")
	handler.ServeStatic(r, "/swagger", http.FS(apiDir))

	staticDir, _ := fs.Sub(staticFS, "static")
	handler.ServeStatic(r, "/", http.FS(staticDir))

	fmt.Println("Ready !")
	processor.Test("https://github.com/playwright-community/playwright-go/blob/main/examples/parallel-scraping/main.go")

	err := http.ListenAndServe(config.Data.Serve, r)
	fmt.Println(err)
}
