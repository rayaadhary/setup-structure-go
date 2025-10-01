package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/rayaadhary/social-go/docs"
	"github.com/rayaadhary/social-go/internal/db"
	"github.com/rayaadhary/social-go/internal/service"
	"github.com/rayaadhary/social-go/internal/store"
)

// @title Social Go API
// @version 1.0
// @description This is a sample server for Social Go API.
// @host localhost:8080
// @BasePath /v1
func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("error loading .env file")
	}

	cfg := config{
		addr: os.Getenv("APP_ADDR"),
		db: dbConfig{
			addr:         os.Getenv("DB_ADDR"),
			maxOpenConns: 30,
			maxIdleConns: 30,
			maxIdleTime:  os.Getenv("DB_MAX_IDDLE_TIME"),
		},
	}

	db, err := db.New(cfg.db.addr,
		cfg.db.maxOpenConns,
		cfg.db.maxIdleConns,
		cfg.db.maxIdleTime,
	)

	if err != nil {
		log.Panic(err)
	}

	defer db.Close()
	log.Printf("db connection pool established")

	store := store.NewRepos(db)

	services := services{
		Posts: service.NewPostService(store.Posts),
	}

	app := &application{
		config:   cfg,
		store:    store,
		services: services,
	}

	mux := app.mount()

	log.Fatal(app.run(mux))
}
