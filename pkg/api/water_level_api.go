package api

type Station struct {
	Key    string `json:"key"`
	Points []struct {
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
	} `json:"points"`
}

type WaterLevelAPI interface {
	GetLevel(station string) (float64, error)
	GetStationList() ([]Station, error)
}
