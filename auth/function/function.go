package userfunction

import (
	"fmt"
	"net/http"
)

func AuthHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain")
	fmt.Fprintln(w, "auth")
}
