package api

import "net/http"

type Storage interface {
	Close() error
}

type Loger interface {
	Insert(id int) int
}

type Api struct {
	Storage Storage
}

func New(db Storage) *Api {
	return &Api{Storage: db}
}

func (h *Api) Home() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}
