package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Movie struct {
	Id       int       `json:"id"`
	Name     string    `json:"name"`
	Rating   int       `json:"title"`
	Director *Director `json:"director`
}

type Director struct {
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
}

var movies []Movie

/*graf 31:30 po jednej wrócić
GET ALL getMovies
GET BY ID getMovie
CREATE createMovie
UPDATE updateMovie
DELETE deleteMovie

*/

func getMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(movies)
}

func deleteMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	for index, item := range movies {

		if strconv.Itoa(item.Id) == params["id"] {
			movies = append(movies[:index], movies[index+1:]...) //przydatna składnia
			break
		}
	}
	json.NewEncoder(w).Encode(movies)
}

func getMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	for _, item := range movies {

		if strconv.Itoa(item.Id) == params["id"] {
			json.NewEncoder(w).Encode(item)
			break
		}
	}
}

func main() {
	r := mux.NewRouter()

	movies = append(movies, Movie{Id: 1, Name: "LOTR", Rating: 10, Director: &Director{FirstName: "John", LastName: "COO"}})
	movies = append(movies, Movie{Id: 2, Name: "Test2", Rating: 6, Director: &Director{FirstName: "Sebastian", LastName: "Co"}})
	movies = append(movies, Movie{Id: 3, Name: "Test3", Rating: 3, Director: &Director{FirstName: "John", LastName: "COO"}})
	r.HandleFunc("movie/get-all", getMovies).Methods("GET")          //GET ALL getMovies
	r.HandleFunc("movie/get-movie/{Id}", getMovie).Methods("GET")    //GET BY ID getMovie
	r.HandleFunc("/movie/create", createMovie).Methods("POST")       //CREATE createMovie
	r.HandleFunc("movie/update/{Id}", updateMovie).Methods("PUT")    //UPDATE updateMovie
	r.HandleFunc("movie/delete/{Id}", deleteMovie).Methods("DELETE") //DELETE deleteMovie

	fmt.Printf("Starting server at port 8000\n")
	log.Fatal(http.ListenAndServe(":8000", r))
}
