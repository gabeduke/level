package router

import (
	"encoding/json"
	"github.com/gabeduke/level/pkg/nws"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/apex/log"
	"github.com/stretchr/testify/assert"
)

const v1 = "/api/v1"

type reading struct {
	Reading float32 `json:"reading"`
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

	level := &reading{}
	err := json.Unmarshal(w.Body.Bytes(), level)
	if err != nil {
		log.Error(err.Error())
	}

	assert.Equal(t, 200, w.Code)
	assert.NotEmpty(t, level.Reading)
}

func TestLevelRoute(t *testing.T) {
	router := GetRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", v1+"/level", nil)
	router.ServeHTTP(w, req)

	level := &reading{}
	err := json.Unmarshal(w.Body.Bytes(), level)
	if err != nil {
		log.Error(err.Error())
	}

	assert.Equal(t, 200, w.Code)
	assert.NotEmpty(t, level.Reading)
}

func TestLevelRouteWithStation(t *testing.T) {
	router := GetRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", v1+"/level", nil)

	q := req.URL.Query()
	q.Add("station", "RICV2")
	req.URL.RawQuery = q.Encode()

	router.ServeHTTP(w, req)

	level := &reading{}
	err := json.Unmarshal(w.Body.Bytes(), level)
	if err != nil {
		log.Error(err.Error())
	}

	assert.Equal(t, 200, w.Code)
	assert.NotEmpty(t, level.Reading)
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

	level := &reading{}
	err := json.Unmarshal(w.Body.Bytes(), level)
	if err != nil {
		log.Error(err.Error())
	}

	assert.Equal(t, level.Message, "XML syntax error on line 95: invalid character entity &nbsp;")
	assert.Equal(t, w.Code, 424)
}

func TestStationRoute(t *testing.T) {
	router := GetRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", v1+"/stations", nil)
	router.ServeHTTP(w, req)

	stations := &nws.StationsList{}
	err := json.Unmarshal(w.Body.Bytes(), stations)
	if err != nil {
		log.Error(err.Error())
	}

	assert.Equal(t, 200, w.Code)
	assert.NotEmpty(t, stations.Points)
}
