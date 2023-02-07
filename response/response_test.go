package response_test

import (
	"errors"
	"go-template/response"
	"net/http/httptest"
	"testing"
)

func TestRespondSuccess(t *testing.T) {
	w := httptest.NewRecorder()

	// create the handler to test, using our custom "next" handler
	response.RespondSuccess(w, 200, nil, nil)
}

func TestRespondSuccessError(t *testing.T) {
	w := httptest.NewRecorder()

	value := make(chan interface{})

	// create the handler to test, using our custom "next" handler
	response.RespondSuccess(w, 200, value, nil)
}

func TestRespondError(t *testing.T) {
	w := httptest.NewRecorder()

	// create the handler to test, using our custom "next" handler
	response.RespondError(w, 400, errors.New("error"))
}

func TestRespondStatMsg(t *testing.T) {
	w := httptest.NewRecorder()

	// create the handler to test, using our custom "next" handler
	response.RespondStatMsg(w, 400, "error")
}

func TestSendFileSuccess(t *testing.T) {
	w := httptest.NewRecorder()

	// create the handler to test, using our custom "next" handler
	response.SendFileSuccess(w, "test.csv", "text/csv", true)
}

func TestSendFileNotFound(t *testing.T) {
	w := httptest.NewRecorder()

	// create the handler to test, using our custom "next" handler
	response.SendFileSuccess(w, "test1.csv", "text/csv", true)
}

func TestSendFileErrorStat(t *testing.T) {
	w := httptest.NewRecorder()

	// create the handler to test, using our custom "next" handler
	response.SendFileSuccess(w, "response.go", "text/csv", true)
}
