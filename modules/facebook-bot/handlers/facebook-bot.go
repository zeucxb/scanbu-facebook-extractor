package handlers

import (
	"fmt"
	"net/http"
)

// FacebookBot is the facebook bot webhook handler
func FacebookBot(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "%s", "scanbu-s3Cr3t")
}
