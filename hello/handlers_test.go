package main

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHelloHandler(t *testing.T) {
	req, err := http.NewRequest(http.MethodGet, "/", nil)

	if err != nil {
		t.Fatalf("Failed to create request %v", err)
	}

	body := bytes.NewBufferString("")
	recorder := httptest.NewRecorder()
	recorder.Body = body

	helloWorldHandler(recorder, req)

	res := recorder.Result()

	if assert.Equal(t, http.StatusOK, res.StatusCode) {
		assert.Equal(t, "Hello, World!\n", body.String())
		res.Body.Close()
	}
}
