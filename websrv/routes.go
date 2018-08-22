package main

import (
	"fmt"
	"net/http"
	"plugin"
	"sftse/test/pkg/routes"

	"github.com/husobee/vestigo"
)

type ModuleRouter interface {
	Routes() []routes.Route
}

var MiddleWare []vestigo.Middleware

// Returns an 'http.Handler' that contains all the route definitions
func Routes() http.Handler {
	MiddleWare = make([]vestigo.Middleware, 0)
	router := vestigo.NewRouter()

	// CORS for testing purpose, should be disabled in production use!
	router.SetGlobalCors(&vestigo.CorsAccessControl{
		AllowOrigin:      []string{"http://localhost:8080"},
		AllowCredentials: true,
		ExposeHeaders:    []string{"X-Header", "X-Y-Header"},
		AllowHeaders:     []string{"authorization", "cache-control", "x-requested-with", "Content-Type", "set-cookie"},
	})

	router.Add("GET", "/api/hello", HelloHandler, DefaultChain...)

	plug, err := plugin.Open("./plugin.so")

	if err != nil {
		fmt.Printf("Could not load plugin")
	}

	rs, err := plug.Lookup("Mrouter")

	if err != nil {
		fmt.Printf("Could not lookup plugin")
	}

	mr, ok := rs.(ModuleRouter)

	if !ok {
		fmt.Printf("Cast fail\n")
	}

	mroutes := mr.Routes()

	for _, t := range mroutes {

		MiddleWare = append(MiddleWare, DefaultChain...)
		MiddleWare = append(MiddleWare, t.Middleware...)

		fmt.Printf("%v\n", len(MiddleWare))
		router.Add(t.Method, t.Path, t.Handler, MiddleWare...)
	}

	return router
}
