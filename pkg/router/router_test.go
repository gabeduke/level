package router

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/apex/log"
	"github.com/stretchr/testify/assert"
)

type message struct {
	Message string `json:"message"`
}

func TestHealthzRoute(t *testing.T) {
	router := GetRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/healthz", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, "{\"message\":\"healthy\"}", w.Body.String())
}

func TestLevelRoute(t *testing.T) {
	router := GetRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/level", nil)
	router.ServeHTTP(w, req)

	level := &message{}
	err := json.Unmarshal(w.Body.Bytes(), level)
	if err != nil {
		log.Error(err.Error())
	}

	assert.Equal(t, 200, w.Code)
	assert.NotEmpty(t, level.Message)
}

func TestLevelRouteWithStation(t *testing.T) {
	router := GetRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/level", nil)
	req.URL.Query().Set("station", "RICV2")
	router.ServeHTTP(w, req)

	level := &message{}
	err := json.Unmarshal(w.Body.Bytes(), level)
	if err != nil {
		log.Error(err.Error())
	}

	assert.Equal(t, 200, w.Code)
	assert.NotEmpty(t, level.Message)
}
