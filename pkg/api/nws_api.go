package api

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

type NwsAPI struct {
	baseurl string
}

func (n *NwsAPI) GetLevel(station string) (float64, error) {
	if n.baseurl == "" {
		n.baseurl = "http://water.weather.gov/ahps2/hydrograph_to_xml.php"
		log.Infof("Get default baseurl: %s", n.baseurl)
	}

	url := fmt.Sprintf("%s?gage=%s&output=xml", n.baseurl, station)
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

func (n *NwsAPI) GetStationList() ([]Station, error) {
	body := strings.NewReader(`key=akq&fcst_type=obs&percent=50&current_type=all`)
	req, err := http.NewRequest("POST", "https://water.weather.gov/ahps/get_map_points.php", body)
	if err != nil {
		log.Error(err.Error())
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded; charset=UTF-8")
	req.Header.Set("Accept", "*/*")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var list []Station
	err = json.Unmarshal([]byte(data), &list)
	if err != nil {
		return nil, err
	}

	return list, nil
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

type NWS struct {
	Observed struct {
		Datum []struct {
			Primary struct {
				Text  string `xml:",chardata"`
				Name  string `xml:"name,attr"`
				Units string `xml:"units,attr"`
			} `xml:"primary"`
		} `xml:"datum"`
	} `xml:"observed"`
}
