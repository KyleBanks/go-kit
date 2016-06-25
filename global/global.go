// The global package manages the initialization and access of various
// global variables, such as database and cache connections.
package global

import (
	"net/http"
	"time"

	"github.com/KyleBanks/go-kit/auth"
	"github.com/KyleBanks/go-kit/cache"
	"github.com/KyleBanks/go-kit/log"
	"github.com/KyleBanks/go-kit/orm"
	"github.com/KyleBanks/go-kit/router"
)

var (
	DB     *orm.ORM
	Server *http.Server
	Cache  auth.Cache
)

// Initializes the ORM and registers models.
func InitORM(username string, password string, database string, models []interface{}) {
	DB = &orm.ORM{
		Username: username,
		Password: password,
		Database: database,
	}
	DB.Open()

	if err := DB.AutoMigrate(models); err != nil {
		panic(err)
	}
}

// InitCache initializes the application cache.
func InitCache(host string) {
	Cache = cache.New(host)
}

// Initializes the server and registers routes
func InitServer(routes []router.Route) {
	Server = &http.Server{
		Addr:         ":8080",
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	// Register our routes with the router
	router.Register(http.DefaultServeMux, routes)

	// Start the server
	log.Info("Application running on", Server.Addr)
	log.Info(Server.ListenAndServe())
}
