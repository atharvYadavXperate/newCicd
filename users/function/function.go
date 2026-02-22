package function

import (
	"fmt"
	"net/http"
)

func usersHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain")
	fmt.Fprintln(w, "Users")
}
