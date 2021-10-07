package main

import (
	"encoding/json"
	"github.com/go-redis/redis"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	"gorm.io/gorm"
)

type App struct {
	Router *mux.Router
	DB     *gorm.DB
	Redis  *redis.Client
}

func (a *App) Initialize() {
	a.Router = mux.NewRouter()
	a.initializeRoutes()

	a.DB = InitDb()
	a.DB.AutoMigrate(&link{})
	a.Redis = InitRedis()
}

func respondWithError(w http.ResponseWriter, code int, message string) {
	respondWithJSON(w, code, map[string]string{"error": message})
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

func (a *App) Run() {
	log.Fatal(http.ListenAndServe(":8010", a.Router))
}

func (a *App) initializeRoutes() {
	a.Router.HandleFunc("/links/{id}", a.getLink).Methods("GET")
	a.Router.HandleFunc("/links", a.createLink).Methods("POST")

	a.Router.HandleFunc("/{id}", a.redirect).Methods("GET")
}
