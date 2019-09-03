package controller

import (
	"net/http"
	"fmt"
)

// index
func Index(w http.ResponseWriter, r *http.Request) {
	_, _ = fmt.Fprintln(w, "Welcome!")
}
