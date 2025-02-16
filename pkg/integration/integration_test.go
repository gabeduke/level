package integration

import (
	"encoding/json"
	"github.com/gabeduke/level/pkg/api"
	"github.com/gabeduke/level/pkg/router"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

const v1 = "/api/v1"

type reading struct {
	Reading float32 `json:"reading"`
	Message string  `json:"message"`
}

func TestIntegrationHealthzRoute(t *testing.T) {
	router := router.GetRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", v1+"/healthz", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, "{\"message\":\"healthy\"}", w.Body.String())
}

func TestIntegrationLevelRoute(t *testing.T) {
	router := router.GetRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", v1+"/level", nil)
	router.ServeHTTP(w, req)

	level := &reading{}
	err := json.Unmarshal(w.Body.Bytes(), level)
	if err != nil {
		t.Error(err)
	}

	assert.Equal(t, 200, w.Code)
	assert.NotEmpty(t, level.Reading)
}

func TestIntegrationLevelRouteWithStation(t *testing.T) {
	router := router.GetRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", v1+"/level", nil)

	q := req.URL.Query()
	q.Add("station", "RICV2")
	req.URL.RawQuery = q.Encode()

	router.ServeHTTP(w, req)

	level := &reading{}
	err := json.Unmarshal(w.Body.Bytes(), level)
	if err != nil {
		t.Error(err)
	}

	assert.Equal(t, 200, w.Code)
	assert.NotEmpty(t, level.Reading)
}

func TestIntegrationLevelRouteWithBadStation(t *testing.T) {
	router := router.GetRouter()

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
		t.Error(err)
	}

	assert.Equal(t, level.Message, "XML syntax error on line 95: invalid character entity &nbsp;")
	assert.Equal(t, w.Code, 424)
}

func TestIntegrationStationRoute(t *testing.T) {
	router := router.GetRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", v1+"/stations", nil)
	router.ServeHTTP(w, req)

	stations := &[]api.Station{}
	err := json.Unmarshal(w.Body.Bytes(), stations)
	if err != nil {
		t.Error(err)
	}

	assert.Equal(t, 200, w.Code)
	assert.NotEmpty(t, stations)
}

func TestIntegrationSlackRoute(t *testing.T) {
	router := router.GetRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", v1+"/slack", nil)
	router.ServeHTTP(w, req)

	slack := &api.Slack{}
	err := json.Unmarshal(w.Body.Bytes(), slack)
	if err != nil {
		t.Error(err)
	}

	assert.Equal(t, 200, w.Code)
	assert.NotEmpty(t, slack.Text)
}
