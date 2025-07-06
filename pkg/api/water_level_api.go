package api

// StationPoint holds details for a single point.
type StationPoint struct {
	Lid          string `json:"lid"`
	Latitude     string `json:"latitude"`
	Longitude    string `json:"longitude"`
	GaugeType    string `json:"gauge_type"`
	ObsStatus    string `json:"obs_status"`
	Name         string `json:"name"`
	Wfo          string `json:"wfo"`
	Inundation   string `json:"inundation"`
	HsaDisplay   string `json:"hsa_display"`
	State        string `json:"state"`
	SuppressFcst string `json:"suppress_fcst"`
	Icon         string `json:"icon"`
}

// Station is the unified station type returned by both APIs.
type Station struct {
	Key    string         `json:"key"`
	Points []StationPoint `json:"points"`
}

// WaterLevelAPI is the interface both our USGS and NWS API implementations must satisfy.
type WaterLevelAPI interface {
	GetLevel(station string) (float64, error)
	GetStationList() ([]Station, error)
}
