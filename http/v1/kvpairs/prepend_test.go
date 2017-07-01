package kvpairs

import (
	"testing"

	"github.com/amsokol/go-ignite-client/http/v1/cache"
	"github.com/amsokol/go-ignite-client/http/v1/exec"
)

func TestKeyValuePairsImpl_Prepend(t *testing.T) {
	t.Log("")
	t.Log("Preparing test data for 'TestKeyValuePairsImpl_Prepend'...")

	e := exec.ExecuterImpl{Servers: []string{"http://localhost:8080/ignite"}, Username: "", Password: ""}
	c := cache.ManagementImpl{}

	_, err := c.DestroyCache(&e, "TestKeyValuePairs")
	if err != nil {
		t.Fatal(err)
	}

	_, err = c.GetOrCreateCache(&e, "TestKeyValuePairs")
	if err != nil {
		t.Fatal(err)
	}

	t.Log("Done")

	type args struct {
		e      exec.Executer
		cache  string
		key    string
		val    string
		destID string
	}
	tests := []struct {
		name       string
		p          *KeyValuePairsImpl
		args       args
		wantOk     bool
		wantNodeID string
		wantToken  string
		wantErr    bool
	}{
		{
			name: "Prepend",
			p:    &KeyValuePairsImpl{},
			args: args{
				e:     &e,
				cache: "TestKeyValuePairs",
				key:   "Prepend",
				val:   "-1",
			},
			wantOk: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotOk, _, _, err := tt.p.Prepend(tt.args.e, tt.args.cache, tt.args.key, tt.args.val, tt.args.destID)
			if (err != nil) != tt.wantErr {
				t.Errorf("KeyValuePairsImpl.Prepend() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotOk != tt.wantOk {
				t.Errorf("KeyValuePairsImpl.Prepend() gotOk = %v, want %v", gotOk, tt.wantOk)
			}
		})
	}
}
