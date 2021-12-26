package service

import (
	"errors"
	"fmt"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"shortlink/mocks/mock_repository"
	"testing"
)

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

func TestGetLink(t *testing.T) {
	mockController := gomock.NewController(t)
	mockRepository := mock_repository.NewMockRepository(mockController)
	defer mockController.Finish()

	t.Run("got link", func(t *testing.T) {
		mockRepository.EXPECT().GetLink("google.com").Return("qweasd55_W", nil)

		service := Service{DB: mockRepository}
		link, err := service.GetLink("google.com")

		assert.NoError(t, err)
		assert.Equal(t, "qweasd55_W", link)
	})

	t.Run("no link", func(t *testing.T) {
		mockRepository.EXPECT().GetLink("google.com").Return("", errors.New("link not found"))

		service := Service{DB: mockRepository}
		link, err := service.GetLink("google.com")

		assert.Error(t, err)
		assert.Equal(t, "", link)
	})
}

func TestGetShortLink(t *testing.T) {
	mockController := gomock.NewController(t)
	mockRepository := mock_repository.NewMockRepository(mockController)
	defer mockController.Finish()

	t.Run("link in db", func(t *testing.T) {
		mockRepository.EXPECT().GetShortLink("google.com").Return("asdfqwe", nil)

		service := Service{DB: mockRepository}
		link, err := service.GetShortLink(("google.com"))

		assert.NoError(t, err)
		assert.Equal(t, "asdfqwe", link)
	})

	t.Run("generate new shortlink", func(t *testing.T) {
		//???!!!
		mockRepository.EXPECT().GetShortLink("google.com").Return("", errors.New("no"))
		mockRepository.EXPECT().GetLink("google.com").Return("", errors.New("no"))
		mockRepository.EXPECT().AddLink("link", "shortlink").Return(nil)

		service := Service{DB: mockRepository}
		link, err := service.GetShortLink("google.com")

		assert.NoError(t, err)
		assert.NotEqual(t, "", link)
	})
}
