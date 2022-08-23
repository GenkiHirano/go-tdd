package context1

import (
	"fmt"
	"net/http"
)

type Store interface {
	Fetch() string
	Cansel()
}

func Server(store Store) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		store.Cansel()
		fmt.Fprint(w, store.Fetch())
	}
}
