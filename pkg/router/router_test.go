package router

import (
	"encoding/json"
	"github.com/gabeduke/level/pkg/api"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/apex/log"
	"github.com/stretchr/testify/assert"
)

const v1 = "/api/v1"

type reading struct {
	Reading float64 `json:"reading"`
	Message string  `json:"message"`
}

func TestHealthzRoute(t *testing.T) {
	router := GetRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", v1+"/healthz", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, "{\"message\":\"healthy\"}", w.Body.String())
}

func TestDefaultRoute(t *testing.T) {
	router := GetRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/", nil)
	router.ServeHTTP(w, req)

	var lvl reading
	err := json.Unmarshal(w.Body.Bytes(), &lvl)
	if err != nil {
		log.Error(err.Error())
	}

	assert.Equal(t, 200, w.Code)
	assert.NotZero(t, lvl.Reading)
}

func TestLevelRoute(t *testing.T) {
	router := GetRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", v1+"/level", nil)
	router.ServeHTTP(w, req)

	var lvl reading
	err := json.Unmarshal(w.Body.Bytes(), &lvl)
	if err != nil {
		log.Error(err.Error())
	}

	assert.Equal(t, 200, w.Code)
	assert.NotZero(t, lvl.Reading)
}

func TestLevelRouteWithStation(t *testing.T) {
	router := GetRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", v1+"/level", nil)

	q := req.URL.Query()
	req.URL.RawQuery = q.Encode()

	router.ServeHTTP(w, req)

	var lvl reading
	err := json.Unmarshal(w.Body.Bytes(), &lvl)
	if err != nil {
		log.Error(err.Error())
	}

	assert.Equal(t, 200, w.Code)
	assert.NotZero(t, lvl.Reading)
}

func TestLevelRouteWithBadStation(t *testing.T) {
	router := GetRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", v1+"/level", nil)

	q := req.URL.Query()
	q.Add("station", "asdf")
	req.URL.RawQuery = q.Encode()
	t.Log(req.URL)

	router.ServeHTTP(w, req)
	t.Log(w.Body.String())

	var lvl reading
	err := json.Unmarshal(w.Body.Bytes(), &lvl)
	if err != nil {
		log.Error(err.Error())
	}

	// Expect an error message if the station is invalid.
	// (Adjust the expected message based on your actual error handling.)
	assert.NotEmpty(t, lvl.Message)
	assert.Equal(t, 424, w.Code)
}

func TestStationRoute(t *testing.T) {
	router := GetRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", v1+"/stations", nil)
	router.ServeHTTP(w, req)

	var stations []api.Station
	err := json.Unmarshal(w.Body.Bytes(), &stations)
	if err != nil {
		log.Error(err.Error())
	}

	assert.Equal(t, 200, w.Code)
	assert.NotEmpty(t, stations)
}
