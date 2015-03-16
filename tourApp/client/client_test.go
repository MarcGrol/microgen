package client

import (
	"fmt"
	"github.com/MarcGrol/microgen/lib/myerrors"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestCreateTourOk(t *testing.T) {
	ts := startServer(200, "{\"status\":true}")
	defer ts.Close()

	client := NewClient(ts.URL)
	err := client.CreateTour(2015)
	assert.Nil(t, err)
}

func TestCreateTourInvalidInputError(t *testing.T) {
	ts := startServer(400, "{\"status\":false,\"error\": {\"code\":4,\"message\":\"invalid input\"}}")
	defer ts.Close()

	client := NewClient(ts.URL)
	err := client.CreateTour(2015)
	assert.NotNil(t, err)
	assert.True(t, myerrors.IsInvalidInputError(err))
	assert.Contains(t, err.Error(), "invalid input")
}

func TestCreateTourInternalError(t *testing.T) {
	ts := startServer(500, "{\"status\":false,\"error\": {\"code\":4,\"message\":\"internal problem\"}}")
	defer ts.Close()

	client := NewClient(ts.URL)
	err := client.CreateTour(2015)
	assert.NotNil(t, err)
	assert.True(t, myerrors.IsInternalError(err))
	assert.Contains(t, err.Error(), "internal problem")
}

func TestCreateTourNotAuthorizedError(t *testing.T) {
	ts := startServer(401, "{\"status\":false,\"error\": {\"code\":4,\"message\":\"not authorized\"}}")
	defer ts.Close()

	client := NewClient(ts.URL)
	err := client.CreateTour(2015)
	assert.NotNil(t, err)
	assert.True(t, myerrors.IsNotAuthorizedError(err))
	assert.Contains(t, err.Error(), "not authorized")
}

func TestCreateTourServiceNotAvailableError(t *testing.T) {
	ts := startServer(503, "{\"status\":false,\"error\": {\"code\":4,\"message\":\"not available\"}}")
	defer ts.Close()

	client := NewClient(ts.URL)
	err := client.CreateTour(2015)
	assert.NotNil(t, err)
	assert.False(t, myerrors.IsInternalError(err))
	assert.False(t, myerrors.IsNotAuthorizedError(err))
	assert.False(t, myerrors.IsNotFoundError(err))
	assert.False(t, myerrors.IsInvalidInputError(err))
	assert.Contains(t, err.Error(), "not available")
}

func TestCreateTourUnspecifiedError(t *testing.T) {

	ts := startServer(500, "Internal error")
	defer ts.Close()

	client := NewClient(ts.URL)
	err := client.CreateTour(2015)
	assert.NotNil(t, err)
	assert.False(t, myerrors.IsInternalError(err))
	assert.False(t, myerrors.IsNotAuthorizedError(err))
	assert.False(t, myerrors.IsNotFoundError(err))
	assert.False(t, myerrors.IsInvalidInputError(err))
	assert.Contains(t, err.Error(), "Error unmarshalling response")
}

func TestCreateCyclistOk(t *testing.T) {
	ts := startServer(200, "{\"status\":true}")
	defer ts.Close()

	client := NewClient(ts.URL)
	err := client.CreateCyclist(2015, 42, "Boogerd", "Rabo")
	assert.Nil(t, err)
}

func TestCreateEtappeOk(t *testing.T) {
	ts := startServer(200, "{\"status\":true}")
	defer ts.Close()

	client := NewClient(ts.URL)
	err := client.CreateEtappe(2015, 2, time.Date(2015, time.July, 14, 9, 0, 0, 0, time.Local),
		"Utrecht", "Valkenburg", 255, 1)
	assert.Nil(t, err)
}

func TestCreateEtappeNotFoundError(t *testing.T) {
	ts := startServer(404, "{\"status\":false,\"error\": {\"code\":4,\"message\":\"tour not found\"}}")
	defer ts.Close()

	client := NewClient(ts.URL)
	err := client.CreateEtappe(2015, 2, time.Date(2015, time.July, 14, 9, 0, 0, 0, time.Local),
		"Utrecht", "Valkenburg", 255, 1)
	assert.NotNil(t, err)
	assert.True(t, myerrors.IsNotFoundError(err))
	assert.Contains(t, err.Error(), "tour not found")
}

func TestCreateGamblerOk(t *testing.T) {
	ts := startServer(200, "{\"status\":true}")
	defer ts.Close()

	client := NewClient(ts.URL)
	err := client.CreateGambler("123", "marc", "mgrol@home.nl")
	assert.Nil(t, err)
}

func TestCreateGamblerTeamOk(t *testing.T) {
	ts := startServer(200, "{\"status\":true}")
	defer ts.Close()

	client := NewClient(ts.URL)
	err := client.CreateGamblerTeam(2015, "marc", []int{1, 2, 3})
	assert.Nil(t, err)
}

func TestCreateEtappeResultOk(t *testing.T) {
	ts := startServer(200, "{\"status\":true}")
	defer ts.Close()

	client := NewClient(ts.URL)
	err := client.CreateEtappeResults(2015, 3, []int{1, 2, 3}, []int{4, 5, 6}, []int{7, 8, 9}, []int{10, 12, 12})
	assert.Nil(t, err)
}

func startServer(httpCode int, httpResponse string) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(httpCode)
		fmt.Fprintln(w, httpResponse)
	}))
}
