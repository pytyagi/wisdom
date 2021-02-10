package api_test

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"sync"
	"testing"
	"time"

	. "github.com/pytyagi/wisdom/lib/api"
	// . "github.com/pytyagi/wisdom/lib/wisdom"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func JSONResponseOK(t *testing.T, rr *httptest.ResponseRecorder) map[string]interface{} {
	return jsonResponse(t, rr, http.StatusOK)
}

func jsonResponse(t *testing.T, rr *httptest.ResponseRecorder, code int) map[string]interface{} {
	assert.Equal(t, code, rr.Code)
	assert.Equal(t, "application/json; charset=utf-8", strings.ToLower(rr.HeaderMap["Content-Type"][0]))
	var response map[string]interface{}
	require.Nil(t, json.NewDecoder(rr.Body).Decode(&response))
	return response
}

func TestPing(t *testing.T) {
	req := get("/ping")
	rr := serve(t, req)
	response := JSONResponseOK(t, rr)
	assert.Equal(t, map[string]interface{}{"ping": "pong"}, response)

}

func TestVersion(t *testing.T) {
	req := get("/version")
	rr := serve(t, req)
	response := JSONResponseOK(t, rr)
	assert.Equal(t, map[string]interface{}{"Version": "latest"}, response)

}

// func TestQuote(t *testing.T) {
// 	quotes := NewDispenser([]Quote{
// 		{Quote: "predictable", Author: "sam"},
// 	})

// 	req := get("/quote")
// 	rr := serveWith(t, req, quotes, DefaultConfig())
// 	response := JSONResponseOK(t, rr)
// 	assert.Equal(t, map[string]interface{}{"quote": "predictable", "author": "sam"}, response)

// }

func TestWithCustomAPIPath(t *testing.T) {

	req := get("/custom/path/ping")
	rr := serveWith(t, req, FixtureQuotes, &Config{APIPath: "/custom/path"})
	response := JSONResponseOK(t, rr)
	assert.Equal(t, map[string]interface{}{"ping": "pong"}, response)

	req = get("/custom/path/version")
	rr = serveWith(t, req, FixtureQuotes, &Config{APIPath: "/custom/path"})
	response = JSONResponseOK(t, rr)
	assert.Equal(t, map[string]interface{}{"Version": "latest"}, response)

}

func TestStartAndStop(t *testing.T) {

	cfg := &Config{Host: "localhost", Port: 9999}
	server := NewServer(cfg, FixtureQuotes)

	wg := &sync.WaitGroup{}
	wg.Add(1)
	go func() {
		defer wg.Done()
		err := server.Start()
		require.NotNil(t, err)
		assert.Equal(t, "http: Server closed", err.Error())
	}()

	time.Sleep(50 * time.Millisecond) // give server a chance to start

	err := server.Stop(context.Background())
	require.Nil(t, err)
	wg.Wait()

}

func TestCors(t *testing.T) {
	req := get("/ping")
	rr := serve(t, req)
	assert.Equal(t, []string{"*"}, rr.HeaderMap["Access-Control-Allow-Origin"])
}
