package function

import (
	"fmt"
	"net/http"
)

func OrdersHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain")
	fmt.Fprintln(w, "Order s")
}
