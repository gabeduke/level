package router

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/apex/log"
	"github.com/beevik/etree"
	"github.com/gabeduke/level/pkg/httputil"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// GetRouter returns a level router
func GetRouter() *gin.Engine {

	r := gin.Default()
	r.Use(cors.Default())
	v1 := r.Group("/api/v1")

	v1.GET("/level", level)
	v1.GET("/healthz", healthz)

	return r
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
		return []byte{}, fmt.Errorf("Status error: %v", resp.StatusCode)
	}

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return []byte{}, fmt.Errorf("Read body: %v", err)
	}

	return data, nil
}
