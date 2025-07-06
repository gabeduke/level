package api

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNwsAPI_GetLevel(t *testing.T) {
	api := &NwsAPI{}
	level, err := api.GetLevel("000SE")
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
