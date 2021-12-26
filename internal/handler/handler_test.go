package handler

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"shortlink/mocks/mock_service"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGetlink(t *testing.T) {
	mockController := gomock.NewController(t)
	mockServices := mock_service.NewMockServices(mockController)
	testHandler := NewHandler(mockServices)

	t.Run("link in db", func(t *testing.T) {
		r, err := http.NewRequest("GET", "http://localhost:8080/yandex", nil)
		require.NoError(t, err)
		rw := httptest.NewRecorder()

		mockServices.EXPECT().GetLink(gomock.Any()).Return("yandex", nil)

		testHandler.getLink(rw, r)

		assert.Equal(t, "<a href=\"/yandex\">Permanent Redirect</a>.\n\n", rw.Body.String())
	})

	t.Run("No link", func(t *testing.T) {
		r, err := http.NewRequest("GET", "http://localhost:8080/google.com/", nil)
		require.NoError(t, err)
		rw := httptest.NewRecorder()

		mockServices.EXPECT().GetLink("google.com").Return("", errors.New("no link"))

		testHandler.getLink(rw, r)

		assert.Equal(t, "404 page not found\n", rw.Body.String())
		assert.Equal(t, http.StatusNotFound, rw.Result().StatusCode)
	})
}

func TestShortLink(t *testing.T) {
	mockController := gomock.NewController(t)
	mockServices := mock_service.NewMockServices(mockController)
	testHandler := NewHandler(mockServices)

	t.Run("link in db or generate new", func(t *testing.T) {
		r, err := http.NewRequest("POST", "http://localhost:8080/?url=yandex", nil)
		require.NoError(t, err)
		rw := httptest.NewRecorder()

		mockServices.EXPECT().GetShortLink(gomock.Any()).Return("qwerty", nil)

		testHandler.getShortLink(rw, r)

		assert.Equal(t, "http://localhost:8000/qwerty", rw.Body.String())
	})

	t.Run("error generate link", func(t *testing.T) {
		r, err := http.NewRequest("POST", "http://localhost:8080/?url=google.com/", nil)
		require.NoError(t, err)
		rw := httptest.NewRecorder()

		mockServices.EXPECT().GetShortLink(gomock.Any()).Return("", errors.New("no link"))

		testHandler.getShortLink(rw, r)

		assert.Equal(t, "\n", rw.Body.String())
		assert.Equal(t, http.StatusInternalServerError, rw.Result().StatusCode)
	})
}
