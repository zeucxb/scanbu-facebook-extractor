package handlers

import (
	"fmt"
	"net/http"

	"github.com/pressly/chi/render"
)

// FacebookBot is the facebook bot webhook handler
func FacebookBot(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()

	verifyToken := query.Get("hub.verify_token")

	if verifyToken != "scanbu-S3cr3t" {
		render.Status(r, http.StatusUnauthorized)
		return
	}

	fmt.Fprintf(w, "%s", query.Get("hub.challenge"))
}
