package router

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func performRequest(r http.Handler, method, path string) *httptest.ResponseRecorder {
	req, _ := http.NewRequest(method, path, nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	return w
}

func TestSetupRouter(t *testing.T) {
	r := SetupRouter()

	t.Run("Test GET /mars", func(t *testing.T) {
		w := performRequest(r, "GET", "/mars")
		var response map[string]string
		err := json.Unmarshal(w.Body.Bytes(), &response)

		value, exists := response["system"]

		body := gin.H{
			"system": "Hello Mars, we're here!!",
		}

		assert.Nil(t, err)
		assert.True(t, exists)
		assert.Equal(t, http.StatusOK, w.Code)
		assert.Equal(t, body["system"], value)
	})

	t.Run("Test GET /mars/explore", func(t *testing.T) {
		w := performRequest(r, "GET", "/mars/explore")

		assert.Equal(t, http.StatusOK, w.Code)
	})

	w := performRequest(r, "GET", "/moon")
	assert.Equal(t, http.StatusNotFound, w.Code)
}
