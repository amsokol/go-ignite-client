package sql

import (
	"testing"

	"github.com/amsokol/go-ignite-client/http/v1/exec"
)

func TestCommands_SQLQueryClose(t *testing.T) {
	t.Log("")
	t.Log("Preparing test data for 'Test_Commands_SQLQueryClose'...")

	e := &exec.ExecuterImpl{Servers: []string{"http://localhost:8080/ignite"}, Username: "", Password: ""}
	c := &Commands{}

	_, _, err := c.SQLFieldsQueryExecute(e, "Person", 1, `DELETE from "Person".Person`, nil)
	if err != nil {
		t.Fatal(err)
	}
	_, _, err = c.SQLFieldsQueryExecute(e, "Person", 1, `DELETE from "Organization".Organization`, nil)
	if err != nil {
		t.Fatal(err)
	}

	_, _, err = c.SQLFieldsQueryExecute(e, "Person", 1, `INSERT INTO "Organization".Organization(_key, name) VALUES(1, 'Org1')`, nil)
	if err != nil {
		t.Fatal(err)
	}
	_, _, err = c.SQLFieldsQueryExecute(e, "Person", 1, `INSERT INTO "Organization".Organization(_key, name) VALUES(2, 'Org2')`, nil)
	if err != nil {
		t.Fatal(err)
	}
	result, _, err := c.SQLFieldsQueryExecute(e, "Person", 1, `SELECT _key, name FROM "Organization".Organization ORDER BY _key`, nil)
	if err != nil {
		t.Fatal(err)
	}

	t.Log("Done")

	type args struct {
		e       exec.Executer
		queryID int64
	}
	tests := []struct {
		name      string
		c         *Commands
		args      args
		wantOk    bool
		wantToken string
		wantErr   bool
	}{
		{
			name: "Close select from Organization",
			c:    c,
			args: args{
				e:       e,
				queryID: result.QueryID,
			},
			wantOk: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotOk, _, err := tt.c.SQLQueryClose(tt.args.e, tt.args.queryID)
			if (err != nil) != tt.wantErr {
				t.Errorf("Commands.SQLQueryClose() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotOk != tt.wantOk {
				t.Errorf("Commands.SQLQueryClose() gotOk = %v, want %v", gotOk, tt.wantOk)
			}
		})
	}
}
