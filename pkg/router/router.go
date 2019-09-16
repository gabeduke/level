package router

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/apex/log"
	"github.com/beevik/etree"
	"github.com/gin-gonic/gin"
)

// GetRouter returns a level router
func GetRouter() *gin.Engine {

	r := gin.Default()
	r.GET("/healthz", healthz)
	r.GET("/level", level)

	return r
}

// healthz is a service healthcheck
func healthz(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "healthy",
	})
}

// level gets the water level for a given station
// queryparams: ?station=[station]
func level(c *gin.Context) {

	station := c.DefaultQuery("station", "RMDV2")
	url := fmt.Sprintf("http://water.weather.gov/ahps2/hydrograph_to_xml.php?gage=%s&output=xml", station)

	xmlBytes, err := getXML(url)
	if err != nil {
		log.Errorf("Failed to get XML: %v", err)
		c.AbortWithError(417, err)
	}

	doc := etree.NewDocument()
	err = doc.ReadFromBytes(xmlBytes)
	if err != nil {
		log.Error(err.Error())
		c.AbortWithError(417, err)
	}

	reading := doc.FindElement("//*/observed/datum[1]/primary").Text()
	if reading == "" {
		err = fmt.Errorf("unable to find root element for url: %s", url)
		c.AbortWithError(417, err)
	}
	log.Debugf("Gauge Reading: %s", reading)

	c.JSON(200, gin.H{
		"message": reading,
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
