package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/fakelozic/go-todo/internal/database"
	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
)

func (apiCfg *ApiConfig) HandlerCreateTasks(w http.ResponseWriter, r *http.Request, user database.User) {
	type parameters struct {
		Task string `json:"task"`
	}
	params := parameters{}

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&params)
	if err != nil {
		ResponseWithError(w, 400, fmt.Sprintf("error parsing JSON: %v", err))
		return
	}

	task, err := apiCfg.DB.CreateTask(r.Context(), database.CreateTaskParams{
		ID:        uuid.New(),
		Task:      params.Task,
		Status:    false,
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		UserID:    user.ID,
	})
	if err != nil {
		ResponseWithError(w, 400, fmt.Sprintf("couldn't create task: %v", err))
	}
	ResponseWithJSON(w, http.StatusCreated, task)
}

func (apiCfg *ApiConfig) HandlerGetAllTasks(w http.ResponseWriter, r *http.Request, user database.User) {
	tasks, err := apiCfg.DB.GetAllTasks(r.Context(), user.ID)
	if err != nil {
		ResponseWithError(w, 400, fmt.Sprintf("couldn't get all task: %v", err))
	}
	ResponseWithJSON(w, http.StatusOK, tasks)
}

func (apiConfig *ApiConfig) HandleUpdateTask(w http.ResponseWriter, r *http.Request, user database.User) {
	taskIDStr := chi.URLParam(r, "taskID")
	taskID, err := uuid.Parse(taskIDStr)
	if err != nil {
		ResponseWithError(w, 400, fmt.Sprintf("error paring task id: %v", err))
		return
	}

	foundTask, err := apiConfig.DB.GetTaskByID(r.Context(), database.GetTaskByIDParams{ID: taskID, UserID: user.ID})

	if foundTask.Task == "" {
		ResponseWithError(w, 400, fmt.Sprintf("task not found: %v", err))
		return
	}
	if err != nil {
		ResponseWithError(w, 400, fmt.Sprint(err.Error()))
		return
	}

	type parameters struct {
		Task string `json:"task"`
	}
	params := parameters{}

	decoder := json.NewDecoder(r.Body)
	err = decoder.Decode(&params)
	if err != nil {
		ResponseWithError(w, 400, fmt.Sprintf("error parsing JSON: %v", err))
		return
	}

	task, err := apiConfig.DB.UpdateTask(r.Context(), database.UpdateTaskParams{
		Task:      params.Task,
		UpdatedAt: time.Now().UTC(),
		ID:        taskID,
		UserID:    user.ID,
	})
	if err != nil {
		ResponseWithError(w, 400, fmt.Sprint(err.Error()))
		return
	}

	ResponseWithJSON(w, http.StatusAccepted, task)
}

func (apiConfig *ApiConfig) HandleToggleTask(w http.ResponseWriter, r *http.Request, user database.User) {
	taskIDStr := chi.URLParam(r, "taskID")
	taskID, err := uuid.Parse(taskIDStr)
	if err != nil {
		ResponseWithError(w, 400, fmt.Sprintf("error paring task id: %v", err))
		return
	}

	foundTask, err := apiConfig.DB.GetTaskByID(r.Context(), database.GetTaskByIDParams{ID: taskID, UserID: user.ID})

	if foundTask.Task == "" {
		ResponseWithError(w, 400, fmt.Sprintf("task not found: %v", err))
		return
	}
	if err != nil {
		ResponseWithError(w, 400, fmt.Sprint(err.Error()))
		return
	}

	task, err := apiConfig.DB.ToggleTask(r.Context(), database.ToggleTaskParams{
		Status:    !foundTask.Status,
		UpdatedAt: time.Now().UTC(),
		ID:        taskID,
		UserID:    user.ID,
	})
	if err != nil {
		ResponseWithError(w, 400, fmt.Sprint(err.Error()))
		return
	}

	ResponseWithJSON(w, http.StatusAccepted, task)
}

func (apiConfig *ApiConfig) HandleDeleteTask(w http.ResponseWriter, r *http.Request, user database.User) {
	taskIDStr := chi.URLParam(r, "taskID")
	taskID, err := uuid.Parse(taskIDStr)
	if err != nil {
		ResponseWithError(w, 400, fmt.Sprintf("error paring task id: %v", err))
		return
	}

	foundTask, err := apiConfig.DB.GetTaskByID(r.Context(), database.GetTaskByIDParams{ID: taskID, UserID: user.ID})

	if foundTask.Task == "" {
		ResponseWithError(w, 400, fmt.Sprintf("task not found: %v", err))
		return
	}
	if err != nil {
		ResponseWithError(w, 400, fmt.Sprint(err.Error()))
		return
	}

	task, err := apiConfig.DB.DeleteTask(r.Context(), database.DeleteTaskParams{
		ID:     taskID,
		UserID: user.ID,
	})
	if err != nil {
		ResponseWithError(w, 400, fmt.Sprint(err.Error()))
		return
	}

	ResponseWithJSON(w, http.StatusAccepted, task)
}
