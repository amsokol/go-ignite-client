package cache

import (
	"testing"

	"github.com/amsokol/go-ignite-client/http/v1/exec"
)

func TestManagementImpl_DestroyCache(t *testing.T) {
	t.Log("")
	t.Log("Preparing test data for 'TestManagementImpl_DestroyCache'...")

	e := exec.ExecuterImpl{Servers: []string{"http://localhost:8080/ignite"}, Username: "", Password: ""}
	c := ManagementImpl{}

	_, err := c.GetOrCreateCache(&e, "Cache4TestDestroyCache")
	if err != nil {
		t.Fatal(err)
	}

	t.Log("Done")

	type args struct {
		e     exec.Executer
		cache string
	}
	tests := []struct {
		name      string
		c         *ManagementImpl
		args      args
		wantToken string
		wantErr   bool
	}{
		{
			name: "Destroy Cache4TestDestroyCache cache",
			c:    &c,
			args: args{
				e:     &e,
				cache: "Cache4TestDestroyCache",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := tt.c.DestroyCache(tt.args.e, tt.args.cache)
			if (err != nil) != tt.wantErr {
				t.Errorf("ManagementImpl.DestroyCache() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}
