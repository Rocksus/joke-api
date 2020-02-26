package main

import (
	"log"
	"math/rand"
	"net/http"
	"time"

	"github.com/Rocksus/joke-api/internal/repository/joke"
	"github.com/gorilla/mux"
)

func init() {
	rand.Seed(time.Now().UTC().UnixNano())
}

func main() {
	r := mux.NewRouter()

	joke.Load("../data/jokes.json")
	jokeHandler := joke.InitHandler()
	randomHandler := joke.InitRandomHandler()

	r.HandleFunc("/jokes/random", randomHandler).Methods("GET")
	r.HandleFunc("/jokes/{category}/random", randomHandler).Methods("GET")
	r.HandleFunc("/jokes/{id:[0-9]+}", jokeHandler).Methods("GET")

	srv := &http.Server{
		Handler:      r,
		Addr:         ":8000",
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 5 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}
