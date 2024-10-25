package api

import (
	"fmt"
	"net/http"
)

func (a *api) healthCheck(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "server is up and running")
}
