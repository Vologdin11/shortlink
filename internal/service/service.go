package service

import (
	"context"
	"shortlink/internal/repository"
)

type Services interface {
	GetLink(shortLink string) (string, error)
	GetShortLink(link string) (string, error)
	AddLink(link string) error
}

type Service struct {
	DB repository.Repository
}

func (s *Service) GetLink(shortLink string) (string, error) {
	//Проверить уникальность ссылки
	link, err := s.DB.GetLink(context.Background(), shortLink)
	if err != nil {
		return "", err
	}
	return link, nil
}

func (s *Service) GetShortLink(link string) (string, error) {
	return "", nil
}

func (s *Service) AddLink(link string) error {
	return nil
}
