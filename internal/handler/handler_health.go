package handler

import (
	"net/http"

	"github.com/fakelozic/go-todo/internal/models"
)

func HandlerHealth(w http.ResponseWriter, _ *http.Request) {
	msg := models.Response{Message: "OK"}
	ResponseWithJSON(w, http.StatusOK, msg)
}
