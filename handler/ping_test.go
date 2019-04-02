package handler

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPingHandler(t *testing.T) {
	setup()

	w := httptest.NewRecorder()
	r, _ := http.NewRequest("GET", "/ping", nil)

	PingHandler(w, r)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "application/json", w.Header().Get("Content-Type"))
	assert.Equal(t, "{\"success\": \"pong\"}", w.Body.String())
}
