package v1

import (
	"testing"
)

func Test_client_GetVersion(t *testing.T) {
	tests := []struct {
		name        string
		c           Client
		wantVersion string
		wantToken   string
		wantErr     bool
	}{
		{
			name: "Get version of server",
			c:    NewClient([]string{"http://localhost:8080/ignite"}, "", ""),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotVersion, gotToken, err := tt.c.GetVersion()
			if (err != nil) != tt.wantErr {
				t.Errorf("client.GetVersion() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			t.Log("")
			t.Logf("http.v1.client.GetVersion returned for '%s':", tt.name)
			t.Log("version =", gotVersion)
			t.Log("sessionToken =", gotToken)
		})
	}
}
