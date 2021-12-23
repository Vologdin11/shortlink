package service

import (
	"fmt"
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

func TestGenerateShortLink(t *testing.T) {
	shortLink := generateShortLink()
	shortLink1 := generateShortLink()

	fmt.Println(string(shortLink), string(shortLink1))
	assert.Len(t, shortLink, 10)
	assert.NotEqual(t, shortLink, shortLink1)
}

func BenchmarkGenerateShortLink(b *testing.B) {
	for i := 0; i < b.N; i++ {
		generateShortLink()
	}
	b.ReportAllocs()
}
