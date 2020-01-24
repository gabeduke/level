package router

import (
	"github.com/gabeduke/level/pkg/httputil"
	"github.com/gabeduke/level/pkg/nws"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"net/http"
)

// GetRouter returns a level router
func GetRouter() *gin.Engine {

	r := gin.Default()
	r.Use(cors.Default())
	r.Use(gin.Recovery())
	r.GET("/", RedirectRootToAPI(r))

	v1 := r.Group("/api/v1")

	v1.GET("/level", level)
	v1.GET("/stations", stations)
	v1.GET("/healthz", healthz)

	return r
}

// RedirectRootToAPI redirects all calls from root endpoint to current API documentation endpoint
func RedirectRootToAPI(r *gin.Engine) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Request.URL.Path = "/api/v1/level"
		r.HandleContext(c)
	}
}

// healthz is a service healthcheck
// @Summary return healthcheck
// @Description get health
// @ID healthz
// @Accept  json
// @Produce  json
// @Success 200
// @Router /healthz [get]
func healthz(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "healthy",
	})
}

// stations gets the list of stations for a region
// @Summary returns list of stations
// @Description get stations
// @ID stations
// @Accept  json
// @Produce  json
// @Success 200
// @Failure 424 {object} httputil.HTTPError
// @Router /stations [get]
func stations(c *gin.Context) {

	stations := nws.StationsList{}

	i := nws.NwsStationAPI{}
	err := i.GetStationList(&stations)
	if err != nil {
		httputil.NewError(c, http.StatusFailedDependency, err)
		return
	}

	c.JSON(200, &stations)
}

// level gets the water level for a given station
// @Summary return water level
// @Description get level by station
// @ID level
// @Accept  json
// @Produce  json
// @Param station path string false "NWS Station to query"
// @Success 200
// @Failure 424 {object} httputil.HTTPError
// @Router /level [get]
func level(c *gin.Context) {

	station := c.DefaultQuery("station", "RMDV2")

	i := nws.NwsStationAPI{}
	lvl, err := i.GetLevel(station)
	if err != nil {
		httputil.NewError(c, http.StatusFailedDependency, err)
		return
	}

	c.JSON(200, gin.H{
		"reading": lvl,
	})
}
