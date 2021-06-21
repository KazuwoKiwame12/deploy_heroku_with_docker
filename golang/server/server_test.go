package server_test

import (
	"app/server"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHandler(t *testing.T) {
	router := server.NewRouter()

	req := httptest.NewRequest("GET", "/", nil)
	rec := httptest.NewRecorder()

	router.ServeHTTP(rec, req)
	if rec.Code != http.StatusOK {
		t.Errorf("error: %+v\n", rec)
	}
	if rec.Body.String() != server.ExportResponseMsg {
		t.Errorf("expected: %+v, result: %+v\n", server.ExportResponseMsg, rec.Body.String())
	}
}

func TestServe(t *testing.T) {
	router := server.NewRouter()
	testServer := httptest.NewServer(router)
	defer testServer.Close()

	req, _ := http.NewRequest("GET", testServer.URL+"/", nil)
	client := new(http.Client)
	resp, _ := client.Do(req)
	respBody, _ := ioutil.ReadAll(resp.Body)

	if resp.StatusCode != http.StatusOK {
		t.Errorf("error: %+v\n", resp)
	}
	if string(respBody) != server.ExportResponseMsg {
		t.Errorf("expected: %+v, result: %+v\n", server.ExportResponseMsg, string(respBody))
	}
}
