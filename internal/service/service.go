package service

import (
	"math/rand"
	"shortlink/internal/repository"
)

type Services interface {
	GetLink(shortLink string) (string, error)
	GetShortLink(link string) (string, error)
}

type Service struct {
	DB repository.Repository
}

func (s *Service) GetLink(shortLink string) (string, error) {
	link, err := s.DB.GetLink(shortLink)
	if err != nil {
		return "", err
	}
	return link, nil
}

func (s *Service) GetShortLink(link string) (string, error) {
	//Проверить уникальность ссылки
	shortlink, err := s.DB.GetShortLink(link)
	//Не уникальна, вернуть ее сокращенную версию
	if err == nil {
		return shortlink, nil
	}

	for i := 0; i < 1000; i++ {
		//Уникальная сгенерировать сокращенную версию
		shortlink = string(generateShortLink())
		//проверить что сгенерированная ссылка уникальна
		_, err := s.DB.GetLink(shortlink)
		//Укикальна остановить
		if err != nil {
			break
		}
	}
	//добавить в БД
	err = s.DB.AddLink(link, shortlink)
	if err != nil {
		return "", err
	}

	return shortlink, nil
}

func generateShortLink() []byte {
	shortLink := make([]byte, 10)
	usedSymbols := make([]bool, 4)
	symbols := []string{
		"abcdefghijklmnopqrstuvwxyz",
		"ABCDEFGHIJKLMNOPQRSTUVWXYZ",
		"0123456789",
		"_",
	}

	for i := range shortLink {
		if i < 7 {
			shortLink[i] = generateSymbol(symbols, usedSymbols)
		} else {
			unusedSymbols := checkUnusedSymbols(usedSymbols)
			shortLink[i] = symbols[unusedSymbols][rand.Intn(len(symbols[unusedSymbols]))]
		}
	}
	return shortLink
}

func checkUnusedSymbols(usedSymbols []bool) int {
	for i := range usedSymbols {
		if !usedSymbols[i] {
			usedSymbols[i] = true
			return i
		}
	}
	return 0
}

func generateSymbol(symbols []string, usedSymbols []bool) byte {
	var symbolType int
	//выбрать что вставлять
	if usedSymbols[3] {
		symbolType = rand.Intn(3)
	} else {
		symbolType = rand.Intn(4)
	}
	//указываем что символ использован
	usedSymbols[symbolType] = true
	//выбираем символ
	return symbols[symbolType][rand.Intn(len(symbols[symbolType]))]
}
