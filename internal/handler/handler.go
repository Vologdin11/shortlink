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
	switch r.Method {
	case "GET":
		h.getLink(w, r)
	case "POST":
		h.getShortLink(w, r)
	default:
		http.Error(w, "", http.StatusMethodNotAllowed)
	}
}

func (h *Handler) getLink(w http.ResponseWriter, r *http.Request) {
	if len(r.URL.Path) < 2 {
		http.Error(w, "no path", http.StatusBadRequest)
		return
	}
	//Проверить ссылку на уникальность и вернуть ее или ошибку
	link, err := h.service.GetLink(strings.Trim(r.URL.Path, "/"))
	if err != nil {
		http.NotFound(w, r)
		return
	}

	http.Redirect(w, r, link, http.StatusPermanentRedirect)
}

func (h *Handler) getShortLink(w http.ResponseWriter, r *http.Request) {
	if r.URL.Query().Get("url") == "" {
		http.Error(w, "no query", http.StatusBadRequest)
		return
	}

	link, err := h.service.GetShortLink(strings.Trim(r.URL.Query().Get("url"), "/"))
	//вывести ошибку что-то пошло не так
	if err != nil {
		http.Error(w, "", http.StatusInternalServerError)
		return
	}
	_, err = fmt.Fprint(w, "http://localhost:8000/"+link)
	//вывести ошибку что-то пошло не так
	if err != nil {
		http.Error(w, "", http.StatusInternalServerError)
		return
	}
}
