package api

import (
	"net/http"
)

func Hello(w http.ResponseWriter, req *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Hello Kitty"))
}
