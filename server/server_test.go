package server_test

import (
	"encoding/json"
	"go-template/config"
	"go-template/server"
	"io/ioutil"
	"net/http/httptest"
	"testing"

	"github.com/ghodss/yaml"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
)

func TestBoot(t *testing.T) {
	yamlFile := "../config.local.yaml"
	buf, _ := ioutil.ReadFile(yamlFile)

	jsonStruct, _ := yaml.YAMLToJSON(buf)

	var conf config.Config
	json.Unmarshal(jsonStruct, &conf)

	route := mux.NewRouter()

	server.InitApp(route, conf, true)
}

func TestServerNew(t *testing.T) {
	res := server.New()
	assert.Equal(t, &server.RestServer{}, res)

}

func TestNotFoundHandler(t *testing.T) {
	// create the handler to test, using our custom "next" handler
	handlerToTest := server.NotFoundHandler()

	// create a mock request to use
	req := httptest.NewRequest("GET", "http://testing", nil)

	// call the handler using a mock response recorder (we'll not use that anyway)
	handlerToTest.ServeHTTP(httptest.NewRecorder(), req)
}

func TestMethodNotAllowedHandler(t *testing.T) {
	// create the handler to test, using our custom "next" handler
	handlerToTest := server.MethodNotAllowedHandler()

	// create a mock request to use
	req := httptest.NewRequest("GET", "http://testing", nil)

	// call the handler using a mock response recorder (we'll not use that anyway)
	handlerToTest.ServeHTTP(httptest.NewRecorder(), req)
}
