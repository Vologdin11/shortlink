package handler

import (
	"net/http"
	"shortlink/internal/service"
	"strings"

	"github.com/gorilla/mux"
)

type Handler struct {
	service service.Services
}

func NewHandler(service service.Services) *Handler {
	return &Handler{service: service}
}

func (h *Handler) InitRouter() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/", h.shortLink).Methods("POST")
	router.HandleFunc("/", h.getLink).Methods("GET").PathPrefix("/")

	return router
}

func (h *Handler) getLink(w http.ResponseWriter, r *http.Request) {
	//вернутьь подсказку как пользаваться сервисом
	if len(r.URL.Path) < 2 {
		return
	}
	//Проверить ссылку на уникальность и вернуть ее или ошибку 404
	link, err := h.service.GetLink(strings.Trim(r.URL.Path, "/"))
	if err != nil {
		http.NotFound(w, r)
		return
	}
	//Если ошибки нет сделать редирект по ссылке
	http.Redirect(w, r, link, http.StatusPermanentRedirect)
}

func (h *Handler) shortLink(w http.ResponseWriter, r *http.Request) {

}
