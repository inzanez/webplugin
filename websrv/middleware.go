package main

import (
	"context"
	"fmt"
	"net/http"

	"github.com/husobee/vestigo"
)

var DefaultChain = []vestigo.Middleware{
	AuthMiddle,
	Logger,
	// Uncomment this or comment previous line and it will break
	//	Logger2,
}

func AuthMiddle(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), "user", "Somebody")
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

// Logger middleware, every request can be logged here
func Logger(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("ACCESS\n")
		next.ServeHTTP(w, r)
	})
}

// Logger middleware, every request can be logged here
func Logger2(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("ACCESS2\n")
		next.ServeHTTP(w, r)
	})
}
