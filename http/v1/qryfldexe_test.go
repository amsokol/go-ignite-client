package v1

import (
	"net/url"
	"testing"

	core "github.com/amsokol/go-ignite-client/http"
)

func Test_client_SQLFieldsQueryExecute(t *testing.T) {
	insertValues := url.Values{}
	insertValues.Add("arg1", "1")
	insertValues.Add("arg2", "Org 1")

	type args struct {
		cache    string
		pageSize int64
		query    string
		args     url.Values
	}
	tests := []struct {
		name       string
		c          Client
		args       args
		wantResult core.SQLQueryResult
		wantToken  string
		wantErr    bool
	}{
		{
			name: "Delete all from Person table",
			c:    NewClient([]string{"http://localhost:8080/ignite"}, "", ""),
			args: args{
				cache:    "Person",
				pageSize: 1000,
				query:    `DELETE from "Person".Person`,
				args:     url.Values{},
			},
		},
		{
			name: "Delete all from Organization table",
			c:    NewClient([]string{"http://localhost:8080/ignite"}, "", ""),
			args: args{
				cache:    "Person",
				pageSize: 1000,
				query:    `DELETE from "Organization".Organization`,
				args:     url.Values{},
			},
		},
		{
			name: "Insert one record into Organization",
			c:    NewClient([]string{"http://localhost:8080/ignite"}, "", ""),
			args: args{
				cache:    "Organization",
				pageSize: 1000,
				query:    `INSERT INTO Organization(_key, name) VALUES(?, ?)`,
				args:     insertValues,
			},
		},
		{
			name: "Select all organizations",
			c:    NewClient([]string{"http://localhost:8080/ignite"}, "", ""),
			args: args{
				cache:    "Organization",
				pageSize: 1000,
				query:    `SELECT _key, name FROM Organization`,
				args:     nil,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotResult, gotToken, err := tt.c.SQLFieldsQueryExecute(tt.args.cache, tt.args.pageSize, tt.args.query, tt.args.args)
			if (err != nil) != tt.wantErr {
				t.Errorf("client.SQLFieldsQueryExecute() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			t.Log("")
			t.Logf("http.v1.client.SQLFieldsQueryExecute returned for '%s':", tt.name)
			t.Log("result =", gotResult)
			t.Log("sessionToken =", gotToken)
		})
	}
}
