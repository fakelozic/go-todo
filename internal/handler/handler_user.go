package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/fakelozic/go-todo/internal/database"
	"github.com/google/uuid"
)

func (apiCfg *ApiConfig) HandlerCreateUsers(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		Username string `json:"username"`
	}
	decoder := json.NewDecoder(r.Body)
	defer r.Body.Close()

	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		ResponseWithError(w, 400, fmt.Sprintf("error parsing JSON: %v", err))
	}

	user, err := apiCfg.DB.CreateUser(r.Context(), database.CreateUserParams{
		ID:        uuid.New(),
		Username:  params.Username,
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
	})
	if err != nil {
		ResponseWithError(w, 400, fmt.Sprintf("couldn't create user: %v", err))
	}

	ResponseWithJSON(w, http.StatusCreated, user)
}

func (apiCfg *ApiConfig) HandlerGetUser(w http.ResponseWriter, r *http.Request, user database.User) {
	ResponseWithJSON(w, http.StatusOK, user)
}
