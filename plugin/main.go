package main

import (
	"net/http"
	"sftse/test/pkg/routes"

	"github.com/husobee/vestigo"
)

func ByeHandler(w http.ResponseWriter, r *http.Request) {
	user := r.Context().Value("user")

	if user == nil {
		w.Write([]byte("Bye there, unknown"))
	} else {
		t, _ := user.(string)
		w.Write([]byte("Bye there, " + t))
	}
}

type Router struct{}

var Mrouter Router

// Returns an 'http.Handler' that contains all the route definitions
func (ro *Router) Routes() []routes.Route {
	myRoutes := make([]routes.Route, 0)

	r := routes.Route{}
	r.Handler = ByeHandler
	r.Method = "GET"
	r.Path = "/api/bye"
	r.Middleware = []vestigo.Middleware{
		PluginLogger,
		PluginLogger2,
	}

	myRoutes = append(myRoutes, r)

	return myRoutes
}
