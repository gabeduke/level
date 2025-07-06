package router

import (
	"fmt"
	"github.com/gabeduke/level/pkg/api"
	"github.com/gabeduke/level/pkg/httputil"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

// @title Level API
// @version 1.0
// @description API for retrieving water level information.
// @BasePath /api/v1
func GetRouter() *gin.Engine {
	r := gin.Default()
	r.Use(cors.Default())
	r.Use(gin.Recovery())
	r.GET("/", RedirectRootToAPI(r))

	v1 := r.Group("/api/v1")
	{
		v1.GET("/level", level)
		v1.GET("/stations", stations)
		v1.GET("/healthz", healthz)
		v1.POST("/slack", slack)
	}

	return r
}

// RedirectRootToAPI sends requests from "/" to "/api/v1/level".
// @Summary Redirect Root
// @Description Redirect requests from root URL to /api/v1/level endpoint
// @Router / [get]
func RedirectRootToAPI(r *gin.Engine) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Request.URL.Path = "/api/v1/level"
		r.HandleContext(c)
	}
}

// healthz is a simple health-check endpoint.
// @Summary Health Check
// @Description Returns service health status.
// @Produce json
// @Success 200 {object} map[string]string "healthy"
// @Router /healthz [get]
func healthz(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "healthy",
	})
}

// slack returns a Slack-compatible payload using the NWS API.
// @Summary Slack Response
// @Description Returns a Slack payload with the water level and an image link.
// @Accept json
// @Produce json
// @Param station query string false "Station identifier" default(RMDV2)
// @Success 200 {object} api.Slack
// @Failure 424 {object} httputil.HTTPError
// @Router /slack [post]
func slack(c *gin.Context) {
	station := c.DefaultQuery("station", "KSEA")

	var i api.WaterLevelAPI = &api.NwsAPI{}
	lvl, err := i.GetLevel(station)
	if err != nil {
		httputil.NewError(c, http.StatusFailedDependency, err)
		return
	}

	slackResp := api.Slack{
		Text:         fmt.Sprintf("%f", lvl),
		ResponseType: "in_channel",
		Parse:        "full",
		UnfurlLinks:  true,
		UnfurlMedia:  true,
		Attachments: []struct {
			ImageURL string `json:"image_url"`
		}{
			{
				ImageURL: fmt.Sprintf("https://water.weather.gov/resources/hydrographs/%s_hg.png", strings.ToLower(station)),
			},
		},
	}

	c.JSON(http.StatusOK, &slackResp)
}

// stations returns the station list using the USGS API.
// @Summary Get Station List
// @Description Returns a list of stations.
// @Produce json
// @Success 200 {array} api.Station
// @Failure 424 {object} httputil.HTTPError
// @Router /stations [get]
func stations(c *gin.Context) {
	var i api.WaterLevelAPI = &api.UsgsAPI{}
	stations, err := i.GetStationList()
	if err != nil {
		httputil.NewError(c, http.StatusFailedDependency, err)
		return
	}

	c.JSON(http.StatusOK, stations)
}

// level returns the water level for a specified station using the USGS API.
// @Summary Get Water Level
// @Description Returns the water level for a given station.
// @Produce json
// @Param station query string false "Station identifier" default(01646500)
// @Success 200 {object} map[string]float64
// @Failure 424 {object} httputil.HTTPError
// @Router /level [get]
func level(c *gin.Context) {
	station := c.DefaultQuery("station", "01646500")

	var i api.WaterLevelAPI = &api.UsgsAPI{}
	lvl, err := i.GetLevel(station)
	if err != nil {
		httputil.NewError(c, http.StatusFailedDependency, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"reading": lvl,
	})
}
