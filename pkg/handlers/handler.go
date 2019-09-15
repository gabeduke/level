package handlers

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/apex/log"
	"github.com/beevik/etree"
	"github.com/gin-gonic/gin"
)

// Healthz is a service healthcheck
func Healthz(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "healthy",
	})
}

// Level gets the water level for a given station
// queryparams: ?station=[station]
func Level(c *gin.Context) {

	station := c.DefaultQuery("station", "rmdv2")
	url := fmt.Sprintf("http://water.weather.gov/ahps2/hydrograph_to_xml.php?gage=%s&output=xml", station)

	xmlBytes, err := getXML(url)
	if err != nil {
		log.Errorf("Failed to get XML: %v", err)
	}

	doc := etree.NewDocument()
	err = doc.ReadFromBytes(xmlBytes)
	if err != nil {
		log.Error(err.Error())
	}

	root := doc.FindElement("/site/observed/datum")

	reading := root.FindElement("primary").Text()
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
