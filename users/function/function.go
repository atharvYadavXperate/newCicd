package userfunction

import (
	"fmt"
	"net/http"
)

func UsersHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Users")
}
