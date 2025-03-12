package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-faker/faker/v4"
	"github.com/mjthecoder65/url-shortener/db"
	"github.com/mjthecoder65/url-shortener/utils"
	"github.com/stretchr/testify/require"
)

type testServer struct {
	router http.Handler
}

func newTestServer() *testServer {
	return &testServer{
		router: router,
	}
}

func (ts *testServer) Request(method, url string, body []byte) *httptest.ResponseRecorder {
	req, err := http.NewRequest(method, url, bytes.NewBuffer(body))
	if err != nil {
		panic(fmt.Sprintf("Failed to create request: %v", err))
	}
	res := httptest.NewRecorder()
	ts.router.ServeHTTP(res, req)
	return res
}

func createRandomShortURL(t *testing.T, ts *testServer) db.URL {
	payload := CreateShortURLRequest{
		URL: faker.URL(),
	}
	data, err := json.Marshal(payload)
	require.NoError(t, err)

	res := ts.Request("POST", "/api/v1/shorten", data)
	require.Equal(t, http.StatusCreated, res.Code)

	var shortURL db.URL
	err = json.NewDecoder(res.Body).Decode(&shortURL)
	require.NoError(t, err)
	require.Equal(t, payload.URL, shortURL.URL)
	return shortURL
}

func TestShortURLAPIs(t *testing.T) {
	ts := newTestServer()

	t.Run("Create and Get Short URL", func(t *testing.T) {
		shortURL := createRandomShortURL(t, ts)

		url := fmt.Sprintf("/api/v1/shorten/%s", shortURL.ShortCode)
		res := ts.Request("GET", url, nil)
		require.Equal(t, http.StatusOK, res.Code)

		var retrievedURL db.URL
		err := json.NewDecoder(res.Body).Decode(&retrievedURL)
		require.NoError(t, err)
		require.Equal(t, shortURL.URL, retrievedURL.URL)
		require.Equal(t, shortURL.ShortCode, retrievedURL.ShortCode)
	})

	t.Run("Update Short URL", func(t *testing.T) {
		shortURL := createRandomShortURL(t, ts)

		newURL := faker.URL()
		payload := UpdateShortURLRequest{URL: newURL}
		data, err := json.Marshal(payload)
		require.NoError(t, err)

		url := fmt.Sprintf("/api/v1/shorten/%s", shortURL.ShortCode)
		res := ts.Request("PUT", url, data)
		require.Equal(t, http.StatusOK, res.Code)

		var updatedURL db.URL
		err = json.NewDecoder(res.Body).Decode(&updatedURL)
		require.NoError(t, err)
		require.Equal(t, newURL, updatedURL.URL)
		require.Equal(t, shortURL.ShortCode, updatedURL.ShortCode)

		// Test updating non-existent URL
		nonExistentCode := utils.GenerateShortCode(testConfig)
		url = fmt.Sprintf("/api/v1/shorten/%s", nonExistentCode)
		res = ts.Request("PUT", url, data)
		require.Equal(t, http.StatusNotFound, res.Code)
	})

	t.Run("Delete Short URL", func(t *testing.T) {
		shortURL := createRandomShortURL(t, ts)

		url := fmt.Sprintf("/api/v1/shorten/%s", shortURL.ShortCode)
		res := ts.Request("DELETE", url, nil)
		require.Equal(t, http.StatusNoContent, res.Code)

		// Test deleting non-existent URL
		nonExistentCode := utils.GenerateShortCode(testConfig)
		url = fmt.Sprintf("/api/v1/shorten/%s", nonExistentCode)
		res = ts.Request("DELETE", url, nil)
		require.Equal(t, http.StatusNotFound, res.Code)
	})

	t.Run("Get URL Stats", func(t *testing.T) {
		shortURL := createRandomShortURL(t, ts)

		// Access URL twice
		url := fmt.Sprintf("/api/v1/shorten/%s", shortURL.ShortCode)
		ts.Request("GET", url, nil)
		ts.Request("GET", url, nil)

		// Check stats
		statsURL := fmt.Sprintf("/api/v1/shorten/%s/stats", shortURL.ShortCode)
		res := ts.Request("GET", statsURL, nil)
		require.Equal(t, http.StatusOK, res.Code)

		var urlStats db.URL
		err := json.NewDecoder(res.Body).Decode(&urlStats)
		require.NoError(t, err)
		require.Equal(t, shortURL.ShortCode, urlStats.ShortCode)
		require.Equal(t, int64(2), urlStats.AccessCount)
	})
}
