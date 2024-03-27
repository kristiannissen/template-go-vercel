package handler

import (
	"net/http"
)

func Handler(w http.ResponseWriter, req *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Hello Kitty"))
}
