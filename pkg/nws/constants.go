package nws

// StationsList is the structure for a list of NWS stations
type StationsList struct {
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

// NWS is the detailed station info
type NWS struct {
	//Name           string `xml:"name,attr"`
	//ID             string `xml:"id,attr"`
	//Generationtime string `xml:"generationtime,attr"`
	//Sigstages      struct {
	//	Text string `xml:",chardata"`
	//	Low  struct {
	//		Text  string `xml:",chardata"`
	//		Units string `xml:"units,attr"`
	//	} `xml:"low"`
	//	Action struct {
	//		Text  string `xml:",chardata"`
	//		Units string `xml:"units,attr"`
	//	} `xml:"action"`
	//	Bankfull struct {
	//		Text  string `xml:",chardata"`
	//		Units string `xml:"units,attr"`
	//	} `xml:"bankfull"`
	//	Flood struct {
	//		Text  string `xml:",chardata"`
	//		Units string `xml:"units,attr"`
	//	} `xml:"flood"`
	//	Moderate struct {
	//		Text  string `xml:",chardata"`
	//		Units string `xml:"units,attr"`
	//	} `xml:"moderate"`
	//	Major struct {
	//		Text  string `xml:",chardata"`
	//		Units string `xml:"units,attr"`
	//	} `xml:"major"`
	//	Record struct {
	//		Text  string `xml:",chardata"`
	//		Units string `xml:"units,attr"`
	//	} `xml:"record"`
	//} `xml:"sigstages"`
	//Sigflows struct {
	//	Text string `xml:",chardata"`
	//	Low  struct {
	//		Text  string `xml:",chardata"`
	//		Units string `xml:"units,attr"`
	//	} `xml:"low"`
	//	Action struct {
	//		Text  string `xml:",chardata"`
	//		Units string `xml:"units,attr"`
	//	} `xml:"action"`
	//	Bankfull struct {
	//		Text  string `xml:",chardata"`
	//		Units string `xml:"units,attr"`
	//	} `xml:"bankfull"`
	//	Flood struct {
	//		Text  string `xml:",chardata"`
	//		Units string `xml:"units,attr"`
	//	} `xml:"flood"`
	//	Moderate struct {
	//		Text  string `xml:",chardata"`
	//		Units string `xml:"units,attr"`
	//	} `xml:"moderate"`
	//	Major struct {
	//		Text  string `xml:",chardata"`
	//		Units string `xml:"units,attr"`
	//	} `xml:"major"`
	//	Record struct {
	//		Text  string `xml:",chardata"`
	//		Units string `xml:"units,attr"`
	//	} `xml:"record"`
	//} `xml:"sigflows"`
	//Zerodatum struct {
	//	Text  string `xml:",chardata"`
	//	Units string `xml:"units,attr"`
	//} `xml:"zerodatum"`
	//Rating struct {
	//	Text    string `xml:",chardata"`
	//	Dignity string `xml:"dignity,attr"`
	//	Datum   []struct {
	//		Text       string `xml:",chardata"`
	//		Stage      string `xml:"stage,attr"`
	//		StageUnits string `xml:"stageUnits,attr"`
	//		Flow       string `xml:"flow,attr"`
	//		FlowUnits  string `xml:"flowUnits,attr"`
	//	} `xml:"datum"`
	//} `xml:"rating"`
	//AltRating struct {
	//	Text    string `xml:",chardata"`
	//	Dignity string `xml:"dignity,attr"`
	//	Datum   []struct {
	//		Text       string `xml:",chardata"`
	//		Stage      string `xml:"stage,attr"`
	//		StageUnits string `xml:"stageUnits,attr"`
	//		Flow       string `xml:"flow,attr"`
	//		FlowUnits  string `xml:"flowUnits,attr"`
	//	} `xml:"datum"`
	//} `xml:"alt_rating"`
	Observed struct {
		//Text  string `xml:",chardata"`
		Datum []struct {
			//Text  string `xml:",chardata"`
			//Valid struct {
			//	Text     string `xml:",chardata"`
			//	Timezone string `xml:"timezone,attr"`
			//} `xml:"valid"`
			Primary struct {
				Text  string `xml:",chardata"`
				Name  string `xml:"name,attr"`
				Units string `xml:"units,attr"`
			} `xml:"primary"`
			//Secondary struct {
			//	Text  string `xml:",chardata"`
			//	Name  string `xml:"name,attr"`
			//	Units string `xml:"units,attr"`
			//} `xml:"secondary"`
			//Pedts string `xml:"pedts"`
		} `xml:"datum"`
	} `xml:"observed"`
	//Forecast struct {
	//	Text     string `xml:",chardata"`
	//	Timezone string `xml:"timezone,attr"`
	//	Issued   string `xml:"issued,attr"`
	//	Datum    []struct {
	//		Text  string `xml:",chardata"`
	//		Valid struct {
	//			Text     string `xml:",chardata"`
	//			Timezone string `xml:"timezone,attr"`
	//		} `xml:"valid"`
	//		Primary struct {
	//			Text  string `xml:",chardata"`
	//			Name  string `xml:"name,attr"`
	//			Units string `xml:"units,attr"`
	//		} `xml:"primary"`
	//		Secondary struct {
	//			Text  string `xml:",chardata"`
	//			Name  string `xml:"name,attr"`
	//			Units string `xml:"units,attr"`
	//		} `xml:"secondary"`
	//		Pedts string `xml:"pedts"`
	//	} `xml:"datum"`
	//} `xml:"forecast"`
}

type Slack struct {
	Parse        string `json:"parse"`
	ResponseType string `json:"response_type"`
	Text         string `json:"text"`
	Attachments  []struct {
		ImageURL string `json:"image_url"`
	} `json:"attachments"`
	UnfurlMedia bool `json:"unfurl_media"`
	UnfurlLinks bool `json:"unfurl_links"`
}
