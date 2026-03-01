package demofunction

import "net/http"

func Register(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Shravani Patil"))
}
