package api

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/NarthurN/TODO-API-web/pkg/loger"
)

var ErrInvalidJSONFormat error = errors.New("invalid JSON format")
var ErrTitleIsEmpty error = errors.New("title is empty")
var ErrInvalidDate error = errors.New("date is in invalid format")

type Storage interface {
	GetTasks(limit int, search string) ([]Task, error)
	AddTask(task Task) (int64, error)
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
			now, err = time.Parse(Layout, nowStr)
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

func (h *Api) AddTaskHandle() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var task Task
		if err := json.NewDecoder(r.Body).Decode(&task); err != nil {
			loger.L.Error(ErrInvalidJSONFormat.Error())
			SendErrorResponse(w, ErrInvalidJSONFormat.Error())
			return
		}

		if task.Title == "" {
			loger.L.Error(ErrTitleIsEmpty.Error())
			SendErrorResponse(w, ErrTitleIsEmpty.Error())
			return
		}

		err := checkDate(&task)
		if err != nil {
			loger.L.Error(ErrTitleIsEmpty.Error())
			SendErrorResponse(w, ErrTitleIsEmpty.Error())
			return
		}

		id, err := h.Storage.AddTask(task)
		loger.L.Info("Отпраляем id", "id", id)
		task.ID = strconv.Itoa(int(id))
		if err != nil {
			loger.L.Error(err.Error())
			SendErrorResponse(w, err.Error())
			return
		}

		if err = json.NewEncoder(w).Encode(task); err != nil {
			loger.L.Info("Отпраляем id", "id", id)
			SendIdResponse(w, id)
			return
		}
	})
}

func (h *Api) GetTasksHandle() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		search := r.URL.Query().Get("search")
		tasks, err := h.Storage.GetTasks(50, search) // в параметре максимальное количество записей
		if err != nil {
			SendErrorResponse(w, err.Error())
			return
		}
		WriteJSON(w, TasksResponse{
			Tasks: tasks,
		})
	})
}
