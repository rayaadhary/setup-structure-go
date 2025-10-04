package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/rayaadhary/social-go/cmd/auth"
	"github.com/rayaadhary/social-go/internal/service"
	"github.com/rayaadhary/social-go/internal/store"
	httpSwagger "github.com/swaggo/http-swagger"
)

type application struct {
	config   config
	store    store.Repos
	services services
}

type services struct {
	Posts *service.PostService
	Users *service.UserService
}

type config struct {
	addr string
	db   dbConfig
}

type dbConfig struct {
	addr         string
	maxOpenConns int
	maxIdleConns int
	maxIdleTime  string
}

func (app *application) mount() *chi.Mux {
	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Use(middleware.Timeout(60 * time.Second))

	r.Get("/health", app.healthCheckHandler)

	// Public route
	authHandler := auth.NewAuthHandler(app.services.Users)
	r.Post("/login", authHandler.Login)

	r.Route("/v1", func(r chi.Router) {
		r.Use(auth.AuthMiddleware)
		r.Get("/me", func(w http.ResponseWriter, r *http.Request) {
			if userID, ok := auth.GetUserID(r.Context()); ok {
				w.Write([]byte("Hello user with ID: " + fmt.Sprint(userID)))
			} else {
				http.Error(w, "no user found", http.StatusUnauthorized)
			}
		})

		r.Route("/posts", func(r chi.Router) {
			r.Post("/", app.createPostHandler)
			r.Get("/", app.listPostsHandler)
			r.Get("/{id}", app.getPostHandler)
			r.Put("/{id}", app.updatePostHandler)
			r.Delete("/{id}", app.deletePostHandler)
		})
	})

	// Swagger docs
	r.Get("/swagger/*", httpSwagger.WrapHandler)

	return r
}

func (app *application) run(mux *chi.Mux) error {

	srv := &http.Server{
		Addr:         app.config.addr,
		Handler:      mux,
		WriteTimeout: time.Second * 30,
		ReadTimeout:  time.Second * 10,
		IdleTimeout:  time.Minute,
	}

	log.Printf("server has started at %s", app.config.addr)

	return srv.ListenAndServe()
}
