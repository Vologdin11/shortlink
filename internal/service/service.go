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
	//вернуть
	//задавать через переменные окружения добавить перед возвращением
	//serviceUrl := "http://localhost:8000/"
	return shortlink, nil
}

func generateShortLink() []byte {
	shortLink := make([]byte, 10)
	//проверки
	symbols := []string{
		"abcdefghijklmnopqrstuvwxyz",
		"ABCDEFGHIJKLMNOPQRSTUVWXYZ",
		"0123456789",
		"_",
	}
	usedSymbols := make([]bool, 4)
	for i := range shortLink {
		//оптимихировать
		if j, check := checkUnusedSymbols(usedSymbols); i > 6 && check {
			shortLink[i] = symbols[j][rand.Intn(len(symbols[j]))]
		} else {
			shortLink[i] = generateSymbol(symbols, usedSymbols)
		}
	}
	//проверить что ссылка уникальна, иначе сгенерировать еще раз
	return shortLink
}

func checkUnusedSymbols(usedSymbols []bool) (int, bool) {
	for i := range usedSymbols {
		if !usedSymbols[i] {
			return i, true
		}
	}
	return 0, false
}

func generateSymbol(symbols []string, usedSymbols []bool) byte {
	var symbolType int
	//выбрать что вставлять
	if usedSymbols[3] {
		//выбрать из 3-х
		symbolType = rand.Intn(3)
	} else {
		//выбрать из всех
		symbolType = rand.Intn(4)
	}
	//указываем что символ использован
	usedSymbols[symbolType] = true
	//выбираем символ
	return symbols[symbolType][rand.Intn(len(symbols[symbolType]))]
}
