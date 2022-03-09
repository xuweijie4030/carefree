package test

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestDemoController_SubmitSaga(t *testing.T) {
	router := setupRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/home", nil)
	router.ServeHTTP(w, req)

	fmt.Println(w.Code)
	fmt.Println(w.Body.String())
}
