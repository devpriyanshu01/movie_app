package main

import (
	"log"
	"movieapp/rating/internal/controller/rating"
	"movieapp/rating/internal/repository/memory"
	"net/http"
	httpHandler "movieapp/rating/internal/handler/http"
)

func main() {
	log.Println("Starting the rating service.")

	repo := memory.New()
	ctrl := rating.New(repo)
	h := httpHandler.New(ctrl)
	http.Handle("/rating", http.HandlerFunc(h.Handle))
	if err := http.ListenAndServe(":8082", nil); err != nil {
		panic(err)
	}
}
