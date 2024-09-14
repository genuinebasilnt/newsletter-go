package test

import (
	"fmt"
	"genuinebasilnt/newsletter-go/internal/config"
	"genuinebasilnt/newsletter-go/internal/router"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestSubscribeHandler(t *testing.T) {
	config, err := config.GetConfiguration("../config")
	if err != nil {
		t.Fatal(err)
	}

	config.DatabaseSettings.DatabaseName = uuid.New().String()
	configuration := configureDatabase(t, config)

	t.Run("subscribe returns 200 for valid form data", func(t *testing.T) {
		r := router.Router(configuration)

		body := "name=genuinebasilnt&email=genuine.basilnt@gmail.com"
		req, _ := http.NewRequest("POST", "/subscriptions", strings.NewReader(body))
		req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)

		assert.Equal(t, http.StatusOK, w.Code)
	})

	t.Run("subscribe returns 400 when data is missing", func(t *testing.T) {
		r := router.Router(configuration)

		testCases := []struct {
			invalidBody  string
			errorMessage string
		}{
			{"name=genuinebasilnt", "missing email"},
			{"email=genuine.basilnt@gmail.com", "missing name"},
			{"", "missing name and email"},
		}

		for _, testCase := range testCases {
			formData, _ := url.ParseQuery(testCase.invalidBody)

			req, _ := http.NewRequest("POST", "/subscriptions", strings.NewReader(formData.Encode()))
			req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

			w := httptest.NewRecorder()
			r.ServeHTTP(w, req)

			assert.Equal(t, http.StatusBadRequest, w.Code, fmt.Sprintf("API did not fail with 400 when payload was %s", testCase.errorMessage))
		}

	})
}
