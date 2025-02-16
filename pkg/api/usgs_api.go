package api

import (
	"encoding/json"
	"fmt"
	"github.com/apex/log"
	"io/ioutil"
	"net/http"
	"strconv"
)

type UsgsAPI struct {
	baseurl string
}

func (u *UsgsAPI) GetLevel(station string) (float64, error) {
	if u.baseurl == "" {
		u.baseurl = "https://waterservices.usgs.gov/nwis/iv/"
		log.Infof("Get default baseurl: %s", u.baseurl)
	}

	url := fmt.Sprintf("%s?sites=%s&parameterCd=00065&format=json", u.baseurl, station)
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

	var usgsData struct {
		Value struct {
			TimeSeries []struct {
				Values []struct {
					Value []struct {
						Value string `json:"value"`
					} `json:"value"`
				} `json:"values"`
			} `json:"timeSeries"`
		} `json:"value"`
	}

	err = json.Unmarshal(data, &usgsData)
	if err != nil {
		return 0, err
	}

	if len(usgsData.Value.TimeSeries) == 0 || len(usgsData.Value.TimeSeries[0].Values) == 0 || len(usgsData.Value.TimeSeries[0].Values[0].Value) == 0 {
		return 0, fmt.Errorf("no data found for station: %s", station)
	}

	reading := usgsData.Value.TimeSeries[0].Values[0].Value[0].Value
	log.Debugf("Gauge Reading: %s", reading)

	f, err := strconv.ParseFloat(reading, 64)
	if err != nil {
		log.Error(err.Error())
	}

	return f, nil
}

func (u *UsgsAPI) GetStationList() ([]Station, error) {
	url := "https://waterservices.usgs.gov/nwis/site/?format=rdb&stateCd=all&siteType=ST&siteStatus=active"
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
