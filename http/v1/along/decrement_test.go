package along

import (
	"testing"

	"github.com/amsokol/go-ignite-client/http/v1/cache"
	"github.com/amsokol/go-ignite-client/http/v1/exec"
)

func TestCommands_Decrement(t *testing.T) {
	t.Log("")
	t.Log("Preparing test data for 'TestCommands_Decrement'...")

	e := exec.ExecuterImpl{Servers: []string{"http://localhost:8080/ignite"}, Username: "", Password: ""}
	c := cache.Commands{}

	_, err := c.DestroyCache(&e, "TestAtomicLongs")
	if err != nil {
		t.Fatal(err)
	}

	_, err = c.GetOrCreateCache(&e, "TestAtomicLongs")
	if err != nil {
		t.Fatal(err)
	}

	t.Log("Done")

	init := int64(0)

	type args struct {
		e     exec.Executer
		cache string
		key   string
		init  *int64
		delta int64
	}
	tests := []struct {
		name       string
		l          *Commands
		args       args
		wantValue  int64
		wantNodeID string
		wantToken  string
		wantErr    bool
	}{
		{
			name: "Decrement atomic long",
			l:    &Commands{},
			args: args{
				e:     &e,
				cache: "TestAtomicLongs",
				key:   "atomicLong1",
				init:  &init,
				delta: 1,
			},
			wantValue: -1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotValue, _, _, err := tt.l.Decrement(tt.args.e, tt.args.cache, tt.args.key, tt.args.init, tt.args.delta)
			if (err != nil) != tt.wantErr {
				t.Errorf("Commands.Decrement() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotValue != tt.wantValue {
				t.Errorf("Commands.Decrement() gotValue = %v, want %v", gotValue, tt.wantValue)
			}
		})
	}
}
