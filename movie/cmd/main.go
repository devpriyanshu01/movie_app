// package main

// import (
// 	"log"
// 	"movieapp/movie/internal/controller/movie"
// 	metadatagateway "movieapp/movie/internal/gateway/metadata/http"
// 	ratinggateway "movieapp/movie/internal/gateway/rating/http"
// 	"movieapp/movie/internal/handler/http"
// 	httphandler "movieapp/movie/internal/handler/http"
// )

// func main() {
// 	log.Println("Starting the movie service")
// 	metadataGateway := metadatagateway.New("localhost:8081")
// 	ratingGateway := ratinggateway.New("localhost:8082")
// 	ctrl := movie.New(ratingGateway, metadataGateway)
// 	h := httphandler.New(ctrl)
// 	http.Handle("/movie", http.HandlerFunc(h.GetMovieDetails))
// 	if err := http.ListenAnsServe(":8083", nil); err != nil {
// 		panic(err)
// 	}
// }

package main

import (
	"log"
	"net/http"

	"movieapp/movie/internal/controller/movie"
	metadatagateway "movieapp/movie/internal/gateway/metadata/http"
	ratinggateway "movieapp/movie/internal/gateway/rating/http"
	httphandler "movieapp/movie/internal/handler/http"
)

func main() {
	log.Println("Starting the movie service")
	metadataGateway := metadatagateway.New("localhost:8081")
	ratingGateway := ratinggateway.New("localhost:8082")
	ctrl := movie.New(ratingGateway, metadataGateway)
	h := httphandler.New(ctrl)
	http.Handle("/movie", http.HandlerFunc(h.GetMovieDetails))
	if err := http.ListenAndServe(":8083", nil); err != nil {
		panic(err)
	}
}
