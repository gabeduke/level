package api

//
//import (
//	"github.com/stretchr/testify/assert"
//	"net/http"
//	"net/http/httptest"
//	"testing"
//)
//
//func TestNwsAPIMock_GetLevel(t *testing.T) {
//	mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
//		w.WriteHeader(http.StatusOK)
//		w.Write([]byte(`<observed><datum><primary name="Stage" units="ft">10.5</primary></datum></observed>`))
//	}))
//	defer mockServer.Close()
//
//	api := &NwsAPI{baseurl: mockServer.URL}
//	level, err := api.GetLevel("test_station")
//	assert.NoError(t, err)
//	assert.Equal(t, 10.5, level)
//}
//
//func TestNwsAPIMock_GetStationList(t *testing.T) {
//	mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
//		w.WriteHeader(http.StatusOK)
//		w.Write([]byte(`[{"key":"test_key","points":[{"lid":"test_lid","latitude":"test_latitude","longitude":"test_longitude","gauge_type":"test_gauge_type","obs_status":"test_obs_status","name":"test_name","wfo":"test_wfo","inundation":"test_inundation","hsa_display":"test_hsa_display","state":"test_state","suppress_fcst":"test_suppress_fcst","icon":"test_icon"}]}]`))
//	}))
//	defer mockServer.Close()
//
//	api := &NwsAPI{baseurl: mockServer.URL}
//	stations, err := api.GetStationList()
//	assert.NoError(t, err)
//	assert.Len(t, stations, 1)
//	assert.Equal(t, "test_key", stations[0].Key)
//}
//
//func TestUsgsAPIMock_GetLevel(t *testing.T) {
//	mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
//		w.WriteHeader(http.StatusOK)
//		w.Write([]byte(`{"value":{"timeSeries":[{"values":[{"value":[{"value":"15.2"}]}]}]}}`))
//	}))
//	defer mockServer.Close()
//
//	api := &UsgsAPI{baseurl: mockServer.URL}
//	level, err := api.GetLevel("test_station")
//	assert.NoError(t, err)
//	assert.Equal(t, 15.2, level)
//}
//
//func TestUsgsAPIMock_GetStationList(t *testing.T) {
//	mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
//		w.WriteHeader(http.StatusOK)
//		w.Write([]byte(`[{"key":"test_key","points":[{"lid":"test_lid","latitude":"test_latitude","longitude":"test_longitude","gauge_type":"test_gauge_type","obs_status":"test_obs_status","name":"test_name","wfo":"test_wfo","inundation":"test_inundation","hsa_display":"test_hsa_display","state":"test_state","suppress_fcst":"test_suppress_fcst","icon":"test_icon"}]}]`))
//	}))
//	defer mockServer.Close()
//
//	api := &UsgsAPI{baseurl: mockServer.URL}
//	stations, err := api.GetStationList()
//	assert.NoError(t, err)
//	assert.Len(t, stations, 1)
//	assert.Equal(t, "test_key", stations[0].Key)
//}
//
//func TestEcwoAPIMock_GetLevel(t *testing.T) {
//	mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
//		w.WriteHeader(http.StatusOK)
//		w.Write([]byte(`{"stationData":[{"level":"20.3"}]}`))
//	}))
//	defer mockServer.Close()
//
//	api := &EcwoAPI{baseurl: mockServer.URL}
//	level, err := api.GetLevel("test_station")
//	assert.NoError(t, err)
//	assert.Equal(t, 20.3, level)
//}
//
//func TestEcwoAPIMock_GetStationList(t *testing.T) {
//	mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
//		w.WriteHeader(http.StatusOK)
//		w.Write([]byte(`[{"key":"test_key","points":[{"lid":"test_lid","latitude":"test_latitude","longitude":"test_longitude","gauge_type":"test_gauge_type","obs_status":"test_obs_status","name":"test_name","wfo":"test_wfo","inundation":"test_inundation","hsa_display":"test_hsa_display","state":"test_state","suppress_fcst":"test_suppress_fcst","icon":"test_icon"}]}]`))
//	}))
//	defer mockServer.Close()
//
//	api := &EcwoAPI{baseurl: mockServer.URL}
//	stations, err := api.GetStationList()
//	assert.NoError(t, err)
//	assert.Len(t, stations, 1)
//	assert.Equal(t, "test_key", stations[0].Key)
//}
