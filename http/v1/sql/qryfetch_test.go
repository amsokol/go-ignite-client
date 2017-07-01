package sql

import (
	"testing"

	core "github.com/amsokol/go-ignite-client/http"
	"github.com/amsokol/go-ignite-client/http/v1/exec"
)

func TestCommands_SQLQueryFetch(t *testing.T) {
	t.Log("")
	t.Log("Preparing test data for 'TestCommands_SQLQueryFetch'...")

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
		e        exec.Executer
		pageSize int64
		queryID  int64
	}
	tests := []struct {
		name       string
		c          *Commands
		args       args
		wantResult core.SQLQueryResult
		wantToken  string
		wantErr    bool
	}{
		{
			name: "Close select from Organization",
			c:    c,
			args: args{
				e:        e,
				pageSize: 1000,
				queryID:  result.QueryID,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotResult, gotToken, err := tt.c.SQLQueryFetch(tt.args.e, tt.args.pageSize, tt.args.queryID)
			if (err != nil) != tt.wantErr {
				t.Errorf("Commands.SQLQueryFetch() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			t.Log("")
			t.Logf("http.v1.sql.SQLQueryFetch returned for '%s':", tt.name)
			t.Log("result =", gotResult)
			t.Log("sessionToken =", gotToken)
		})
	}
}
