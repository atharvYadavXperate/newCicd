package function

import (
	"fmt"
	"net/http"
)

func MainHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain")
	fmt.Fprintln(w, "Main")
}
