package nws

import (
	"encoding/xml"
	"github.com/apex/log"
	"net/http"
)

const baseUrl = "http://water.weather.gov/ahps2/hydrograph_to_xml.php"

type nwsActions interface {
	GetStationInfo(url string, data *NWS) error
}

//nolint
type nwsConfig struct {
	baseurl string
}

type nwsStationInfo struct {
	nwsActions
	nwsConfig
}

// NwsStationAPI is the consumable API for accessing NWS station data
type NwsStationAPI struct {
	nwsStationInfo
}

func (i *nwsStationInfo) getStationConfig() {
	i.baseurl = baseUrl
	log.Infof("Get default baseurl: %s", i.baseurl)
}

func getStationInfo(url string, data *NWS) error {
	log.Debug(url)

	resp, err := http.Get(url)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	err = xml.Unmarshal(body, data)
	if err != nil {
		return err
	}

	return nil
}
