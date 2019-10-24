package router

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"

	"github.com/apex/log"
	"github.com/beevik/etree"
	"github.com/gabeduke/level/pkg/httputil"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type stationsList struct {
	Key    string `json:"key"`
	Points []struct {
		Lid          string `json:"lid"`
		Latitude     string `json:"latitude"`
		Longitude    string `json:"longitude"`
		GaugeType    string `json:"gauge_type"`
		ObsStatus    string `json:"obs_status"`
		Name         string `json:"name"`
		Wfo          string `json:"wfo"`
		Inundation   string `json:"inundation"`
		HsaDisplay   string `json:"hsa_display"`
		State        string `json:"state"`
		SuppressFcst string `json:"suppress_fcst"`
		Icon         string `json:"icon"`
	} `json:"points"`
}

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
// @Failure 417 {object} httputil.HTTPError
// @Router /stations [get]
func stations(c *gin.Context) {

	body := strings.NewReader(`key=akq&fcst_type=obs&percent=50&current_type=all`)
	req, err := http.NewRequest("POST", "https://water.weather.gov/ahps/get_map_points.php", body)
	if err != nil {
		log.Error(err.Error())
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded; charset=UTF-8")
	req.Header.Set("Accept", "*/*")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		httputil.NewError(c, http.StatusExpectationFailed, err)
		return
	}
	defer resp.Body.Close()

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Error(err.Error())
	}

	var t stationsList
	err = json.Unmarshal([]byte(data), &t)
	if err != nil {
		httputil.NewError(c, http.StatusUnprocessableEntity, err)
		return
	}

	c.JSON(200, &t)
}

// level gets the water level for a given station
// @Summary return water level
// @Description get level by station
// @ID level
// @Accept  json
// @Produce  json
// @Param station path string false "NWS Station to query"
// @Success 200
// @Failure 417 {object} httputil.HTTPError
// @Router /level [get]
func level(c *gin.Context) {

	station := c.DefaultQuery("station", "RMDV2")
	url := fmt.Sprintf("http://water.weather.gov/ahps2/hydrograph_to_xml.php?gage=%s&output=xml", station)

	xmlBytes, err := getXML(url)
	if err != nil {
		log.Errorf("Failed to get XML: %v", err)
		httputil.NewError(c, http.StatusExpectationFailed, err)
		return
	}

	doc := etree.NewDocument()
	err = doc.ReadFromBytes(xmlBytes)
	if err != nil {
		log.Error(err.Error())
		httputil.NewError(c, http.StatusExpectationFailed, err)
		return
	}

	reading := doc.FindElement("//*/observed/datum[1]/primary").Text()
	if reading == "" {
		err = fmt.Errorf("unable to find root element for url: %s", url)
		httputil.NewError(c, http.StatusExpectationFailed, err)
		return
	}
	log.Debugf("Gauge Reading: %s", reading)
	f, err := strconv.ParseFloat(reading, 64)
	if err != nil {
		httputil.NewError(c, http.StatusExpectationFailed, err)
	}

	c.JSON(200, gin.H{
		"reading": f,
	})
}

// tweaked from: https://stackoverflow.com/a/42718113/1170664
func getXML(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		return []byte{}, fmt.Errorf("GET error: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return []byte{}, fmt.Errorf("status error: %v", resp.StatusCode)
	}

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return []byte{}, fmt.Errorf("read body: %v", err)
	}

	return data, nil
}
