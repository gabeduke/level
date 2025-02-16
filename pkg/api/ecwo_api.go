package api

import (
	"encoding/json"
	"fmt"
	"github.com/apex/log"
	"io/ioutil"
	"net/http"
	"strconv"
)

type EcwoAPI struct {
	baseurl string
}

func (e *EcwoAPI) GetLevel(station string) (float64, error) {
	if e.baseurl == "" {
		e.baseurl = "https://wateroffice.ec.gc.ca/services/real_time_data/csv/inline"
		log.Infof("Get default baseurl: %s", e.baseurl)
	}

	url := fmt.Sprintf("%s?station=%s", e.baseurl, station)
	log.Debugf("GetLevel URL: %s", url)

	resp, err := http.Get(url)
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return 0, err
	}

	var ecwoData struct {
		StationData []struct {
			Level string `json:"level"`
		} `json:"stationData"`
	}

	err = json.Unmarshal(data, &ecwoData)
	if err != nil {
		return 0, err
	}

	if len(ecwoData.StationData) == 0 {
		return 0, fmt.Errorf("no data found for station: %s", station)
	}

	reading := ecwoData.StationData[0].Level
	log.Debugf("Gauge Reading: %s", reading)

	f, err := strconv.ParseFloat(reading, 64)
	if err != nil {
		log.Error(err.Error())
	}

	return f, nil
}

func (e *EcwoAPI) GetStationList() ([]Station, error) {
	url := "https://wateroffice.ec.gc.ca/services/real_time_data/stations/csv/inline"
	log.Debugf("GetStationList URL: %s", url)

	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var stations []Station
	err = json.Unmarshal(data, &stations)
	if err != nil {
		return nil, err
	}

	return stations, nil
}
