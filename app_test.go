package main

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHelloWorldHandler(t *testing.T) {

	content, err := ioutil.ReadFile("index.html")
	if err != nil {
		t.Fatal("Can not read file index.html")
	}

	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(HelloWorldHandler)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
	respContent := string(content)
	if rr.Body.String() != respContent {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), respContent)
	}
}