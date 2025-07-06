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

// reading is the expected payload from the /level endpoint.
type reading struct {
	Reading float64 `json:"reading"`
	Message string  `json:"message"`
}

func TestIntegrationHealthzRoute(t *testing.T) {
	r := router.GetRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", v1+"/healthz", nil)
	r.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, "{\"message\":\"healthy\"}", w.Body.String())
}

func TestIntegrationLevelRoute(t *testing.T) {
	r := router.GetRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", v1+"/level", nil)
	r.ServeHTTP(w, req)

	var lvl reading
	err := json.Unmarshal(w.Body.Bytes(), &lvl)
	if err != nil {
		t.Error(err)
	}

	assert.Equal(t, 200, w.Code)
	assert.NotZero(t, lvl.Reading)
}

func TestIntegrationLevelRouteWithStation(t *testing.T) {
	r := router.GetRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", v1+"/level", nil)
	// Use a valid USGS station ID; for example, "01646500"
	q := req.URL.Query()
	q.Add("station", "01646500")
	req.URL.RawQuery = q.Encode()

	r.ServeHTTP(w, req)

	var lvl reading
	err := json.Unmarshal(w.Body.Bytes(), &lvl)
	if err != nil {
		t.Error(err)
	}

	assert.Equal(t, 200, w.Code)
	assert.NotZero(t, lvl.Reading)
}

func TestIntegrationLevelRouteWithBadStation(t *testing.T) {
	r := router.GetRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", v1+"/level", nil)
	q := req.URL.Query()
	q.Add("station", "asdf")
	req.URL.RawQuery = q.Encode()

	r.ServeHTTP(w, req)

	var lvl reading
	err := json.Unmarshal(w.Body.Bytes(), &lvl)
	if err != nil {
		t.Error(err)
	}

	// Expect the USGS API to return an error with status code 424 and the message below.
	assert.Equal(t, "received non-200 response code: 400", lvl.Message)
	assert.Equal(t, 424, w.Code)
}

func TestIntegrationStationRoute(t *testing.T) {
	r := router.GetRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", v1+"/stations", nil)
	r.ServeHTTP(w, req)

	var stations []api.Station
	err := json.Unmarshal(w.Body.Bytes(), &stations)
	if err != nil {
		t.Error(err)
	}

	assert.Equal(t, 200, w.Code)
	assert.NotEmpty(t, stations)
}

func TestIntegrationSlackRoute(t *testing.T) {
	r := router.GetRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", v1+"/slack", nil)
	r.ServeHTTP(w, req)

	var slackResp api.Slack
	err := json.Unmarshal(w.Body.Bytes(), &slackResp)
	if err != nil {
		t.Error(err)
	}

	assert.Equal(t, 200, w.Code)
	assert.NotEmpty(t, slackResp.Text)
}
