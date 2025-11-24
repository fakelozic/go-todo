package main

import (
	"context"
	"log"
	"net/http"

	"github.com/fakelozic/go-todo/internal/config"
	"github.com/fakelozic/go-todo/internal/database"
	"github.com/fakelozic/go-todo/internal/handler"
	"github.com/fakelozic/go-todo/internal/middleware"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"github.com/jackc/pgx/v5/pgxpool"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatal("error loading config", err)
	}
	pgxPoolConfig, err := pgxpool.ParseConfig(cfg.Database.DSN)
	if err != nil {
		log.Fatalf("failed to parse pgx pool config: %v", err)
	}
	dbPool, err := pgxpool.NewWithConfig(context.Background(), pgxPoolConfig)
	if err != nil {
		log.Fatalf("failed to create pgx pool: %v", err)
	}
	defer dbPool.Close()

	handlerApi := handler.ApiConfig{
		DB: database.New(dbPool),
	}
	middlewareApi := middleware.ApiConfig{
		DB: database.New(dbPool),
	}

	router := chi.NewRouter()

	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{cfg.Server.CORSAllowedOrigins[0]},
		AllowedMethods:   []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"*"},
		AllowCredentials: false,
		MaxAge:           300,
	}))

	v1Router := chi.NewRouter()

	v1Router.Get("/health", handler.HandlerHealth)

	v1Router.Post("/users", handlerApi.HandlerCreateUsers)
	v1Router.Get("/users", middlewareApi.MiddlewareAuth(handlerApi.HandlerGetUser))

	v1Router.Post("/tasks", middlewareApi.MiddlewareAuth(handlerApi.HandlerCreateTasks))
	v1Router.Get("/tasks", middlewareApi.MiddlewareAuth(handlerApi.HandlerGetAllTasks))
	v1Router.Put("/task/{taskID}", middlewareApi.MiddlewareAuth(handlerApi.HandleUpdateTask))
	v1Router.Patch("/task/{taskID}", middlewareApi.MiddlewareAuth(handlerApi.HandleToggleTask))
	v1Router.Delete("/task/{taskID}", middlewareApi.MiddlewareAuth(handlerApi.HandleDeleteTask))

	router.Mount("/v1", v1Router)

	srv := &http.Server{
		Handler: router,
		Addr:    ":" + cfg.Server.Port,
	}

	log.Printf("server running on port %v", cfg.Server.Port)
	err = srv.ListenAndServe()
	if err != nil {
		log.Fatal("error starting server", err)
	}
}
