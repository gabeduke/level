package api

import (
	"encoding/json"
	"fmt"
	"github.com/apex/log"
	"io"
	"net/http"
	"strconv"
)

type NwsAPI struct {
	baseurl string
}

// GetLevel returns the temperature value (or water level if available)
// for the specified station based on the latest observation.
func (n *NwsAPI) GetLevel(station string) (float64, error) {
	if n.baseurl == "" {
		n.baseurl = "https://api.weather.gov/stations"
		log.Infof("Using default baseurl: %s", n.baseurl)
	}

	url := fmt.Sprintf("%s/%s/observations/latest", n.baseurl, station)
	log.Debugf("GetLevel URL: %s", url)

	resp, err := http.Get(url)
	if err != nil {
		return 0, fmt.Errorf("error fetching data: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		log.Errorf("Non-200 response: %d, Body: %s", resp.StatusCode, string(body))
		return 0, fmt.Errorf("received non-200 response code: %d", resp.StatusCode)
	}

	var data struct {
		Properties struct {
			Temperature struct {
				Value float64 `json:"value"`
			} `json:"temperature"`
		} `json:"properties"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return 0, fmt.Errorf("error decoding JSON: %v", err)
	}

	return data.Properties.Temperature.Value, nil
}

// GetStationList retrieves a list of stations from the NWS API and converts them into the common Station type.
func (n *NwsAPI) GetStationList() ([]Station, error) {
	url := "https://api.weather.gov/stations"
	log.Debugf("GetStationList URL: %s", url)

	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("error fetching station list: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		log.Errorf("Non-200 response: %d, Body: %s", resp.StatusCode, string(body))
		return nil, fmt.Errorf("received non-200 response code: %d", resp.StatusCode)
	}

	var stationResponse struct {
		Features []struct {
			Properties struct {
				ID        string  `json:"stationIdentifier"`
				Name      string  `json:"name"`
				Latitude  float64 `json:"latitude"`
				Longitude float64 `json:"longitude"`
			} `json:"properties"`
		} `json:"features"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&stationResponse); err != nil {
		return nil, fmt.Errorf("error decoding station list JSON: %v", err)
	}

	var stations []Station
	for _, feature := range stationResponse.Features {
		props := feature.Properties
		s := Station{
			Key: props.ID,
			Points: []StationPoint{
				{
					Lid:       props.ID,
					Latitude:  strconv.FormatFloat(props.Latitude, 'f', 6, 64),
					Longitude: strconv.FormatFloat(props.Longitude, 'f', 6, 64),
					Name:      props.Name,
				},
			},
		}
		stations = append(stations, s)
	}

	return stations, nil
}
