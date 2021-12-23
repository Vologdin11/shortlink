package service

import (
	"shortlink/mocks/mock_repository"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestGetLink(t *testing.T) {
	mockController := gomock.NewController(t)
	defer mockController.Finish()

	mockRepository := mock_repository.NewMockRepository(mockController)
	mockRepository.EXPECT().GetLink("google.com").Return("qweasd55_W", nil)

	service := Service{DB: mockRepository}
	link, err := service.GetLink("google.com")

	assert.NoError(t, err)
	assert.Equal(t, "qweasd55_W", link)
}
