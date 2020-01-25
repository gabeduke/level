package nws

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"github.com/apex/log"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
)

const baseUrl = "http://water.weather.gov/ahps2/hydrograph_to_xml.php"

type nwsActions interface {
	GetLevel(url string) (float64, error)
	GetStationList(list *StationsList) error
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

// GetLevel returns the level for a given station
func (i *nwsStationInfo) GetLevel(station string) (float64, error) {

	if i.baseurl == "" {
		i.getStationConfig()
	}

	url := fmt.Sprintf("%s?gage=%s&output=xml", i.baseurl, station)
	log.Debugf("GetLevel URL: %s", url)

	nwsData := NWS{}

	err := getStationInfo(url, &nwsData)
	if err != nil {
		return 0, err
	}

	reading := nwsData.Observed.Datum[0].Primary.Text
	if reading == "" {
		err = fmt.Errorf("unable to find observed datum element for url: %s", url)
		return 0, err
	}
	log.Debugf("Gauge Reading: %s", reading)

	f, err := strconv.ParseFloat(reading, 64)
	if err != nil {
		log.Error(err.Error())
	}

	return f, nil
}

func (i *nwsStationInfo) GetStationList(list *StationsList) error {

	body := strings.NewReader(`key=akq&fcst_type=obs&percent=50&current_type=all`)
	req, err := http.NewRequest("POST", "https://water.weather.gov/ahps/get_map_points.php", body)
	if err != nil {
		log.Error(err.Error())
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded; charset=UTF-8")
	req.Header.Set("Accept", "*/*")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	err = json.Unmarshal([]byte(data), &list)
	if err != nil {
		return err
	}

	return nil
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
