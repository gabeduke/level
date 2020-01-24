package nws

import "testing"

func TestGetNwsData_ScrapeNws(t *testing.T) {
	type args struct {
		url  string
		data *NWS
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name:    "bad url",
			args:    args{"", &NWS{}},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := getStationInfo(tt.args.url, tt.args.data); (err != nil) != tt.wantErr {
				t.Errorf("ScrapeNws() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
