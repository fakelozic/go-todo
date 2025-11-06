package models

import (
	"time"

	"github.com/fakelozic/go-todo/internal/database"
	"github.com/google/uuid"
)

type Task struct {
	ID        uuid.UUID `json:"id"`
	Task      string    `json:"task"`
	Status    bool      `json:"status"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	UserId    uuid.UUID `json:"user_id"`
}

type User struct {
	ID        uuid.UUID `json:"id"`
	Username  string    `json:"username"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	APIKey    string    `json:"api_key"`
}

type Response struct {
	Message string `json:"message"`
}

func DatabaseUserToUser(dbUser database.User) User {
	return User{
		ID:        dbUser.ID,
		Username:  dbUser.Username,
		CreatedAt: dbUser.CreatedAt,
		UpdatedAt: dbUser.UpdatedAt,
		APIKey:    dbUser.ApiKey,
	}
}

func DatabaseTaskToTask(dbTask database.Task) Task {
	return Task{
		ID:        dbTask.ID,
		Task:      dbTask.Task,
		Status:    dbTask.Status,
		CreatedAt: dbTask.CreatedAt,
		UpdatedAt: dbTask.UpdatedAt,
		UserId:    dbTask.UserID,
	}
}

func DatabaseTasksToTasks(dbTasks []database.Task) []Task {
	tasks := []Task{}

	for _, dbFeed := range dbTasks {
		tasks = append(tasks, DatabaseTaskToTask(dbFeed))
	}
	return tasks
}
