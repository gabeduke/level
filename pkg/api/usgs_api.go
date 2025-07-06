package api

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"github.com/apex/log"
	"io"
	"net/http"
	"strconv"
	"strings"
)

type UsgsAPI struct {
	baseurl string
	// Default state code if not provided
	stateCd string
}

// GetLevel returns the water level (gage reading) for the given station.
func (u *UsgsAPI) GetLevel(station string) (float64, error) {
	if u.baseurl == "" {
		u.baseurl = "https://waterservices.usgs.gov/nwis/iv/"
		log.Infof("Using default baseurl: %s", u.baseurl)
	}
	url := fmt.Sprintf("%s?sites=%s&parameterCd=00065&format=json", u.baseurl, station)
	log.Debugf("GetLevel URL: %s", url)

	resp, err := http.Get(url)
	if err != nil {
		return 0, fmt.Errorf("error fetching data: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return 0, fmt.Errorf("received non-200 response code: %d", resp.StatusCode)
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

	if err := json.NewDecoder(resp.Body).Decode(&usgsData); err != nil {
		return 0, fmt.Errorf("error decoding JSON response: %v", err)
	}

	if len(usgsData.Value.TimeSeries) == 0 ||
		len(usgsData.Value.TimeSeries[0].Values) == 0 ||
		len(usgsData.Value.TimeSeries[0].Values[0].Value) == 0 {
		return 0, fmt.Errorf("no data found for station: %s", station)
	}

	reading := usgsData.Value.TimeSeries[0].Values[0].Value[0].Value
	log.Debugf("Gauge Reading: %s", reading)

	f, err := strconv.ParseFloat(reading, 64)
	if err != nil {
		return 0, fmt.Errorf("error parsing gauge reading: %v", err)
	}

	return f, nil
}

// GetStationList retrieves a list of stations from USGS and converts them into the common Station type.
func (u *UsgsAPI) GetStationList() ([]Station, error) {
	// Use default state code if not set.
	stateCd := u.stateCd
	if stateCd == "" {
		stateCd = "va"
	}

	url := fmt.Sprintf("https://waterservices.usgs.gov/nwis/site/?format=rdb&stateCd=%s&siteType=ST&siteStatus=active", stateCd)
	log.Debugf("GetStationList URL: %s", url)

	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("error fetching station list: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		log.Errorf("Received non-200 response code: %d, Body: %s", resp.StatusCode, string(body))
		return nil, fmt.Errorf("received non-200 response code: %d", resp.StatusCode)
	}

	reader := csv.NewReader(resp.Body)
	reader.Comma = '\t'
	reader.Comment = '#'

	records, err := reader.ReadAll()
	if err != nil {
		return nil, fmt.Errorf("error reading RDB data: %v", err)
	}
	if len(records) < 2 {
		return nil, fmt.Errorf("no station data found")
	}

	headers := records[0]
	var stations []Station
	for _, record := range records[1:] {
		if len(record) != len(headers) {
			log.Warnf("record length mismatch: %v", record)
			continue
		}

		agencyCode := record[getIndex(headers, "agency_cd")]
		siteNumber := record[getIndex(headers, "site_no")]
		stationName := record[getIndex(headers, "station_nm")]
		siteType := record[getIndex(headers, "site_tp_cd")]
		decLat := record[getIndex(headers, "dec_lat_va")]
		decLong := record[getIndex(headers, "dec_long_va")]

		s := Station{
			Key: agencyCode + "-" + siteNumber,
			Points: []StationPoint{
				{
					Lid:       siteNumber,
					Latitude:  decLat,
					Longitude: decLong,
					GaugeType: siteType,
					Name:      stationName,
				},
			},
		}
		stations = append(stations, s)
	}

	return stations, nil
}

func getIndex(headers []string, header string) int {
	for i, h := range headers {
		if strings.TrimSpace(h) == header {
			return i
		}
	}
	return -1
}
