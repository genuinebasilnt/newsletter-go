package test

import (
	"genuinebasilnt/newsletter-go/internal/config"
	"genuinebasilnt/newsletter-go/internal/router"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHealthCheckHandler(t *testing.T) {
	config, err := config.GetConfiguration("../config")
	if err != nil {
		t.Fatal(err)
	}

	env := configureDatabase(t, config)
	r := router.Router(env)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/health_check", nil)
	r.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
}
