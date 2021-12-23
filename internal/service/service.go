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
	//проблема с БД
	if err != nil {
		return "", nil
	}
	//Не уникальна, вернуть ее сокращенную версию
	if shortlink != "" {
		return shortlink, nil
	}
	//Уникальная сгенерировать сокращенную версию

	//добавить в БД

	//вернуть
	//задавать через переменные окружения добавить перед возвращением
	//serviceUrl := "http://localhost:8000/"
	return "", nil
}

func generateShortLink() []rune {
	shortLink := make([]rune, 10)
	//проверки
	symbols := []string{
		"abcdefghijklmnopqrstuvwxyz",
		"ABCDEFGHIJKLMNOPQRSTUVWXYZ",
		"0123456789",
		"_",
	}
	checkSymbols := make([]bool, 4)
	var symbol rune
	for i := range shortLink {
		if i > 6 {
			for j := range checkSymbols {
				if checkSymbols[j] == false {
					symbol = rune(symbols[j][rand.Intn(len(symbols[j]))])
					break
				}
			}
		} else {
			symbol = generateSymbol(symbols, checkSymbols)
		}
		shortLink[i] = symbol
	}
	//проверить что ссылка уникальна, иначе сгенерировать еще раз
	return shortLink
}

func generateSymbol(symbols []string, checkSymbols []bool) rune {
	var symbolType int
	//выбрать что вставлять
	if checkSymbols[3] {
		//выбрать из 3-х
		symbolType = rand.Intn(3)
	} else {
		//выбрать из всех
		symbolType = rand.Intn(4)
	}
	//указываем что символ использован
	checkSymbols[symbolType] = true
	//выбираем символ
	return rune(symbols[symbolType][rand.Intn(len(symbols[symbolType]))])
}
