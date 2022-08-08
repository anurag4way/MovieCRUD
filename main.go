package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Movie struct {
	ID string `json:"id"`
	Isbn string `json:"isbn"`
	Title string `json:"title"`
	Director *Director `json:"director"`
}

type Director struct {
	Firstname string `json:"firstname"`
	Last_Name string `json:"lastName"`
}

var movies []Movie
func getMovies(w http.ResponseWriter, r *http.Request){
	w.Header().Set("content-type", "application/json")
	json.NewEncoder(w).Encode(movies)
}

func getMovie(w http.ResponseWriter, r *http.Request){
w.Header().Set("content-type", "application/json")
params := mux.Vars(r)

for _,item := range movies {

	if item.ID == params["id"]{
		json.NewEncoder(w).Encode(item)
	}

 }
 }
func createMovie(w http.ResponseWriter, r *http.Request){
	w.Header().Set("content-type", "application/json")
	var movie Movie
	_= json.NewDecoder(r.Body).Decode(&movie)
	movie.ID = strconv.Itoa(rand.Intn(10000000))
movies = append(movies, movie)
json.NewEncoder(w).Encode(movie)
}
func updateMovie(w http.ResponseWriter, r *http.Request){
w.Header().Set("content-type", "application/json")
params :=mux.Vars(r)

for index, item := range movies {

	if item.ID == params["id"] {
		movies = append(movies[:index], movies[index+1:]...)
		var movie Movie
	_= json.NewDecoder(r.Body).Decode(&movie)
	movie.ID=params["id"]
	movies=append(movies, movie)
	json.NewEncoder(w).Encode(movie)
	return
	}
}
}
func deleteMovie(w http.ResponseWriter, r *http.Request){
	w.Header().Set("content-type", "application/json")
	params := mux.Vars(r)
	for index, item := range movies {
 if item.ID==params["id"] {

	movies = append(movies[:index], movies[index+1:]...)
	break	
 }

 }
 json.NewEncoder(w).Encode(movies)
}

func main(){
movies = append(movies, Movie{ID:"1",Isbn: "3234",Title:"loc",Director:&Director{Firstname: "Anurag",Last_Name: "gupta"}})
movies = append(movies, Movie{ID:"2",Isbn: "2234",Title:"pyarloc",Director:&Director{Firstname: "ram",Last_Name: "jha"}})
movies = append(movies, Movie{ID:"3",Isbn: "4234",Title:"4way",Director:&Director{Firstname: "hassan",Last_Name: "ahmed"}})
	r:=mux.NewRouter()
	r.HandleFunc("/movies", getMovies).Methods("GET")
	r.HandleFunc("/movies/{id}", getMovie).Methods("GET")
	r.HandleFunc("/movie", createMovie).Methods("POST")
	r.HandleFunc("/movies/{id}", updateMovie).Methods("PUT")
	r.HandleFunc("/movies/{id}", deleteMovie).Methods("DELETE")
	fmt.Printf("starting server at port 8000\n")
	log.Fatal(http.ListenAndServe(":8000", r))
}

