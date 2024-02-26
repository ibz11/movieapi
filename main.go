package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"


	"strconv"
	"math/rand"

	"github.com/gorilla/mux"
)

type Movie struct {
	ID       string    `json:"id"`
	Isbn     string    `json:"isbn"`
	Title    string    `json:"title"`
	Director *Director `json:"director"`
}

type Director struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}

var movies []Movie

// Route function
func getMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	fmt.Println("Endpoint hit: GetMovies")
	json.NewEncoder(w).Encode(movies)
}

func getAMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for _, item := range movies {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	fmt.Println("Endpoint hit: Get A Movie")
}

func createMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var movie Movie
	_ = json.NewDecoder(r.Body).Decode(&movie)
	movie.ID = strconv.Itoa(rand.Intn(10000000))
	movies = append(movies, movie)
	json.NewEncoder(w).Encode(movie)

}

func updateMovie(w http.ResponseWriter, r *http.Request) {
	//set json content type
	w.Header().Set("Content-Type", "application/json")

	//params
	params := mux.Vars(r)

	//loop
	var movie Movie
	for index, item := range movies {
		if item.ID == params["id"] {
			//Delete the movie with the id you sent
			movies = append(movies[:index], movies[index+1:]...)
			
			_ = json.NewDecoder(r.Body).Decode(&movie)
			movie.ID = params["id"]
			movies = append(movies, movie)
		
			return

		}
		json.NewEncoder(w).Encode("Movie is updated")
	}
	fmt.Println("Endpoint hit: Update A Movie")

}

func deleteMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range movies {
		if item.ID == params["id"] {
			movies = append(movies[:index], movies[index+1:]...)
			break
		}
	}
	fmt.Println("Endpoint hit: Delete A Movie")
	json.NewEncoder(w).Encode(movies)
}

func main() {
	r := mux.NewRouter().StrictSlash(true)

	movies = append(movies, Movie{ID: "1", Isbn: "345227", Title: "King Kong", Director: &Director{Firstname: "John", Lastname: "Baldwin"}})
	movies = append(movies, Movie{ID: "2", Isbn: "455227", Title: "Star Wars", Director: &Director{Firstname: "Racheal", Lastname: "John"}})

	r.HandleFunc("/movies", getMovies).Methods("GET")
	r.HandleFunc("/movies/{id}", getAMovie).Methods("GET")
	r.HandleFunc("/movie", createMovie).Methods("POST")
	r.HandleFunc("/movies/{id}", updateMovie).Methods("PUT")
	r.HandleFunc("/movies/{id}", deleteMovie).Methods("DELETE")

	fmt.Printf("Starting server at http://localhost:8000\n")
	log.Fatal(http.ListenAndServe(":8000", r))
}
