package handler

import (
	"encoding/json"
	"log"
	"net/http"
)

func ResponseWithJSON(w http.ResponseWriter, code int, payload any) {
	data, err := json.Marshal(payload)
	if err != nil {
		log.Printf("filed to marshal json respond: %v", err)
		w.WriteHeader(500)
		return
	}
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(data)
}

func ResponseWithError(w http.ResponseWriter, code int, msg string) {
	if code > 499 {
		log.Println("responding with 5xx error:", msg)
	}

	type errResponse struct {
		Error string `json:"error"`
	}

	ResponseWithJSON(w, code, errResponse{Error: msg})
}
