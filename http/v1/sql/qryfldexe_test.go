package sql

import (
	"net/url"
	"testing"

	core "github.com/amsokol/go-ignite-client/http"
	"github.com/amsokol/go-ignite-client/http/v1/exec"
)

func TestCommands_SQLFieldsQueryExecute(t *testing.T) {
	insertValues := url.Values{}
	insertValues.Add("arg1", "1")
	insertValues.Add("arg2", "Org 1")

	type args struct {
		e        exec.Executer
		cache    string
		pageSize int64
		query    string
		args     url.Values
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
			name: "Delete all from Person table",
			c:    &Commands{},
			args: args{
				e:        &exec.ExecuterImpl{Servers: []string{"http://localhost:8080/ignite"}, Username: "", Password: ""},
				cache:    "Person",
				pageSize: 1000,
				query:    `DELETE from "Person".Person`,
				args:     url.Values{},
			},
		},
		{
			name: "Delete all from Organization table",
			c:    &Commands{},
			args: args{
				e:        &exec.ExecuterImpl{Servers: []string{"http://localhost:8080/ignite"}, Username: "", Password: ""},
				cache:    "Person",
				pageSize: 1000,
				query:    `DELETE from "Organization".Organization`,
				args:     url.Values{},
			},
		},
		{
			name: "Insert one record into Organization",
			c:    &Commands{},
			args: args{
				e:        &exec.ExecuterImpl{Servers: []string{"http://localhost:8080/ignite"}, Username: "", Password: ""},
				cache:    "Organization",
				pageSize: 1000,
				query:    `INSERT INTO Organization(_key, name) VALUES(?, ?)`,
				args:     insertValues,
			},
		},
		{
			name: "Select all organizations",
			c:    &Commands{},
			args: args{
				e:        &exec.ExecuterImpl{Servers: []string{"http://localhost:8080/ignite"}, Username: "", Password: ""},
				cache:    "Organization",
				pageSize: 1000,
				query:    `SELECT _key, name FROM Organization`,
				args:     nil,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotResult, gotToken, err := tt.c.SQLFieldsQueryExecute(tt.args.e, tt.args.cache, tt.args.pageSize, tt.args.query, tt.args.args)
			if (err != nil) != tt.wantErr {
				t.Errorf("Commands.SQLFieldsQueryExecute() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			t.Log("")
			t.Logf("http.v1.sql.SQLFieldsQueryExecute returned for '%s':", tt.name)
			t.Log("result =", gotResult)
			t.Log("sessionToken =", gotToken)
		})
	}
}
