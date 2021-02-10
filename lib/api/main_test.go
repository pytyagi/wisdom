package api_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo"
	. "github.com/pytyagi/wisdom/lib/api"
	. "github.com/pytyagi/wisdom/lib/wisdom"
)

var FixtureQuotes = NewDispenser([]Quote{
	{Quote: "first", Author: "amy"},
	{Quote: "second", Author: "bob"},
	{Quote: "third", Author: "carol"},
})

func get(path string) *http.Request {
	return httptest.NewRequest(echo.GET, path, nil)
}

func serve(t *testing.T, request *http.Request) *httptest.ResponseRecorder {
	return serveWith(t, request, FixtureQuotes, DefaultConfig())
}

func serveWith(t *testing.T, request *http.Request, dispenser Dispenser, cfg *Config) *httptest.ResponseRecorder {
	rr := httptest.NewRecorder()
	server := NewServer(cfg, dispenser)
	server.ServeHTTP(rr, request)
	return rr
}
