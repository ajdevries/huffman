// +build OMIT
package main

import (
	"fmt"
	"net/http"
)

// START0 OMIT

func main() {
	http.Handle("/", Middleware(IndexHandler()))
}

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello")
}

func IndexHandler() http.Handler { return http.HandlerFunc(Index) }

// STOP0 OMIT

// START1 OMIT
func Middleware(next http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// before
		next.ServeHTTP(w, r) // HL
		// after
	})
}

// STOP1 OMIT
