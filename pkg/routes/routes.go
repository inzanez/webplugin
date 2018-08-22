package routes

import (
	"net/http"

	"github.com/husobee/vestigo"
)

type Route struct {
	Handler    http.HandlerFunc
	Path       string
	Method     string
	Middleware []vestigo.Middleware
}
