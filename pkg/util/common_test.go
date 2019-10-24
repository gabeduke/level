package util

import (
	"os"
	"testing"
)

func setPortEnv(t *testing.T, port string) {
	os.Setenv("PORT", port)
}

func TestGetPort(t *testing.T) {
	tests := []struct {
		name    string
		portGen string
		want    string
	}{
		{name: "defaultPort", portGen: "", want: "8081"},
		{name: "overridePort", portGen: "8080", want: "8080"},
	}
	for _, tt := range tests {
		setPortEnv(t, tt.portGen)
		t.Run(tt.name, func(t *testing.T) {
			if got := GetPort(); got != tt.want {
				t.Errorf("GetPort() = %v, want %v", got, tt.want)
			}
		})
	}
}
