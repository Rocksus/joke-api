package main

import (
	"log"
	"math/rand"
	"net/http"
	"os"
	"time"

	"github.com/Rocksus/joke-api/internal/repository/joke"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func init() {
	rand.Seed(time.Now().UTC().UnixNano())
}

func main() {
	r := mux.NewRouter()

	joke.Load("data/jokes.json")
	jokeHandler := joke.InitHandler()
	randomHandler := joke.InitRandomHandler()

	r.HandleFunc("/jokes", randomHandler).Methods("GET")
	r.HandleFunc("/jokes/{category}", randomHandler).Methods("GET")
	r.HandleFunc("/joke/{id:[0-9]+}", jokeHandler).Methods("GET")

	godotenv.Load()

	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}

	srv := &http.Server{
		Handler:      r,
		Addr:         ":" + port,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 5 * time.Second,
	}

	log.Printf("Server listening to port %s", port)
	log.Fatal(srv.ListenAndServe())
}
