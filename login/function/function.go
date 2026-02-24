package login

import (
	"net/http"

	customerror "github.com/atharvYadavXperate/newCicd/kapad-app/domain/errors"
)

func Login(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, customerror.ErrMethodNotAllows.Error(), http.StatusBadRequest)
		return
	}
	w.Write([]byte("Hello World"))
}
