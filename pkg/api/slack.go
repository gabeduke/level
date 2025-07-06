package api

type Slack struct {
	Text         string `json:"text"`
	ResponseType string `json:"response_type"`
	Parse        string `json:"parse"`
	UnfurlLinks  bool   `json:"unfurl_links"`
	UnfurlMedia  bool   `json:"unfurl_media"`
	Attachments  []struct {
		ImageURL string `json:"image_url"`
	} `json:"attachments"`
}
