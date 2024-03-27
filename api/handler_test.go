package handler

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHandler(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/hello", nil)
	w := httptest.NewRecorder()

	Handler(w, req)

	res := w.Result()
	defer res.Body.Close()

	resp, _ := io.ReadAll(res.Body)

	if string(resp) != "Hello Kitty" {
		t.Errorf("Want Hello Kitty got %s", resp)
	}
}
