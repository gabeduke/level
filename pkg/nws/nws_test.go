package nws

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_getStationInfo(t *testing.T) {
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
		{
			name:    "good url",
			args:    args{baseUrl, &NWS{}},
			wantErr: false,
		},
		{
			name:    "good url bad xml",
			args:    args{"https://google.com", &NWS{}},
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

func Test_nwsStationInfo_GetLevel(t *testing.T) {
	type fields struct {
		nwsActions nwsActions
		nwsConfig  nwsConfig
	}
	type args struct {
		station string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    float64
		wantErr bool
	}{
		//{
		//	name: "happy path",
		//	fields: fields{nwsConfig:nwsConfig{baseurl:},},
		//	args: args{"RMDV2"},
		//	want: 0,
		//	wantErr: false,
		//},
		{
			name:    "bad url",
			fields:  fields{nwsConfig: nwsConfig{"asdf"}},
			args:    args{},
			want:    0,
			wantErr: true,
		},
		{
			name:    "good url bad xml",
			fields:  fields{nwsConfig: nwsConfig{"https://google.com"}},
			args:    args{},
			want:    0,
			wantErr: true,
		},
		{
			name:    "err no arg",
			fields:  fields{},
			args:    args{},
			want:    0,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &nwsStationInfo{
				nwsActions: tt.fields.nwsActions,
				nwsConfig:  tt.fields.nwsConfig,
			}
			got, err := i.GetLevel(tt.args.station)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetLevel() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("GetLevel() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_nwsStationInfo_GetStationList(t *testing.T) {
	type fields struct {
		nwsActions nwsActions
		nwsConfig  nwsConfig
	}
	type args struct {
		list *StationsList
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name:    "happy path",
			fields:  fields{},
			args:    args{},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &nwsStationInfo{
				nwsActions: tt.fields.nwsActions,
				nwsConfig:  tt.fields.nwsConfig,
			}
			if err := i.GetStationList(tt.args.list); (err != nil) != tt.wantErr {
				t.Errorf("GetStationList() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_nwsStationInfo_getStationConfig(t *testing.T) {
	type fields struct {
		nwsActions nwsActions
		nwsConfig  nwsConfig
	}
	tests := []struct {
		name   string
		fields fields
	}{
		{
			name:   "populated baseUrl should equal",
			fields: fields{nwsConfig: nwsConfig{baseurl: baseUrl}},
		},
		{
			name:   "unpopulated baseUrl should equal",
			fields: fields{nwsConfig: nwsConfig{}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &nwsStationInfo{
				nwsActions: tt.fields.nwsActions,
				nwsConfig:  tt.fields.nwsConfig,
			}
			i.getStationConfig()
			assert.Equal(t, i.baseurl, baseUrl)
		})
	}
}
