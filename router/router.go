// Defines the Route interface, and registers routes to a server
package router

import (
	"net/http"
	"time"

	"github.com/KyleBanks/go-kit/log"
)

// Interface for the provided server to comply with
type Server interface {
	HandleFunc(string, func(http.ResponseWriter, *http.Request))
}

// Defines an executable Route
type Route struct {
	Path   string // The URL path to listen for (i.e. "/api")
	Handle func(w http.ResponseWriter, r *http.Request)
}

// Register registers each Route with the Server provided.
//
// Each Route will be wrapped in a middleware function that adds trace logging.
func Register(s Server, routes []Route) {
	for _, route := range routes {
		log.Info("Registering route:", route.Path)
		s.HandleFunc(route.Path, handleWrapper(route))
	}
}

// handleWrapper returns a request handling function that wraps the provided route.
func handleWrapper(route Route) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		start := time.Now().Unix()

		log.Info("START:", r.URL.Path, r.URL.RawQuery, r.PostForm)
		route.Handle(w, r)
		log.Info("END:", r.URL.Path, r.URL.RawQuery, r.PostForm, time.Now().Unix()-start)
	}
}

// Param returns a POST/GET parameter from the request.
//
// If the parameter is found in the POST and the GET parameter set, the POST parameter
// will be given priority.
func Param(r *http.Request, key string) string {
	val := r.PostForm.Get(key)
	if len(val) != 0 {
		return val
	}

	return r.URL.Query().Get(key)
}
