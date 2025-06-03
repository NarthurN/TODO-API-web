package api

import (
	"fmt"
	"net/http"
	"time"

	"github.com/NarthurN/TODO-API-web/pkg/loger"
)

type Storage interface {
	Close() error
}

type Api struct {
	Storage Storage
}

func New(db Storage) *Api {
	return &Api{Storage: db}
}

func (h *Api) NextDayHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		nowStr := r.FormValue("now")
		dateStr := r.FormValue("date")
		repeat := r.FormValue("repeat")
		loger.L.Info("FormValue:", "nowStr", nowStr)
		loger.L.Info("FormValue:", "dateStr", dateStr)
		loger.L.Info("FormValue:", "repeat", repeat)
		var now time.Time
		if nowStr == "" {
			now = time.Now().UTC()
		} else {
			var err error
			now, err = time.Parse(layout, nowStr)
			if err != nil {
				http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
				return
			}
		}

		if dateStr == "" || repeat == "" {
			http.Error(w, "Missing parameters: now, date or repeat", http.StatusBadRequest)
			return
		}

		newDate, err := NextDate(now, dateStr, repeat)
		if err != nil {
			http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
			return
		}

		loger.L.Info("Ответ newDate:", "newDate", newDate)
		fmt.Fprint(w, newDate)
	})
}
