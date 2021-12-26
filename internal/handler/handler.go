package handler

import (
	"fmt"
	"net/http"
	"shortlink/internal/service"
	"strings"
)

type Handler struct {
	service service.Services
}

func NewHandler(service service.Services) *Handler {
	return &Handler{service: service}
}

func (h Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if len(r.URL.Path) < 2 {
		return
	}

	switch r.Method {
	case "GET":
		h.getLink(w, r)
	case "POST":
		h.shortLink(w, r)
	default:
		http.NotFound(w, r)
	}
}

func (h *Handler) getLink(w http.ResponseWriter, r *http.Request) {
	//Проверить ссылку на уникальность и вернуть ее или ошибку 404
	link, err := h.service.GetLink(strings.Trim(r.URL.Path, "/"))
	if err != nil {
		http.NotFound(w, r)
		return
	}

	http.Redirect(w, r, link, http.StatusPermanentRedirect)
}

func (h *Handler) shortLink(w http.ResponseWriter, r *http.Request) {
	link, err := h.service.GetShortLink(strings.Trim(r.URL.Query().Get("url"), "/"))
	//вывести ошибку что-то пошло не так
	if err != nil {
		return
	}
	_, err = fmt.Fprint(w, "http://localhost:8000/"+link)
	//вывести ошибку что-то пошло не так
	if err != nil {
		return
	}
}
