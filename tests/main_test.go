package tests

import (
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"olahcitra/main"
	"testing"
)

func router() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/", main.HomePage).Methods("GET")
	return router
}

func TestHomePage(t *testing.T) {
	req, _ := http.NewRequest("GET", "/", nil)
	res := httptest.NewRecorder()
	router().ServeHTTP(res, req)
	assert.Equal(t, 200, res.Code, "Expect OK Response")
}
