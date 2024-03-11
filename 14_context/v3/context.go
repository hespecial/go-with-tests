package main

import (
	"context"
	"fmt"
	"net/http"
)

// Store fetches data.
type Store interface {
	Fetch(context.Context) (string, error)
}

// Server returns a handler for calling Store.
func Server(store Store) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		data, _ := store.Fetch(r.Context())

		if data == "" {
			return
		}

		fmt.Fprint(w, data)
	}
}
