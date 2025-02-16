package api

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNwsAPI_GetLevel(t *testing.T) {
	api := &NwsAPI{}
	level, err := api.GetLevel("RMDV2")
	assert.NoError(t, err)
	assert.NotZero(t, level)
}

func TestNwsAPI_GetStationList(t *testing.T) {
	api := &NwsAPI{}
	stations, err := api.GetStationList()
	assert.NoError(t, err)
	assert.NotEmpty(t, stations)
}

func TestUsgsAPI_GetLevel(t *testing.T) {
	api := &UsgsAPI{}
	level, err := api.GetLevel("01646500")
	assert.NoError(t, err)
	assert.NotZero(t, level)
}

func TestUsgsAPI_GetStationList(t *testing.T) {
	api := &UsgsAPI{}
	stations, err := api.GetStationList()
	assert.NoError(t, err)
	assert.NotEmpty(t, stations)
}

func TestEcwoAPI_GetLevel(t *testing.T) {
	api := &EcwoAPI{}
	level, err := api.GetLevel("02HC001")
	assert.NoError(t, err)
	assert.NotZero(t, level)
}

func TestEcwoAPI_GetStationList(t *testing.T) {
	api := &EcwoAPI{}
	stations, err := api.GetStationList()
	assert.NoError(t, err)
	assert.NotEmpty(t, stations)
}
