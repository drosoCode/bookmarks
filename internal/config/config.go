package config

import (
	"flag"
	"strings"

	"github.com/drosocode/bookmarks/internal/database"
	"github.com/jamiealquiza/envy"
)

type Config struct {
	Serve              string
	RemoteUserHeader   string
	RemoteGroupHeader  string
	RemoteNameHeader   string
	TrustedSources     []string
	AllowedGroups      []string
	EnableRegistration bool
	CachePath          string
	DB                 database.DBMSConn
}

type CtxKey string

var Tokens map[string]int64

var Data Config

var CtxUserKey CtxKey

func ParseConfig() {
	serve := flag.String("serve", "0.0.0.0:9000", "bind address")
	remoteUser := flag.String("user-header", "Remote-User", "name of the username header")
	remoteGroup := flag.String("group-header", "Remote-Group", "name of the groups header")
	remoteName := flag.String("name-header", "Remote-Name", "name of the name header")
	allowedGroups := flag.String("allowed-groups", "bookmark", "name of the allowed groups (separated by ;)")
	trustedSources := flag.String("trusted-sources", "0.0.0.0/0", "allowed sources (separated by ;)")
	registration := flag.Bool("registration", false, "enable registration")
	cachePath := flag.String("cache", "./cache", "path to the cache")
	dbHost := flag.String("db-host", "", "address to your postgresql server")
	dbUser := flag.String("db-user", "", "db username")
	dbPass := flag.String("db-password", "", "db password")
	dbPort := flag.Int("db-port", 5432, "postgres port")
	dbName := flag.String("db-name", "bookmarks", "database name")

	envy.Parse("BM")
	flag.Parse()

	Data = Config{
		Serve:              *serve,
		RemoteUserHeader:   *remoteUser,
		RemoteGroupHeader:  *remoteGroup,
		RemoteNameHeader:   *remoteName,
		TrustedSources:     strings.Split(*trustedSources, ";"),
		AllowedGroups:      strings.Split(*allowedGroups, ";"),
		EnableRegistration: *registration,
		CachePath:          *cachePath,
		DB: database.DBMSConn{
			Host:     *dbHost,
			Port:     *dbPort,
			User:     *dbUser,
			Password: *dbPass,
			Name:     *dbName,
		},
	}
}
