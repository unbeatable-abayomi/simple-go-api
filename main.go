package main

import (
	"encoding/json"
	"fmt"
	//"encoding/json"
	"log"
	"net/http"

	"math/rand"
	"strconv"
	"github.com/gorilla/mux"
)

//Book Struct {Model}

type Book struct{
	ID string  `json:"id"`
	Isbn string  `json:"isbn"`
	Title string  `json:"title"`
	Author *Author  `json:"author"`
}


//Author Struct

type Author struct {
	Firstname string `json:"firstname"`
	Lastname string `json:"lastname"`

}

//Init books var as a slice Book struct

var books []Book 

//Get All Books
func getBooks (w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
   json.NewEncoder(w).Encode(books)
}
//Get a single book
func getBook (w http.ResponseWriter, r *http.Request){
   params := mux.Vars(r) //Get params
   //loop through books
   for _, item := range books{
	if item.ID == params["id"]{
		json.NewEncoder(w).Encode(item)
		return
	}
   }
   json.NewEncoder(w).Encode(&Book{})
}

//Create a New Book
func createBook (w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	var book Book
	_ = json.NewDecoder(r.Body).Decode(&book)
	book.ID = strconv.Itoa(rand.Intn(10000000)) // Mock ID
	books = append(books, book)
	json.NewEncoder(w).Encode(book)
}

//Update a book
func updateBook (w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	for index ,item := range books {
		if item.ID == params["id"] {
			books = append(books[:index],books[index+1:]...)
			var book Book
	_ = json.NewDecoder(r.Body).Decode(&book)
	book.ID = params["id"] // Mock ID
	books = append(books, book)
	json.NewEncoder(w).Encode(book)
	return
		}
		
	}
	json.NewEncoder(w).Encode(books)

}
//Delete a booj
func deleteBook (w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	for index ,item := range books {
		if item.ID == params["id"] {
			books = append(books[:index],books[index+1:]...)
			break
		}
		
	}
	json.NewEncoder(w).Encode(books)
}


func main (){
	fmt.Println("Hello World")
	//Init Router
	r := mux.NewRouter()
//Mock Data -@todo - implement
	books = append(books, Book{ID: "1", Isbn: "448793",Title: "Book One",Author: &Author{Firstname: "Jane",Lastname: "Doe"}})
	books = append(books, Book{ID: "2", Isbn: "400793",Title: "Book Two",Author: &Author{Firstname: "Steve",Lastname: "Smith"}})
	books = append(books, Book{ID: "3", Isbn: "4910793",Title: "Book Three",Author: &Author{Firstname: "Ola",Lastname: "Park"}})
	//Route Handlers / Endpoints

	r.HandleFunc("/api/books", getBooks).Methods("GET")
	r.HandleFunc("/api/books/{id}", getBook).Methods("GET")
	r.HandleFunc("/api/books", createBook).Methods("POST")
	r.HandleFunc("/api/books/{id}", updateBook).Methods("PUT")
	r.HandleFunc("/api/books/{id}", deleteBook).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8080",r))

	


}