package webserver

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHelloRequest(t *testing.T) {
	req, err := http.NewRequest("GET", "/hello", nil)
	if err != nil {
		t.Fatal(err)
	}

	recorder := httptest.NewRecorder()
	handler := http.HandlerFunc(helloRequest)

	handler.ServeHTTP(recorder, req)

	if status := recorder.Code; status != http.StatusOK {
		t.Errorf("The Handler return code is wrong: return %v, have %v",
			status, http.StatusOK)
	}

	expected := "Hello World"
	if body := recorder.Body.String(); body != expected {
		t.Errorf("The Body of response is wrong: return %v, have this %v",
			body, expected)
	}
}

func TestGetRequest(t *testing.T) {
	req, err := http.NewRequest("GET", "/somefile.txt", nil)
	if err != nil {
		t.Fatal(err)
	}

	recorder := httptest.NewRecorder()
	handler := http.HandlerFunc(getRequest)

	handler.ServeHTTP(recorder, req)

	if status := recorder.Code; status != http.StatusNotFound {
		t.Errorf("The Handler return code is wrong: return %v, have this %v",
			status, http.StatusOK)
	}
}

func TestGetRabbit(t *testing.T) {
	req, err := http.NewRequest("GET", "/rabbit", nil)
	if err != nil {
		t.Fatal(err)
	}

	recorder := httptest.NewRecorder()
	handler := http.HandlerFunc(getRabbit)

	handler.ServeHTTP(recorder, req)

	if status := recorder.Code; status != http.StatusOK {
		t.Errorf("The Handler return code is wrong: return %v, have this %v",
			status, http.StatusOK)
	}

	expected := "Rabbit Evil"
	if body := recorder.Body.String(); body != expected {
		t.Errorf("The Handler return code is wrong: return %v, have this %v",
			body, expected)
	}
}
