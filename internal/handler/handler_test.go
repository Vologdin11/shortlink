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
	t.Run("No path", func(t *testing.T) {
		r, err := http.NewRequest("GET", "http://localhost:8080/", nil)
		require.NoError(t, err)
		rw := httptest.NewRecorder()

		handlers := Handler{}
		handlers.getLink(rw, r)

		assert.Empty(t, rw.Body.String())
	})

	t.Run("No link", func(t *testing.T) {
		r, err := http.NewRequest("GET", "http://localhost:8080/google.com/", nil)
		require.NoError(t, err)
		rw := httptest.NewRecorder()

		mockController := gomock.NewController(t)
		defer mockController.Finish()

		mockService := mock_service.NewMockServices(mockController)
		mockService.EXPECT().GetLink("google.com").Return("", errors.New("no link"))

		handlers := Handler{service: mockService}
		handlers.getLink(rw, r)

		assert.Equal(t, "404 page not found\n", rw.Body.String())
		assert.Equal(t, http.StatusNotFound, rw.Result().StatusCode)
	})
}
