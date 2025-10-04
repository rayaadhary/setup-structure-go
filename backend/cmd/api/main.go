package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/rayaadhary/social-go/docs"
	"github.com/rayaadhary/social-go/internal/db"
	"github.com/rayaadhary/social-go/internal/service"
	"github.com/rayaadhary/social-go/internal/store"
	"golang.org/x/crypto/bcrypt"
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

	password := "123456"
	hash, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	_, err = db.Exec(`
		INSERT INTO users (username, password)
		VALUES ($1, $2)
		ON CONFLICT (username) DO NOTHING
	`, "superadmin", string(hash))

	if err != nil {
		log.Panic(err)
	}

	defer db.Close()
	log.Printf("db connection pool established")

	store := store.NewRepos(db)

	services := services{
		Posts: service.NewPostService(store.Posts),
		Users: service.NewUserService(store.Users),
	}

	app := &application{
		config:   cfg,
		store:    store,
		services: services,
	}

	mux := app.mount()

	log.Fatal(app.run(mux))
}
