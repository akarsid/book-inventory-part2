package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func initializeRouter() {
	r := mux.NewRouter()

	r.HandleFunc("/books",GetBooks).Methods("GET")
	r.HandleFunc("/books/{bookId}",GetBook).Methods("GET")
	r.HandleFunc("/bookscount", getBooksCount).Methods("GET")
	r.HandleFunc("/storeBooks",CreateBook).Methods("POST")
	r.HandleFunc("/authorbooks/{authname}",getBooksByAuthor).Methods("GET")
	r.HandleFunc("/authors",GetAuthors).Methods("GET")
	log.Fatal(http.ListenAndServe(":9000",r))
}

func main() {
	
	InitialMigration()
	initializeRouter()
}