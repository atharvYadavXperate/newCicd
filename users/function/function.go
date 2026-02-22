package function

import (
	"fmt"
	"net/http"
)

func usersHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Users")
}
