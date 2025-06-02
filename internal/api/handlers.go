package api

import "net/http"

type Storage interface {
	Insert(id int) int
}

type Loger interface {
	Insert(id int) int
}

type Api struct {
	Storage Storage
	Loger   Loger
}

func New() *Api {
	return &Api{}
}

func (h *Api) Home() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}
