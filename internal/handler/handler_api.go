package handler

import (
	"github.com/fakelozic/go-todo/internal/database"
)

type ApiConfig struct {
	DB *database.Queries
}
