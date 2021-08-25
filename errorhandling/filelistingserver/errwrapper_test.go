package main

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func errPanic(writer http.ResponseWriter, request *http.Request) error {
	panic(123)
}

func errUserErr(writer http.ResponseWriter, request *http.Request) error {
	return testingUserError("user error")
}

type testingUserError string

func (err testingUserError) Error() string {
	return err.Message()
}

func (err testingUserError) Message() string {
	return string(err)
}

var test = []struct {
	h       appHandler
	code    int
	message string
}{
	{errPanic, 500, "Internal Server Error"},
	{errUserErr, 400, "user error"},
}

func TestErrWrapper(t *testing.T) {

	for _, tt := range test {
		f := errWrapper(tt.h)
		response := httptest.NewRecorder()
		request := httptest.NewRequest(http.MethodGet, "https://www.jaluik.com", nil)
		f(response, request)
		verifyResponse(t, response.Result(), tt.code, tt.message)

	}
}

func TestErrWrapperInServer(t *testing.T) {

	for _, tt := range test {
		f := errWrapper(tt.h)
		server := httptest.NewServer(http.HandlerFunc(f))
		response, _ := http.Get(server.URL)
		verifyResponse(t, response, tt.code, tt.message)

	}

}

func verifyResponse(t *testing.T, response *http.Response, code int, message string) {
	b, _ := ioutil.ReadAll(response.Body)
	body := strings.Trim(string(b), "\n")
	if response.StatusCode != code || body != message {
		t.Errorf("expect (%d, %s)  but got (%d, %s)", code, message, response.StatusCode, body)
	}
}
