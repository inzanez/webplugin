package main

import (
	"fmt"
	"net/http"
)

// Logger middleware, every request can be logged here
func PluginLogger(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("PLUGIN LOGGER\n")
		next.ServeHTTP(w, r)
	})
}

// Logger middleware, every request can be logged here
func PluginLogger2(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("PLUGIN LOGGER2\n")
		next.ServeHTTP(w, r)
	})
}
