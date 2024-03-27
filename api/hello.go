package api

import (
	"net/http"

	p "template-go-vercel/api/_pkg"
)

func Hello(w http.ResponseWriter, req *http.Request) {
	w.WriteHeader(http.StatusOK)

	model := *p.NewModel()

	resp := "Hello "+ model.Name
	w.Write([]byte(resp))
}
